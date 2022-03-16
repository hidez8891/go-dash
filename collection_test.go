package go_dash

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
