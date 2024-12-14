package slices

import (
	"cmp"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskat/aoc/pkg/utils/types"
)

type numbers []int

func TestAny(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		predicate     func(T) bool
		expectedValue bool
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i int) bool { return true }, false},
		{"Empty slice", []int{}, func(i int) bool { return true }, false},
		{"Any even numbers", []int{1, 2, 3}, func(i int) bool { return i%2 == 0 }, true},
		{"Any odd numbers", []int{1, 2, 3}, func(i int) bool { return i%2 == 1 }, true},
		{"Any no numbers", []int{1, 2, 3}, func(i int) bool { return false }, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Any(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		testName      string
		slice         []T
		item          int
		expectedValue bool
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, 1, false},
		{"Empty slice", []int{}, 1, false},
		{"Item not in slice", []int{2, 3, 4}, 1, false},
		{"Item in slice", []int{2, 3, 4}, 3, true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Contains(tc.slice, tc.item)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestCopy(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		expectedValue []T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, []int{}},
		{"Empty slice", []int{}, []int{}},
		{"Copy slice", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Copy(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, Copy)
}

func TestCount(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		predicate     func(T) bool
		expectedValue int
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i int) bool { return true }, 0},
		{"Empty slice", []int{}, func(i int) bool { return true }, 0},
		{"Count even numbers", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }, 2},
		{"Count odd numbers", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 1 }, 2},
		{"Count all numbers", []int{1, 2, 3, 4}, func(i int) bool { return true }, 4},
		{"Count no numbers", []int{1, 2, 3, 4}, func(i int) bool { return false }, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Count(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestFilter(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		predicate     func(T) bool
		expectedValue []T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i int) bool { return true }, []int{}},
		{"Empty slice", []int{}, func(i int) bool { return true }, []int{}},
		{"Filter even numbers", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }, []int{2, 4}},
		{"Filter odd numbers", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 1 }, []int{1, 3}},
		{"Filter all numbers", []int{1, 2, 3, 4}, func(i int) bool { return true }, []int{1, 2, 3, 4}},
		{"Filter no numbers", []int{1, 2, 3, 4}, func(i int) bool { return false }, []int{}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Filter(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestFind(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		predicate     func(T) bool
		expectedValue T
		expectedFound bool
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i int) bool { return true }, 0, false},
		{"Empty slice", []int{}, func(i int) bool { return true }, 0, false},
		{"Find even number", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }, 2, true},
		{"Find odd number", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 1 }, 1, true},
		{"Find first number", []int{1, 2, 3, 4}, func(i int) bool { return true }, 1, true},
		{"Find no number", []int{1, 2, 3, 4}, func(i int) bool { return false }, 0, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, found := Find(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedValue, result)
			assert.Equal(t, tc.expectedFound, found)
		})
	}
}

func TestFindIndex(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		predicate     func(T) bool
		expectedValue int
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i int) bool { return true }, -1},
		{"Empty slice", []int{}, func(i int) bool { return true }, -1},
		{"Find even number", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }, 1},
		{"Find odd number", []int{1, 2, 3, 4}, func(i int) bool { return i%2 == 1 }, 0},
		{"Find first number", []int{1, 2, 3, 4}, func(i int) bool { return true }, 0},
		{"Find no number", []int{1, 2, 3, 4}, func(i int) bool { return false }, -1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := FindIndex(tc.slice, tc.predicate)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestFirst(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"First element", []int{1, 2, 3}, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := First(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { First(tc.slice) })
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, First)
}

func TestFor(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		f             func(int) T
		i             int
		expectedSlice []T
	}
	testCases := []testCase[int]{
		{"No repeat", func(i int) int { return i }, 0, []int{}},
		{"Repeat 3 times", func(i int) int { return i }, 3, []int{0, 1, 2}},
		{"Repeat 5 times", func(i int) int { return i * 2 }, 5, []int{0, 2, 4, 6, 8}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := For(tc.i, tc.f)
			assert.Equal(t, tc.expectedSlice, actual)
		})
	}
}

func TestForEach(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		f             func(T)
		expectedSlice []T
	}
	testCases := []testCase[[]int]{
		{"Nil slice", nil, func(i []int) {}, nil},
		{"Empty slice", [][]int{}, func(i []int) {}, [][]int{}},
		{"For each element", [][]int{{1}, {2, 3}}, func(i []int) { i[0] += 1 }, [][]int{{2}, {3, 3}}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ForEach(tc.slice, tc.f)
			assert.Equal(t, tc.expectedSlice, tc.slice)
		})
	}
}

func TestForEach_m(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		f             func(*T)
		expectedSlice []T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(i *int) {}, nil},
		{"Empty slice", []int{}, func(i *int) {}, []int{}},
		{"For each element", []int{1, 2, 3}, func(i *int) { *i += 1 }, []int{2, 3, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ForEach_m(tc.slice, tc.f)
			assert.Equal(t, tc.expectedSlice, tc.slice)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		expectedValue bool
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, true},
		{"Empty slice", []int{}, true},
		{"Non-empty slice", []int{1, 2, 3}, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := IsEmpty(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, IsEmpty)
}

// TestIsInBounds tests the IsInBounds function
func TestIsInBounds(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		index         int
		expectedValue bool
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, 0, false},
		{"Empty slice", []int{}, 0, false},
		{"Index out of bounds", []int{1, 2, 3}, 3, false},
		{"Index in bounds", []int{1, 2, 3}, 2, true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := IsInBounds(tc.slice, tc.index)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestLast(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Last element", []int{1, 2, 3}, 3},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Last(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { Last(tc.slice) })
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, Last)
}

func TestMap(t *testing.T) {
	type testCase[T, U any] struct {
		testName      string
		slice         []T
		f             func(T) U
		expectedValue []U
	}
	testCases := []testCase[int, string]{
		{"Nil slice", nil, func(i int) string { return "a" }, []string{}},
		{"Empty slice", []int{}, func(i int) string { return "a" }, []string{}},
		{"Map from int to string", []int{1, 2, 3}, func(i int) string { return fmt.Sprintf("%d", i) }, []string{"1", "2", "3"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Map(tc.slice, tc.f)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestMap_i(t *testing.T) {
	type testCase[T, U any] struct {
		testName      string
		slice         []T
		f             func(T, int) U
		expectedValue []U
	}
	testCases := []testCase[int, string]{
		{"Nil slice", nil, func(i int, j int) string { return "a" }, []string{}},
		{"Empty slice", []int{}, func(i int, j int) string { return "a" }, []string{}},
		{"Map from int to string", []int{1, 2, 3}, func(i int, j int) string { return fmt.Sprintf("%d", i+j) }, []string{"1", "3", "5"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Map_i(tc.slice, tc.f)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestMax(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		testName      string
		slice         []T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Max of positive numbers", []int{1, 2, 3}, 3},
		{"Max of negative numbers", []int{-1, -2, -3}, -1},
		{"Max of mixed numbers", []int{1, -2, 3}, 3},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Max(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { Max(tc.slice) })
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, Max)
}

func TestMax_i(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		testName      string
		slice         []T
		expectedValue T
		expectedIndex int
	}
	testCases := []testCase[int]{
		{"Max of positive numbers", []int{1, 2, 3}, 3, 2},
		{"Max of negative numbers", []int{-1, -2, -3}, -1, 0},
		{"Max of mixed numbers", []int{1, -2, 3}, 3, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, index := Max_i(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
			assert.Equal(t, tc.expectedIndex, index)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0, 0},
		{"Empty slice", []int{}, 0, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { Max_i(tc.slice) })
		})
	}
}

func TestMiddle(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Middle of odd length slice", []int{1, 2, 3}, 2},
		{"Middle of even length slice", []int{1, 2, 3, 4}, 3},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Middle(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { Middle(tc.slice) })
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, Middle)
}

func TestMin(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		testName      string
		slice         []T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Min of positive numbers", []int{1, 2, 3}, 1},
		{"Min of negative numbers", []int{-1, -2, -3}, -3},
		{"Min of mixed numbers", []int{1, -2, 3}, -2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Min(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { Min(tc.slice) })
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, Min)
}

func TestMin_i(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		testName      string
		slice         []T
		expectedValue T
		expectedIndex int
	}
	testCases := []testCase[int]{
		{"Min of positive numbers", []int{1, 2, 3}, 1, 0},
		{"Min of negative numbers", []int{-1, -2, -3}, -3, 2},
		{"Min of mixed numbers", []int{1, -2, 3}, -2, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, index := Min_i(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
			assert.Equal(t, tc.expectedIndex, index)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0, 0},
		{"Empty slice", []int{}, 0, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { Min_i(tc.slice) })
		})
	}
}

func TestProduct(t *testing.T) {
	type testCase[T types.Number] struct {
		testName      string
		slice         []T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, 1},
		{"Empty slice", []int{}, 1},
		{"Product of positive numbers", []int{1, 2, 3}, 6},
		{"Product of negative numbers", []int{-1, -2, -3}, -6},
		{"Product of mixed numbers", []int{1, -2, 3}, -6},
		{"Product of zeros", []int{1, 0, 3}, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Product(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, Product)
}

func TestReduce(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		f             func(T, T) T
		initialValue  T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(a, b int) int { return a + b }, 0, 0},
		{"Empty slice", []int{}, func(a, b int) int { return a + b }, 0, 0},
		{"Sum of positive numbers", []int{1, 2, 3}, func(a, b int) int { return a + b }, 0, 6},
		{"Sum of negative numbers", []int{-1, -2, -3}, func(a, b int) int { return a + b }, 0, -6},
		{"Sum of mixed numbers", []int{1, -2, 3}, func(a, b int) int { return a + b }, 0, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Reduce(tc.slice, tc.f, tc.initialValue)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestReduce_i(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		f             func(T, T, int) T
		initialValue  T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, func(a, b, i int) int { return a + b + i }, 0, 0},
		{"Empty slice", []int{}, func(a, b, i int) int { return a + b + i }, 0, 0},
		{"Sum of positive numbers", []int{1, 2, 3}, func(a, b, i int) int { return a + b + i }, 0, 9},
		{"Sum of negative numbers", []int{-1, -2, -3}, func(a, b, i int) int { return a + b + i }, 0, -3},
		{"Sum of mixed numbers", []int{1, -2, 3}, func(a, b, i int) int { return a + b + i }, 0, 5},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Reduce_i(tc.slice, tc.f, tc.initialValue)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestRemoveNth(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		index         int
		expectedValue []T
	}
	testCases := []testCase[int]{
		{"Remove first element", []int{1, 2, 3}, 0, []int{2, 3}},
		{"Remove middle element", []int{1, 2, 3}, 1, []int{1, 3}},
		{"Remove last element", []int{1, 2, 3}, 2, []int{1, 2}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := RemoveNth(tc.slice, tc.index)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0, []int{}},
		{"Empty slice", []int{}, 0, []int{}},
		{"Remove out of bounds", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"Remove negative index", []int{1, 2, 3}, -1, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { RemoveNth(tc.slice, tc.index) })
		})
	}
}

func TestRepeat(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		item          T
		times         int
		expectedValue []T
	}
	testCases := []testCase[int]{
		{"Repeat 0 times", 1, 0, []int{}},
		{"Repeat 1 time", 1, 1, []int{1}},
		{"Repeat 3 times", 1, 3, []int{1, 1, 1}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Repeat(tc.item, tc.times)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}

func TestSum(t *testing.T) {
	type testCase[T types.Summable] struct {
		testName      string
		slice         []T
		expectedValue T
	}
	testCases := []testCase[int]{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
		{"Sum of positive numbers", []int{1, 2, 3}, 6},
		{"Sum of negative numbers", []int{-1, -2, -3}, -6},
		{"Sum of mixed numbers", []int{1, -2, 3}, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Sum(tc.slice)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test that it works with a custom type
	Map([]numbers{}, Sum)
}

func TestSwap(t *testing.T) {
	type testCase[T any] struct {
		testName      string
		slice         []T
		i             int
		j             int
		expectedValue []T
	}
	testCases := []testCase[int]{
		{"Swap first and last elements", []int{1, 2, 3}, 0, 2, []int{3, 2, 1}},
		{"Swap two elements", []int{1, 2, 3}, 0, 1, []int{2, 1, 3}},
		{"Swap same element", []int{1, 2, 3}, 1, 1, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			Swap(tc.slice, tc.i, tc.j)
			assert.Equal(t, tc.expectedValue, tc.slice)
		})
	}
	// Test for panics
	testCases = []testCase[int]{
		{"Nil slice", nil, 0, 0, []int{}},
		{"Empty slice", []int{}, 0, 0, []int{}},
		{"Swap out of bounds", []int{1, 2, 3}, 3, 0, []int{1, 2, 3}},
		{"Swap negative index", []int{1, 2, 3}, -1, 0, []int{1, 2, 3}},
		{"Swap out of bounds", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
		{"Swap negative index", []int{1, 2, 3}, 0, -1, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { Swap(tc.slice, tc.i, tc.j) })
		})
	}
}

func TestToMap(t *testing.T) {
	type testCase[T comparable, U any] struct {
		testName      string
		keys          []T
		values        []U
		expectedValue map[T]U
	}
	testCases := []testCase[int, string]{
		{"Nil slices", nil, nil, map[int]string{}},
		{"Empty slices", []int{}, []string{}, map[int]string{}},
		{"To map", []int{1, 2, 3}, []string{"a", "b", "c"}, map[int]string{1: "a", 2: "b", 3: "c"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := ToMap(tc.keys, tc.values)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
	// Test for panics
	testCases = []testCase[int, string]{
		{"Different lengths", []int{1, 2, 3}, []string{"a", "b"}, map[int]string{}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() { ToMap(tc.keys, tc.values) })
		})
	}
}

func TestZipWith(t *testing.T) {
	type testCase[T, U, V any] struct {
		testName      string
		slice1        []T
		slice2        []U
		f             func(T, U) V
		expectedValue []V
	}
	testCases := []testCase[int, int, string]{
		{"Nil slices", nil, nil, func(i int, j int) string { return "a" }, []string{}},
		{"Empty slices", []int{}, []int{}, func(i int, j int) string { return "a" }, []string{}},
		{"Empty and non-empty slices", []int{}, []int{1, 2, 3}, func(i int, j int) string { return "a" }, []string{}},
		{"Non-empty and empty slices", []int{1, 2, 3}, []int{}, func(i int, j int) string { return "a" }, []string{}},
		{"Zip with sum", []int{1, 2, 3}, []int{4, 5, 6}, func(i int, j int) string { return fmt.Sprintf("%d", i+j) }, []string{"5", "7", "9"}},
		{"Zip with different lengths", []int{1, 2, 3}, []int{4, 5}, func(i int, j int) string { return fmt.Sprintf("%d", i+j) }, []string{"5", "7"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := ZipWith(tc.slice1, tc.slice2, tc.f)
			assert.Equal(t, tc.expectedValue, result)
		})
	}
}
