// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: status.proto

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

// Validate checks the field values on Status with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Status) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Status

	// no validation rules for RecipientId

	// no validation rules for Timestamp

	return nil
}

// StatusValidationError is the validation error returned by Status.Validate if
// the designated constraints aren't met.
type StatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusValidationError) ErrorName() string { return "StatusValidationError" }

// Error satisfies the builtin error interface
func (e StatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusValidationError{}
