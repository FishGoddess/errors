// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import "fmt"

// Code unwraps error and gets code of it.
// It returns 0 if err is nil.
// It returns defaultCode if err doesn't have a code.
func Code(err error, defaultCode int32) int32 {
	if err == nil {
		return 0
	}

	cerr, ok := err.(interface {
		Code() int32
	})
	if ok {
		return cerr.Code()
	}

	var xerr *Error
	if As(err, &xerr) {
		return xerr.Code()
	}

	return defaultCode
}

// Message unwraps error and gets message of it.
// It returns "" if err is nil.
// It returns defaultMessage if err doesn't have a message.
func Message(err error, defaultMessage string, args ...any) string {
	if err == nil {
		return ""
	}

	merr, ok := err.(interface {
		Message() string
	})
	if ok {
		return merr.Message()
	}

	var xerr *Error
	if As(err, &xerr) {
		return xerr.Message()
	}

	if len(args) > 0 {
		defaultMessage = fmt.Sprintf(defaultMessage, args...)
	}

	return defaultMessage
}

// CodeMessage unwraps error and gets code & message of it.
// It returns 0 & "" if err is nil.
// It returns defaultCode if err doesn't have a code.
// It returns defaultMessage if err doesn't have a message.
func CodeMessage(err error, defaultCode int32, defaultMessage string, args ...any) (int32, string) {
	code := Code(err, defaultCode)
	message := Message(err, defaultMessage, args...)

	return code, message
}

// Match unwraps error and check if its code equals to code.
func Match(err error, code int32) bool {
	if err == nil {
		return code == 0
	}

	cerr, ok := err.(interface {
		Code() int32
	})
	if ok {
		return cerr.Code() == code
	}

	var xerr *Error
	if As(err, &xerr) {
		return xerr.Code() == code
	}

	return false
}
