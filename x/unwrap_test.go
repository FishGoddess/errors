// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package x

import (
	"io"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestCode$
func TestCode(t *testing.T) {
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
			err:         New(1000, "wow"),
			defaultCode: 999,
			code:        1000,
		},
		{
			err:         Wrap(io.EOF, 1000, "eof"),
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
			err:            New(1000, "wow"),
			defaultMessage: "xxx",
			message:        "wow",
		},
		{
			err:            Wrap(io.EOF, 1000, "eof"),
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
