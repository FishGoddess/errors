// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"io"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWrap$
func TestWrap(t *testing.T) {
	testCases := []struct {
		code        int32
		message     string
		args        []any
		wantCode    int32
		wantMessage string
	}{
		{code: -1000, message: "io timeout", args: nil, wantCode: -1000, wantMessage: "io timeout"},
		{code: 1000, message: "need login", args: nil, wantCode: 1000, wantMessage: "need login"},
		{code: 2000, message: "with args %d %s %+v", args: []any{1, "x", true}, wantCode: 2000, wantMessage: "with args 1 x true"},
	}

	for _, testCase := range testCases {
		err := Wrap(testCase.code, testCase.message, testCase.args...)
		if err.Code() != testCase.wantCode {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}

		if err.Message() != testCase.wantMessage {
			t.Errorf("err.Message() %s != testCase.code %s", err.Message(), testCase.message)
		}

		if err.Unwrap() != nil {
			t.Errorf("err.Unwrap() %+v != nil", err.Unwrap())
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWrapWith$
func TestWrapWith(t *testing.T) {
	testCases := []struct {
		code    int32
		message string
		cause   error
	}{
		{code: -1000, message: "eof", cause: io.EOF},
		{code: 1000, message: "need login", cause: nil},
	}

	for _, testCase := range testCases {
		err := Wrap(testCase.code, testCase.message).With(testCase.cause)
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
		args        map[string]any
		errorString string
	}{
		{code: -1000, message: "eof", cause: io.EOF, args: nil, errorString: "-1000: eof (EOF)"},
		{code: -1000, message: "eof", cause: io.EOF, args: map[string]any{"x": 666, "y": "300ms", "z": false}, errorString: "-1000: eof {\"x\":666,\"y\":\"300ms\",\"z\":false} (EOF)"},
		{code: 1000, message: "need login", args: nil, errorString: "1000: need login"},
		{code: 1000, message: "need login", args: map[string]any{"x": 123, "y": "100ms", "z": true}, errorString: "1000: need login {\"x\":123,\"y\":\"100ms\",\"z\":true}"},
	}

	for _, testCase := range testCases {
		err := Wrap(testCase.code, testCase.message).With(testCase.cause)

		for key, value := range testCase.args {
			err = err.WithArgs(key, value)
		}

		if len(err.args) != len(testCase.args) {
			t.Errorf("len(err.args) %d != len(testCase.args) %d", len(err.args), len(testCase.args))
		}

		if err.Error() != testCase.errorString {
			t.Errorf("err.Error() %s != testCase.errorString %s", err.Error(), testCase.errorString)
		}
	}
}
