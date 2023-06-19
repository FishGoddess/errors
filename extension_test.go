// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"testing"
)

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

// go test -v -cover -run=^TestPageTokenInvalid$
func TestPageTokenInvalid(t *testing.T) {
	err := PageTokenInvalid(nil)
	if err != nil {
		t.Error("PageTokenInvalid is wrong", err)
	}

	err = PageTokenInvalid(errors.New("page token invalid"))
	if !IsPageTokenInvalid(err) {
		t.Error("IsPageTokenInvalid is wrong", err)
	}

	if e, ok := UnwrapPageTokenInvalid(err); !ok || e.Error() != "page token invalid" {
		t.Error("PageTokenInvalid or UnwrapPageTokenInvalid is wrong", err)
	}
}
