package gonum

import "runtime"

// Get the name of the caller. The argument skip is the number of stack frames
// to ascend, with 0 identifying the caller of [CallerName].
func CallerName(skip int) string {
	if pc, _, _, ok := runtime.Caller(skip + 1); ok {
		return runtime.FuncForPC(pc).Name()
	} else {
		return ""
	}
}

// Assert panics the message msg if b is false
//go:inline
func Assert(b bool, msg any) {
	if !b {
		panic(msg)
	}
}
