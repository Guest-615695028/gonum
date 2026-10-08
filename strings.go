package gonum

func RepeatRune(r rune, n int) string {
	if n <= 0 {
		return ""
	} else if n == 1 {
		return string(r)
	}
	s := make([]rune, n)
	for i := range n {
		s[i] = r
	}
	return string(s)
}
func RepeatByte(b byte, n int) string {
	if n <= 0 {
		return ""
	} else if n == 1 {
		return string(b)
	}
	s := make([]byte, n)
	for i := range n {
		s[i] = b
	}
	return string(s)
}

func ToLower(s string) string {
	b := []byte(s)
	for i, v := range b {
		if v >= 'A' && v <= 'Z' {
			b[i] += 32
		}
	}
	return string(b)
}

func ToUpper(s string) string {
	b := []byte(s)
	for i, v := range b {
		if v >= 'a' && v <= 'z' {
			b[i] -= 32
		}
	}
	return string(b)
}
