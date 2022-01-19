// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/20 00:07:35

package main

import (
	"fmt"

	"github.com/FishGoddess/errors"
)

const (
	codeTestError = 10000 // Your error's code
)

// TestError returns a test error.
func TestError(err error) error {
	return errors.Wrap(err, codeTestError)
}

// IsTestError if err is test error.
func IsTestError(err error) (error, bool) {
	return errors.Is(err, codeTestError)
}

func main() {
	// We provide two graceful ways to handle error in Go: Wrap() and Is().
	// Wrap wraps error with a code and Is returns one error is with this code.
	// As you can see, we define two functions above, and this is the basic way to use this lib.
	err := TestError(errors.New("something wrong"))
	if e, ok := IsTestError(err); ok {
		fmt.Printf("I got a test error which says: \"%s\"\n", e.Error())
	}

	// Also, we provide some basic errors for you:
	err = errors.NotFound(nil) // Classic enough! Ah :)
	err = errors.Timeout(nil)
	err = errors.NetworkError(nil)
	err = errors.DBError(nil)
}
