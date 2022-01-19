// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/19 23:43:15

package errors

import (
	stderrors "errors"
	"net/http"
)

const (
	_                = http.StatusTeapot   // :)
	codeNotFound     = http.StatusNotFound // Classic...
	codeTimeout      = 1000
	codeNetworkError = 1100
	codeDBError      = 1200
)

// New returns a string error.
func New(text string) error {
	return stderrors.New(text)
}

// NotFound returns a not found error.
func NotFound(err error) error {
	return Wrap(err, codeNotFound)
}

// IsNotFound if err is not found.
func IsNotFound(err error) (error, bool) {
	return Is(err, codeNotFound)
}

// Timeout returns a timeout error.
func Timeout(err error) error {
	return Wrap(err, codeTimeout)
}

// IsTimeout if err is timeout.
func IsTimeout(err error) (error, bool) {
	return Is(err, codeTimeout)
}

// NetworkError returns a network error.
func NetworkError(err error) error {
	return Wrap(err, codeNetworkError)
}

// IsNetworkError if err is network error.
func IsNetworkError(err error) (error, bool) {
	return Is(err, codeNetworkError)
}

// DBError returns a db error.
func DBError(err error) error {
	return Wrap(err, codeDBError)
}

// IsDBError if err is db error.
func IsDBError(err error) (error, bool) {
	return Is(err, codeDBError)
}
