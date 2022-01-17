package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	codeBadRequest          = http.StatusBadRequest
	codeUnauthorized        = http.StatusUnauthorized
	codeForbidden           = http.StatusForbidden
	codeNotFound            = http.StatusNotFound
	codeRequestTimeout      = http.StatusRequestTimeout
	codeTeapot              = http.StatusTeapot // :)
	codeTooManyRequests     = http.StatusTooManyRequests
	codeServerInternalError = http.StatusInternalServerError
	codeServiceUnavailable  = http.StatusServiceUnavailable

	codeTokenInvalid = 1000
	codeDBError      = 1100
)

var (
	// FormatError formats e and returns a string.
	FormatError = func(e *Error) string {
		return fmt.Sprintf("%d (%s)", e.code, e.err.Error())
	}
)

// Error wraps err with code.
type Error struct {
	err  error
	code int32
}

// Error returns the msg of e.
func (e *Error) Error() string {
	if e == nil || e.err == nil {
		return ""
	}
	return FormatError(e)
}

// Is returns if e has the same type of target.
func (e *Error) Is(target error) bool {
	if e == nil {
		return e == target
	}

	err, ok := target.(*Error)
	if !ok {
		return e.err == target
	}

	return e.code == err.code
}

// Unwrap returns err inside.
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.err
}

// WithCode wraps err with code
func WithCode(err error, code int32) error {
	if err == nil {
		return nil
	}

	return &Error{
		err:  err,
		code: code,
	}
}

// Is returns if err is Error and its code == code.
func Is(err error, code int32) bool {
	for {
		if err == nil {
			return false
		}

		e, ok := err.(*Error)
		if !ok {
			return false
		}

		if e.code == code {
			return true
		}

		err = errors.Unwrap(err)
	}
}

// BadRequest returns a bad request error.
func BadRequest(err error) error {
	return WithCode(err, codeBadRequest)
}

// IsBadRequest if err is bad request.
func IsBadRequest(err error) bool {
	return Is(err, codeBadRequest)
}

// Unauthorized returns a unauthorized error.
func Unauthorized(err error) error {
	return WithCode(err, codeUnauthorized)
}

// IsUnauthorized if err is unauthorized.
func IsUnauthorized(err error) bool {
	return Is(err, codeUnauthorized)
}

// Forbidden returns a forbidden error.
func Forbidden(err error) error {
	return WithCode(err, codeForbidden)
}

// IsForbidden if err is forbidden.
func IsForbidden(err error) bool {
	return Is(err, codeForbidden)
}

// NotFound returns a not found error.
func NotFound(err error) error {
	return WithCode(err, codeNotFound)
}

// IsNotFound if err is not found.
func IsNotFound(err error) bool {
	return Is(err, codeNotFound)
}

// RequestTimeout returns a request timeout error.
func RequestTimeout(err error) error {
	return WithCode(err, codeRequestTimeout)
}

// IsRequestTimeout if err is request timeout.
func IsRequestTimeout(err error) bool {
	return Is(err, codeRequestTimeout)
}

// TooManyRequests returns a too many requests error.
func TooManyRequests(err error) error {
	return WithCode(err, codeTooManyRequests)
}

// IsTooManyRequests if err is too many requests.
func IsTooManyRequests(err error) bool {
	return Is(err, codeTooManyRequests)
}

// ServerInternalError returns a server internal error.
func ServerInternalError(err error) error {
	return WithCode(err, codeServerInternalError)
}

// IsServerInternalError if err is server internal request.
func IsServerInternalError(err error) bool {
	return Is(err, codeServerInternalError)
}

// ServiceUnavailable returns a service unavailable error.
func ServiceUnavailable(err error) error {
	return WithCode(err, codeServiceUnavailable)
}

// IsServiceUnavailable if err is service unavailable.
func IsServiceUnavailable(err error) bool {
	return Is(err, codeServiceUnavailable)
}

// TokenInvalid returns a token invalid error.
func TokenInvalid(err error) error {
	return WithCode(err, codeTokenInvalid)
}

// IsTokenInvalid if err is token invalid.
func IsTokenInvalid(err error) bool {
	return Is(err, codeTokenInvalid)
}

// DBError returns a db error.
func DBError(err error) error {
	return WithCode(err, codeDBError)
}

// IsDBError if err is db error.
func IsDBError(err error) bool {
	return Is(err, codeDBError)
}
