package slices

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountBy(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4}
		expect := map[string]int{
			"even": 3,
			"odd":  2,
		}
		f := func(v int) string {
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
		input := []int{0, 2, 4}
		expect := map[string]int{
			"even": 3,
		}
		f := func(v int) string {
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
}

func TestEvery(t *testing.T) {
	{
		input := []int{0, 2, 4}
		pred := func(v int) bool { return v%2 == 0 }
		expect := true

		output := every(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4, 5}
		pred := func(v int) bool { return v%2 == 0 }
		expect := false

		output := every(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFilter(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		expect := []int{0, 2, 4}

		output := filter(input, func(v int) bool { return v%2 == 0 })
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a1", "b1", "a2", "b2", "a3", "b3"}
		expect := []string{"a1", "a2", "a3"}

		output := filter(input, func(v string) bool {
			return strings.HasPrefix(v, "a")
		})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFind(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		output, exist := find(input, func(v int) bool { return v%2 != 0 })
		if diff := cmp.Diff(true, exist); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
		if diff := cmp.Diff(1, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4, 5}
		_, exist := find(input, func(v int) bool { return v%2 == 3 })
		if diff := cmp.Diff(false, exist); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFindLast(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		output, exist := findLast(input, func(v int) bool { return v%2 == 0 })
		if diff := cmp.Diff(true, exist); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
		if diff := cmp.Diff(4, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4, 5}
		_, exist := findLast(input, func(v int) bool { return v%2 == 3 })
		if diff := cmp.Diff(false, exist); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestForEach(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		output := make([]int, 0, cap(input))

		forEach(input, func(v int) {
			output = append(output, v)
		})
		if diff := cmp.Diff(input, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "d", "e"}
		output := make([]string, 0, cap(input))

		forEach(input, func(v string) {
			output = append(output, v)
		})
		if diff := cmp.Diff(input, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestForEachRight(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		expect := []int{5, 4, 3, 2, 1, 0}
		output := make([]int, 0, cap(input))

		forEachRight(input, func(v int) {
			output = append(output, v)
		})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "d", "e"}
		expect := []string{"e", "d", "c", "b", "a"}
		output := make([]string, 0, cap(input))

		forEachRight(input, func(v string) {
			output = append(output, v)
		})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestGroupBy(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4}
		expect := map[string][]int{
			"even": {0, 2, 4},
			"odd":  {1, 3},
		}
		f := func(v int) string {
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
		input := []int{0, 2, 4}
		expect := map[string][]int{
			"even": {0, 2, 4},
		}
		f := func(v int) string {
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
}

func TestIncludes(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4}

		output := includes(input, 3)
		if diff := cmp.Diff(true, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}

		output := includes(input, -1)
		if diff := cmp.Diff(false, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestMapBy(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		expect := []int{0, 2, 4, 6, 8, 10}

		output := mapBy(input, func(v int) int { return v * 2 })
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4, 5}
		expect := []string{"0", "1", "2", "3", "4", "5"}

		output := mapBy(input, func(v int) string {
			return fmt.Sprintf("%d", v)
		})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestPartition(t *testing.T) {
	trans := cmp.Transformer("sort", func(in [][]int) [][]int {
		sort.Slice(in, func(i, j int) bool {
			if in[i][0] == in[j][0] {
				return len(in[i]) < len(in[j])
			}
			return in[i][0] < in[j][0]
		})
		return in
	})

	{
		input := []int{0, 1, 2, 3, 4}
		expect := [][]int{{0, 2, 4}, {1, 3}}
		f := func(v int) string {
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
		input := []int{0, 2, 4}
		expect := [][]int{{0, 2, 4}}
		f := func(v int) string {
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
}

func TestReduce(t *testing.T) {
	{
		input := []string{"0", "1", "2", "3", "4"}
		expect := "01234"
		f := func(v, acc string) string { return acc + v }

		output := reduce(input, f, "")
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := "01234"
		f := func(v int, acc string) string { return acc + fmt.Sprintf("%d", v) }

		output := reduce(input, f, "")
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestReduceRight(t *testing.T) {
	{
		input := []string{"0", "1", "2", "3", "4"}
		expect := "43210"
		f := func(v, acc string) string { return acc + v }

		output := reduceRight(input, f, "")
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := "43210"
		f := func(v int, acc string) string { return acc + fmt.Sprintf("%d", v) }

		output := reduceRight(input, f, "")
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSome(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		pred := func(v int) bool { return v%2 == 0 }
		expect := true

		output := some(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{1, 3, 5}
		pred := func(v int) bool { return v%2 == 0 }
		expect := false

		output := some(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestReject(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		pred := func(v int) bool { return v%2 == 0 }
		expect := []int{1, 3, 5}

		output := reject(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{1, 3, 5}
		pred := func(v int) bool { return v%2 == 0 }
		expect := []int{1, 3, 5}

		output := reject(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}
