package go_dash

func filter[T any](ary []T, pred func(T) bool) []T {
	n := make([]T, 0)
	for _, a := range ary {
		if pred(a) {
			n = append(n, a)
		}
	}
	return n
}
