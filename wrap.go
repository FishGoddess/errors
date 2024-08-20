// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	keyBad     = "errors.bad_key"
	keyCaller  = "errors.caller"
	keyCallers = "errors.callers"
)

type Error struct {
	code    int32
	message string
	cause   error
	args    map[string]any
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

// With carries the cause err for *Error.
func (e *Error) With(err error) *Error {
	e.cause = err
	return e
}

// WithArgs carries some args for *Error.
func (e *Error) WithArgs(args ...any) *Error {
	if e.args == nil {
		e.args = make(map[string]any, 4)
	}

	badKeys := 0
	badKey := func() string {
		badKeys++
		return keyBad + "_" + strconv.Itoa(badKeys)
	}

	for len(args) > 0 {
		// One last arg, tag it as bad key.
		if len(args) == 1 {
			e.args[badKey()] = args[0]
			break
		}

		// len(kvs) >= 2, the key should be a string
		k, v := args[0], args[1]
		if ks, ok := k.(string); ok {
			e.args[ks] = v
		} else {
			e.args[badKey()] = k
			e.args[badKey()] = v
		}

		args = args[2:]
	}

	return e
}

// WithCaller carries the top caller for *Error.
func (e *Error) WithCaller() *Error {
	return e.WithArgs(keyCaller, Caller())
}

// WithCallers carries all callers for *Error.
func (e *Error) WithCallers() *Error {
	return e.WithArgs(keyCallers, Callers())
}

// Code returns the code of *Error.
func (e *Error) Code() int32 {
	return e.code
}

// Message returns the message of *Error.
func (e *Error) Message() string {
	return e.message
}

// Unwrap returns the cause of *Error.
func (e *Error) Unwrap() error {
	return e.cause
}

// Args returns the args of *Error.
func (e *Error) Args() map[string]any {
	if e.args == nil {
		return nil
	}

	args := make(map[string]any, len(e.args))
	for k, v := range e.args {
		args[k] = v
	}

	return args
}

// Error returns *Error as string.
func (e *Error) Error() string {
	return e.String()
}

func (e *Error) marshalArgs() []byte {
	args, err := json.Marshal(e.args)
	if err == nil {
		return args
	}

	argsString := fmt.Sprintf(`{"args":"%+v","marshal_error":"%+v"}`, e.args, err)
	return []byte(argsString)
}

// String returns *Error as string.
func (e *Error) String() string {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "%d: %s", e.code, e.message)

	if len(e.args) > 0 {
		buff.WriteByte(' ')
		buff.Write(e.marshalArgs())
	}

	if e.cause != nil {
		fmt.Fprintf(&buff, " (%+v)", e.cause)
	}

	return buff.String()
}
