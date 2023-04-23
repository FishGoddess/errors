// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"net/http"
	"testing"
)

// go test -v -cover -run=^TestOptionApply$
func TestOptionApply(t *testing.T) {
	opt := Option(func(e *Error) {
		e.code = http.StatusOK
	})

	got := &Error{code: 0}
	expect := &Error{code: http.StatusOK}

	opt.Apply(got)
	if got.String() != expect.String() {
		t.Errorf("got %s != expect %s", got.String(), expect.String())
	}
}

// go test -v -cover -run=^TestWithMsg$
func TestWithMsg(t *testing.T) {
	got := &Error{msg: ""}
	expect := &Error{msg: "ok"}

	WithMsg("ok")(got)
	if got.String() != expect.String() {
		t.Errorf("got %s != expect %s", got.String(), expect.String())
	}

	got = &Error{msg: ""}
	expect = &Error{msg: "ok123"}

	WithMsg("%s%d", "ok", 123)(got)
	if got.String() != expect.String() {
		t.Errorf("got %s != expect %s", got.String(), expect.String())
	}
}
