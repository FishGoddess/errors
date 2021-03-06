# ð§ Errors

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/errors)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**Errors** æ¯ä¸ä¸ªç¨äºä¼éå°å¤ç Go ä¸­éè¯¯çåºã

[Read me in English](./README.en.md)

### ðâ åè½ç¹æ§

* ä¼éå°å¤ç errorï¼å¯ï¼æ²¡äºããã

_åå²çæ¬çç¹æ§è¯·æ¥ç [HISTORY.md](./HISTORY.md)ãæªæ¥çæ¬çæ°ç¹æ§åè®¡åè¯·æ¥ç [FUTURE.md](./FUTURE.md)ã_

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
	err = errors.Forbidden(nil)           // Classic enough! Ah :)
	err = errors.NotFound(nil)            // Classic enough! Ah :)
	err = errors.InternalServerError(nil) // Classic enough! Ah :)
	err = errors.Timeout(nil)
	err = errors.NetworkError(nil)
	err = errors.DBError(nil)
}
```

### ð¥ è´¡ç®è

å¦ææ¨è§å¾ **Errors** ç¼ºå°æ¨éè¦çåè½ï¼é£å°± fork å°èªå·±ä»åºéä¾¿ç©ï¼ä¸è¦æ _**issue**_ï¼ä¸è¦æ _**issue**_ï¼ä¸è¦æ _**issue**_ :)ã
