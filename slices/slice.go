package slices

func chunk[T any](ary []T, size int) [][]T {
	if len(ary) == 0 || size < 1 || len(ary) <= size {
		return [][]T{ary}
	}

	out := make([][]T, 0)
	l := len(ary)
	for i := 0; i < l; i += size {
		var v []T
		if l-i < size {
			v = make([]T, l-i)
		} else {
			v = make([]T, size)
		}

		copy(v, ary[i:i+len(v)])
		out = append(out, v)
	}

	return out
}

func concat[T any](a1 []T, a2 []T) []T {
	n := make([]T, len(a1)+len(a2))
	copy(n[0:], a1)
	copy(n[len(a1):], a2)
	return n
}

func difference[T comparable](a1 []T, a2 []T) []T {
	n := make([]T, 0)
	for _, v := range a1 {
		if !includes(a2, v) {
			n = append(n, v)
		}
	}
	return n
}

func differenceBy[T, U comparable](a1 []T, a2 []T, f func(T) U) []T {
	b2 := make([]U, len(a2))
	for i, v := range a2 {
		b2[i] = f(v)
	}

	n := make([]T, 0)
	for _, v := range a1 {
		if !includes(b2, f(v)) {
			n = append(n, v)
		}
	}
	return n
}

func differenceWith[T comparable](a1 []T, a2 []T, pred func(T, T) bool) []T {
	n := make([]T, 0)
	for _, v1 := range a1 {
		f := func(v2 T) bool { return pred(v1, v2) }
		if !some(a2, f) {
			n = append(n, v1)
		}
	}
	return n
}

func drop[T any](ary []T, size int) []T {
	if len(ary) == 0 || size < 1 {
		return ary
	}
	if len(ary) <= size {
		return []T{}
	}

	out := make([]T, len(ary)-size)
	copy(out, ary[size:])
	return out
}

func dropRight[T any](ary []T, size int) []T {
	if len(ary) == 0 || size < 1 {
		return ary
	}
	if len(ary) <= size {
		return []T{}
	}

	out := make([]T, len(ary)-size)
	copy(out, ary[:len(ary)-size])
	return out
}

func dropRightWhile[T any](ary []T, pred func(T) bool) []T {
	index := 0
	for i := len(ary); i > 0; i-- {
		if !pred(ary[i-1]) {
			index = i
			break
		}
	}

	if index == 0 {
		return []T{}
	}

	out := make([]T, index)
	copy(out, ary[:index])
	return out
}

func dropWhile[T any](ary []T, pred func(T) bool) []T {
	index := len(ary)
	for i, v := range ary {
		if !pred(v) {
			index = i
			break
		}
	}

	if index == len(ary) {
		return []T{}
	}

	out := make([]T, len(ary)-index)
	copy(out, ary[index:])
	return out
}

func fill[T any](ary []T, v T) []T {
	for i, _ := range ary {
		ary[i] = v
	}
	return ary
}

func findIndex[T any](ary []T, pred func(T) bool) int {
	for i, a := range ary {
		if pred(a) {
			return i
		}
	}
	return -1
}

func findLastIndex[T any](ary []T, pred func(T) bool) int {
	for i := len(ary); i > 0; i-- {
		if pred(ary[i-1]) {
			return i - 1
		}
	}
	return -1
}

func flatten[T any](ary [][]T) []T {
	size := 0
	for _, v := range ary {
		size += len(v)
	}

	out := make([]T, size)
	index := 0
	for _, v := range ary {
		copy(out[index:], v)
		index += len(v)
	}

	return out
}

func intersection[T comparable](a1 []T, a2 []T) []T {
	n := make([]T, 0)
	for _, v := range a1 {
		if includes(a2, v) {
			n = append(n, v)
		}
	}
	return n
}

func intersectionBy[T, U comparable](a1 []T, a2 []T, f func(T) U) []T {
	b2 := make([]U, len(a2))
	for i, v := range a2 {
		b2[i] = f(v)
	}

	n := make([]T, 0)
	for _, v := range a1 {
		if includes(b2, f(v)) {
			n = append(n, v)
		}
	}
	return n
}

func intersectionWith[T comparable](a1 []T, a2 []T, pred func(T, T) bool) []T {
	n := make([]T, 0)
	for _, v1 := range a1 {
		f := func(v2 T) bool { return pred(v1, v2) }
		if some(a2, f) {
			n = append(n, v1)
		}
	}
	return n
}

func lastIndexOf[T comparable](ary []T, v T) int {
	for i := len(ary); i > 0; i-- {
		if ary[i-1] == v {
			return i - 1
		}
	}
	return -1
}

func indexOf[T comparable](ary []T, v T) int {
	for i, a := range ary {
		if a == v {
			return i
		}
	}
	return -1
}

func pull[T comparable](ary []T, values ...T) []T {
	out := make([]T, 0)
	for _, v := range ary {
		if !includes(values, v) {
			out = append(out, v)
		}
	}
	return out
}

func pullAll[T comparable](ary []T, values []T) []T {
	out := make([]T, 0)
	for _, v := range ary {
		if !includes(values, v) {
			out = append(out, v)
		}
	}
	return out
}
