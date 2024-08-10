// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBadRequest$
func TestBadRequest(t *testing.T) {
	testCases := []struct {
		message string
		code    int32
	}{
		{
			message: "xxx",
			code:    codeBadRequest,
		},
	}

	for _, testCase := range testCases {
		err := BadRequest(testCase.message)
		if err.Code() != testCase.code {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestForbidden$
func TestForbidden(t *testing.T) {
	testCases := []struct {
		message string
		code    int32
	}{
		{
			message: "xxx",
			code:    codeForbidden,
		},
	}

	for _, testCase := range testCases {
		err := Forbidden(testCase.message)
		if err.Code() != testCase.code {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNotFound$
func TestNotFound(t *testing.T) {
	testCases := []struct {
		message string
		code    int32
	}{
		{
			message: "xxx",
			code:    codeNotFound,
		},
	}

	for _, testCase := range testCases {
		err := NotFound(testCase.message)
		if err.Code() != testCase.code {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestRequireLogin$
func TestRequireLogin(t *testing.T) {
	testCases := []struct {
		message string
		code    int32
	}{
		{
			message: "xxx",
			code:    codeRequireLogin,
		},
	}

	for _, testCase := range testCases {
		err := RequireLogin(testCase.message)
		if err.Code() != testCase.code {
			t.Errorf("err.Code() %d != testCase.code %d", err.Code(), testCase.code)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestMatchBadRequest$
func TestMatchBadRequest(t *testing.T) {
	testCases := []struct {
		message string
	}{
		{
			message: "xxx",
		},
	}

	for _, testCase := range testCases {
		err := BadRequest(testCase.message)
		if !MatchBadRequest(err) {
			t.Errorf("err %+v not match code %d", err, err.Code())
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestMatchForbidden$
func TestMatchForbidden(t *testing.T) {
	testCases := []struct {
		message string
	}{
		{
			message: "xxx",
		},
	}

	for _, testCase := range testCases {
		err := Forbidden(testCase.message)
		if !MatchForbidden(err) {
			t.Errorf("err %+v not match code %d", err, err.Code())
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestMatchNotFound$
func TestMatchNotFound(t *testing.T) {
	testCases := []struct {
		message string
	}{
		{
			message: "xxx",
		},
	}

	for _, testCase := range testCases {
		err := NotFound(testCase.message)
		if !MatchNotFound(err) {
			t.Errorf("err %+v not match code %d", err, err.Code())
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestMatchRequireLogin$
func TestMatchRequireLogin(t *testing.T) {
	testCases := []struct {
		message string
	}{
		{
			message: "xxx",
		},
	}

	for _, testCase := range testCases {
		err := RequireLogin(testCase.message)
		if !MatchRequireLogin(err) {
			t.Errorf("err %+v not match code %d", err, err.Code())
		}
	}
}
