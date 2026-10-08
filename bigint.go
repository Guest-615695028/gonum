package gonum

import (
	"encoding/binary"
	"slices"
)

//The type of arbitary-length signed integers, different to [math/big.BigInt]
//The lowest bit is the sign bit.
type BigInt []byte

// Construct an instance of an arbitary-length signed integer.
// This is different from standard [encoding/binary.Varint] series.
func NewInt(x int64) BigInt {
	i := make(BigInt, 8)
	for j := range 8 {
		i[j] = byte(x)
		x >>= 8
	}
	return i
}

//Crop unneeded bytes into a new slice
func (i BigInt) Crop() BigInt {
	j := len(i)
	for j > 1 && i[j-1] == 0 {
		j--
	}
	k := make(BigInt, j)
	copy(k, i)
	return k
}

func (i BigInt) BigInt() (x int64) {
	x = -int64(i[0]&1) ^ int64(i[0]>>1) // -1 or 0
	for k := range min(6, len(i)-1) {
		x ^= int64(i[k+1]) << (8*k + 7)
	}
	return x
}

func (i BigInt) AppendBinary(b []byte) (data []byte, err error) {
	data, err = i.MarshalBinary()
	return append(b, data...), err
}

func (i BigInt) MarshalBinary() (data []byte, err error) {
	var b [8]byte
	i = i.Crop()
	if len(i) == 1 && i[0] < 0x80 {
		return []byte{i[0]}, nil // deep copy
	}
	data = make([]byte, (len(i)*7+7)/8)
	m := 0
	for k := 0; k < len(i); k += 7 {
		j := copy(b[:7], i)
		x := uint64(b[0])
		for kj := 1; kj < j; kj++ {
			x |= uint64(b[kj]) << (kj * 7)
		}
		for j = 0; x > 0; x >>= 7 {
			b[j] = byte(x | 0x80)
		}
		m += copy(data[m:m+8], b[:])
	}
	for data[m-1] &^= 0x80; m >= 2 && data[m-1] == 0; m-- {
		data[m-2] &^= 0x80
	}
	return data[:m], nil
}

func (i *BigInt) UnmarshalBinary(data []byte) error {
	n := len(data)
	for k, v := range data {
		if v <= 0x7f {
			n = k
			break
		}
	}
	if n == len(data) {
		*i = nil
		binary.Uvarint(nil)
		return ErrorString("gonum: BigInt: data not ending")
	}
	var b [8]byte
	*i = make(BigInt, 0, n)
	for k := 0; k <= n; k += 8 {
		_ = copy(b[:], data[k:])
		for i := range 7 {
			if b[i]&0x80 == 0 {
				break
			}
			b[i] &= 0x7f

		}
		b[0] &= 0x7f
		b[0] |= b[1] & 1
		b[1] >>= 1
		b[1] &= 0x3f
		b[1] |= b[1] & 1
		b[1] >>= 1
	}
	*i = i.Crop()
	return nil
}

func (i BigInt) ByteAt(p int) byte {
	if p >= len(i) || p < 0 {
		return 0
	}
	return i[p]
}

func (i BigInt) Equal(j BigInt) bool {
	for k := range max(len(i), len(j)) {
		if i.ByteAt(k) != j.ByteAt(k) {
			return false
		}
	}
	return true
}

func (i BigInt) Compare(j BigInt) int {
	if (i[0]^j[0])&1 == 1 { //different sign
		return int(i[0]&1) - int(j[0]-1)
	}
	s := int(i[0]&j[0]&1<<1) - 1 // 1 or -1
	return s * slices.Compare(i, j)
}

func CompareInt(i, j BigInt) int {
	return i.Compare(j)
}

func (i BigInt) Not() BigInt {
	i = i.Crop()
	i[0] ^= 1
	return i
}

func (i BigInt) Neg() BigInt {
	return i.Not().Inc()
}

// Increment operation: i++
func (i BigInt) Inc() BigInt {
	i = i.Crop()
	switch i[0] {
	case 0xfe:
		i[0] = 0
	case 0x01:
		if len(i) == 1 {
			i[0] = 0
			return i
		}
		i[0] = 0xff
	default:
		i[0] += i[0]&1<<2 - 2
		return i
	}
	for k, v := range i[1:] {
		i[k+1] += i[0] | 1
		if i[0]+v != 255 {
			return i
		}
	}
	if i[len(i)-1] == 0 {
		i = append(i, 1)
	}
	return i
}

func (i BigInt) Add(j BigInt) BigInt {
	return add(i, j)
}

// used in the case (i[0]^j[0])&1 == 0
func add(i, j BigInt) BigInt {
	r := make(BigInt, max(len(i), len(j)))
	ij := int(i[0]) + int(j[0]) + int(i[0]&j[0]&1)
	for k := range r {
		ij += int(i.ByteAt(k)) + int(j.ByteAt(k))
		r[k] = byte(ij)
		ij >>= 8
	}
	if ij > 0 {
		r = append(r, byte(ij))
	}
	return r
}

func (i BigInt) Sub(j BigInt) (r BigInt) {
	return add(i, j.Neg())
}
