// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package status

import (
	"net/http"
	"sync"

	"github.com/FishGoddess/errors"
)

var (
	CodeOK      int32 = 0
	CodeUnknown int32 = http.StatusInternalServerError
)

var (
	// MsgUnknown is the default msg returned from Parse if err passed to Parse doesn't have a msg.
	MsgUnknown = "Server is too busy, please try again later."
)

var (
	statuses     = make([]*Status, 0, 8)
	statusesLock sync.RWMutex
)

// Status wraps code and msg for returning.
type Status struct {
	code    int32
	msg     string
	isError func(err error) bool
}

// New returns a new with given params.
// isError returns if err belongs to status.
func New(code int32, msg string, isError func(err error) bool) *Status {
	return &Status{
		code:    code,
		msg:     msg,
		isError: isError,
	}
}

func (s *Status) Code() int32 {
	if s == nil {
		return CodeOK
	}

	return s.code
}

func (s *Status) Msg() string {
	if s == nil {
		return ""
	}

	return s.msg
}

func (s *Status) IsError(err error) bool {
	return s.isError != nil && s.isError(err)
}

// RegisterStatus registers a new status to local statuses so that we can use it in Parse.
func RegisterStatus(status *Status) {
	statusesLock.Lock()
	statuses = append(statuses, status)
	statusesLock.Unlock()
}

// RegisterStatuses registers some statuses to local statuses so that we can use it in Parse.
func RegisterStatuses(registerStatus ...*Status) {
	statusesLock.Lock()
	defer statusesLock.Unlock()

	for _, status := range registerStatus {
		statuses = append(statuses, status)
	}
}

// Parse parses err and returns its code and error created from msg.
func Parse(err error) (int32, string) {
	return ParseOrDefault(err, CodeUnknown, MsgUnknown)
}

// ParseOrDefault parses err and returns its code and error created from msg.
func ParseOrDefault(err error, defaultCode int32, defaultMsg string) (int32, string) {
	if err == nil {
		return CodeOK, ""
	}

	statusesLock.RLock()
	defer statusesLock.RUnlock()

	for _, status := range statuses {
		if status != nil && status.IsError(err) {
			code := status.Code()
			msg := errors.MsgOrDefault(err, status.Msg())
			return code, msg
		}
	}

	msg := errors.MsgOrDefault(err, defaultMsg)
	return defaultCode, msg
}
