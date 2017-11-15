package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = objects.EpochTime(0)

// UsersProfileGetCall is created by UsersProfileService.Get method call
type UsersProfileGetCall struct {
	service       *UsersProfileService
	includeLabels bool
	user          string
}

// UsersProfileSetCall is created by UsersProfileService.Set method call
type UsersProfileSetCall struct {
	service *UsersProfileService
	name    string
	profile *objects.UserProfile
	user    string
	value   string
}

// Get creates a UsersProfileGetCall object in preparation for accessing the users.profile.get endpoint
func (s *UsersProfileService) Get() *UsersProfileGetCall {
	var call UsersProfileGetCall
	call.service = s
	return &call
}

// IncludeLabels sets the value for optional includeLabels parameter
func (c *UsersProfileGetCall) IncludeLabels(includeLabels bool) *UsersProfileGetCall {
	c.includeLabels = includeLabels
	return c
}

// User sets the value for optional user parameter
func (c *UsersProfileGetCall) User(user string) *UsersProfileGetCall {
	c.user = user
	return c
}

// ValidateArgs checks that all required fields are set in the UsersProfileGetCall object
func (c *UsersProfileGetCall) ValidateArgs() error {
	return nil
}

// Values returns the UsersProfileGetCall object as url.Values
func (c *UsersProfileGetCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeLabels {
		v.Set("include_labels", "true")
	}

	if len(c.user) > 0 {
		v.Set("user", c.user)
	}
	return v, nil
}

// Do executes the call to access users.profile.get endpoint
func (c *UsersProfileGetCall) Do(ctx context.Context) (*objects.UserProfile, error) {
	const endpoint = "users.profile.get"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*objects.UserProfile
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.profile.get`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserProfile, nil
}

// FromValues parses the data in v and populates `c`
func (c *UsersProfileGetCall) FromValues(v url.Values) error {
	var tmp UsersProfileGetCall
	if raw := strings.TrimSpace(v.Get("include_labels")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_labels"`)
		}
		tmp.includeLabels = parsed
	}
	if raw := strings.TrimSpace(v.Get("user")); len(raw) > 0 {
		tmp.user = raw
	}
	*c = tmp
	return nil
}

// Set creates a UsersProfileSetCall object in preparation for accessing the users.profile.set endpoint
func (s *UsersProfileService) Set() *UsersProfileSetCall {
	var call UsersProfileSetCall
	call.service = s
	return &call
}

// Name sets the value for optional name parameter
func (c *UsersProfileSetCall) Name(name string) *UsersProfileSetCall {
	c.name = name
	return c
}

// Profile sets the value for optional profile parameter
func (c *UsersProfileSetCall) Profile(profile *objects.UserProfile) *UsersProfileSetCall {
	c.profile = profile
	return c
}

// User sets the value for optional user parameter
func (c *UsersProfileSetCall) User(user string) *UsersProfileSetCall {
	c.user = user
	return c
}

// Value sets the value for optional value parameter
func (c *UsersProfileSetCall) Value(value string) *UsersProfileSetCall {
	c.value = value
	return c
}

// ValidateArgs checks that all required fields are set in the UsersProfileSetCall object
func (c *UsersProfileSetCall) ValidateArgs() error {
	return nil
}

// Values returns the UsersProfileSetCall object as url.Values
func (c *UsersProfileSetCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.name) > 0 {
		v.Set("name", c.name)
	}

	if c.profile != nil {
		profileEncoded, err := c.profile.Encode()
		if err != nil {
			return nil, errors.Wrap(err, `failed to encode field`)
		}
		v.Set("profile", profileEncoded)
	}

	if len(c.user) > 0 {
		v.Set("user", c.user)
	}

	if len(c.value) > 0 {
		v.Set("value", c.value)
	}
	return v, nil
}

// Do executes the call to access users.profile.set endpoint
func (c *UsersProfileSetCall) Do(ctx context.Context) (*objects.UserProfile, error) {
	const endpoint = "users.profile.set"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*objects.UserProfile
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.profile.set`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserProfile, nil
}

// FromValues parses the data in v and populates `c`
func (c *UsersProfileSetCall) FromValues(v url.Values) error {
	var tmp UsersProfileSetCall
	if raw := strings.TrimSpace(v.Get("name")); len(raw) > 0 {
		tmp.name = raw
	}
	if raw := strings.TrimSpace(v.Get("profile")); len(raw) > 0 {
		if err := tmp.profile.Decode(raw); err != nil {
			return errors.Wrap(err, `failed to decode value "profile"`)
		}
	}
	if raw := strings.TrimSpace(v.Get("user")); len(raw) > 0 {
		tmp.user = raw
	}
	if raw := strings.TrimSpace(v.Get("value")); len(raw) > 0 {
		tmp.value = raw
	}
	*c = tmp
	return nil
}
