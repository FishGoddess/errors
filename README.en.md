# 🧯 Errors

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/errors)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/errors/actions/workflows/test.yml/badge.svg)

**Errors** is a library for handling errors gracefully in Go.

[阅读中文版的 Read me](./README.md)

### 🙋‍ Features

* Handling errors gracefully, yep, that's all...

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

```go
package main

import (
	"fmt"
	"io"

	"github.com/FishGoddess/errors"
)

func main() {
	// Use wrap function to create an *Error error which has code and message.
	// You can get code and message of err anytime.
	err := errors.Wrap(1000, "need login")
	fmt.Println(err)
	fmt.Println(err.Code(), err.Message())

	// Try these ways to get code and message!
	// You will get default code or message if err doesn't have a code or message.
	code := errors.Code(err, 6699)
	message := errors.Message(err, "default message")
	fmt.Println(code, message)

	code = errors.Code(io.EOF, 6699)
	message = errors.Message(io.EOF, "default message")
	fmt.Println(code, message)

	// Also, we provide some useful information carrier for you.
	// For examples, you can carry an error or caller information.
	err = errors.Wrap(9999, "io timeout").With(io.EOF).WithCaller()
	fmt.Println(err)
	fmt.Println(errors.CodeMessage(err, 6666, "default message"))

	// What's more, we provide some shortcuts for you.
	// All these ways are returning *Error and you are free to use all methods on *Error.
	berr := errors.BadRequest("id is wrong")
	ferr := errors.Forbidden("user isn't allowed")
	nerr := errors.NotFound("book not found")
	rerr := errors.RequireLogin("user requires login")
	fmt.Printf("%+v\n%+v\n%+v\n%+v\n", berr, ferr, nerr, rerr)

	isBadRequest := errors.MatchBadRequest(berr)
	isForbidden := errors.MatchForbidden(ferr)
	isNotFound := errors.MatchNotFound(nerr)
	isRequireLogin := errors.MatchRequireLogin(rerr)
	fmt.Printf("isBadRequest: %+v\nisForbidden: %+v\nisNotFound: %+v\nisRequireLogin: %+v\n", isBadRequest, isForbidden, isNotFound, isRequireLogin)
}

```

* [basic](_examples/basic.go)

### 👥 Contributing

If you find that something is not working as expected, just fork and fix by yourself or open an _**issue**_ :).
