// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"net/http"
)

const (
	_ = http.StatusTeapot // :)
)

const (
	codeBadRequest          = http.StatusBadRequest          // Classic...
	codeForbidden           = http.StatusForbidden           // Classic...
	codeNotFound            = http.StatusNotFound            // Classic...
	codeRequestTimeout      = http.StatusRequestTimeout      // Classic...
	codeInternalServerError = http.StatusInternalServerError // Classic...
)

// BadRequest returns a bad request error.
func BadRequest(err error, opts ...Option) error {
	return Wrap(err, codeBadRequest, opts...)
}

// IsBadRequest if err is bad request.
func IsBadRequest(err error) bool {
	return Is(err, codeBadRequest)
}

// UnwrapBadRequest if err is bad request.
func UnwrapBadRequest(err error) (error, bool) {
	return Unwrap(err, codeBadRequest)
}

// Forbidden returns a forbidden error.
func Forbidden(err error, opts ...Option) error {
	return Wrap(err, codeForbidden, opts...)
}

// IsForbidden if err is forbidden.
func IsForbidden(err error) bool {
	return Is(err, codeForbidden)
}

// UnwrapForbidden if err is forbidden.
func UnwrapForbidden(err error) (error, bool) {
	return Unwrap(err, codeForbidden)
}

// NotFound returns a not found error.
func NotFound(err error, opts ...Option) error {
	return Wrap(err, codeNotFound, opts...)
}

// IsNotFound if err is not found.
func IsNotFound(err error) bool {
	return Is(err, codeNotFound)
}

// UnwrapNotFound if err is not found.
func UnwrapNotFound(err error) (error, bool) {
	return Unwrap(err, codeNotFound)
}

// RequestTimeout returns a request timeout error.
func RequestTimeout(err error, opts ...Option) error {
	return Wrap(err, codeRequestTimeout, opts...)
}

// IsRequestTimeout if err is request timeout.
func IsRequestTimeout(err error) bool {
	return Is(err, codeRequestTimeout)
}

// UnwrapRequestTimeout if err is request timeout.
func UnwrapRequestTimeout(err error) (error, bool) {
	return Unwrap(err, codeRequestTimeout)
}

// InternalServerError returns an internal server error.
func InternalServerError(err error, opts ...Option) error {
	return Wrap(err, codeInternalServerError, opts...)
}

// IsInternalServerError if err is an internal server.
func IsInternalServerError(err error) bool {
	return Is(err, codeInternalServerError)
}

// UnwrapInternalServerError if err is an internal server.
func UnwrapInternalServerError(err error) (error, bool) {
	return Unwrap(err, codeInternalServerError)
}
