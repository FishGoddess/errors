# 🧯 Errors

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/errors)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/errors/actions/workflows/test.yml/badge.svg)

**Errors** 是一个用于优雅地处理 Go 中错误的库。

[Read me in English](./README.en.md)

### 🙋‍ 功能特性

* 优雅地处理 error，嗯，没了。。。

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

```go
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
}

```

* [basic](_examples/basic.go)

### 👥 贡献者

如果您觉得 **Errors** 缺少您需要的功能，那就 fork 到自己仓库随便玩。当然，也可以提 _**issue**_ :)。
