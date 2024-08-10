// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

const (
	codeBadRequest   = 400
	codeForbidden    = 403
	codeNotFound     = 404
	codeRequireLogin = 1000
)

// BadRequest returns *Error with bad request code.
func BadRequest(message string) *Error {
	return Wrap(codeBadRequest, message)
}

// Forbidden returns *Error with forbidden code.
func Forbidden(message string) *Error {
	return Wrap(codeForbidden, message)
}

// NotFound returns *Error with not found code.
func NotFound(message string) *Error {
	return Wrap(codeNotFound, message)
}

// RequireLogin returns *Error with require login code.
func RequireLogin(message string) *Error {
	return Wrap(codeRequireLogin, message)
}

// MatchBadRequest matches err with bad request code.
func MatchBadRequest(err error) bool {
	return Match(err, codeBadRequest)
}

// MatchForbidden matches err with forbidden code.
func MatchForbidden(err error) bool {
	return Match(err, codeForbidden)
}

// MatchNotFound matches err with not found code.
func MatchNotFound(err error) bool {
	return Match(err, codeNotFound)
}

// MatchRequireLogin matches err with require login code.
func MatchRequireLogin(err error) bool {
	return Match(err, codeRequireLogin)
}
