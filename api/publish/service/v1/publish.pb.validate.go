// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: publish/service/v1/publish.proto

package v1

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

// Validate checks the field values on VideoListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *VideoListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VideoListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VideoListRequestMultiError, or nil if none found.
func (m *VideoListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VideoListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LatestTime

	// no validation rules for UserId

	// no validation rules for Number

	if len(errors) > 0 {
		return VideoListRequestMultiError(errors)
	}

	return nil
}

// VideoListRequestMultiError is an error wrapping multiple validation errors
// returned by VideoListRequest.ValidateAll() if the designated constraints
// aren't met.
type VideoListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VideoListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VideoListRequestMultiError) AllErrors() []error { return m }

// VideoListRequestValidationError is the validation error returned by
// VideoListRequest.Validate if the designated constraints aren't met.
type VideoListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoListRequestValidationError) ErrorName() string { return "VideoListRequestValidationError" }

// Error satisfies the builtin error interface
func (e VideoListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideoListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoListRequestValidationError{}

// Validate checks the field values on VideoListReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VideoListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VideoListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VideoListReplyMultiError,
// or nil if none found.
func (m *VideoListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *VideoListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NextTime

	for idx, item := range m.GetVideoList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, VideoListReplyValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, VideoListReplyValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return VideoListReplyValidationError{
					field:  fmt.Sprintf("VideoList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return VideoListReplyMultiError(errors)
	}

	return nil
}

// VideoListReplyMultiError is an error wrapping multiple validation errors
// returned by VideoListReply.ValidateAll() if the designated constraints
// aren't met.
type VideoListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VideoListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VideoListReplyMultiError) AllErrors() []error { return m }

// VideoListReplyValidationError is the validation error returned by
// VideoListReply.Validate if the designated constraints aren't met.
type VideoListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoListReplyValidationError) ErrorName() string { return "VideoListReplyValidationError" }

// Error satisfies the builtin error interface
func (e VideoListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideoListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoListReplyValidationError{}

// Validate checks the field values on VideoListByVideoIdsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VideoListByVideoIdsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VideoListByVideoIdsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VideoListByVideoIdsRequestMultiError, or nil if none found.
func (m *VideoListByVideoIdsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VideoListByVideoIdsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return VideoListByVideoIdsRequestMultiError(errors)
	}

	return nil
}

// VideoListByVideoIdsRequestMultiError is an error wrapping multiple
// validation errors returned by VideoListByVideoIdsRequest.ValidateAll() if
// the designated constraints aren't met.
type VideoListByVideoIdsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VideoListByVideoIdsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VideoListByVideoIdsRequestMultiError) AllErrors() []error { return m }

// VideoListByVideoIdsRequestValidationError is the validation error returned
// by VideoListByVideoIdsRequest.Validate if the designated constraints aren't met.
type VideoListByVideoIdsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoListByVideoIdsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoListByVideoIdsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoListByVideoIdsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoListByVideoIdsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoListByVideoIdsRequestValidationError) ErrorName() string {
	return "VideoListByVideoIdsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VideoListByVideoIdsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideoListByVideoIdsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoListByVideoIdsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoListByVideoIdsRequestValidationError{}

// Validate checks the field values on UpdateFavoriteCountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateFavoriteCountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateFavoriteCountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateFavoriteCountRequestMultiError, or nil if none found.
func (m *UpdateFavoriteCountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateFavoriteCountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for VideoId

	// no validation rules for FavoriteChange

	if len(errors) > 0 {
		return UpdateFavoriteCountRequestMultiError(errors)
	}

	return nil
}

// UpdateFavoriteCountRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateFavoriteCountRequest.ValidateAll() if
// the designated constraints aren't met.
type UpdateFavoriteCountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateFavoriteCountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateFavoriteCountRequestMultiError) AllErrors() []error { return m }

// UpdateFavoriteCountRequestValidationError is the validation error returned
// by UpdateFavoriteCountRequest.Validate if the designated constraints aren't met.
type UpdateFavoriteCountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateFavoriteCountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateFavoriteCountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateFavoriteCountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateFavoriteCountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateFavoriteCountRequestValidationError) ErrorName() string {
	return "UpdateFavoriteCountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateFavoriteCountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateFavoriteCountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateFavoriteCountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateFavoriteCountRequestValidationError{}

// Validate checks the field values on UpdateCommentCountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCommentCountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCommentCountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCommentCountRequestMultiError, or nil if none found.
func (m *UpdateCommentCountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCommentCountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for VideoId

	// no validation rules for CommentChange

	if len(errors) > 0 {
		return UpdateCommentCountRequestMultiError(errors)
	}

	return nil
}

// UpdateCommentCountRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateCommentCountRequest.ValidateAll() if the
// designated constraints aren't met.
type UpdateCommentCountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCommentCountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCommentCountRequestMultiError) AllErrors() []error { return m }

// UpdateCommentCountRequestValidationError is the validation error returned by
// UpdateCommentCountRequest.Validate if the designated constraints aren't met.
type UpdateCommentCountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCommentCountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCommentCountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCommentCountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCommentCountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCommentCountRequestValidationError) ErrorName() string {
	return "UpdateCommentCountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCommentCountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCommentCountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCommentCountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCommentCountRequestValidationError{}

// Validate checks the field values on PublishActionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublishActionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishActionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishActionRequestMultiError, or nil if none found.
func (m *PublishActionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishActionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetToken()) < 1 {
		err := PublishActionRequestValidationError{
			field:  "Token",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Data

	// no validation rules for Title

	if len(errors) > 0 {
		return PublishActionRequestMultiError(errors)
	}

	return nil
}

// PublishActionRequestMultiError is an error wrapping multiple validation
// errors returned by PublishActionRequest.ValidateAll() if the designated
// constraints aren't met.
type PublishActionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishActionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishActionRequestMultiError) AllErrors() []error { return m }

// PublishActionRequestValidationError is the validation error returned by
// PublishActionRequest.Validate if the designated constraints aren't met.
type PublishActionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishActionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishActionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishActionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishActionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishActionRequestValidationError) ErrorName() string {
	return "PublishActionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PublishActionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishActionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishActionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishActionRequestValidationError{}

// Validate checks the field values on PublishActionReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublishActionReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishActionReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishActionReplyMultiError, or nil if none found.
func (m *PublishActionReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishActionReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for StatusMsg

	if len(errors) > 0 {
		return PublishActionReplyMultiError(errors)
	}

	return nil
}

// PublishActionReplyMultiError is an error wrapping multiple validation errors
// returned by PublishActionReply.ValidateAll() if the designated constraints
// aren't met.
type PublishActionReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishActionReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishActionReplyMultiError) AllErrors() []error { return m }

// PublishActionReplyValidationError is the validation error returned by
// PublishActionReply.Validate if the designated constraints aren't met.
type PublishActionReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishActionReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishActionReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishActionReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishActionReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishActionReplyValidationError) ErrorName() string {
	return "PublishActionReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PublishActionReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishActionReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishActionReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishActionReplyValidationError{}

// Validate checks the field values on PublishListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublishListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishListRequestMultiError, or nil if none found.
func (m *PublishListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Token

	if len(errors) > 0 {
		return PublishListRequestMultiError(errors)
	}

	return nil
}

// PublishListRequestMultiError is an error wrapping multiple validation errors
// returned by PublishListRequest.ValidateAll() if the designated constraints
// aren't met.
type PublishListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishListRequestMultiError) AllErrors() []error { return m }

// PublishListRequestValidationError is the validation error returned by
// PublishListRequest.Validate if the designated constraints aren't met.
type PublishListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishListRequestValidationError) ErrorName() string {
	return "PublishListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PublishListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishListRequestValidationError{}

// Validate checks the field values on PublishListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PublishListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishListReplyMultiError, or nil if none found.
func (m *PublishListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for StatusMsg

	for idx, item := range m.GetVideoList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PublishListReplyValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PublishListReplyValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PublishListReplyValidationError{
					field:  fmt.Sprintf("VideoList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PublishListReplyMultiError(errors)
	}

	return nil
}

// PublishListReplyMultiError is an error wrapping multiple validation errors
// returned by PublishListReply.ValidateAll() if the designated constraints
// aren't met.
type PublishListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishListReplyMultiError) AllErrors() []error { return m }

// PublishListReplyValidationError is the validation error returned by
// PublishListReply.Validate if the designated constraints aren't met.
type PublishListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishListReplyValidationError) ErrorName() string { return "PublishListReplyValidationError" }

// Error satisfies the builtin error interface
func (e PublishListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishListReplyValidationError{}

// Validate checks the field values on Video with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Video) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Video with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in VideoMultiError, or nil if none found.
func (m *Video) ValidateAll() error {
	return m.validate(true)
}

func (m *Video) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetAuthor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VideoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VideoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VideoValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PlayUrl

	// no validation rules for CoverUrl

	// no validation rules for FavoriteCount

	// no validation rules for CommentCount

	// no validation rules for IsFavorite

	// no validation rules for Title

	if len(errors) > 0 {
		return VideoMultiError(errors)
	}

	return nil
}

// VideoMultiError is an error wrapping multiple validation errors returned by
// Video.ValidateAll() if the designated constraints aren't met.
type VideoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VideoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VideoMultiError) AllErrors() []error { return m }

// VideoValidationError is the validation error returned by Video.Validate if
// the designated constraints aren't met.
type VideoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoValidationError) ErrorName() string { return "VideoValidationError" }

// Error satisfies the builtin error interface
func (e VideoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for FollowCount

	// no validation rules for FollowerCount

	// no validation rules for IsFollow

	// no validation rules for Avatar

	// no validation rules for BackgroundImage

	// no validation rules for Signature

	// no validation rules for TotalFavorited

	// no validation rules for WorkCount

	// no validation rules for FavoriteCount

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

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