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

		fill(input, 0)
		if diff := cmp.Diff(expect, input); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		fill(input, 0)
		if diff := cmp.Diff(expect, input); diff != "" {
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

func TestHead(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := "a"

		output := head(input)
		if diff := cmp.Diff(expect, output); diff != "" {
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

func TestInitial(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := []string{"a", "b"}

		output := initial(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a"}
		expect := []string{}

		output := initial(input)
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

func TestJoin(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := "a-b-c"

		output := join(input, "-")
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestLast(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := "c"

		output := last(input)
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

func TestNth(t *testing.T) {
	{
		input := []string{"a", "b", "c", "d", "e", "f"}
		expect := "b"

		output := nth(input, 1)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "d", "e", "f"}
		expect := "e"

		output := nth(input, -2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestReverse(t *testing.T) {
	{
		input := []string{"a", "b", "c", "d", "e", "f"}
		expect := []string{"f", "e", "d", "c", "b", "a"}

		reverse(input)
		if diff := cmp.Diff(expect, input); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "d", "e"}
		expect := []string{"e", "d", "c", "b", "a"}

		reverse(input)
		if diff := cmp.Diff(expect, input); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a"}
		expect := []string{"a"}

		reverse(input)
		if diff := cmp.Diff(expect, input); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{}
		expect := []string{}

		reverse(input)
		if diff := cmp.Diff(expect, input); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSlice(t *testing.T) {
	{
		input := []string{"a", "b", "c", "d", "e", "f"}
		expect := []string{"c", "d"}

		output := slice(input, 2, 4)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "d", "e", "f"}
		expect := []string{"c", "d", "e", "f"}

		output := slice(input, 2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c", "d", "e", "f"}
		expect := []string{"a", "b", "c", "d", "e", "f"}

		output := slice(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSortedIndex(t *testing.T) {
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 2

		output := sortedIndex(input, 25)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 0

		output := sortedIndex(input, 5)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 5

		output := sortedIndex(input, 55)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 2

		output := sortedIndex(input, 30)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 20, 20, 30}
		expect := 1

		output := sortedIndex(input, 20)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSortedIndexBy(t *testing.T) {
	conv := func(a int) int { return a }

	{
		input := []int{10, 20, 30, 40, 50}
		expect := 2

		output := sortedIndexBy(input, 25, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 0

		output := sortedIndexBy(input, 5, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 5

		output := sortedIndexBy(input, 55, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 2

		output := sortedIndexBy(input, 30, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 20, 20, 30}
		expect := 1

		output := sortedIndexBy(input, 20, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSortedLastIndex(t *testing.T) {
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 2

		output := sortedLastIndex(input, 25)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 0

		output := sortedLastIndex(input, 5)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 5

		output := sortedLastIndex(input, 55)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 3

		output := sortedLastIndex(input, 30)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 20, 20, 30}
		expect := 4

		output := sortedLastIndex(input, 20)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSortedLastIndexBy(t *testing.T) {
	conv := func(a int) int { return a }

	{
		input := []int{10, 20, 30, 40, 50}
		expect := 2

		output := sortedLastIndexBy(input, 25, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 0

		output := sortedLastIndexBy(input, 5, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 5

		output := sortedLastIndexBy(input, 55, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 40, 50}
		expect := 3

		output := sortedLastIndexBy(input, 30, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 20, 20, 30}
		expect := 4

		output := sortedLastIndexBy(input, 20, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSortedUniq(t *testing.T) {
	{
		input := []int{10, 20, 20, 20, 30}
		expect := []int{10, 20, 30}

		output := sortedUniq(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 50, 20, 30}
		expect := []int{10, 20, 50, 20, 30}

		output := sortedUniq(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := sortedUniq(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestSortedUniqBy(t *testing.T) {
	conv := func(v int) int { return v }

	{
		input := []int{10, 20, 20, 20, 30}
		expect := []int{10, 20, 30}

		output := sortedUniqBy(input, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 50, 20, 30}
		expect := []int{10, 20, 50, 20, 30}

		output := sortedUniqBy(input, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := sortedUniqBy(input, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestTail(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := []string{"b", "c"}

		output := tail(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a"}
		expect := []string{}

		output := tail(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestTake(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := []string{"a", "b"}

		output := take(input, 2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{"a", "b", "c"}

		output := take(input, 5)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{}

		output := take(input, 0)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestTakeRight(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := []string{"b", "c"}

		output := takeRight(input, 2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{"a", "b", "c"}

		output := takeRight(input, 5)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{}

		output := takeRight(input, 0)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestTakeRightWhile(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := []string{"b", "c"}
		pred := func(v string) bool { return v != "a" }

		output := takeRightWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{"a", "b", "c"}
		pred := func(v string) bool { return v != "GGG" }

		output := takeRightWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{}
		pred := func(v string) bool { return v != "c" }

		output := takeRightWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestTakeWhile(t *testing.T) {
	{
		input := []string{"a", "b", "c"}
		expect := []string{"a", "b"}
		pred := func(v string) bool { return v != "c" }

		output := takeWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{"a", "b", "c"}
		pred := func(v string) bool { return v != "GGG" }

		output := takeWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []string{"a", "b", "c"}
		expect := []string{}
		pred := func(v string) bool { return v != "a" }

		output := takeWhile(input, pred)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestUnion(t *testing.T) {
	{
		input1 := []int{0, 1}
		input2 := []int{2, 1, 4}
		expect := []int{0, 1, 2, 4}

		output := union(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{0, 1}
		input2 := []int{2, 4}
		expect := []int{0, 1, 2, 4}

		output := union(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestUnionBy(t *testing.T) {
	floor := func(v float64) int { return int(math.Floor(v)) }

	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 1.5, 4.5}
		expect := []float64{0.5, 1.4, 2.3, 4.5}

		output := unionBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 4.5}
		expect := []float64{0.5, 1.4, 2.3, 4.5}

		output := unionBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestUnionWith(t *testing.T) {
	equals := func(v1, v2 float64) bool { return int(math.Floor(v1)) == int(math.Floor(v2)) }

	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 1.5, 4.5}
		expect := []float64{0.5, 1.4, 2.3, 4.5}

		output := unionWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{0.5, 1.4}
		input2 := []float64{2.3, 4.5}
		expect := []float64{0.5, 1.4, 2.3, 4.5}

		output := unionWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestUniq(t *testing.T) {
	{
		input := []int{10, 20, 20, 20, 30}
		expect := []int{10, 20, 30}

		output := uniq(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 50, 20, 30}
		expect := []int{10, 20, 50, 30}

		output := uniq(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := uniq(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestUniqBy(t *testing.T) {
	conv := func(v int) int { return v }

	{
		input := []int{10, 20, 20, 20, 30}
		expect := []int{10, 20, 30}

		output := uniqBy(input, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 50, 20, 30}
		expect := []int{10, 20, 50, 30}

		output := uniqBy(input, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := uniqBy(input, conv)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestUniqWith(t *testing.T) {
	equals := func(v, u int) bool { return v == u }

	{
		input := []int{10, 20, 20, 20, 30}
		expect := []int{10, 20, 30}

		output := uniqWith(input, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 50, 20, 30}
		expect := []int{10, 20, 50, 30}

		output := uniqWith(input, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{}
		expect := []int{}

		output := uniqWith(input, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestUnzip(t *testing.T) {
	{
		input := [][]int{{10, 1}, {20, 2}, {30, 3}}
		expect := [][]int{
			{10, 20, 30},
			{1, 2, 3},
		}

		output := unzip(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := [][]int{}
		expect := [][]int{}

		output := unzip(input)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestWithout(t *testing.T) {
	{
		input := []int{10, 20, 30, 20, 40}
		expect := []int{10, 30, 40}

		output := without(input, 20)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input := []int{10, 20, 30, 20, 40}
		expect := []int{10, 30}

		output := without(input, 20, 40)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestXor(t *testing.T) {
	{
		input1 := []int{1, 2}
		input2 := []int{2, 3}
		expect := []int{1, 3}

		output := xor(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{1, 2}
		input2 := []int{3, 4}
		expect := []int{1, 2, 3, 4}

		output := xor(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{1, 2}
		input2 := []int{2, 1}
		expect := []int{}

		output := xor(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestXorBy(t *testing.T) {
	floor := func(v float64) int { return int(math.Floor(v)) }

	{
		input1 := []float64{1.1, 2.1}
		input2 := []float64{2.2, 3.2}
		expect := []float64{1.1, 3.2}

		output := xorBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{1.1, 2.1}
		input2 := []float64{3.2, 4.2}
		expect := []float64{1.1, 2.1, 3.2, 4.2}

		output := xorBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{1.1, 2.1}
		input2 := []float64{2.2, 1.2}
		expect := []float64{}

		output := xorBy(input1, input2, floor)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestXorWith(t *testing.T) {
	equals := func(v1, v2 float64) bool { return int(math.Floor(v1)) == int(math.Floor(v2)) }

	{
		input1 := []float64{1.1, 2.1}
		input2 := []float64{2.2, 3.2}
		expect := []float64{1.1, 3.2}

		output := xorWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{1.1, 2.1}
		input2 := []float64{3.2, 4.2}
		expect := []float64{1.1, 2.1, 3.2, 4.2}

		output := xorWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []float64{1.1, 2.1}
		input2 := []float64{2.2, 1.2}
		expect := []float64{}

		output := xorWith(input1, input2, equals)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}

func TestZip(t *testing.T) {
	{
		input1 := []int{10, 20, 30}
		input2 := []int{1, 2, 3}
		expect := [][]int{{10, 1}, {20, 2}, {30, 3}}

		output := zip(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{10, 20, 30, 40}
		input2 := []int{1, 2, 3}
		expect := [][]int{{10, 1}, {20, 2}, {30, 3}}

		output := zip(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{10, 20, 30}
		input2 := []int{1, 2, 3, 4}
		expect := [][]int{{10, 1}, {20, 2}, {30, 3}}

		output := zip(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{}
		input2 := []int{1, 2, 3}
		expect := [][]int{}

		output := zip(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
	{
		input1 := []int{10, 20, 30}
		input2 := []int{}
		expect := [][]int{}

		output := zip(input1, input2)
		if diff := cmp.Diff(expect, output); diff != "" {
			t.Errorf("result is missmatch (-expect, +result):\n%s", diff)
		}
	}
}
