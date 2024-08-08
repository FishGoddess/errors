// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"runtime"
	"strconv"
)

func Caller() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	return file + ":" + strconv.Itoa(line)
}

func Callers() []string {
	var pcs [16]uintptr
	n := runtime.Callers(2, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	var callers []string
	for {
		frame, more := frames.Next()

		caller := frame.File + ":" + strconv.Itoa(frame.Line)
		callers = append(callers, caller)

		if !more {
			break
		}
	}

	return callers
}
