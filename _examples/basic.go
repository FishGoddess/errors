// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"

	"github.com/FishGoddess/errors"
)

func main() {
	// Use wrap function to create an *Error error which has code and message.
	err := errors.Wrap(1000, "need login")
	fmt.Println(err)

	// You can get code and message of err anytime.
	fmt.Println(err.Code(), err.Message())

	// Try these ways to get code and message!
	// You will get default code or message if err doesn't have a code or message.
	fmt.Println(errors.Code(err, 6699), errors.Message(err, "default message"))
	fmt.Println(errors.Code(io.EOF, 6699), errors.Message(io.EOF, "default message"))

	// Also, we provide some useful information carrier for you.
	err = errors.Wrap(9999, "io timeout").With(io.EOF).WithCaller()
	fmt.Println(err)

	// What's more, we provide some shortcuts for you.
	// All these ways are returning a *Error and you are free to use all methods on *Error.
	berr := errors.BadRequest("id is wrong")
	ferr := errors.Forbidden("user isn't allowed")
	nerr := errors.NotFound("book not found")
	rerr := errors.RequireLogin("user requires login")
	fmt.Printf("%+v\n%+v\n%+v\n%+v\n", berr, ferr, nerr, rerr)
}
