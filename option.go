// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

type Option func(e *Error)

func (o Option) Apply(e *Error) {
	o(e)
}

func applyOptions(e *Error, opts []Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

// WithMsg sets msg to e.
func WithMsg(msg string) Option {
	return func(e *Error) {
		e.msg = msg
	}
}
