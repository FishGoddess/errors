// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/errors/status"
)

var (
	errAuthFailed = errors.New("auth failed")
)

func isAuthFailed(err error) bool {
	return err == errAuthFailed
}

func main() {
	// We provide a way to transfer errors to status.
	// A status can be used by a server like grpc server with google.status.
	// For example, we have an auth failed error which will be returned by service.
	// However, we usually return a status code represents it in server,
	// and we usually return a human-being readable msg instead of the error msg.
	// So it has a gap between service error and server status.
	// You can register a status with server code and msg, then use Parse to restore them.
	authFailedStatus := status.New(1000, "you should check auth", isAuthFailed)
	status.RegisterStatus(authFailedStatus)

	// Get code and msg from error.
	code, msg := status.Parse(errAuthFailed)
	fmt.Println(code, msg)

	// Of course, you can use errors.Wrap to get an error with msg.
	// Then register a status about it.
	errCode := int32(123456)
	err := errors.Wrap(io.EOF, errCode, errors.WithMsg("hello"))

	isError := func(err error) bool {
		return errors.Is(err, errCode)
	}

	// As you can see, the error code and status code don't need to be the same.
	statusCode := int32(654321)
	status.RegisterStatus(status.New(statusCode, "i am status", isError))

	// The msg will be the one set to error using WithMsg.
	// The registered msg would be used only if the error doesn't have a set msg.
	// Try to remove errors.WithMsg("hello") above and run again.
	code, msg = status.Parse(err)
	fmt.Println(code, msg)
}
