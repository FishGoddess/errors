// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"testing"
)

// go test -v -cover -run=^TestError$
func TestError(t *testing.T) {
	code := int32(500)

	err := Wrap(nil, code)
	if err != nil {
		t.Error("Wrap is wrong", err)
	}

	err = Wrap(errors.New("500"), code)
	if e, ok := Unwrap(err, code); !ok || e.Error() != "500" {
		t.Error("Wrap or Unwrap is wrong", err, e)
	}
}
