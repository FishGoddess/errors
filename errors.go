// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"fmt"
)

// New returns a string error.
func New(text string) error {
	return errors.New(text)
}

// Newf returns a string error formatted with params.
func Newf(text string, params ...interface{}) error {
	return fmt.Errorf(text, params...)
}
