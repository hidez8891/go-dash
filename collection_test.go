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
