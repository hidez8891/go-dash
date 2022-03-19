package maps

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountBy(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		expect := map[string]int{
			"even": 3,
			"odd":  2,
		}
		f := func(_ string, v int) string {
			if v%2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		}

		output := countBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "b": 2, "c": 4}
		expect := map[string]int{
			"even": 3,
		}
		f := func(_ string, v int) string {
			if v%2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		}

		output := countBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a0": 0, "b0": 1, "a1": 2, "b1": 3, "a2": 4}
		expect := map[string]int{
			"a": 3,
			"b": 2,
		}
		f := func(k string, _ int) string { return string(k[0]) }

		output := countBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestEvery(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 2, "c": 4}
		pred := func(_ string, v int) bool { return v%2 == 0 }
		expect := true

		output := every(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		pred := func(_ string, v int) bool { return v%2 == 0 }
		expect := false

		output := every(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a0": 0, "a1": 2, "a2": 4}
		pred := func(k string, _ int) bool { return k[0] == 'a' }
		expect := true

		output := every(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFilter(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		expect := map[string]int{"a": 0, "c": 2, "e": 4}
		f := func(_ string, v int) bool { return v%2 == 0 }

		output := filter(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[int]string{0: "a1", 1: "b1", 2: "a2", 3: "b2", 4: "a3", 5: "b3"}
		expect := map[int]string{0: "a1", 2: "a2", 4: "a3"}
		f := func(_ int, v string) bool {
			return strings.HasPrefix(v, "a")
		}

		output := filter(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[int]string{0: "a1", 1: "b1", 2: "a2", 3: "b2", 4: "a3", 5: "b3"}
		expect := map[int]string{0: "a1", 2: "a2", 4: "a3"}
		f := func(k int, _ string) bool { return k%2 == 0 }

		output := filter(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFind(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		f := func(_ string, v int) bool { return v%2 != 0 }

		key, value, exist := find(input, f)
		if diff := cmp.Diff(true, exist); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
		if diff := cmp.Diff(1, value); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
		if diff := cmp.Diff("b", key); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		f := func(_ string, v int) bool { return v%2 == 3 }

		_, _, exist := find(input, f)
		if diff := cmp.Diff(false, exist); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		f := func(k string, _ int) bool { return k == "b" }

		key, value, exist := find(input, f)
		if diff := cmp.Diff(true, exist); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
		if diff := cmp.Diff(1, value); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
		if diff := cmp.Diff("b", key); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestForEach(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		output := map[string]int{}

		forEach(input, func(k string, v int) {
			output[k] = v
		})
		if diff := cmp.Diff(input, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		expect := map[int]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e"}
		output := map[int]string{}

		forEach(input, func(k string, v int) {
			output[v] = k
		})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestGroupBy(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		expect := map[string]map[string]int{
			"even": {"a": 0, "c": 2, "e": 4},
			"odd":  {"b": 1, "d": 3},
		}
		f := func(_ string, v int) string {
			if v%2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		}

		output := groupBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "c": 2, "e": 4}
		expect := map[string]map[string]int{
			"even": {"a": 0, "c": 2, "e": 4},
		}
		f := func(_ string, v int) string {
			if v%2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		}

		output := groupBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a0": 0, "b0": 1, "a1": 2, "b1": 3, "a2": 4}
		expect := map[string]map[string]int{
			"a": {"a0": 0, "a1": 2, "a2": 4},
			"b": {"b0": 1, "b1": 3},
		}
		f := func(k string, _ int) string { return string(k[0]) }

		output := groupBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestIncludes(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}

		output := includes(input, 3)
		if diff := cmp.Diff(true, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}

		output := includes(input, -1)
		if diff := cmp.Diff(false, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestMapBy(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		expect := map[string]int{"a": 0, "b": 2, "c": 4, "d": 6, "e": 8}
		f := func(_ string, v int) int { return v * 2 }

		output := mapBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		expect := map[string]string{"a": "0", "b": "1", "c": "2", "d": "3", "e": "4"}
		f := func(_ string, v int) string {
			return fmt.Sprintf("%d", v)
		}

		output := mapBy(input, f)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestPartition(t *testing.T) {
	trans := cmp.Transformer("sort", func(in []map[string]int) []map[string]int {
		sort.Slice(in, func(i, j int) bool {
			return len(in[i]) > len(in[j])
		})
		return in
	})

	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		expect := []map[string]int{{"a": 0, "c": 2, "e": 4}, {"b": 1, "d": 3}}
		f := func(_ string, v int) string {
			if v%2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		}

		output := partition(input, f)
		if diff := cmp.Diff(expect, output, trans); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a": 0, "c": 2, "e": 4}
		expect := []map[string]int{{"a": 0, "c": 2, "e": 4}}
		f := func(_ string, v int) string {
			if v%2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		}

		output := partition(input, f)
		if diff := cmp.Diff(expect, output, trans); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a0": 0, "b0": 1, "a1": 2, "b1": 3, "a2": 4}
		expect := []map[string]int{{"a0": 0, "a1": 2, "a2": 4}, {"b0": 1, "b1": 3}}
		f := func(k string, _ int) string { return string(k[0]) }

		output := partition(input, f)
		if diff := cmp.Diff(expect, output, trans); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestReduce(t *testing.T) {
	{
		input := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4}
		expect := 10
		f := func(_ string, v int, acc int) int { return acc + v }

		output := reduce(input, f, 0)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"}
		expect := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4}
		f := func(k int, v string, acc map[string]int) map[string]int {
			acc[v] = k
			return acc
		}

		output := reduce(input, f, map[string]int{})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSome(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		pred := func(_ string, v int) bool { return v%2 == 0 }
		expect := true

		output := some(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"b": 1, "d": 3}
		pred := func(_ string, v int) bool { return v%2 == 0 }
		expect := false

		output := some(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestReject(t *testing.T) {
	{
		input := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}
		pred := func(_ string, v int) bool { return v%2 == 0 }
		expect := map[string]int{"b": 1, "d": 3}

		output := reject(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"b": 1, "d": 3}
		pred := func(_ string, v int) bool { return v%2 == 0 }
		expect := map[string]int{"b": 1, "d": 3}

		output := reject(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := map[string]int{"a0": 0, "b0": 1, "a1": 2, "b1": 3, "a2": 4}
		pred := func(k string, _ int) bool { return k[0] == 'a' }
		expect := map[string]int{"b0": 1, "b1": 3}

		output := reject(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}
