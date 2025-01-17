// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: internal.proto

package model

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on InternalContact with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *InternalContact) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InternalContact with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InternalContactMultiError, or nil if none found.
func (m *InternalContact) ValidateAll() error {
	return m.validate(true)
}

func (m *InternalContact) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return InternalContactMultiError(errors)
	}
	return nil
}

// InternalContactMultiError is an error wrapping multiple validation errors
// returned by InternalContact.ValidateAll() if the designated constraints
// aren't met.
type InternalContactMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InternalContactMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InternalContactMultiError) AllErrors() []error { return m }

// InternalContactValidationError is the validation error returned by
// InternalContact.Validate if the designated constraints aren't met.
type InternalContactValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InternalContactValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InternalContactValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InternalContactValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InternalContactValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InternalContactValidationError) ErrorName() string { return "InternalContactValidationError" }

// Error satisfies the builtin error interface
func (e InternalContactValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInternalContact.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InternalContactValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InternalContactValidationError{}

// Validate checks the field values on InternalConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InternalConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InternalConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InternalConfigMultiError,
// or nil if none found.
func (m *InternalConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *InternalConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	// no validation rules for Status

	for idx, item := range m.GetContacts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InternalConfigValidationError{
						field:  fmt.Sprintf("Contacts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InternalConfigValidationError{
						field:  fmt.Sprintf("Contacts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InternalConfigValidationError{
					field:  fmt.Sprintf("Contacts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for UploadDir

	// no validation rules for Users

	// no validation rules for InboundMedia

	if all {
		switch v := interface{}(m.GetApplicationSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalConfigValidationError{
					field:  "ApplicationSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalConfigValidationError{
					field:  "ApplicationSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApplicationSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalConfigValidationError{
				field:  "ApplicationSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetProfileAbout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalConfigValidationError{
					field:  "ProfileAbout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalConfigValidationError{
					field:  "ProfileAbout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProfileAbout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalConfigValidationError{
				field:  "ProfileAbout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBusinessProfile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalConfigValidationError{
					field:  "BusinessProfile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalConfigValidationError{
					field:  "BusinessProfile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBusinessProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalConfigValidationError{
				field:  "BusinessProfile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ProfilePhotoFilename

	// no validation rules for Verified

	// no validation rules for WebhookCA

	if len(errors) > 0 {
		return InternalConfigMultiError(errors)
	}
	return nil
}

// InternalConfigMultiError is an error wrapping multiple validation errors
// returned by InternalConfig.ValidateAll() if the designated constraints
// aren't met.
type InternalConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InternalConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InternalConfigMultiError) AllErrors() []error { return m }

// InternalConfigValidationError is the validation error returned by
// InternalConfig.Validate if the designated constraints aren't met.
type InternalConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InternalConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InternalConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InternalConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InternalConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InternalConfigValidationError) ErrorName() string { return "InternalConfigValidationError" }

// Error satisfies the builtin error interface
func (e InternalConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInternalConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InternalConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InternalConfigValidationError{}

// Validate checks the field values on WebhookRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WebhookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WebhookRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WebhookRequestMultiError,
// or nil if none found.
func (m *WebhookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WebhookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetContacts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Contacts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Contacts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WebhookRequestValidationError{
					field:  fmt.Sprintf("Contacts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WebhookRequestValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Statuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Statuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WebhookRequestValidationError{
					field:  fmt.Sprintf("Statuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetErrors() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Errors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WebhookRequestValidationError{
						field:  fmt.Sprintf("Errors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WebhookRequestValidationError{
					field:  fmt.Sprintf("Errors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ErrorCounter

	if len(errors) > 0 {
		return WebhookRequestMultiError(errors)
	}
	return nil
}

// WebhookRequestMultiError is an error wrapping multiple validation errors
// returned by WebhookRequest.ValidateAll() if the designated constraints
// aren't met.
type WebhookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WebhookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WebhookRequestMultiError) AllErrors() []error { return m }

// WebhookRequestValidationError is the validation error returned by
// WebhookRequest.Validate if the designated constraints aren't met.
type WebhookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WebhookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WebhookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WebhookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WebhookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WebhookRequestValidationError) ErrorName() string { return "WebhookRequestValidationError" }

// Error satisfies the builtin error interface
func (e WebhookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWebhookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WebhookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WebhookRequestValidationError{}
