package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

type stringList []string

func (s *stringList) String() string {
	return strings.Join(*s, ",")
}

func (s *stringList) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return errors.Wrap(err, `failed to unmarshal string for stringList`)
		}
		*s = []string{v}
		return nil
	}

	var v []string
	if err := json.Unmarshal(data, &v); err != nil {
		return errors.Wrap(err, `failed to unmarshal string list for stringList`)
	}
	*s = v
	return nil
}

type Endpoint struct {
	file       string
	methodName string
	Group      string     `json:"group,omitempty"`
	Name       string     `json:"name"` // e.g. "chat.PostMessage"
	JSON       stringList `json:"json"`
	Args       []Argument `json:"args,omitempty"`
	ReturnType stringList `json:"return,omitempty"`
	SkipToken  bool       `json:"skip_token,omitempty"`
	Proxy bool `json:"proxy"`
}

type Argument struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Required  bool   `json:"required,omitempty"`
	Default   string `json:"default,omitempty"`
	QueryName string `json:"query_name,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

func camelit(s string) string {
	var buf bytes.Buffer

	var upnext bool
	for i, r := range s {
		if i == 0 || upnext {
			buf.WriteRune(unicode.ToUpper(r))
			upnext = false
			continue
		}

		if !unicode.IsLetter(r) {
			upnext = true
			continue
		}

		buf.WriteRune(r)
	}
	return buf.String()
}

func _main() error {
	var endpoints []*Endpoint

	fh, err := os.Open("endpoints.json")
	if err != nil {
		return errors.Wrap(err, `failed to open endpoints.json`)
	}
	defer fh.Close()

	if err := json.NewDecoder(fh).Decode(&endpoints); err != nil {
		return errors.Wrap(err, `failed to decode endpoints.json`)
	}

	sort.Slice(endpoints, func(i, j int) bool {
		return strings.Compare(endpoints[i].Name, endpoints[j].Name) < 0
	})

	group := map[string][]*Endpoint{}
	for _, endpoint := range endpoints {
		i := strings.LastIndexByte(endpoint.Name, '.')
		endpoint.file = filepath.Join("server", strings.Replace(endpoint.Name[:i], ".", "_", -1)+"_gen.go")
		if len(endpoint.Group) == 0 {
			endpoint.Group = camelit(endpoint.Name[:i])
		}
		endpoint.methodName = camelit(endpoint.Name[i+1:])
		group[endpoint.file] = append(group[endpoint.file], endpoint)
	}

	if err := generateServerFile(endpoints); err != nil {
		return errors.Wrap(err, `failed to generate server file`)
	}

	if err := generateMockServerFile(endpoints); err != nil {
		return errors.Wrap(err, `failed to generate mock server file`)
	}

	if err := generateProxyServerFile(endpoints); err != nil {
		return errors.Wrap(err, `failed to generate proxy server file`)
	}

	return nil
}

func generateServerFile(endpoints []*Endpoint) error {
	var buf bytes.Buffer
	buf.WriteString("package server")
	buf.WriteString("\n\n// Auto-generated by internal/cmd/genserver/genserver.go. DO NOT EDIT!")
	buf.WriteString("\n\nimport (")
	for _, pkg := range []string{"net/http", "strings"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	buf.WriteString("\n")
	for _, pkg := range []string{} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	buf.WriteString("\n)")

	fmt.Fprintf(&buf, "\n\nfunc unimplemented(w http.ResponseWriter, r *http.Request) {")
	fmt.Fprintf(&buf, "\nhttp.Error(w, `method ` + r.URL.Path[1:] + ` is unimplemented on this server`, http.StatusBadRequest)")
	fmt.Fprintf(&buf, "\n}") // end func Unimplemented

	fmt.Fprintf(&buf, "\n\nfunc New(options ...Option) *Server {")
	fmt.Fprintf(&buf, "\nprefix := \"/api\"")
	fmt.Fprintf(&buf, "\nfor _, option := range options {")
	fmt.Fprintf(&buf, "\nswitch option.Name() {")
	fmt.Fprintf(&buf, "\ncase optkeyPrefix:")
	fmt.Fprintf(&buf, "\nprefix = option.Value().(string)")
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\n}")

	fmt.Fprintf(&buf, "\n\nreturn &Server{")
	fmt.Fprintf(&buf, "\nprefix: prefix,")
	fmt.Fprintf(&buf, "\nhandlers: map[string]http.Handler{")
	for _, endpoint := range endpoints {
		fmt.Fprintf(&buf, "\n%s: http.HandlerFunc(unimplemented),", strconv.Quote(endpoint.Name))
	}
	fmt.Fprintf(&buf, "\n},") // end of map
	fmt.Fprintf(&buf, "\n}")  // end of Server
	fmt.Fprintf(&buf, "\n}")  // end of New

	fmt.Fprintf(&buf, "\n\n// Handle sets the http.Handler for the given slack method.")
	fmt.Fprintf(&buf, "\nfunc (s *Server) Handle(method string, h http.Handler) {")
	fmt.Fprintf(&buf, "\ns.muHandlers.Lock()")
	fmt.Fprintf(&buf, "\ndefer s.muHandlers.Unlock()")
	fmt.Fprintf(&buf, "\ns.handlers[method] = h")
	fmt.Fprintf(&buf, "\n}") // end of func Handler

	fmt.Fprintf(&buf, "\n\nfunc (s *Server) GetHandler(method string) (http.Handler, bool) {")
	fmt.Fprintf(&buf, "\ns.muHandlers.RLock()")
	fmt.Fprintf(&buf, "\ndefer s.muHandlers.RUnlock()")
	fmt.Fprintf(&buf, "\nh, ok := s.handlers[method]")
	fmt.Fprintf(&buf, "\nreturn h, ok")
	fmt.Fprintf(&buf, "\n}") // end of GetHandler

	fmt.Fprintf(&buf, "\nfunc (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {")
	fmt.Fprintf(&buf, "\nr.URL.Path = strings.TrimPrefix(r.URL.Path, s.prefix)")
	fmt.Fprintf(&buf, "\nh, ok := s.GetHandler(r.URL.Path[1:])")
	fmt.Fprintf(&buf, "\nif !ok {")
	fmt.Fprintf(&buf, "\nhttp.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)")
	fmt.Fprintf(&buf, "\nreturn")
	fmt.Fprintf(&buf, "\n}") // end if
	fmt.Fprintf(&buf, "\nh.ServeHTTP(w, r)")
	fmt.Fprintf(&buf, "\n}") // end func ServeHTTP

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return errors.Wrap(err, `failed to format code`)
	}

	file := filepath.Join("server", "server.go")
	fh, err := os.Create(file)
	if err != nil {
		return errors.Wrapf(err, `failed to open file %s for writing`, file)
	}
	defer fh.Close()

	fh.Write(formatted)
	return nil
}

func generateMockServerFile(endpoints []*Endpoint) error {
	var buf bytes.Buffer
	buf.WriteString("package mockserver")
	buf.WriteString("\n\n// Auto-generated by internal/cmd/genserver/genserver.go. DO NOT EDIT!")
	buf.WriteString("\n\nimport (")
	for _, pkg := range []string{"bytes", "encoding/json", "net/http", "strconv", "sync", "time"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	buf.WriteString("\n")
	for _, pkg := range []string{"github.com/brontolinux/slack", "github.com/brontolinux/slack/objects", "github.com/brontolinux/slack/server"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	buf.WriteString("\n)")

	fmt.Fprintf(&buf, "\n\ntype Handler struct {")
	fmt.Fprintf(&buf, "\nmuTokens sync.RWMutex")
	fmt.Fprintf(&buf, "\ntokens map[string]struct{}")
	fmt.Fprintf(&buf, "\n}") // end type Handler

	fmt.Fprintf(&buf, "\n\ntype Option interface {")
	fmt.Fprintf(&buf, "\nName() string")
	fmt.Fprintf(&buf, "\nValue() interface{}")
	fmt.Fprintf(&buf, "\n}") // end type Option

	fmt.Fprintf(&buf, "\n\n type option struct {")
	fmt.Fprintf(&buf, "\nname string")
	fmt.Fprintf(&buf, "\nvalue interface{}")
	fmt.Fprintf(&buf, "\n}") // end type option

	fmt.Fprintf(&buf, "\n\nfunc (o *option) Name() string {")
	fmt.Fprintf(&buf, "\nreturn o.name")
	fmt.Fprintf(&buf, "\n}") // end Name()

	fmt.Fprintf(&buf, "\n\nfunc (o *option) Value() interface{} {")
	fmt.Fprintf(&buf, "\nreturn o.value")
	fmt.Fprintf(&buf, "\n}") // end Value()

	fmt.Fprintf(&buf, "\n\nconst (")
	fmt.Fprintf(&buf, "\noptTokenKey = `token`")
	fmt.Fprintf(&buf, "\n)")

	fmt.Fprintf(&buf, "\n\nfunc WithToken(s string) Option {")
	fmt.Fprintf(&buf, "\nreturn &option{")
	fmt.Fprintf(&buf, "\nname: optTokenKey,")
	fmt.Fprintf(&buf, "\nvalue: s,")
	fmt.Fprintf(&buf, "\n}") // end return option
	fmt.Fprintf(&buf, "\n}") // end WithToken()


	fmt.Fprintf(&buf, "\n// New creates a new mock Slack API Handler object. You may pass optional")
	fmt.Fprintf(&buf, "\n// parameters, which can be one of the following:")
	fmt.Fprintf(&buf, "\n//")
	fmt.Fprintf(&buf, "\n// `mockserver.WithToken(string)`: specifies the token to accept. Multiple")
	fmt.Fprintf(&buf, "\n// tokens may be specified, and given any one of them, the server will")
	fmt.Fprintf(&buf, "\n// accept the request.")
	fmt.Fprintf(&buf, "\nfunc New(options ...Option) *Handler {")
	fmt.Fprintf(&buf, "\ntokens := make(map[string]struct{})")
	fmt.Fprintf(&buf, "\nfor _, option := range options {")
	fmt.Fprintf(&buf, "\nswitch option.Name() {")
	fmt.Fprintf(&buf, "\ncase optTokenKey:")
	fmt.Fprintf(&buf, "\ntokens[option.Value().(string)] = struct{}{}")
	fmt.Fprintf(&buf, "\n}") // end switch option.Name()
	fmt.Fprintf(&buf, "\n}") // end for _, option := range options
	fmt.Fprintf(&buf, "\nreturn &Handler{")
	fmt.Fprintf(&buf, "\ntokens: tokens,")
	fmt.Fprintf(&buf, "\n}") // end return Handler
	fmt.Fprintf(&buf, "\n}") // end New

	fmt.Fprintf(&buf, "\n\nfunc (h *Handler) validateToken(r *http.Request) bool {")
	fmt.Fprintf(&buf, "\nif len(h.tokens) == 0 { // only check token if specified to do so")
	fmt.Fprintf(&buf, "\nreturn true")
	fmt.Fprintf(&buf, "\n}") // end if len(h.tokens)
	fmt.Fprintf(&buf, "\ntoken := r.FormValue(`token`)")
	fmt.Fprintf(&buf, "\n\nif len(token) == 0 {")
	fmt.Fprintf(&buf, "\nreturn false")
	fmt.Fprintf(&buf, "\n}") // end if len(token) == 0
	fmt.Fprintf(&buf, "\n\n_, ok := h.tokens[token]")
	fmt.Fprintf(&buf, "\nreturn ok")
	fmt.Fprintf(&buf, "\n}") // end validateToken

	fmt.Fprintf(&buf, "\n\nfunc (h *Handler) InstallHandlers(s *server.Server) {")
	for _, endpoint := range endpoints {
		fmt.Fprintf(&buf, "\ns.Handle(%s, http.HandlerFunc(h.Handle%s%s))", strconv.Quote(endpoint.Name), endpoint.Group, endpoint.methodName)
	}
	fmt.Fprintf(&buf, "\n}") // end Server

	returnTypes := make(map[string][]string)
	for _, endpoint := range endpoints {
		fmt.Fprintf(&buf, "\n\n// Handle%s%s is the default handler method for the Slack %s API", endpoint.Group, endpoint.methodName, endpoint.Name)
		fmt.Fprintf(&buf, "\nfunc (h *Handler) Handle%s%s(w http.ResponseWriter, r *http.Request) {", endpoint.Group, endpoint.methodName)
		fmt.Fprintf(&buf, "\nif err := r.ParseForm(); err != nil {")
		fmt.Fprintf(&buf, "\nhttp.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)")
		fmt.Fprintf(&buf, "\nreturn")
		fmt.Fprintf(&buf, "\n}") // end if err := r.ParseForm()

		if !endpoint.SkipToken {
			fmt.Fprintf(&buf, "\nif !h.validateToken(r) {")
			fmt.Fprintf(&buf, "\nhttp.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)")
			fmt.Fprintf(&buf, "\nreturn")
			fmt.Fprintf(&buf, "\n}") // end if !h.validateToken()
		}

		fmt.Fprintf(&buf, "\nvar c slack.%s%sCall", endpoint.Group, endpoint.methodName)
		fmt.Fprintf(&buf, "\nif err := c.FromValues(r.Form); err != nil {")
		fmt.Fprintf(&buf, "\nhttp.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)")
		fmt.Fprintf(&buf, "\nreturn")
		fmt.Fprintf(&buf, "\n}") // end if err := c.FormValues

		fmt.Fprintf(&buf, "\nvar buf bytes.Buffer")
		fmt.Fprintf(&buf, "\nif err := json.NewEncoder(&buf).Encode(StockResponse(%s)); err != nil {", strconv.Quote(endpoint.Name))
		fmt.Fprintf(&buf, "\nhttp.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)")
		fmt.Fprintf(&buf, "\nreturn")
		fmt.Fprintf(&buf, "\n}") // end if err := json.NewEncoder
		fmt.Fprintf(&buf, "\nw.Header().Set(`Content-Type`, `application/json; charset=utf-8`)")
		fmt.Fprintf(&buf, "\nw.WriteHeader(http.StatusOK)")
		fmt.Fprintf(&buf, "\nbuf.WriteTo(w)")
		fmt.Fprintf(&buf, "\n}") // end func Handle%s%s

		if rt := endpoint.ReturnType.String(); rt != "" {
			returnTypes[rt] = append(returnTypes[rt], endpoint.Name)
		}
	}

	var returnTypeNames []string
	for rt := range returnTypes {
		returnTypeNames = append(returnTypeNames, rt)
	}
	sort.Strings(returnTypeNames)

	fmt.Fprintf(&buf, "\n\nfunc StockResponse(method string) interface{} {")
	fmt.Fprintf(&buf, "\nswitch method {")
	for _, typ := range returnTypeNames {
		methods := returnTypes[typ]
		sort.Strings(methods)

		fmt.Fprintf(&buf, "\ncase ")
		for i, method := range methods {
			fmt.Fprintf(&buf, "%s", strconv.Quote(method))
			if i < len(methods)-1 {
				fmt.Fprintf(&buf, ",")
			} else {
				fmt.Fprintf(&buf, ":")
			}
		}
		fmt.Fprintf(&buf, "\nreturn stock%s()", camelit(typ))
	}
	fmt.Fprintf(&buf, "\ndefault:")
	fmt.Fprintf(&buf, "\nreturn objects.GenericResponse{")
	fmt.Fprintf(&buf, "\nOK: true,")
	fmt.Fprintf(&buf, "\nTimestamp: strconv.FormatInt(time.Now().Unix(), 10),")
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\n}") // end switch
	fmt.Fprintf(&buf, "\n}") // end func StockResponse

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return errors.Wrap(err, `failed to format code`)
	}

	file := filepath.Join("server", "mockserver", "mockserver.go")
	fh, err := os.Create(file)
	if err != nil {
		return errors.Wrapf(err, `failed to open file %s for writing`, file)
	}
	defer fh.Close()

	fh.Write(formatted)
	return nil
}


func generateProxyServerFile(endpoints []*Endpoint) error {
	var buf bytes.Buffer
	buf.WriteString("package proxyserver")
	buf.WriteString("\n\n// Auto-generated by internal/cmd/genserver/genserver.go. DO NOT EDIT!")
	buf.WriteString("\n\nimport (")
	for _, pkg := range []string{"io", "net/http"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	buf.WriteString("\n")
	for _, pkg := range []string{"github.com/brontolinux/slack", "github.com/brontolinux/slack/server", "github.com/brontolinux/slack/server/mockserver"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	buf.WriteString("\n)")

	fmt.Fprintf(&buf, "\ntype Handler struct {")
	fmt.Fprintf(&buf, "\nmock *mockserver.Handler")
	fmt.Fprintf(&buf, "\ntoken string")
	fmt.Fprintf(&buf, "\n}") // end type Handler

	fmt.Fprintf(&buf, "\n\nfunc New(token string) *Handler {")
	fmt.Fprintf(&buf, "\nreturn &Handler{")
	fmt.Fprintf(&buf, "\nmock: mockserver.New(mockserver.WithToken(token)),")
	fmt.Fprintf(&buf, "\ntoken: token,")
	fmt.Fprintf(&buf, "\n}") // end return Handler
	fmt.Fprintf(&buf, "\n}") // end New

	fmt.Fprintf(&buf, "\n\nfunc (h *Handler) ProxyHandler(w http.ResponseWriter, r *http.Request) {")
	fmt.Fprintf(&buf, "\nres, err := http.Post(slack.DefaultAPIEndpoint + r.URL.Path[1:], r.Header.Get(`Content-Type`), r.Body)")
	fmt.Fprintf(&buf, "\nif err != nil {")
	fmt.Fprintf(&buf, "\nhttp.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)")
	fmt.Fprintf(&buf, "\nreturn")
	fmt.Fprintf(&buf, "\n}") // end if err != nil
	fmt.Fprintf(&buf, "\nfor hdrname, hdrvals := range res.Header {")
	fmt.Fprintf(&buf, "\nfor _, hdrval := range hdrvals {")
	fmt.Fprintf(&buf, "\nw.Header().Add(hdrname, hdrval)")
	fmt.Fprintf(&buf, "\n}") // end for _, hdrval
	fmt.Fprintf(&buf, "\n}") // end for hdrname, hdrvals
	fmt.Fprintf(&buf, "\nw.WriteHeader(res.StatusCode)")
	fmt.Fprintf(&buf, "\nio.Copy(w, res.Body)")
	fmt.Fprintf(&buf, "\ndefer res.Body.Close()")
	fmt.Fprintf(&buf, "\n}") // end ProxyHandler

	fmt.Fprintf(&buf, "\n\nfunc (h *Handler) InstallHandlers(s *server.Server) {")
	for _, endpoint := range endpoints {
		if endpoint.Proxy {
			fmt.Fprintf(&buf, "\ns.Handle(%s, http.HandlerFunc(h.ProxyHandler))", strconv.Quote(endpoint.Name))
		} else {
			fmt.Fprintf(&buf, "\ns.Handle(%s, http.HandlerFunc(h.mock.Handle%s%s))", strconv.Quote(endpoint.Name), endpoint.Group, endpoint.methodName)
		}
	}
	fmt.Fprintf(&buf, "\n}") // end Server

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return errors.Wrap(err, `failed to format code`)
	}

	file := filepath.Join("server", "proxyserver", "proxyserver.go")
	fh, err := os.Create(file)
	if err != nil {
		return errors.Wrapf(err, `failed to open file %s for writing`, file)
	}
	defer fh.Close()

	fh.Write(formatted)
	return nil
}

