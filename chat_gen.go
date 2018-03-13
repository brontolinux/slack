package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/lestrrat-go/slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = objects.EpochTime(0)

// ChatDeleteCall is created by ChatService.Delete method call
type ChatDeleteCall struct {
	service   *ChatService
	asUser    bool
	channel   string
	timestamp string
}

// ChatGetPermalinkCall is created by ChatService.GetPermalink method call
type ChatGetPermalinkCall struct {
	service          *ChatService
	channel          string
	messageTimestamp string
}

// ChatMeMessageCall is created by ChatService.MeMessage method call
type ChatMeMessageCall struct {
	service *ChatService
	channel string
	text    string
}

// ChatPostMessageCall is created by ChatService.PostMessage method call
type ChatPostMessageCall struct {
	service     *ChatService
	asUser      bool
	attachments objects.AttachmentList
	channel     string
	escapeText  bool
	iconEmoji   string
	iconURL     string
	linkNames   bool
	markdown    bool
	parse       string
	text        string
	unfurlLinks bool
	unfurlMedia bool
	username    string
}

// ChatUnfurlCall is created by ChatService.Unfurl method call
type ChatUnfurlCall struct {
	service          *ChatService
	channel          string
	timestamp        string
	unfurls          string
	userAuthRequired bool
}

// ChatUpdateCall is created by ChatService.Update method call
type ChatUpdateCall struct {
	service     *ChatService
	asUser      bool
	attachments objects.AttachmentList
	channel     string
	linkNames   bool
	parse       string
	text        string
	timestamp   string
}

// Delete creates a ChatDeleteCall object in preparation for accessing the chat.delete endpoint
func (s *ChatService) Delete(channel string) *ChatDeleteCall {
	var call ChatDeleteCall
	call.service = s
	call.channel = channel
	return &call
}

// AsUser sets the value for optional asUser parameter
func (c *ChatDeleteCall) AsUser(asUser bool) *ChatDeleteCall {
	c.asUser = asUser
	return c
}

// Timestamp sets the value for optional timestamp parameter
func (c *ChatDeleteCall) Timestamp(timestamp string) *ChatDeleteCall {
	c.timestamp = timestamp
	return c
}

// ValidateArgs checks that all required fields are set in the ChatDeleteCall object
func (c *ChatDeleteCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChatDeleteCall object as url.Values
func (c *ChatDeleteCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.asUser {
		v.Set("as_user", "true")
	}

	v.Set("channel", c.channel)

	if len(c.timestamp) > 0 {
		v.Set("ts", c.timestamp)
	}
	return v, nil
}

// Do executes the call to access chat.delete endpoint
func (c *ChatDeleteCall) Do(ctx context.Context) (*objects.ChatResponse, error) {
	const endpoint = "chat.delete"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.ChatResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.delete`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChatResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChatDeleteCall) FromValues(v url.Values) error {
	var tmp ChatDeleteCall
	if raw := strings.TrimSpace(v.Get("as_user")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "as_user"`)
		}
		tmp.asUser = parsed
	}
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("ts")); len(raw) > 0 {
		tmp.timestamp = raw
	}
	*c = tmp
	return nil
}

// GetPermalink creates a ChatGetPermalinkCall object in preparation for accessing the chat.getPermalink endpoint
func (s *ChatService) GetPermalink(channel string, messageTimestamp string) *ChatGetPermalinkCall {
	var call ChatGetPermalinkCall
	call.service = s
	call.channel = channel
	call.messageTimestamp = messageTimestamp
	return &call
}

// ValidateArgs checks that all required fields are set in the ChatGetPermalinkCall object
func (c *ChatGetPermalinkCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.messageTimestamp) <= 0 {
		return errors.New(`required field messageTimestamp not initialized`)
	}
	return nil
}

// Values returns the ChatGetPermalinkCall object as url.Values
func (c *ChatGetPermalinkCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("message_ts", c.messageTimestamp)
	return v, nil
}

// Do executes the call to access chat.getPermalink endpoint
func (c *ChatGetPermalinkCall) Do(ctx context.Context) (*objects.PermalinkResponse, error) {
	const endpoint = "chat.getPermalink"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.PermalinkResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.getPermalink`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.PermalinkResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChatGetPermalinkCall) FromValues(v url.Values) error {
	var tmp ChatGetPermalinkCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("message_ts")); len(raw) > 0 {
		tmp.messageTimestamp = raw
	}
	*c = tmp
	return nil
}

// MeMessage creates a ChatMeMessageCall object in preparation for accessing the chat.meMessage endpoint
func (s *ChatService) MeMessage(channel string) *ChatMeMessageCall {
	var call ChatMeMessageCall
	call.service = s
	call.channel = channel
	return &call
}

// Text sets the value for optional text parameter
func (c *ChatMeMessageCall) Text(text string) *ChatMeMessageCall {
	c.text = text
	return c
}

// ValidateArgs checks that all required fields are set in the ChatMeMessageCall object
func (c *ChatMeMessageCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChatMeMessageCall object as url.Values
func (c *ChatMeMessageCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	if len(c.text) > 0 {
		v.Set("text", c.text)
	}
	return v, nil
}

// Do executes the call to access chat.meMessage endpoint
func (c *ChatMeMessageCall) Do(ctx context.Context) (*objects.ChatResponse, error) {
	const endpoint = "chat.meMessage"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.ChatResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.meMessage`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChatResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChatMeMessageCall) FromValues(v url.Values) error {
	var tmp ChatMeMessageCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("text")); len(raw) > 0 {
		tmp.text = raw
	}
	*c = tmp
	return nil
}

// PostMessage creates a ChatPostMessageCall object in preparation for accessing the chat.postMessage endpoint
func (s *ChatService) PostMessage(channel string) *ChatPostMessageCall {
	var call ChatPostMessageCall
	call.service = s
	call.channel = channel
	return &call
}

// AsUser sets the value for optional asUser parameter
func (c *ChatPostMessageCall) AsUser(asUser bool) *ChatPostMessageCall {
	c.asUser = asUser
	return c
}

// SetAttachments sets the attachments list
func (c *ChatPostMessageCall) SetAttachments(attachments objects.AttachmentList) *ChatPostMessageCall {
	c.attachments = attachments
	return c
}

// Attachment appends to the attachments list
func (c *ChatPostMessageCall) Attachment(attachment *objects.Attachment) *ChatPostMessageCall {
	c.attachments.Append(attachment)
	return c
}

// EscapeText sets the value for optional escapeText parameter
func (c *ChatPostMessageCall) EscapeText(escapeText bool) *ChatPostMessageCall {
	c.escapeText = escapeText
	return c
}

// IconEmoji sets the value for optional iconEmoji parameter
func (c *ChatPostMessageCall) IconEmoji(iconEmoji string) *ChatPostMessageCall {
	c.iconEmoji = iconEmoji
	return c
}

// IconURL sets the value for optional iconURL parameter
func (c *ChatPostMessageCall) IconURL(iconURL string) *ChatPostMessageCall {
	c.iconURL = iconURL
	return c
}

// LinkNames sets the value for optional linkNames parameter
func (c *ChatPostMessageCall) LinkNames(linkNames bool) *ChatPostMessageCall {
	c.linkNames = linkNames
	return c
}

// Markdown sets the value for optional markdown parameter
func (c *ChatPostMessageCall) Markdown(markdown bool) *ChatPostMessageCall {
	c.markdown = markdown
	return c
}

// Parse sets the value for optional parse parameter
func (c *ChatPostMessageCall) Parse(parse string) *ChatPostMessageCall {
	c.parse = parse
	return c
}

// Text sets the value for optional text parameter
func (c *ChatPostMessageCall) Text(text string) *ChatPostMessageCall {
	c.text = text
	return c
}

// UnfurlLinks sets the value for optional unfurlLinks parameter
func (c *ChatPostMessageCall) UnfurlLinks(unfurlLinks bool) *ChatPostMessageCall {
	c.unfurlLinks = unfurlLinks
	return c
}

// UnfurlMedia sets the value for optional unfurlMedia parameter
func (c *ChatPostMessageCall) UnfurlMedia(unfurlMedia bool) *ChatPostMessageCall {
	c.unfurlMedia = unfurlMedia
	return c
}

// Username sets the value for optional username parameter
func (c *ChatPostMessageCall) Username(username string) *ChatPostMessageCall {
	c.username = username
	return c
}

// ValidateArgs checks that all required fields are set in the ChatPostMessageCall object
func (c *ChatPostMessageCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChatPostMessageCall object as url.Values
func (c *ChatPostMessageCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.asUser {
		v.Set("as_user", "true")
	}

	if len(c.attachments) > 0 {
		attachmentsEncoded, err := c.attachments.Encode()
		if err != nil {
			return nil, errors.Wrap(err, `failed to encode field`)
		}
		v.Set("attachments", attachmentsEncoded)
	}

	v.Set("channel", c.channel)

	if c.escapeText {
		v.Set("escapeText", "true")
	}

	if len(c.iconEmoji) > 0 {
		v.Set("iconEmoji", c.iconEmoji)
	}

	if len(c.iconURL) > 0 {
		v.Set("iconURL", c.iconURL)
	}

	if c.linkNames {
		v.Set("linkNames", "true")
	}

	if c.markdown {
		v.Set("markdown", "true")
	}

	if len(c.parse) > 0 {
		v.Set("parse", c.parse)
	}

	if len(c.text) > 0 {
		v.Set("text", c.text)
	}

	if c.unfurlLinks {
		v.Set("unfurlLinks", "true")
	}

	if c.unfurlMedia {
		v.Set("unfurlMedia", "true")
	}

	if len(c.username) > 0 {
		v.Set("username", c.username)
	}
	return v, nil
}

// Do executes the call to access chat.postMessage endpoint
func (c *ChatPostMessageCall) Do(ctx context.Context) (*objects.ChatResponse, error) {
	const endpoint = "chat.postMessage"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.ChatResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.postMessage`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChatResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChatPostMessageCall) FromValues(v url.Values) error {
	var tmp ChatPostMessageCall
	if raw := strings.TrimSpace(v.Get("as_user")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "as_user"`)
		}
		tmp.asUser = parsed
	}
	if raw := strings.TrimSpace(v.Get("attachments")); len(raw) > 0 {
		if err := tmp.attachments.Decode(raw); err != nil {
			return errors.Wrap(err, `failed to decode value "attachments"`)
		}
	}
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("escapeText")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "escapeText"`)
		}
		tmp.escapeText = parsed
	}
	if raw := strings.TrimSpace(v.Get("iconEmoji")); len(raw) > 0 {
		tmp.iconEmoji = raw
	}
	if raw := strings.TrimSpace(v.Get("iconURL")); len(raw) > 0 {
		tmp.iconURL = raw
	}
	if raw := strings.TrimSpace(v.Get("linkNames")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "linkNames"`)
		}
		tmp.linkNames = parsed
	}
	if raw := strings.TrimSpace(v.Get("markdown")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "markdown"`)
		}
		tmp.markdown = parsed
	}
	if raw := strings.TrimSpace(v.Get("parse")); len(raw) > 0 {
		tmp.parse = raw
	}
	if raw := strings.TrimSpace(v.Get("text")); len(raw) > 0 {
		tmp.text = raw
	}
	if raw := strings.TrimSpace(v.Get("unfurlLinks")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "unfurlLinks"`)
		}
		tmp.unfurlLinks = parsed
	}
	if raw := strings.TrimSpace(v.Get("unfurlMedia")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "unfurlMedia"`)
		}
		tmp.unfurlMedia = parsed
	}
	if raw := strings.TrimSpace(v.Get("username")); len(raw) > 0 {
		tmp.username = raw
	}
	*c = tmp
	return nil
}

// Unfurl creates a ChatUnfurlCall object in preparation for accessing the chat.unfurl endpoint
func (s *ChatService) Unfurl(channel string, timestamp string, unfurls string) *ChatUnfurlCall {
	var call ChatUnfurlCall
	call.service = s
	call.channel = channel
	call.timestamp = timestamp
	call.unfurls = unfurls
	return &call
}

// UserAuthRequired sets the value for optional userAuthRequired parameter
func (c *ChatUnfurlCall) UserAuthRequired(userAuthRequired bool) *ChatUnfurlCall {
	c.userAuthRequired = userAuthRequired
	return c
}

// ValidateArgs checks that all required fields are set in the ChatUnfurlCall object
func (c *ChatUnfurlCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	if len(c.timestamp) <= 0 {
		return errors.New(`required field timestamp not initialized`)
	}
	if len(c.unfurls) <= 0 {
		return errors.New(`required field unfurls not initialized`)
	}
	return nil
}

// Values returns the ChatUnfurlCall object as url.Values
func (c *ChatUnfurlCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("channel", c.channel)

	v.Set("ts", c.timestamp)

	v.Set("unfurls", c.unfurls)

	if c.userAuthRequired {
		v.Set("user_auth_required", "true")
	}
	return v, nil
}

// Do executes the call to access chat.unfurl endpoint
func (c *ChatUnfurlCall) Do(ctx context.Context) error {
	const endpoint = "chat.unfurl"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to chat.unfurl`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *ChatUnfurlCall) FromValues(v url.Values) error {
	var tmp ChatUnfurlCall
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("ts")); len(raw) > 0 {
		tmp.timestamp = raw
	}
	if raw := strings.TrimSpace(v.Get("unfurls")); len(raw) > 0 {
		tmp.unfurls = raw
	}
	if raw := strings.TrimSpace(v.Get("user_auth_required")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "user_auth_required"`)
		}
		tmp.userAuthRequired = parsed
	}
	*c = tmp
	return nil
}

// Update creates a ChatUpdateCall object in preparation for accessing the chat.update endpoint
func (s *ChatService) Update(channel string) *ChatUpdateCall {
	var call ChatUpdateCall
	call.service = s
	call.channel = channel
	return &call
}

// AsUser sets the value for optional asUser parameter
func (c *ChatUpdateCall) AsUser(asUser bool) *ChatUpdateCall {
	c.asUser = asUser
	return c
}

// SetAttachments sets the attachments list
func (c *ChatUpdateCall) SetAttachments(attachments objects.AttachmentList) *ChatUpdateCall {
	c.attachments = attachments
	return c
}

// Attachment appends to the attachments list
func (c *ChatUpdateCall) Attachment(attachment *objects.Attachment) *ChatUpdateCall {
	c.attachments.Append(attachment)
	return c
}

// LinkNames sets the value for optional linkNames parameter
func (c *ChatUpdateCall) LinkNames(linkNames bool) *ChatUpdateCall {
	c.linkNames = linkNames
	return c
}

// Parse sets the value for optional parse parameter
func (c *ChatUpdateCall) Parse(parse string) *ChatUpdateCall {
	c.parse = parse
	return c
}

// Text sets the value for optional text parameter
func (c *ChatUpdateCall) Text(text string) *ChatUpdateCall {
	c.text = text
	return c
}

// Timestamp sets the value for optional timestamp parameter
func (c *ChatUpdateCall) Timestamp(timestamp string) *ChatUpdateCall {
	c.timestamp = timestamp
	return c
}

// ValidateArgs checks that all required fields are set in the ChatUpdateCall object
func (c *ChatUpdateCall) ValidateArgs() error {
	if len(c.channel) <= 0 {
		return errors.New(`required field channel not initialized`)
	}
	return nil
}

// Values returns the ChatUpdateCall object as url.Values
func (c *ChatUpdateCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.asUser {
		v.Set("as_user", "true")
	}

	if len(c.attachments) > 0 {
		attachmentsEncoded, err := c.attachments.Encode()
		if err != nil {
			return nil, errors.Wrap(err, `failed to encode field`)
		}
		v.Set("attachments", attachmentsEncoded)
	}

	v.Set("channel", c.channel)

	if c.linkNames {
		v.Set("linkNames", "true")
	}

	if len(c.parse) > 0 {
		v.Set("parse", c.parse)
	}

	if len(c.text) > 0 {
		v.Set("text", c.text)
	}

	if len(c.timestamp) > 0 {
		v.Set("ts", c.timestamp)
	}
	return v, nil
}

// Do executes the call to access chat.update endpoint
func (c *ChatUpdateCall) Do(ctx context.Context) (*objects.ChatResponse, error) {
	const endpoint = "chat.update"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.ChatResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.update`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChatResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *ChatUpdateCall) FromValues(v url.Values) error {
	var tmp ChatUpdateCall
	if raw := strings.TrimSpace(v.Get("as_user")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "as_user"`)
		}
		tmp.asUser = parsed
	}
	if raw := strings.TrimSpace(v.Get("attachments")); len(raw) > 0 {
		if err := tmp.attachments.Decode(raw); err != nil {
			return errors.Wrap(err, `failed to decode value "attachments"`)
		}
	}
	if raw := strings.TrimSpace(v.Get("channel")); len(raw) > 0 {
		tmp.channel = raw
	}
	if raw := strings.TrimSpace(v.Get("linkNames")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "linkNames"`)
		}
		tmp.linkNames = parsed
	}
	if raw := strings.TrimSpace(v.Get("parse")); len(raw) > 0 {
		tmp.parse = raw
	}
	if raw := strings.TrimSpace(v.Get("text")); len(raw) > 0 {
		tmp.text = raw
	}
	if raw := strings.TrimSpace(v.Get("ts")); len(raw) > 0 {
		tmp.timestamp = raw
	}
	*c = tmp
	return nil
}
