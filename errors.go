// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/19 23:43:15

package errors

import (
	"errors"
	"fmt"
)

var (
	// ErrorFormat formats e and returns a string in Error().
	ErrorFormat = func(e *Error) string {
		return fmt.Sprintf("%s", e.err.Error())
	}

	// StringFormat formats e and returns a string in String().
	StringFormat = func(e *Error) string {
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
		return "<nil>"
	}
	return ErrorFormat(e)
}

func (e *Error) String() string {
	if e == nil || e.err == nil {
		return "<nil>"
	}
	return StringFormat(e)
}

// Unwrap returns if e has the same type of target.
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
func Wrap(err error, code int32) error {
	if err == nil {
		return nil
	}

	return &Error{
		err:  err,
		code: code,
	}
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
