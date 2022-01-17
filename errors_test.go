package errors

import (
	"errors"
	"testing"
)

// go test -v -cover -run=^TestError$
func TestError(t *testing.T) {
	code := int32(500)

	err := WithCode(nil, code)
	if err != nil {
		t.Error("WithCode is wrong", err)
	}

	err = WithCode(errors.New("500"), code)
	if !Is(err, code) {
		t.Error("WithCode or Is is wrong", err)
	}
}

// go test -v -cover -run=^TestNotFound$
func TestNotFound(t *testing.T) {
	err := NotFound(nil)
	if err != nil {
		t.Error("NotFound is wrong", err)
	}

	err = NotFound(errors.New("500"))
	if !IsNotFound(err) {
		t.Error("NotFound or IsNotFound is wrong", err)
	}
}
