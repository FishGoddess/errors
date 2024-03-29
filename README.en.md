# 🧯 Errors

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/errors)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/errors/actions/workflows/test.yml/badge.svg)

**Errors** is a lib for handling error gracefully in Go.

[阅读中文版的 Read me](./README.md)

### 🙋‍ Features

* Handling error gracefully, yep, that's all...

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

```go
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

```

* [basic](_examples/basic.go)
* [status](_examples/status.go)

### 👥 Contributing

If you find that something is not working as expected, just fork and fix by yourself or open an _**issue**_ :).
