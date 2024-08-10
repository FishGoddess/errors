// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBadRequest$
func TestBadRequest(t *testing.T) {
	testCases := []struct {
		message     string
		args        []any
		wantCode    int32
		wantMessage string
	}{
		{
			message:     "xxx",
			args:        nil,
			wantCode:    codeBadRequest,
			wantMessage: "xxx",
		},
		{
			message:     "xxx %d%.2f",
			args:        nil,
			wantCode:    codeBadRequest,
			wantMessage: "xxx %d%.2f",
		},
		{
			message:     "xxx %d%s%+v",
			args:        []any{1, ".", true},
			wantCode:    codeBadRequest,
			wantMessage: "xxx 1.true",
		},
	}

	for _, testCase := range testCases {
		err := BadRequest(testCase.message, testCase.args...)
		if err.Code() != testCase.wantCode {
			t.Errorf("err.Code() %d != testCase.wantCode %d", err.Code(), testCase.wantCode)
		}

		if err.Message() != testCase.wantMessage {
			t.Errorf("err.Message() %s != testCase.wantMessage %s", err.Message(), testCase.wantMessage)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestForbidden$
func TestForbidden(t *testing.T) {
	testCases := []struct {
		message     string
		args        []any
		wantCode    int32
		wantMessage string
	}{
		{
			message:     "xxx",
			args:        nil,
			wantCode:    codeForbidden,
			wantMessage: "xxx",
		},
		{
			message:     "xxx %d%.2f",
			args:        nil,
			wantCode:    codeForbidden,
			wantMessage: "xxx %d%.2f",
		},
		{
			message:     "xxx %d%s%+v",
			args:        []any{1, ".", true},
			wantCode:    codeForbidden,
			wantMessage: "xxx 1.true",
		},
	}

	for _, testCase := range testCases {
		err := Forbidden(testCase.message, testCase.args...)
		if err.Code() != testCase.wantCode {
			t.Errorf("err.Code() %d != testCase.wantCode %d", err.Code(), testCase.wantCode)
		}

		if err.Message() != testCase.wantMessage {
			t.Errorf("err.Message() %s != testCase.wantMessage %s", err.Message(), testCase.wantMessage)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNotFound$
func TestNotFound(t *testing.T) {
	testCases := []struct {
		message     string
		args        []any
		wantCode    int32
		wantMessage string
	}{
		{
			message:     "xxx",
			args:        nil,
			wantCode:    codeNotFound,
			wantMessage: "xxx",
		},
		{
			message:     "xxx %d%.2f",
			args:        nil,
			wantCode:    codeNotFound,
			wantMessage: "xxx %d%.2f",
		},
		{
			message:     "xxx %d%s%+v",
			args:        []any{1, ".", true},
			wantCode:    codeNotFound,
			wantMessage: "xxx 1.true",
		},
	}

	for _, testCase := range testCases {
		err := NotFound(testCase.message, testCase.args...)
		if err.Code() != testCase.wantCode {
			t.Errorf("err.Code() %d != testCase.wantCode %d", err.Code(), testCase.wantCode)
		}

		if err.Message() != testCase.wantMessage {
			t.Errorf("err.Message() %s != testCase.wantMessage %s", err.Message(), testCase.wantMessage)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestRequireLogin$
func TestRequireLogin(t *testing.T) {
	testCases := []struct {
		message     string
		args        []any
		wantCode    int32
		wantMessage string
	}{
		{
			message:     "xxx",
			args:        nil,
			wantCode:    codeRequireLogin,
			wantMessage: "xxx",
		},
		{
			message:     "xxx %d%.2f",
			args:        nil,
			wantCode:    codeRequireLogin,
			wantMessage: "xxx %d%.2f",
		},
		{
			message:     "xxx %d%s%+v",
			args:        []any{1, ".", true},
			wantCode:    codeRequireLogin,
			wantMessage: "xxx 1.true",
		},
	}

	for _, testCase := range testCases {
		err := RequireLogin(testCase.message, testCase.args...)
		if err.Code() != testCase.wantCode {
			t.Errorf("err.Code() %d != testCase.wantCode %d", err.Code(), testCase.wantCode)
		}

		if err.Message() != testCase.wantMessage {
			t.Errorf("err.Message() %s != testCase.wantMessage %s", err.Message(), testCase.wantMessage)
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
