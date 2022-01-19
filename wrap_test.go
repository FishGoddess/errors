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
	if e, ok := IsNotFound(err); !ok || e.Error() != "404" {
		t.Error("NotFound or IsNotFound is wrong", err)
	}
}

// go test -v -cover -run=^TestTimeout$
func TestTimeout(t *testing.T) {
	err := Timeout(nil)
	if err != nil {
		t.Error("Timeout is wrong", err)
	}

	err = Timeout(errors.New("Timeout"))
	if e, ok := IsTimeout(err); !ok || e.Error() != "Timeout" {
		t.Error("Timeout or IsTimeout is wrong", err)
	}
}

// go test -v -cover -run=^TestNetworkError$
func TestNetworkError(t *testing.T) {
	err := NetworkError(nil)
	if err != nil {
		t.Error("NetworkError is wrong", err)
	}

	err = NetworkError(errors.New("NetworkError"))
	if e, ok := IsNetworkError(err); !ok || e.Error() != "NetworkError" {
		t.Error("NetworkError or IsNetworkError is wrong", err)
	}
}

// go test -v -cover -run=^TestDBError$
func TestDBError(t *testing.T) {
	err := DBError(nil)
	if err != nil {
		t.Error("DBError is wrong", err)
	}

	err = DBError(errors.New("DBError"))
	if e, ok := IsDBError(err); !ok || e.Error() != "DBError" {
		t.Error("DBError or IsDBError is wrong", err)
	}
}
