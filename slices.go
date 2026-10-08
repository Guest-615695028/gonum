package gonum

import (
	"reflect"
	"slices"
)

func ValuesOf(a ...any) (vs []reflect.Value) {
	vs = make([]reflect.Value, len(a))
	for i, v := range a {
		vs[i] = reflect.ValueOf(v)
	}
	return vs
}

// Fill generates a slice of n-times x.
func Fill[E any](x E, n int) (s []E) {
	if n <= 0 {
		return []E{}
	}
	s = make([]E, n)
	for n--; n >= 0; n-- {
		s[n] = x
	}
	return
}

// FillPointer generate a slice of n-times pointer to values equal to x,
// isolated to each other.
func FillPointer[E any](x E, n int) (s []*E) {
	s = make([]*E, n)
	for n--; n >= 0; n-- {
		s[n] = new(E)
		*s[n] = x
	}
	return
}

func Combine[T any](t ...T) []T { return t }

func CloneSlice[S ~[]E, E any](s S) S {
	r := make(S, len(s))
	copy(r, s)
	return r
}

// Step a slice by n into independent storage
func Step[S ~[]E, E any](s S, n int) S {
	if n <= 0 || len(s) <= 0 {
		return S{}
	} else if len(s) <= n {
		return S{s[0]}
	} else if n == 1 {
		return append(S{}, s...)
	}
	r := make(S, (len(s)-1)/n+1)
	for i := range r {
		r[i] = s[i*n]
	}
	return r
}

// Make a reversed copy
func Reverse[S ~[]E, E any](s S) S {
	c := CloneSlice(s)
	slices.Reverse(c)
	return c
}

// Convert a slice of values to another type of slice
func ConvertSlice[R ~[]E1, S ~[]E2, E1, E2 any](s S) R {
	r, tr := make(R, len(s)), reflect.TypeFor[R]()
	for i, v := range s {
		if vv := reflect.ValueOf(v); vv.Type().ConvertibleTo(tr) {
			r[i] = vv.Convert(tr).Interface().(E1)
		}
	}
	return r
}

// Convert a slice of real numbers to another type of slice
func ConvertSliceFunc[S ~[]E, E, R any](s S, f func(E) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Connect slices
func Connect[S ~[]E, E any](ss ...S) (s S) {
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// Crop a slice into slices of equal length
func Crop[S ~[]E, E any](s S, n int) (r []S) {
	m := len(s) / n
	r = make([]S, m)
	for i := range m {
		c := n * (i + 1)
		r[i] = s[n*i : c : c]
	}
	return r
}

// Safe algorithm to resize a slice (s) to its desired length (n),
// different from standard library [slices.Grow].
func Resize[S ~[]E, E any](s S, n int) (r S) {
	if m := n - len(s); n > 0 {
		r = append(s, make(S, m)...)
	} else {
		r = s[:n]
	}
	return r
}

// N-ary operator without starting value specially defined because
// Reduce(s, op) may behave different to [ReduceFrom](z, s, op), where
// 'z' is the zero value of type E.
func Reduce[S ~[]E, E any](s S, op func(E, E) E) (r E) {
	if len(s) <= 0 {
		return
	}
	r = s[0]
	for _, v := range s[1:] {
		r = op(r, v)
	}
	return r
}

// N-ary operator with starting value
func ReduceFrom[S ~[]E, E, R any](r R, s S, op func(R, E) R) R {
	if len(s) <= 0 {
		return r
	}
	for _, v := range s[1:] {
		r = op(r, v)
	}
	return r
}

func Merge[X ~[]A, Y ~[]B, A, B, C any](x X, y Y, op func(A, B) C) []C {
	r := make([]C, min(len(x), len(y)))
	for i := range r {
		r[i] = op(x[i], y[i])
	}
	return r
}

func All[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func Any[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func None[S ~[]E, E any](s S, f func(E) bool) bool {
	return !Any(s, f)
}

func Count[S ~[]E, E any](s S, f func(E) bool) (n int) {
	for _, e := range s {
		if f(e) {
			n++
		}
	}
	return n
}
