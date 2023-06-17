// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"io"
	"testing"
)

// go test -v -cover -run=^TestError$
func TestError(t *testing.T) {
	code := int32(500)

	err := Wrap(nil, code)
	if err != nil {
		t.Errorf("err %+v is wrong", err)
	}

	err = Wrap(errors.New("500"), code)
	if e, ok := Unwrap(err, code); !ok || e.Error() != "500" {
		t.Errorf("err %+v is wrong", err)
	}
}

// go test -v -cover -run=^TestCode$
func TestCode(t *testing.T) {
	code := int32(500)

	got := Code(nil)
	if got != codeNil {
		t.Errorf("got %d is wrong", got)
	}

	got = Code(Wrap(nil, code))
	if got != codeNil {
		t.Errorf("got %d is wrong", got)
	}

	got = Code(io.EOF)
	if got != codeUnknown {
		t.Errorf("got %d is wrong", got)
	}

	got = Code(Wrap(errors.New("500"), code))
	if got != code {
		t.Errorf("got %d is wrong", got)
	}
}

// go test -v -cover -run=^TestMsg$
func TestMsg(t *testing.T) {
	code := int32(500)

	msg := Msg(Wrap(nil, code))
	if msg != nilString {
		t.Errorf("msg %s is wrong", msg)
	}

	msg = Msg(Wrap(errors.New("500"), code))
	if msg != "500" {
		t.Errorf("msg %s is wrong", msg)
	}

	msg = Msg(Wrap(errors.New("500"), code, WithMsg("internal")))
	if msg != "internal" {
		t.Errorf("msg %s is wrong", msg)
	}
}
