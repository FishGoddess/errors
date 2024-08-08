// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"testing"
)

// BenchmarkCaller-2   	 1390318	       876.7 ns/op	     296 B/op	       3 allocs/op
func BenchmarkCaller(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Caller()
	}
}

// BenchmarkCallers-2   	  443419	      2573 ns/op	     732 B/op	      12 allocs/op
func BenchmarkCallers(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Callers()
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestCaller$
func TestCaller(t *testing.T) {
	caller := Caller()
	t.Log(caller)
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestCallers$
func TestCallers(t *testing.T) {
	callers := Callers()
	t.Log(callers)
}
