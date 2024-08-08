// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package x

import (
	"io"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNew$
func TestNew(t *testing.T) {
	testCases := []struct {
		code    int32
		message string
	}{
		{code: -1000, message: "io timeout"},
		{code: 1000, message: "need login"},
	}

	for _, testCase := range testCases {
		err := New(testCase.code, testCase.message)
		if err.Code() != testCase.code {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}

		if err.Message() != testCase.message {
			t.Errorf("err.Message() %s != testCase.code %s", err.Message(), testCase.message)
		}

		if err.Unwrap() != nil {
			t.Errorf("err.Unwrap() %+v != nil", err.Unwrap())
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWrap$
func TestWrap(t *testing.T) {
	testCases := []struct {
		code    int32
		message string
		cause   error
	}{
		{code: -1000, message: "eof", cause: io.EOF},
		{code: 1000, message: "need login"},
	}

	for _, testCase := range testCases {
		err := Wrap(testCase.cause, testCase.code, testCase.message)
		if err.Code() != testCase.code {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}

		if err.Message() != testCase.message {
			t.Errorf("err.Message() %s != testCase.code %s", err.Message(), testCase.message)
		}

		if err.Unwrap() != testCase.cause {
			t.Errorf("err.Unwrap() %+v != testCase.code %+v", err.Unwrap(), testCase.cause)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestErrorError$
func TestErrorError(t *testing.T) {
	testCases := []struct {
		code        int32
		message     string
		cause       error
		errorString string
	}{
		{code: -1000, message: "eof", cause: io.EOF, errorString: "-1000: eof (EOF)"},
		{code: 1000, message: "need login", errorString: "1000: need login"},
	}

	for _, testCase := range testCases {
		err := Wrap(testCase.cause, testCase.code, testCase.message)
		if err.Error() != testCase.errorString {
			t.Errorf("err.Error() %s != testCase.errorString %s", err.Error(), testCase.errorString)
		}
	}
}
