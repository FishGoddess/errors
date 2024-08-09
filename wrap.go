// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"bytes"
	"fmt"
	"strings"
)

type Error struct {
	code    int32
	message string
	cause   error
	caller  string
}

// Wrap returns *Error with code and message formatted with args.
func Wrap(code int32, message string, args ...any) *Error {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	err := &Error{
		code:    code,
		message: message,
	}

	return err
}

func (e *Error) With(err error) *Error {
	e.cause = err
	return e
}

func (e *Error) WithCaller() *Error {
	e.caller = Caller()
	return e
}

func (e *Error) WithCallers() *Error {
	callers := Callers()
	e.caller = strings.Join(callers, "¦")
	return e
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
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "%d: %s", e.code, e.message)

	if e.caller != "" {
		fmt.Fprintf(&buff, " [%s]", e.caller)
	}

	if e.cause != nil {
		fmt.Fprintf(&buff, " (%+v)", e.cause)
	}

	return buff.String()
}
