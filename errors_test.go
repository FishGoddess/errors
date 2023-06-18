// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"testing"
)

// go test -v -cover -run=^TestNew$
func TestNew(t *testing.T) {
	err := New("test")
	if err.Error() != "test" {
		t.Errorf("err.Error() %s != 'test'", err.Error())
	}
}

// go test -v -cover -run=^TestNewF$
func TestNewF(t *testing.T) {
	err := NewF("test %d %.2f", 123, 3.14)
	if err.Error() != "test 123 3.14" {
		t.Errorf("err.Error() %s != 'test 123 3.14'", err.Error())
	}
}
