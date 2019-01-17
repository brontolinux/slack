package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/brontolinux/slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = objects.EpochTime(0)

// UsergroupsCreateCall is created by UsergroupsService.Create method call
type UsergroupsCreateCall struct {
	service      *UsergroupsService
	channels     string // Comma-separated list of channels IDs
	description  string
	handle       string
	includeCount bool
	name         string
}

// UsergroupsDisableCall is created by UsergroupsService.Disable method call
type UsergroupsDisableCall struct {
	service      *UsergroupsService
	includeCount bool
	usergroup    string
}

// UsergroupsEnableCall is created by UsergroupsService.Enable method call
type UsergroupsEnableCall struct {
	service      *UsergroupsService
	includeCount bool
	usergroup    string
}

// UsergroupsListCall is created by UsergroupsService.List method call
type UsergroupsListCall struct {
	service         *UsergroupsService
	includeCount    bool
	includeDisabled bool
	includeUsers    bool
}

// UsergroupsUpdateCall is created by UsergroupsService.Update method call
type UsergroupsUpdateCall struct {
	service      *UsergroupsService
	channels     string // Comma-separated list of channels IDs
	description  string
	handle       string
	includeCount bool
	name         string
	usergroup    string
}

// Create creates a UsergroupsCreateCall object in preparation for accessing the usergroups.create endpoint
func (s *UsergroupsService) Create(name string) *UsergroupsCreateCall {
	var call UsergroupsCreateCall
	call.service = s
	call.name = name
	return &call
}

// Channels sets the value for optional channels parameter
func (c *UsergroupsCreateCall) Channels(channels string) *UsergroupsCreateCall {
	c.channels = channels
	return c
}

// Description sets the value for optional description parameter
func (c *UsergroupsCreateCall) Description(description string) *UsergroupsCreateCall {
	c.description = description
	return c
}

// Handle sets the value for optional handle parameter
func (c *UsergroupsCreateCall) Handle(handle string) *UsergroupsCreateCall {
	c.handle = handle
	return c
}

// IncludeCount sets the value for optional includeCount parameter
func (c *UsergroupsCreateCall) IncludeCount(includeCount bool) *UsergroupsCreateCall {
	c.includeCount = includeCount
	return c
}

// ValidateArgs checks that all required fields are set in the UsergroupsCreateCall object
func (c *UsergroupsCreateCall) ValidateArgs() error {
	if len(c.name) <= 0 {
		return errors.New(`required field name not initialized`)
	}
	return nil
}

// Values returns the UsergroupsCreateCall object as url.Values
func (c *UsergroupsCreateCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.channels) > 0 {
		v.Set("channels", c.channels)
	}

	if len(c.description) > 0 {
		v.Set("description", c.description)
	}

	if len(c.handle) > 0 {
		v.Set("handle", c.handle)
	}

	if c.includeCount {
		v.Set("include_count", "true")
	}

	v.Set("name", c.name)
	return v, nil
}

// Do executes the call to access usergroups.create endpoint
func (c *UsergroupsCreateCall) Do(ctx context.Context) (*objects.Usergroup, error) {
	const endpoint = "usergroups.create"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Usergroup `json:"usergroup"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.create`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Usergroup, nil
}

// FromValues parses the data in v and populates `c`
func (c *UsergroupsCreateCall) FromValues(v url.Values) error {
	var tmp UsergroupsCreateCall
	if raw := strings.TrimSpace(v.Get("channels")); len(raw) > 0 {
		tmp.channels = raw
	}
	if raw := strings.TrimSpace(v.Get("description")); len(raw) > 0 {
		tmp.description = raw
	}
	if raw := strings.TrimSpace(v.Get("handle")); len(raw) > 0 {
		tmp.handle = raw
	}
	if raw := strings.TrimSpace(v.Get("include_count")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_count"`)
		}
		tmp.includeCount = parsed
	}
	if raw := strings.TrimSpace(v.Get("name")); len(raw) > 0 {
		tmp.name = raw
	}
	*c = tmp
	return nil
}

// Disable creates a UsergroupsDisableCall object in preparation for accessing the usergroups.disable endpoint
func (s *UsergroupsService) Disable(usergroup string) *UsergroupsDisableCall {
	var call UsergroupsDisableCall
	call.service = s
	call.usergroup = usergroup
	return &call
}

// IncludeCount sets the value for optional includeCount parameter
func (c *UsergroupsDisableCall) IncludeCount(includeCount bool) *UsergroupsDisableCall {
	c.includeCount = includeCount
	return c
}

// ValidateArgs checks that all required fields are set in the UsergroupsDisableCall object
func (c *UsergroupsDisableCall) ValidateArgs() error {
	if len(c.usergroup) <= 0 {
		return errors.New(`required field usergroup not initialized`)
	}
	return nil
}

// Values returns the UsergroupsDisableCall object as url.Values
func (c *UsergroupsDisableCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeCount {
		v.Set("include_count", "true")
	}

	v.Set("usergroup", c.usergroup)
	return v, nil
}

// Do executes the call to access usergroups.disable endpoint
func (c *UsergroupsDisableCall) Do(ctx context.Context) (*objects.Usergroup, error) {
	const endpoint = "usergroups.disable"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Usergroup `json:"usergroup"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.disable`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Usergroup, nil
}

// FromValues parses the data in v and populates `c`
func (c *UsergroupsDisableCall) FromValues(v url.Values) error {
	var tmp UsergroupsDisableCall
	if raw := strings.TrimSpace(v.Get("include_count")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_count"`)
		}
		tmp.includeCount = parsed
	}
	if raw := strings.TrimSpace(v.Get("usergroup")); len(raw) > 0 {
		tmp.usergroup = raw
	}
	*c = tmp
	return nil
}

// Enable creates a UsergroupsEnableCall object in preparation for accessing the usergroups.enable endpoint
func (s *UsergroupsService) Enable(usergroup string) *UsergroupsEnableCall {
	var call UsergroupsEnableCall
	call.service = s
	call.usergroup = usergroup
	return &call
}

// IncludeCount sets the value for optional includeCount parameter
func (c *UsergroupsEnableCall) IncludeCount(includeCount bool) *UsergroupsEnableCall {
	c.includeCount = includeCount
	return c
}

// ValidateArgs checks that all required fields are set in the UsergroupsEnableCall object
func (c *UsergroupsEnableCall) ValidateArgs() error {
	if len(c.usergroup) <= 0 {
		return errors.New(`required field usergroup not initialized`)
	}
	return nil
}

// Values returns the UsergroupsEnableCall object as url.Values
func (c *UsergroupsEnableCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeCount {
		v.Set("include_count", "true")
	}

	v.Set("usergroup", c.usergroup)
	return v, nil
}

// Do executes the call to access usergroups.enable endpoint
func (c *UsergroupsEnableCall) Do(ctx context.Context) (*objects.Usergroup, error) {
	const endpoint = "usergroups.enable"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Usergroup `json:"usergroup"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.enable`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Usergroup, nil
}

// FromValues parses the data in v and populates `c`
func (c *UsergroupsEnableCall) FromValues(v url.Values) error {
	var tmp UsergroupsEnableCall
	if raw := strings.TrimSpace(v.Get("include_count")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_count"`)
		}
		tmp.includeCount = parsed
	}
	if raw := strings.TrimSpace(v.Get("usergroup")); len(raw) > 0 {
		tmp.usergroup = raw
	}
	*c = tmp
	return nil
}

// List creates a UsergroupsListCall object in preparation for accessing the usergroups.list endpoint
func (s *UsergroupsService) List() *UsergroupsListCall {
	var call UsergroupsListCall
	call.service = s
	return &call
}

// IncludeCount sets the value for optional includeCount parameter
func (c *UsergroupsListCall) IncludeCount(includeCount bool) *UsergroupsListCall {
	c.includeCount = includeCount
	return c
}

// IncludeDisabled sets the value for optional includeDisabled parameter
func (c *UsergroupsListCall) IncludeDisabled(includeDisabled bool) *UsergroupsListCall {
	c.includeDisabled = includeDisabled
	return c
}

// IncludeUsers sets the value for optional includeUsers parameter
func (c *UsergroupsListCall) IncludeUsers(includeUsers bool) *UsergroupsListCall {
	c.includeUsers = includeUsers
	return c
}

// ValidateArgs checks that all required fields are set in the UsergroupsListCall object
func (c *UsergroupsListCall) ValidateArgs() error {
	return nil
}

// Values returns the UsergroupsListCall object as url.Values
func (c *UsergroupsListCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeCount {
		v.Set("include_count", "true")
	}

	if c.includeDisabled {
		v.Set("include_disabled", "true")
	}

	if c.includeUsers {
		v.Set("include_users", "true")
	}
	return v, nil
}

// Do executes the call to access usergroups.list endpoint
func (c *UsergroupsListCall) Do(ctx context.Context) (objects.UsergroupList, error) {
	const endpoint = "usergroups.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		objects.UsergroupList `json:"usergroups"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.list`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UsergroupList, nil
}

// FromValues parses the data in v and populates `c`
func (c *UsergroupsListCall) FromValues(v url.Values) error {
	var tmp UsergroupsListCall
	if raw := strings.TrimSpace(v.Get("include_count")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_count"`)
		}
		tmp.includeCount = parsed
	}
	if raw := strings.TrimSpace(v.Get("include_disabled")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_disabled"`)
		}
		tmp.includeDisabled = parsed
	}
	if raw := strings.TrimSpace(v.Get("include_users")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_users"`)
		}
		tmp.includeUsers = parsed
	}
	*c = tmp
	return nil
}

// Update creates a UsergroupsUpdateCall object in preparation for accessing the usergroups.update endpoint
func (s *UsergroupsService) Update(usergroup string) *UsergroupsUpdateCall {
	var call UsergroupsUpdateCall
	call.service = s
	call.usergroup = usergroup
	return &call
}

// Channels sets the value for optional channels parameter
func (c *UsergroupsUpdateCall) Channels(channels string) *UsergroupsUpdateCall {
	c.channels = channels
	return c
}

// Description sets the value for optional description parameter
func (c *UsergroupsUpdateCall) Description(description string) *UsergroupsUpdateCall {
	c.description = description
	return c
}

// Handle sets the value for optional handle parameter
func (c *UsergroupsUpdateCall) Handle(handle string) *UsergroupsUpdateCall {
	c.handle = handle
	return c
}

// IncludeCount sets the value for optional includeCount parameter
func (c *UsergroupsUpdateCall) IncludeCount(includeCount bool) *UsergroupsUpdateCall {
	c.includeCount = includeCount
	return c
}

// Name sets the value for optional name parameter
func (c *UsergroupsUpdateCall) Name(name string) *UsergroupsUpdateCall {
	c.name = name
	return c
}

// ValidateArgs checks that all required fields are set in the UsergroupsUpdateCall object
func (c *UsergroupsUpdateCall) ValidateArgs() error {
	if len(c.usergroup) <= 0 {
		return errors.New(`required field usergroup not initialized`)
	}
	return nil
}

// Values returns the UsergroupsUpdateCall object as url.Values
func (c *UsergroupsUpdateCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.channels) > 0 {
		v.Set("channels", c.channels)
	}

	if len(c.description) > 0 {
		v.Set("description", c.description)
	}

	if len(c.handle) > 0 {
		v.Set("handle", c.handle)
	}

	if c.includeCount {
		v.Set("include_count", "true")
	}

	if len(c.name) > 0 {
		v.Set("name", c.name)
	}

	v.Set("usergroup", c.usergroup)
	return v, nil
}

// Do executes the call to access usergroups.update endpoint
func (c *UsergroupsUpdateCall) Do(ctx context.Context) (*objects.Usergroup, error) {
	const endpoint = "usergroups.update"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Usergroup `json:"usergroup"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.update`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Usergroup, nil
}

// FromValues parses the data in v and populates `c`
func (c *UsergroupsUpdateCall) FromValues(v url.Values) error {
	var tmp UsergroupsUpdateCall
	if raw := strings.TrimSpace(v.Get("channels")); len(raw) > 0 {
		tmp.channels = raw
	}
	if raw := strings.TrimSpace(v.Get("description")); len(raw) > 0 {
		tmp.description = raw
	}
	if raw := strings.TrimSpace(v.Get("handle")); len(raw) > 0 {
		tmp.handle = raw
	}
	if raw := strings.TrimSpace(v.Get("include_count")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_count"`)
		}
		tmp.includeCount = parsed
	}
	if raw := strings.TrimSpace(v.Get("name")); len(raw) > 0 {
		tmp.name = raw
	}
	if raw := strings.TrimSpace(v.Get("usergroup")); len(raw) > 0 {
		tmp.usergroup = raw
	}
	*c = tmp
	return nil
}
