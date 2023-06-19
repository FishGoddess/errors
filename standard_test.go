// Copyright 2022 FishGoddess. All rights reserved.
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

	err = BadRequest(errors.New("bad request"))
	if !IsBadRequest(err) {
		t.Error("IsBadRequest is wrong", err)
	}

	if e, ok := UnwrapBadRequest(err); !ok || e.Error() != "bad request" {
		t.Error("BadRequest or UnwrapBadRequest is wrong", err)
	}
}

// go test -v -cover -run=^TestForbidden$
func TestForbidden(t *testing.T) {
	err := Forbidden(nil)
	if err != nil {
		t.Error("Forbidden is wrong", err)
	}

	err = Forbidden(errors.New("forbidden"))
	if !IsForbidden(err) {
		t.Error("IsForbidden is wrong", err)
	}

	if e, ok := UnwrapForbidden(err); !ok || e.Error() != "forbidden" {
		t.Error("Forbidden or UnwrapForbidden is wrong", err)
	}
}

// go test -v -cover -run=^TestNotFound$
func TestNotFound(t *testing.T) {
	err := NotFound(nil)
	if err != nil {
		t.Error("NotFound is wrong", err)
	}

	err = NotFound(errors.New("not found"))
	if !IsNotFound(err) {
		t.Error("IsNotFound is wrong", err)
	}

	if e, ok := UnwrapNotFound(err); !ok || e.Error() != "not found" {
		t.Error("NotFound or UnwrapNotFound is wrong", err)
	}
}

// go test -v -cover -run=^TestRequestTimeout$
func TestRequestTimeout(t *testing.T) {
	err := RequestTimeout(nil)
	if err != nil {
		t.Error("RequestTimeout is wrong", err)
	}

	err = RequestTimeout(errors.New("request timeout"))
	if !IsRequestTimeout(err) {
		t.Error("IsRequestTimeout is wrong", err)
	}

	if e, ok := UnwrapRequestTimeout(err); !ok || e.Error() != "request timeout" {
		t.Error("RequestTimeout or UnwrapRequestTimeout is wrong", err)
	}
}

// go test -v -cover -run=^TestInternalServerError$
func TestInternalServerError(t *testing.T) {
	err := InternalServerError(nil)
	if err != nil {
		t.Error("InternalServerError is wrong", err)
	}

	err = InternalServerError(errors.New("internal server error"))
	if !IsInternalServerError(err) {
		t.Error("IsInternalServerError is wrong", err)
	}

	if e, ok := UnwrapInternalServerError(err); !ok || e.Error() != "internal server error" {
		t.Error("InternalServerError or UnwrapInternalServerError is wrong", err)
	}
}
