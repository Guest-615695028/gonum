package gonum

import (
	"reflect"
)

// Analogous to C/C++ conditional operator
//
//	r = c ? t : f
//
// but both sides are evaluated
func Cond[T any](b bool, t, f T) (r T) {
	if b {
		return t
	} else {
		return f
	}
}

// Ignore the returned error
func NoError[T any](t T, err error) T { return t }

// Must OK
func OK[T any](t T, ok bool) T { return t }

// (C/C++/JavaScript)-style boolean
func Bool(a any) bool {
	switch b := a.(type) {
	case nil:
		return false
	case bool:
		return b
	case string:
		return b != ""
	}
	defer func() { recover() }() //Avoid potential panic
	v := reflect.ValueOf(a)
	return v.IsValid() && !v.IsZero()
}
