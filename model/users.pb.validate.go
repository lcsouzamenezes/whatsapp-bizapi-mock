// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: users.proto

package model

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on TokenResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TokenResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	// no validation rules for ExpiresAfter

	return nil
}

// TokenResponseValidationError is the validation error returned by
// TokenResponse.Validate if the designated constraints aren't met.
type TokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenResponseValidationError) ErrorName() string { return "TokenResponseValidationError" }

// Error satisfies the builtin error interface
func (e TokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenResponseValidationError{}

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LoginResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginResponseValidationError{
				field:  "Meta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LoginResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on ChangePwdRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ChangePwdRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetNewPassword()) < 8 {
		return ChangePwdRequestValidationError{
			field:  "NewPassword",
			reason: "value length must be at least 8 runes",
		}
	}

	return nil
}

// ChangePwdRequestValidationError is the validation error returned by
// ChangePwdRequest.Validate if the designated constraints aren't met.
type ChangePwdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChangePwdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChangePwdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChangePwdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChangePwdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChangePwdRequestValidationError) ErrorName() string { return "ChangePwdRequestValidationError" }

// Error satisfies the builtin error interface
func (e ChangePwdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChangePwdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChangePwdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChangePwdRequestValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUsername()) < 5 {
		return UserValidationError{
			field:  "Username",
			reason: "value length must be at least 5 runes",
		}
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		return UserValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}
