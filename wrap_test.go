// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"testing"
)

// go test -v -cover -run=^TestBadRequest$
func TestBadRequest(t *testing.T) {
	err := BadRequest(nil)
	if err != nil {
		t.Error("BadRequest is wrong", err)
	}

	err = BadRequest(errors.New("400"))
	if !IsBadRequest(err) {
		t.Error("IsBadRequest is wrong", err)
	}

	if e, ok := UnwrapBadRequest(err); !ok || e.Error() != "400" {
		t.Error("BadRequest or UnwrapBadRequest is wrong", err)
	}
}

// go test -v -cover -run=^TestForbidden$
func TestForbidden(t *testing.T) {
	err := Forbidden(nil)
	if err != nil {
		t.Error("Forbidden is wrong", err)
	}

	err = Forbidden(errors.New("403"))
	if !IsForbidden(err) {
		t.Error("IsForbidden is wrong", err)
	}

	if e, ok := UnwrapForbidden(err); !ok || e.Error() != "403" {
		t.Error("Forbidden or UnwrapForbidden is wrong", err)
	}
}

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

// go test -v -cover -run=^TestInternalServerError$
func TestInternalServerError(t *testing.T) {
	err := InternalServerError(nil)
	if err != nil {
		t.Error("InternalServerError is wrong", err)
	}

	err = InternalServerError(errors.New("400"))
	if !IsInternalServerError(err) {
		t.Error("IsInternalServerError is wrong", err)
	}

	if e, ok := UnwrapInternalServerError(err); !ok || e.Error() != "400" {
		t.Error("InternalServerError or UnwrapInternalServerError is wrong", err)
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
