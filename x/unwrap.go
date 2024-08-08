// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

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
func Message(err error, defaultMessage string) string {
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

	return defaultMessage
}

// CodeMessage unwraps error and gets code & message of it.
// It returns 0 & "" if err is nil.
// It returns defaultCode if err doesn't have a code.
// It returns defaultMessage if err doesn't have a message.
func CodeMessage(err error, defaultCode int32, defaultMessage string) (int32, string) {
	code := Code(err, defaultCode)
	message := Message(err, defaultMessage)

	return code, message
}
