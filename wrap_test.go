// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"io"
	"testing"
)

func compareMaps(m1 map[string]any, m2 map[string]any) error {
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok {
			return fmt.Errorf("key %s not found in m2 %+v", k1, m2)
		}

		if v1 != v2 {
			return fmt.Errorf("v1 %+v != v2 %+v", v1, v2)
		}
	}

	for k2, v2 := range m2 {
		v1, ok := m1[k2]
		if !ok {
			return fmt.Errorf("key %s not found in m1 %+v", k2, m1)
		}

		if v1 != v2 {
			return fmt.Errorf("v1 %+v != v2 %+v", v1, v2)
		}
	}

	return nil
}

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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWrapWithArgs$
func TestWrapWithArgs(t *testing.T) {
	testCases := []struct {
		code     int32
		message  string
		cause    error
		args     []any
		wantArgs map[string]any
	}{
		{code: -1000, message: "eof", cause: io.EOF, args: nil, wantArgs: nil},
		{code: 1000, message: "need login", cause: nil, args: nil, wantArgs: nil},
		{code: 1000, message: "need login", cause: nil, args: []any{io.EOF}, wantArgs: map[string]any{keyBad + "_1": io.EOF}},
		{code: -1000, message: "eof", cause: io.EOF, args: []any{1, true, 3.14, io.EOF}, wantArgs: map[string]any{keyBad + "_1": 1, keyBad + "_2": true, keyBad + "_3": 3.14, keyBad + "_4": io.EOF}},
		{code: -1000, message: "eof", cause: io.EOF, args: []any{1, true, 3.14, "abc"}, wantArgs: map[string]any{keyBad + "_1": 1, keyBad + "_2": true, keyBad + "_3": 3.14, keyBad + "_4": "abc"}},
		{code: 1000, message: "need login", cause: io.EOF, args: []any{"k1", 1, "k2", true, "k3", 3.14, "key", "abc"}, wantArgs: map[string]any{"k1": 1, "k2": true, "k3": 3.14, "key": "abc"}},
	}

	for _, testCase := range testCases {
		err := Wrap(testCase.code, testCase.message).With(testCase.cause).WithArgs(testCase.args...)
		if err.Code() != testCase.code {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}

		if err.Message() != testCase.message {
			t.Errorf("err.Message() %s != testCase.code %s", err.Message(), testCase.message)
		}

		if err.Unwrap() != testCase.cause {
			t.Errorf("err.Unwrap() %+v != testCase.code %+v", err.Unwrap(), testCase.cause)
		}

		args := err.Args()
		if merr := compareMaps(args, testCase.wantArgs); merr != nil {
			t.Error(merr)
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
