# 🧯 Errors

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/errors)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/errors/actions/workflows/test.yml/badge.svg)

**Errors** 是一个用于优雅地处理 Go 中错误的库。

[Read me in English](./README.en.md)

### 🙋‍ 功能特性

* 优雅地处理错误，嗯，没了。。。

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
	// For examples, you can carry an error, caller information or some args.
	// Remember the count of args should be even and their keys should be a string.
	err = errors.Wrap(9999, "io timeout").With(io.EOF).WithCaller().WithArgs("user_id", 123, "timeout", "200ms")
	fmt.Println(err)
	fmt.Println(errors.CodeMessage(err, 6666, "default message"))

	// What's more, we provide some shortcuts for you.
	// All these ways are returning *Error and you are free to use all methods on *Error.
	berr := errors.BadRequest("id is wrong")
	ferr := errors.Forbidden("user isn't allowed")
	nerr := errors.NotFound("book not found")
	fmt.Printf("%+v\n%+v\n%+v\n", berr, ferr, nerr)

	isBadRequest := errors.IsBadRequest(berr)
	isForbidden := errors.IsForbidden(ferr)
	isNotFound := errors.IsNotFound(nerr)
	fmt.Printf("isBadRequest: %+v\nisForbidden: %+v\nisNotFound: %+v\n", isBadRequest, isForbidden, isNotFound)
}
```

* [basic](_examples/basic.go)

### 👥 贡献者

如果您觉得 **errors** 缺少您需要的功能，那就 fork 到自己仓库随便玩。当然，也可以提 _**issue**_ :)。
