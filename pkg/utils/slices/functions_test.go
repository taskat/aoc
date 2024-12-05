package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		testName      string
		slice         []int
		item          int
		expectedValue bool
	}{
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
}

func TestRemoveNthPanic(t *testing.T) {
	testCases := []struct {
		testName string
		slice    []int
		index    int
	}{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
		{"Remove out of bounds", []int{1, 2, 3}, 3},
		{"Remove negative index", []int{1, 2, 3}, -1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Panics(t, func() {
				RemoveNth(tc.slice, tc.index)
			})
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		testName      string
		slice         []int
		expectedValue int
	}{
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
