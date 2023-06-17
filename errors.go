// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"fmt"
)

const (
	codeNil     = 0
	codeUnknown = 1
)

const (
	nilString = "<nil>"
)

var (
	// FormatError formats e and returns a string in Error().
	FormatError = func(e *Error) string {
		return e.err.Error()
	}

	// FormatString formats e and returns a string in String().
	FormatString = func(e *Error) string {
		return fmt.Sprintf("%d (%s)", e.code, e.msg)
	}
)

// Error wraps err with some information.
type Error struct {
	err error

	code int32
	msg  string
}

func (e *Error) Error() string {
	if e == nil || e.err == nil {
		return nilString
	}

	return FormatError(e)
}

func (e *Error) String() string {
	if e == nil || e.err == nil {
		return nilString
	}

	return FormatString(e)
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

// Wrap wraps err with code
func Wrap(err error, code int32, opts ...Option) error {
	if err == nil {
		return nil
	}

	e := &Error{
		err:  err,
		code: code,
		msg:  err.Error(),
	}

	applyOptions(e, opts)
	return e
}

// Unwrap returns if err is Error and its code == code, and the original error will be returned, too.
func Unwrap(err error, code int32) (error, bool) {
	for {
		if err == nil {
			return nil, false
		}

		e, ok := err.(*Error)
		if !ok {
			return err, false
		}

		if e.code == code {
			return err, true
		}

		err = errors.Unwrap(err)
	}
}

// Is returns if err is Error and its code == code.
func Is(err error, code int32) bool {
	_, ok := Unwrap(err, code)
	return ok
}

// Code returns the code of err.
func Code(err error) int32 {
	if err == nil {
		return codeNil
	}

	e, ok := err.(*Error)
	if !ok {
		return codeUnknown
	}

	return e.code
}

// Msg returns the msg of err.
func Msg(err error) string {
	if err == nil {
		return nilString
	}

	e, ok := err.(*Error)
	if !ok {
		return err.Error()
	}

	return e.msg
}
