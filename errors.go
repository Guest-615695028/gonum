package gonum

type ErrorString string

func (s ErrorString) Error() string {
	return string(s)
}

func Recover[T any](t *T) {
	if a, ok := recover().(T); ok && t != nil {
		*t = a
		return
	}
}
