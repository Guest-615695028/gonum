package gonum

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Integer interface{ Int | UInt }
type Float interface{ ~float32 | ~float64 }
type Complex interface{ ~complex64 | ~complex128 }
type Real interface{ Integer | Float }
type Number interface{ Real | Complex }

// Ordered is a constraint that permits any ordered type:
// any type that supports the operators < <= >= >
type Ordered interface{ Real | ~string }

type Iterator2[K, V any] interface{ Iter(yield func(K, V) bool) }
type Iterator1[K any] interface{ Iter(yield func(K) bool) }
type Iterator0 interface{ Iter(yield func() bool) }
