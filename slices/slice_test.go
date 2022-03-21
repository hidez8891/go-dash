package slices

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestChunk(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4}
		expect := [][]int{{0, 1}, {2, 3}, {4}}

		output := chunk(input, 2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := [][]int{{0, 1, 2, 3, 4}}

		output := chunk(input, 10)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := [][]int{{}}

		output := chunk(input, 10)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := [][]int{{0, 1, 2, 3, 4}}

		output := chunk(input, -1)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestConcat(t *testing.T) {
	{
		input1 := []int{0, 1}
		input2 := []int{2, 3, 4}
		expect := []int{0, 1, 2, 3, 4}

		output := concat(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{}
		input2 := []int{2, 3, 4}
		expect := []int{2, 3, 4}

		output := concat(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{}
		input2 := []int{}
		expect := []int{}

		output := concat(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestDifference(t *testing.T) {
	{
		input1 := []int{0, 1}
		input2 := []int{2, 0, 4}
		expect := []int{1}

		output := difference(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{0, 1}
		input2 := []int{2, 0, 4, 1}
		expect := []int{}

		output := difference(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestDifferenceBy(t *testing.T) {
	floor := func(v float64) int { return int(math.Floor(v)) }

	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 0.4, 4.5}
		expect := []float64{1.4}

		output := differenceBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 0.4, 4.5, 1.3}
		expect := []float64{}

		output := differenceBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	equals := func(v1, v2 float64) bool { return int(math.Floor(v1)) == int(math.Floor(v2)) }

	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 0.4, 4.5}
		expect := []float64{1.4}

		output := differenceWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 0.4, 4.5, 1.3}
		expect := []float64{}

		output := differenceWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestDrop(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4}
		expect := []int{2, 3, 4}

		output := drop(input, 2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := []int{}

		output := drop(input, 10)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := drop(input, 10)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := []int{0, 1, 2, 3, 4}

		output := drop(input, -1)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestDropRight(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4}
		expect := []int{0, 1, 2}

		output := dropRight(input, 2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := []int{}

		output := dropRight(input, 10)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := dropRight(input, 10)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4}
		expect := []int{0, 1, 2, 3, 4}

		output := dropRight(input, -1)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestDropRightWhile(t *testing.T) {
	pred := func(v int) bool { return v%2 != 0 }

	{
		input := []int{0, 2, 4, 1, 3}
		expect := []int{0, 2, 4}

		output := dropRightWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{1, 3, 5}
		expect := []int{}

		output := dropRightWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 2, 4}
		expect := []int{0, 2, 4}

		output := dropRightWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := dropRightWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestDropWhile(t *testing.T) {
	pred := func(v int) bool { return v%2 == 0 }

	{
		input := []int{0, 2, 4, 1, 3}
		expect := []int{1, 3}

		output := dropWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 2, 4}
		expect := []int{}

		output := dropWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{1, 3, 5}
		expect := []int{1, 3, 5}

		output := dropWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := dropWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFill(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4}
		expect := []int{0, 0, 0, 0, 0}

		output := fill(input, 0)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := fill(input, 0)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFindIndex(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		pred := func(v int) bool { return v%2 != 0 }

		output := findIndex(input, pred)
		if diff := cmp.Diff(1, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4, 5}
		pred := func(v int) bool { return v%2 == 3 }

		output := findIndex(input, pred)
		if diff := cmp.Diff(-1, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFindLastIndex(t *testing.T) {
	{
		input := []int{0, 1, 2, 3, 4, 5}
		pred := func(v int) bool { return v%2 == 0 }

		output := findLastIndex(input, pred)
		if diff := cmp.Diff(4, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{0, 1, 2, 3, 4, 5}
		pred := func(v int) bool { return v%2 == 3 }

		output := findLastIndex(input, pred)
		if diff := cmp.Diff(-1, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestFlatten(t *testing.T) {
	{
		input := [][]int{{0, 1, 2}, {3, 4}, {5}}
		expect := []int{0, 1, 2, 3, 4, 5}

		output := flatten(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := [][][]int{{{0, 1, 2}}, {{3, 4}, {5}}}
		expect := [][]int{{0, 1, 2}, {3, 4}, {5}}

		output := flatten(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := [][]int{{}, {3, 4}, {5}}
		expect := []int{3, 4, 5}

		output := flatten(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := [][]int{{}, {}, {}}
		expect := []int{}

		output := flatten(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestIntersection(t *testing.T) {
	{
		input1 := []int{0, 1}
		input2 := []int{2, 1, 4}
		expect := []int{1}

		output := intersection(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{0, 1}
		input2 := []int{2, 4}
		expect := []int{}

		output := intersection(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestIntersectionBy(t *testing.T) {
	floor := func(v float64) int { return int(math.Floor(v)) }

	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 1.5, 4.5}
		expect := []float64{1.4}

		output := intersectionBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 4.5}
		expect := []float64{}

		output := intersectionBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestIntersectionWith(t *testing.T) {
	equals := func(v1, v2 float64) bool { return int(math.Floor(v1)) == int(math.Floor(v2)) }

	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 1.5, 4.5}
		expect := []float64{1.4}

		output := intersectionWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 4.5}
		expect := []float64{}

		output := intersectionWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestLastIndexOf(t *testing.T) {
	{
		input := []string{"a", "b", "c", "a", "b", "c"}

		output := lastIndexOf(input, "a")
		if diff := cmp.Diff(3, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "a", "b", "c"}

		output := lastIndexOf(input, "GGG")
		if diff := cmp.Diff(-1, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestIndexOf(t *testing.T) {
	{
		input := []string{"a", "b", "c", "a", "b", "c"}

		output := indexOf(input, "c")
		if diff := cmp.Diff(2, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "a", "b", "c"}

		output := indexOf(input, "GGG")
		if diff := cmp.Diff(-1, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestPull(t *testing.T) {
	{
		input := []string{"a", "b", "c", "a", "b", "c"}
		expect := []string{"a", "c", "a", "c"}

		output := pull(input, "b")
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "a", "b", "c"}
		expect := []string{"a", "a"}

		output := pull(input, "c", "b")
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "a", "b", "c"}
		expect := []string{"a", "b", "c", "a", "b", "c"}

		output := pull(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestPullAll(t *testing.T) {
	{
		input := []string{"a", "b", "c", "a", "b", "c"}
		expect := []string{"a", "c", "a", "c"}

		output := pullAll(input, []string{"b"})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "a", "b", "c"}
		expect := []string{"a", "a"}

		output := pullAll(input, []string{"c", "b"})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "a", "b", "c"}
		expect := []string{"a", "b", "c", "a", "b", "c"}

		output := pullAll(input, []string{})
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}
