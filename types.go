// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

const (
	CodeBadRequest   = 400
	CodeForbidden    = 403
	CodeNotFound     = 404
	CodeRequireLogin = 1000
)

func BadRequest(message string) *Error {
	return Wrap(CodeBadRequest, message)
}

func Forbidden(message string) *Error {
	return Wrap(CodeForbidden, message)
}

func NotFound(message string) *Error {
	return Wrap(CodeNotFound, message)
}

func RequireLogin(message string) *Error {
	return Wrap(CodeRequireLogin, message)
}
