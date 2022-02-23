// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	stderrors "errors"
	"net/http"
)

const (
	_                       = http.StatusTeapot              // :)
	codeBadRequest          = http.StatusBadRequest          // Classic...
	codeNotFound            = http.StatusNotFound            // Classic...
	codeInternalServerError = http.StatusInternalServerError // Classic...
	codeTimeout             = 1000
	codeNetworkError        = 1100
	codeDBError             = 1200
)

// New returns a string error.
func New(text string) error {
	return stderrors.New(text)
}

// BadRequest returns a bad request error.
func BadRequest(err error) error {
	return Wrap(err, codeBadRequest)
}

// IsBadRequest if err is bad request.
func IsBadRequest(err error) bool {
	return Is(err, codeBadRequest)
}

// UnwrapBadRequest if err is bad request.
func UnwrapBadRequest(err error) (error, bool) {
	return Unwrap(err, codeBadRequest)
}

// NotFound returns a not found error.
func NotFound(err error) error {
	return Wrap(err, codeNotFound)
}

// IsNotFound if err is not found.
func IsNotFound(err error) bool {
	return Is(err, codeNotFound)
}

// UnwrapNotFound if err is not found.
func UnwrapNotFound(err error) (error, bool) {
	return Unwrap(err, codeNotFound)
}

// InternalServerError returns an internal server error.
func InternalServerError(err error) error {
	return Wrap(err, codeInternalServerError)
}

// IsInternalServerError if err is an internal server.
func IsInternalServerError(err error) bool {
	return Is(err, codeInternalServerError)
}

// UnwrapInternalServerError if err is an internal server.
func UnwrapInternalServerError(err error) (error, bool) {
	return Unwrap(err, codeInternalServerError)
}

// Timeout returns a timeout error.
func Timeout(err error) error {
	return Wrap(err, codeTimeout)
}

// IsTimeout if err is timeout.
func IsTimeout(err error) bool {
	return Is(err, codeTimeout)
}

// UnwrapTimeout if err is timeout.
func UnwrapTimeout(err error) (error, bool) {
	return Unwrap(err, codeTimeout)
}

// NetworkError returns a network error.
func NetworkError(err error) error {
	return Wrap(err, codeNetworkError)
}

// IsNetworkError if err is network error.
func IsNetworkError(err error) bool {
	return Is(err, codeNetworkError)
}

// UnwrapNetworkError if err is network error.
func UnwrapNetworkError(err error) (error, bool) {
	return Unwrap(err, codeNetworkError)
}

// DBError returns a db error.
func DBError(err error) error {
	return Wrap(err, codeDBError)
}

// IsDBError if err is db error.
func IsDBError(err error) bool {
	return Is(err, codeDBError)
}

// UnwrapDBError if err is db error.
func UnwrapDBError(err error) (error, bool) {
	return Unwrap(err, codeDBError)
}
