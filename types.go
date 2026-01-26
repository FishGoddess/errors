// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

const (
	codeBadRequest = 400
	codeForbidden  = 403
	codeNotFound   = 404
)

// BadRequest returns *Error with bad request code.
func BadRequest(message string, args ...any) *Error {
	return Wrap(codeBadRequest, message, args...)
}

// Forbidden returns *Error with forbidden code.
func Forbidden(message string, args ...any) *Error {
	return Wrap(codeForbidden, message, args...)
}

// NotFound returns *Error with not found code.
func NotFound(message string, args ...any) *Error {
	return Wrap(codeNotFound, message, args...)
}

// IsBadRequest checks err with bad request code.
func IsBadRequest(err error) bool {
	return IsCode(err, codeBadRequest)
}

// IsForbidden checks err with forbidden code.
func IsForbidden(err error) bool {
	return IsCode(err, codeForbidden)
}

// IsNotFound checks err with not found code.
func IsNotFound(err error) bool {
	return IsCode(err, codeNotFound)
}

// MatchBadRequest matches err with bad request code.
// Deprecated: Use IsBadRequest instead because this name is 'ugly' to me.
func MatchBadRequest(err error) bool {
	return IsBadRequest(err)
}

// MatchForbidden matches err with forbidden code.
// Deprecated: Use IsForbidden instead because this name is 'ugly' to me.
func MatchForbidden(err error) bool {
	return IsForbidden(err)
}

// MatchNotFound matches err with not found code.
// Deprecated: Use IsNotFound instead because this name is 'ugly' to me.
func MatchNotFound(err error) bool {
	return IsNotFound(err)
}
