// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"io"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestCode$
func TestCode(t *testing.T) {
	testErr := &testError{}

	testCases := []struct {
		err         error
		defaultCode int32
		code        int32
	}{
		{
			err:         nil,
			defaultCode: 999,
			code:        0,
		},
		{
			err:         io.EOF,
			defaultCode: 999,
			code:        999,
		},
		{
			err:         testErr,
			defaultCode: 999,
			code:        testErr.Code(),
		},
		{
			err:         Wrap(1000, "wow"),
			defaultCode: 999,
			code:        1000,
		},
		{
			err:         Wrap(1000, "eof").With(io.EOF),
			defaultCode: 999,
			code:        1000,
		},
	}

	for _, testCase := range testCases {
		code := Code(testCase.err, testCase.defaultCode)
		if code != testCase.code {
			t.Errorf("code %d != testCase.code %d", code, testCase.code)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestMessage$
func TestMessage(t *testing.T) {
	testCases := []struct {
		err            error
		defaultMessage string
		message        string
	}{
		{
			err:            nil,
			defaultMessage: "xxx",
			message:        "",
		},
		{
			err:            io.EOF,
			defaultMessage: "eof",
			message:        "eof",
		},
		{
			err:            Wrap(1000, "wow"),
			defaultMessage: "xxx",
			message:        "wow",
		},
		{
			err:            Wrap(1000, "eof").With(io.EOF),
			defaultMessage: "xxx",
			message:        "eof",
		},
	}

	for _, testCase := range testCases {
		message := Message(testCase.err, testCase.defaultMessage)
		if message != testCase.message {
			t.Errorf("message %s != testCase.message %s", message, testCase.message)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestMatch$
func TestMatch(t *testing.T) {
	testErr := &testError{}

	testCases := []struct {
		err   error
		code  int32
		match bool
	}{
		{
			err:   nil,
			code:  0,
			match: true,
		},
		{
			err:   nil,
			code:  999,
			match: false,
		},
		{
			err:   io.EOF,
			code:  999,
			match: false,
		},
		{
			err:   testErr,
			code:  testErr.Code(),
			match: true,
		},
		{
			err:   Wrap(1000, "wow"),
			code:  1000,
			match: true,
		},
		{
			err:   Wrap(1000, "eof").With(io.EOF),
			code:  1000,
			match: true,
		},
	}

	for _, testCase := range testCases {
		match := Match(testCase.err, testCase.code)
		if match != testCase.match {
			t.Errorf("match %+v != testCase.match %+v", match, testCase.match)
		}
	}
}
