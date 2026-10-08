package gonum

// analogous to [maps.Clone], but with simpler assignment
func CloneMap[M ~map[K]V, K comparable, V any](m M) M {
	if m == nil {
		return nil
	}
	n := make(M, len(m))
	CopyMap(n, m)
	return n
}

//Copy a map into another, and return the number of elements duplicated
func CopyMap[M1, M2 ~map[K]V, K comparable, V any](m1 M1, m2 M2) (n int) {
	for k, v := range m2 {
		if _, ok := m1[k]; ok {
			n++
		}
		m1[k] = v
	}
	return n
}

//Copy a map into another, and return the number of elements added
func UnionMap[M1, M2 ~map[K]V, K comparable, V any](m1 M1, m2 M2) (n int) {
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			n++
			m1[k] = v
		}
	}
	return n
}

//Subtract a map from another, and return the number of elements deleted
func SubtractMap[M1, M2 ~map[K]V, K comparable, V any](m1 M1, m2 M2) (n int) {
	for k := range m2 {
		if _, ok := m1[k]; ok {
			n++
			delete(m1, k)
		}
	}
	return n
}
