package maps

func countBy[T any, K, U comparable](m map[K]T, f func(K, T) U) map[U]int {
	group := map[U]int{}
	for k, v := range m {
		group[f(k, v)] += 1
	}
	return group
}

func every[T any, K comparable](m map[K]T, pred func(K, T) bool) bool {
	for k, v := range m {
		if !pred(k, v) {
			return false
		}
	}
	return true
}

func filter[T any, K comparable](m map[K]T, pred func(K, T) bool) map[K]T {
	n := map[K]T{}
	for k, v := range m {
		if pred(k, v) {
			n[k] = v
		}
	}
	return n
}

func find[T any, K comparable](m map[K]T, pred func(K, T) bool) (K, T, bool) {
	for k, v := range m {
		if pred(k, v) {
			return k, v, true
		}
	}

	var k K
	var t T
	return k, t, false
}

func forEach[T any, K comparable](m map[K]T, f func(K, T)) {
	for k, v := range m {
		f(k, v)
	}
}

func groupBy[T any, K, U comparable](m map[K]T, f func(K, T) U) map[U]map[K]T {
	group := map[U]map[K]T{}
	for k, v := range m {
		key := f(k, v)
		if _, ok := group[key]; !ok {
			group[key] = map[K]T{}
		}
		group[key][k] = v
	}
	return group
}

func includes[T, K comparable](m map[K]T, t T) bool {
	for _, v := range m {
		if v == t {
			return true
		}
	}
	return false
}

func mapBy[T, U any, K comparable](m map[K]T, conv func(K, T) U) map[K]U {
	n := map[K]U{}
	for k, v := range m {
		n[k] = conv(k, v)
	}
	return n
}

func partition[T any, K, U comparable](m map[K]T, f func(K, T) U) []map[K]T {
	parts := make([]map[K]T, 0)
	group := groupBy(m, f)
	for _, v := range group {
		parts = append(parts, v)
	}
	return parts
}

func reduce[T, U any, K comparable](m map[K]T, f func(K, T, U) U, acc U) U {
	for k, v := range m {
		acc = f(k, v, acc)
	}
	return acc
}

func some[T any, K comparable](m map[K]T, pred func(K, T) bool) bool {
	for k, v := range m {
		if pred(k, v) {
			return true
		}
	}
	return false
}

func reject[T any, K comparable](m map[K]T, pred func(K, T) bool) map[K]T {
	ret := map[K]T{}
	for k, v := range m {
		if !pred(k, v) {
			ret[k] = v
		}
	}
	return ret
}
