// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

const (
	codeTimeout      = 1000
	codeNetworkError = 1100
	codeDBError      = 1200
)

// Timeout returns a timeout error.
func Timeout(err error, opts ...Option) error {
	return Wrap(err, codeTimeout, opts...)
}

// IsTimeout if err is timeout.
func IsTimeout(err error) bool {
	return Is(err, codeTimeout)
}

// UnwrapTimeout if err is timeout.
func UnwrapTimeout(err error) (error, bool) {
	return Unwrap(err, codeTimeout)
}

// NetworkError returns a network error.
func NetworkError(err error, opts ...Option) error {
	return Wrap(err, codeNetworkError, opts...)
}

// IsNetworkError if err is network error.
func IsNetworkError(err error) bool {
	return Is(err, codeNetworkError)
}

// UnwrapNetworkError if err is network error.
func UnwrapNetworkError(err error) (error, bool) {
	return Unwrap(err, codeNetworkError)
}

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
