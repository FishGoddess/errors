// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"testing"
)

// go test -v -cover -run=^TestWrapAndUnwrap$
func TestWrapAndUnwrap(t *testing.T) {
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

// go test -v -cover -run=^TestMsg$
func TestMsg(t *testing.T) {
	code := int32(500)

	msg, ok := Msg(Wrap(nil, code))
	if ok {
		t.Error("msg should be not ok")
	}

	if msg != "" {
		t.Errorf("msg %s is wrong", msg)
	}

	msg, ok = Msg(Wrap(errors.New("500"), code))
	if ok {
		t.Error("msg should be not ok")
	}

	if msg != "" {
		t.Errorf("msg %s is wrong", msg)
	}

	msg, ok = Msg(Wrap(errors.New("500"), code, WithMsg("internal")))
	if !ok {
		t.Error("msg should be ok")
	}

	if msg != "internal" {
		t.Errorf("msg %s is wrong", msg)
	}
}

// go test -v -cover -run=^TestMsgOrDefault$
func TestMsgOrDefault(t *testing.T) {
	code := int32(500)

	msg := MsgOrDefault(Wrap(nil, code), "xxx")
	if msg != "xxx" {
		t.Errorf("msg %s is wrong", msg)
	}

	msg = MsgOrDefault(Wrap(errors.New("500"), code), "abc")
	if msg != "abc" {
		t.Errorf("msg %s is wrong", msg)
	}

	msg = MsgOrDefault(Wrap(errors.New("500"), code, WithMsg("internal")), "ignore")
	if msg != "internal" {
		t.Errorf("msg %s is wrong", msg)
	}
}
