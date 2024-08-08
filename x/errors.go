// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
)

func New(text string) error {
	return errors.New(text)
}

// Is is a shortcut of errors.Is.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As is a shortcut of errors.As.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Unwrap is a shortcut of errors.Unwrap.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Join is a shortcut of errors.Join.
func Join(errs ...error) error {
	return errors.Join(errs...)
}
