package gonum

import "unsafe"

// Bit size of integer, either 32 or 64
const IntSize = 32 << (^uint(0) >> 63)

func RawBytes[S ~[]E, E any](s S) []byte {
	if len(s) <= 0 {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])),
		unsafe.Sizeof(s[0])*uintptr(len(s)))
}

func FromRawBytes[E any](b []byte) []E {
	if len(b) <= 0 {
		return nil
	}
	var e E
	return unsafe.Slice((*E)(unsafe.Pointer(&b[0])),
		uintptr(len(b))/unsafe.Sizeof(e))
}
