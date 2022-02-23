# 🔧 Errors

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/errors)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**Errors** 是一个用于优雅地处理 Go 中错误的库。

[Read me in English](./README.en.md)

### 🙋‍ 功能特性

* 优雅地处理 error，嗯，没了。。。

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

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
		fmt.Printf("I got a test error")
	}

	// Also, we provide some basic errors for you:
	err = errors.BadRequest(nil)          // Classic enough! Ah :)
	err = errors.NotFound(nil)            // Classic enough! Ah :)
	err = errors.InternalServerError(nil) // Classic enough! Ah :)
	err = errors.Timeout(nil)
	err = errors.NetworkError(nil)
	err = errors.DBError(nil)
}
```

### 👥 贡献者

如果您觉得 **Errors** 缺少您需要的功能，那就 fork 到自己仓库随便玩，不要提 _**issue**_，不要提 _**issue**_，不要提 _**issue**_ :)。
