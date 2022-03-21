package slices

import (
	"strings"

	"golang.org/x/exp/constraints"
)

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

func differenceBy[T any, U comparable](a1 []T, a2 []T, f func(T) U) []T {
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

func differenceWith[T any](a1 []T, a2 []T, pred func(T, T) bool) []T {
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

func fill[T any](ary []T, v T) {
	for i, _ := range ary {
		ary[i] = v
	}
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

func head[T any](ary []T) T {
	if len(ary) == 0 {
		panic("slice is empty (slice head is out of range)")
	}
	return ary[0]
}

func indexOf[T comparable](ary []T, v T) int {
	for i, a := range ary {
		if a == v {
			return i
		}
	}
	return -1
}

func initial[T any](ary []T) []T {
	if len(ary) == 0 {
		panic("slice is empty (slice initial is out of range)")
	}
	return ary[:len(ary)-1]
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

func join(ary []string, separator string) string {
	return strings.Join(ary, separator)
}

func last[T any](ary []T) T {
	if len(ary) == 0 {
		panic("slice is empty (slice last is out of range)")
	}
	return ary[len(ary)-1]
}

func lastIndexOf[T comparable](ary []T, v T) int {
	for i := len(ary); i > 0; i-- {
		if ary[i-1] == v {
			return i - 1
		}
	}
	return -1
}

func nth[T any](ary []T, index int) T {
	if index < 0 {
		index = len(ary) + index
	}
	if index < 0 || len(ary) <= index {
		panic("slice index is out of range")
	}

	return ary[index]
}

func reverse[T any](ary []T) {
	limit := len(ary) / 2
	for i := 0; i < limit; i++ {
		j := len(ary) - i - 1
		ary[i], ary[j] = ary[j], ary[i]
	}
}

func slice[T any](ary []T, ranges ...int) []T {
	start := 0
	if len(ranges) >= 1 {
		start = ranges[0]
	}

	end := len(ary)
	if len(ranges) >= 2 {
		end = ranges[1]
	}

	return ary[start:end]
}

func sortedIndex[T constraints.Ordered](ary []T, value T) int {
	identity := func(v T) T { return v }
	return sortedIndexBy(ary, value, identity)
}

func sortedIndexBy[T any, U constraints.Ordered](ary []T, value T, conv func(T) U) int {
	if len(ary) == 0 {
		return 0
	}

	u := conv(value)
	if len(ary) == 1 {
		if conv(ary[0]) >= u {
			return 0
		} else {
			return 1
		}
	}

	middle := len(ary) / 2
	if conv(ary[middle]) >= u {
		return sortedIndexBy(ary[:middle], value, conv)
	} else {
		return sortedIndexBy(ary[middle:], value, conv) + middle
	}
}

func sortedLastIndex[T constraints.Ordered](ary []T, value T) int {
	identity := func(v T) T { return v }
	return sortedLastIndexBy(ary, value, identity)
}

func sortedLastIndexBy[T any, U constraints.Ordered](ary []T, value T, conv func(T) U) int {
	if len(ary) == 0 {
		return 0
	}

	u := conv(value)
	if len(ary) == 1 {
		if conv(ary[0]) > u {
			return 0
		} else {
			return 1
		}
	}

	middle := len(ary) / 2
	if conv(ary[middle]) > u {
		return sortedLastIndexBy(ary[:middle], value, conv)
	} else {
		return sortedLastIndexBy(ary[middle:], value, conv) + middle
	}
}

func sortedUniq[T comparable](ary []T) []T {
	identity := func(v T) T { return v }
	return sortedUniqBy(ary, identity)
}

func sortedUniqBy[T any, U comparable](ary []T, conv func(T) U) []T {
	if len(ary) == 0 {
		return ary
	}

	out := make([]T, 0)
	last := conv(ary[0])
	out = append(out, ary[0])

	for _, v := range ary[1:] {
		u := conv(v)
		if last != u {
			last = u
			out = append(out, v)
		}
	}
	return out
}

func tail[T any](ary []T) []T {
	return ary[1:]
}

func take[T any](ary []T, n int) []T {
	if len(ary) <= n {
		return ary
	}
	return ary[:n]
}

func takeRight[T any](ary []T, n int) []T {
	if len(ary) <= n {
		return ary
	}
	return ary[len(ary)-n:]
}

func takeRightWhile[T any](ary []T, pred func(T) bool) []T {
	for i := len(ary); i > 0; i-- {
		if !pred(ary[i-1]) {
			return ary[i:]
		}
	}
	return ary
}

func takeWhile[T any](ary []T, pred func(T) bool) []T {
	for i, v := range ary {
		if !pred(v) {
			return ary[:i]
		}
	}
	return ary
}

func union[T comparable](a1 []T, a2 []T) []T {
	n := make([]T, len(a1))
	copy(n, a1)

	for _, v := range a2 {
		if !includes(n, v) {
			n = append(n, v)
		}
	}
	return n
}

func unionBy[T any, U comparable](a1 []T, a2 []T, f func(T) U) []T {
	b1 := make([]U, len(a1))
	for i, v := range a1 {
		b1[i] = f(v)
	}

	n := make([]T, len(a1))
	copy(n, a1)
	for _, v := range a2 {
		if !includes(b1, f(v)) {
			n = append(n, v)
		}
	}
	return n
}

func unionWith[T any](a1 []T, a2 []T, pred func(T, T) bool) []T {
	n := make([]T, len(a1))
	copy(n, a1)

	for _, v := range a2 {
		f := func(v2 T) bool { return pred(v, v2) }
		if !some(n, f) {
			n = append(n, v)
		}
	}
	return n
}

func uniq[T comparable](ary []T) []T {
	out := make([]T, 0)
	for _, v := range ary {
		if !includes(out, v) {
			out = append(out, v)
		}
	}
	return out
}

func uniqBy[T any, U comparable](ary []T, conv func(T) U) []T {
	out := make([]T, 0)
	memo := make([]U, 0)

	for _, v := range ary {
		u := conv(v)
		if !includes(memo, u) {
			out = append(out, v)
			memo = append(memo, u)
		}
	}
	return out
}

func uniqWith[T any](ary []T, pred func(T, T) bool) []T {
	out := make([]T, 0)

	for _, v := range ary {
		f := func(v2 T) bool { return pred(v, v2) }
		if !some(out, f) {
			out = append(out, v)
		}
	}
	return out
}

func unzip[T any](as [][]T) [][]T {
	return zip(as...)
}

func without[T comparable](ary []T, values ...T) []T {
	out := make([]T, 0)
	for _, v := range ary {
		if !includes(values, v) {
			out = append(out, v)
		}
	}
	return out
}

func xor[T comparable](a1, a2 []T) []T {
	out1 := difference(a1, a2)
	out2 := difference(a2, a1)
	return concat(out1, out2)
}

func xorBy[T any, U comparable](a1, a2 []T, conv func(T) U) []T {
	out1 := differenceBy(a1, a2, conv)
	out2 := differenceBy(a2, a1, conv)
	return concat(out1, out2)
}

func xorWith[T any](a1, a2 []T, pred func(T, T) bool) []T {
	out1 := differenceWith(a1, a2, pred)
	out2 := differenceWith(a2, a1, pred)
	return concat(out1, out2)
}

func zip[T any](as ...[]T) [][]T {
	if len(as) == 0 {
		return [][]T{}
	}

	out := make([][]T, 0)
	for index := 0; ; index++ {
		tmp := make([]T, len(as))
		for i, a := range as {
			if len(a) <= index {
				return out
			}
			tmp[i] = a[index]
		}
		out = append(out, tmp)
	}
}
