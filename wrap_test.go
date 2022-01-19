// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/01/19 23:52:40

package errors

import (
	"errors"
	"testing"
)

// go test -v -cover -run=^TestNotFound$
func TestNotFound(t *testing.T) {
	err := NotFound(nil)
	if err != nil {
		t.Error("NotFound is wrong", err)
	}

	err = NotFound(errors.New("404"))
	if !IsNotFound(err) {
		t.Error("IsNotFound is wrong", err)
	}

	if e, ok := UnwrapNotFound(err); !ok || e.Error() != "404" {
		t.Error("NotFound or UnwrapNotFound is wrong", err)
	}
}

// go test -v -cover -run=^TestTimeout$
func TestTimeout(t *testing.T) {
	err := Timeout(nil)
	if err != nil {
		t.Error("Timeout is wrong", err)
	}

	err = Timeout(errors.New("timeout"))
	if !IsTimeout(err) {
		t.Error("IsTimeout is wrong", err)
	}

	if e, ok := UnwrapTimeout(err); !ok || e.Error() != "timeout" {
		t.Error("Timeout or UnwrapTimeout is wrong", err)
	}
}

// go test -v -cover -run=^TestNetworkError$
func TestNetworkError(t *testing.T) {
	err := NetworkError(nil)
	if err != nil {
		t.Error("NetworkError is wrong", err)
	}

	err = NetworkError(errors.New("network error"))
	if !IsNetworkError(err) {
		t.Error("IsNetworkError is wrong", err)
	}

	if e, ok := UnwrapNetworkError(err); !ok || e.Error() != "network error" {
		t.Error("NetworkError or UnwrapNetworkError is wrong", err)
	}
}

// go test -v -cover -run=^TestDBError$
func TestDBError(t *testing.T) {
	err := DBError(nil)
	if err != nil {
		t.Error("DBError is wrong", err)
	}

	err = DBError(errors.New("db error"))
	if !IsDBError(err) {
		t.Error("IsDBError is wrong", err)
	}

	if e, ok := UnwrapDBError(err); !ok || e.Error() != "db error" {
		t.Error("DBError or UnwrapDBError is wrong", err)
	}
}
