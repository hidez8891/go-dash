package go_dash

func every[T any](ary []T, pred func(T) bool) bool {
	for _, a := range ary {
		if !pred(a) {
			return false
		}
	}
	return true
}

func filter[T any](ary []T, pred func(T) bool) []T {
	n := make([]T, 0)
	for _, a := range ary {
		if pred(a) {
			n = append(n, a)
		}
	}
	return n
}

func mapBy[T, U any](ary []T, conv func(T) U) []U {
	n := make([]U, len(ary), cap(ary))
	for i, a := range ary {
		n[i] = conv(a)
	}
	return n
}

func some[T any](ary []T, pred func(T) bool) bool {
	for _, a := range ary {
		if pred(a) {
			return true
		}
	}
	return false
}
