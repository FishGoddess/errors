// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"

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
func IsTestError(err error) bool {
	return errors.Is(err, codeTestError)
}

func main() {
	// We provide three graceful ways to handle error in Go: Wrap() and Unwrap() and Is().
	// Wrap wraps error with a code and Unwrap returns one error with this code.
	// Is returns one error is with the same code or not.
	// As you can see, we define two functions above, and this is the basic way to use this lib.
	err := TestError(errors.New("something wrong"))
	if IsTestError(err) {
		fmt.Println("I got a test error")
	}

	// Also, we provide some basic errors for you:
	err = errors.BadRequest(nil)          // Classic enough! Ah :)
	err = errors.Forbidden(nil)           // Classic enough! Ah :)
	err = errors.NotFound(nil)            // Classic enough! Ah :)
	err = errors.RequestTimeout(nil)      // Classic enough! Ah :)
	err = errors.InternalServerError(nil) // Classic enough! Ah :)
	err = errors.DBError(nil)
	err = errors.PageTokenInvalid(nil)

	// Use WithMsg to carry a message.
	err = errors.Wrap(io.EOF, codeTestError, errors.WithMsg("test"))
	fmt.Println(err.Error())
	fmt.Println(errors.Msg(err))
	fmt.Println(errors.MsgOrDefault(io.EOF, "default error message"))
}
