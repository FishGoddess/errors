// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

const (
	codeDBError          = 1100
	codePageTokenInvalid = 1200
)

// DBError returns a db error.
func DBError(err error, opts ...Option) error {
	return Wrap(err, codeDBError, opts...)
}

// IsDBError if err is db error.
func IsDBError(err error) bool {
	return Is(err, codeDBError)
}

// UnwrapDBError if err is db error.
func UnwrapDBError(err error) (error, bool) {
	return Unwrap(err, codeDBError)
}

// PageTokenInvalid returns a page token invalid error.
func PageTokenInvalid(err error, opts ...Option) error {
	return Wrap(err, codePageTokenInvalid, opts...)
}

// IsPageTokenInvalid if err is page token invalid.
func IsPageTokenInvalid(err error) bool {
	return Is(err, codePageTokenInvalid)
}

// UnwrapPageTokenInvalid if err is page token invalid.
func UnwrapPageTokenInvalid(err error) (error, bool) {
	return Unwrap(err, codePageTokenInvalid)
}
