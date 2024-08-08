// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package x

import (
	"fmt"
)

type Error struct {
	code    int32
	message string
	cause   error
}

// New returns *Error with code and message.
func New(code int32, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

// Newf returns *Error with code and formatted message.
func Newf(code int32, format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	return New(code, message)
}

// Wrap wraps err to *Error with code and message.
func Wrap(err error, code int32, message string) *Error {
	return &Error{
		code:    code,
		message: message,
		cause:   err,
	}
}

// Wrapf wraps err to *Error with code and formatted message.
func Wrapf(err error, code int32, format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	return Wrap(err, code, message)
}

// Code returns the code of *Error.
func (e *Error) Code() int32 {
	return e.code
}

// Message returns the message of *Error.
func (e *Error) Message() string {
	return e.message
}

// Unwrap unwraps *Error and returns its cause error.
func (e *Error) Unwrap() error {
	return e.cause
}

// Error returns *Error as string.
func (e *Error) Error() string {
	return e.String()
}

// String returns *Error as string.
func (e *Error) String() string {
	if e.cause == nil {
		return fmt.Sprintf("%d: %s", e.code, e.message)
	}

	return fmt.Sprintf("%d: %s (%+v)", e.code, e.message, e.cause)
}
