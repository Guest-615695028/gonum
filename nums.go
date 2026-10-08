package gonum

import (
	"reflect"
	"strconv"
)

func Abs[T Real](t T) T {
	if t < 0 {
		return -t
	} else {
		return t + 0 //avoid -0
	}
}

func Sign[T Real](t T) int {
	if t < 0 {
		return -1
	} else if t > 0 {
		return 1
	} else {
		return 0
	}
}

func IsNan[T Number](t T) bool {
	return t != t
}
func IsInf[T Number](t T) bool {
	return t != 0 && t == t/2
}
func IsFinite[T Number](t T) bool {
	return t == t && !IsInf(t)
}

// Format number into string, parameters:
//   - t: the number
//   - f: the formatter character as for [fmt.Printf]()
//   - prec: the precision, or base
//   - pos: show plus sign or space
func FormatNumber[T Number](t T, f byte, prec int) (s string) {
	defer func() { recover() }()
	switch v := reflect.ValueOf(t); v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		s = strconv.FormatInt(v.Int(), base(f))
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		s = strconv.FormatUint(v.Uint(), base(f))
	case reflect.Float32:
		s = strconv.FormatFloat(v.Float(), floatfmt(f), prec, 32)
	case reflect.Float64:
		s = strconv.FormatFloat(v.Float(), floatfmt(f), prec, 64)
	case reflect.Complex64:
		s = strconv.FormatComplex(v.Complex(), floatfmt(f), prec, 64)
	case reflect.Complex128:
		s = strconv.FormatComplex(v.Complex(), floatfmt(f), prec, 128)
	default:
		s = v.String()
	}
	if f >= 'A' && f <= 'Z' {
		s = ToUpper(s)
	}
	return s
}

func base(f byte) int {
	switch f {
	case 0, 1, 'B', 'b':
		return 2
	case 'O', 'o':
		return 8
	case 'X', 'x':
		return 16
	default:
		return Cond(f > 36, 10, int(f))
	}
}

func floatfmt(f byte) byte {
	switch f {
	case 'b', 'e', 'E', 'f', 'g', 'G', 'x', 'X':
		return f
	case 'B', 'F':
		return f + 32
	default:
		return 'g'
	}
}

func CompareNumber[X, Y Real](x X, y Y) int {
	if x != x && y != y || x == 0 && y == 0 {
		return 0
	} else if x != x || x <= 0 && y >= 0 {
		return -1
	} else if y != y || x >= 0 && y <= 0 {
		return 1
	} else { // x and y of the same sign
		return Compare(float64(x), float64(y))<<1 +
			Compare(uint64(x), uint64(y))
	}
}

// Compare number, -1 if x<y, 0 if x==y, 1 if x>y
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
func Compare[T Ordered](x, y T) int {
	if x != x && y != y || x == y {
		return 0
	}
	if x != x || x < y {
		return -1
	}
	return 1
}

// Same as [slices.CompareFunc], but with length-first option - @param l1
func CompareSlicesFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](
	x S1, y S2, op func(E1, E2) int, l1 bool,
) int {
	c := Compare(len(x), len(y))
	if l1 && c != 0 {
		return c
	}
	for i := range min(len(x), len(y)) {
		if c1 := op(x[i], y[i]); c1 != 0 {
			return c1
		}
	}
	return c
}

func CompareSlices[S ~[]E, E Ordered](x, y S) int {
	return CompareSlicesFunc(x, y, Compare, false)
}
func CompareSlicesL1[S ~[]E, E Ordered](x, y S) int {
	return CompareSlicesFunc(x, y, Compare, true)
}

func ConvertFromReal[X Number, Y Real](y Y) (x X) {
	v := reflect.ValueOf(&x).Elem()
	defer Recover(&v)
	if v.CanComplex() {
		v.SetComplex(complex(float64(y), 0))
	} else if v.CanFloat() {
		v.SetFloat(float64(y))
	} else if v.CanInt() {
		v.SetInt(int64(y))
	} else if v.CanUint() {
		v.SetUint(uint64(y))
	}
	return
}

func Add[T Number](x, y T) T       { return x + y }
func Sub[T Number](x, y T) T       { return x - y }
func Mul[T Number](x, y T) T       { return x * y }
func Div[T Number](x, y T) T       { return x / y }
func Mod[T Integer](x, y T) T      { return x % y }
func BitAnd[T Integer](x, y T) T   { return x & y }
func BitOr[T Integer](x, y T) T    { return x | y }
func BitXor[T Integer](x, y T) T   { return x ^ y }
func BitClr[T Integer](x, y T) T   { return x &^ y }
func ShL[X, Y Integer](x X, y Y) X { return x << y }
func ShR[X, Y Integer](x X, y Y) X { return x >> y }

func Iota[T Integer](x T) (i []T) {
	i = make([]T, x)
	for j := range i {
		i[j] = T(j)
	}
	return i
}
