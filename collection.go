package go_dash

func countBy[T any, U comparable](ary []T, f func(T) U) map[U]int {
	group := map[U]int{}
	for _, a := range ary {
		group[f(a)] += 1
	}
	return group

}

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

func find[T any](ary []T, pred func(T) bool) (T, bool) {
	for _, a := range ary {
		if pred(a) {
			return a, true
		}
	}

	var v T
	return v, false
}

func findLast[T any](ary []T, pred func(T) bool) (T, bool) {
	for i := len(ary); i > 0; i-- {
		a := ary[i-1]
		if pred(a) {
			return a, true
		}
	}

	var v T
	return v, false
}

func forEach[T any](ary []T, f func(T)) {
	for _, a := range ary {
		f(a)
	}
}

func forEachRight[T any](ary []T, f func(T)) {
	for i := len(ary); i > 0; i-- {
		f(ary[i-1])
	}
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
