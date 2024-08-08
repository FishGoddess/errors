// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package x

import (
	"os"
	"testing"
)

type testError struct {
	reason string
}

func (te *testError) Code() int32 {
	return 500
}

func (te *testError) Error() string {
	return te.reason
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestIs$
func TestIs(t *testing.T) {
	testErr := &testError{reason: "test error"}
	wrapErr := Wrap(testErr, -1000, "wrap test error")

	testCases := []struct {
		err   error
		cause error
	}{
		{
			err:   Wrap(testErr, 1000, "wow"),
			cause: testErr,
		},
		{
			err:   Wrap(wrapErr, 1000, "wow too"),
			cause: testErr,
		},
	}

	for _, testCase := range testCases {
		if ok := Is(testCase.err, testCase.cause); !ok {
			t.Errorf("testCase.err %+v isn't testCase.cause %+v", testCase.err, testCase.cause)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestAs$
func TestAs(t *testing.T) {
	testErr := &testError{reason: "test error"}
	wrapErr := Wrap(testErr, -1000, "wrap test error")

	targetTestErr := &testError{}
	targetTestErr2 := &testError{}
	targetPathErr := &os.PathError{}

	testCases := []struct {
		err    error
		target any
		ok     bool
		want   any
	}{
		{
			err:    Wrap(testErr, 1000, "wow"),
			target: &targetTestErr,
			ok:     true,
			want:   &testErr,
		},
		{
			err:    Wrap(wrapErr, 1000, "wow too"),
			target: &targetTestErr2,
			ok:     true,
			want:   &testErr,
		},
		{
			err:    New(1000, "no"),
			target: &targetPathErr,
			ok:     false,
			want:   nil,
		},
	}

	for _, testCase := range testCases {
		ok := As(testCase.err, testCase.target)
		if ok != testCase.ok {
			t.Errorf("err %+v ok %+v != err %+v testCase.ok %+v", testCase.err, testCase.target, ok, testCase.ok)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestUnwrap$
func TestUnwrap(t *testing.T) {
	testErr := &testError{reason: "test error"}
	wrapErr := Wrap(testErr, -1000, "wrap test error")

	testCases := []struct {
		err   error
		cause error
	}{
		{
			err:   testErr,
			cause: nil,
		},
		{
			err:   Wrap(testErr, 1000, "wow"),
			cause: testErr,
		},
		{
			err:   Wrap(wrapErr, 1000, "wow too"),
			cause: wrapErr,
		},
	}

	for _, testCase := range testCases {
		if Unwrap(testCase.err) != testCase.cause {
			t.Errorf("Unwrap(testCase.err) %+v != testCase.cause %+v", Unwrap(testCase.err), testCase.cause)
		}
	}
}
