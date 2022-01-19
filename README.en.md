# 🔧 Errors

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/errors)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**Errors** is a lib for handling error gracefully in Go.

[阅读中文版的 Read me](./README.md)

### 🙋‍ Features

* Handling error gracefully, yep, that's all...

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

```go
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
```

### 👥 Contributing

If you find that something is not working as expected, just fork and fix by yourself, don't open an _**issue**_ :).
