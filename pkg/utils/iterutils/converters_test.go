package iterutils

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskat/aoc/pkg/utils/types"
)

func TestNewFromFunc(t *testing.T) {
	init := 0
	f := func() int {
		init++
		return init
	}
	n := 5
	iter := NewFromFunc(f, n)
	expected := []int{1, 2, 3, 4, 5}
	actual := ToSlice(iter)
	assert.Equal(t, expected, actual)
	// Read all elements
	actual = ToSliceN(iter, 5)
}

func TestNewFromFunc2(t *testing.T) {
	init := 0
	f := func() (int, int) {
		init++
		return init, init
	}
	n := 5
	iter := NewFromFunc2(f, n)
	expected := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	actual := ToMap(iter)
	assert.Equal(t, expected, actual)
	// Read all elements
	actual = ToMapN(iter, 5)
}

func TestNewFromFuncIterations(t *testing.T) {
	f := func(i int) int {
		return i
	}
	n := 5
	iter := NewFromFuncIterations(f, n)
	expected := []int{0, 1, 2, 3, 4}
	actual := ToSlice(iter)
	assert.Equal(t, expected, actual)
	// Read all elements
	actual = ToSliceN(iter, 5)
}

func TestNewFromFuncIterations2(t *testing.T) {
	f := func(i int) (int, int) {
		return i, i
	}
	n := 5
	iter := NewFromFuncIterations2(f, n)
	expected := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}
	actual := ToMap(iter)
	assert.Equal(t, expected, actual)
	// Read all elements
	actual = ToMapN(iter, 5)
}

func TestNewFromMap(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		m        map[K]V
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"nil map", nil, map[int]int{}},
		{"empty map", map[int]int{}, map[int]int{}},
		{"single element", map[int]int{42: 42}, map[int]int{42: 42}},
		{"multiple elements", map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 2: 2, 3: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			iter := NewFromMap(tc.m)
			actual := ToMap(iter)
			assert.Equal(t, tc.expected, actual)
			// Read all elements
			actual = ToMapN(iter, uint(len(tc.expected)))
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNewFromRepeat(t *testing.T) {
	value := 42
	n := 5
	iter := NewFromRepeat(value, n)
	expected := []int{42, 42, 42, 42, 42}
	actual := ToSlice(iter)
	assert.Equal(t, expected, actual)
	// Read all elements
	actual = ToSliceN(iter, 5)
}

func TestNewFromSet(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		set      map[T]struct{}
		expected []T
	}
	testCases := []testCase[int]{
		{"nil set", nil, []int{}},
		{"empty set", map[int]struct{}{}, []int{}},
		{"single element", map[int]struct{}{42: {}}, []int{42}},
		{"multiple elements", map[int]struct{}{1: {}, 2: {}, 3: {}}, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			iter := NewFromSet(tc.set)
			actual := ToSlice(iter)
			assert.ElementsMatch(t, tc.expected, actual)
			// Read all elements
			actual = ToSliceN(iter, uint(len(tc.expected)))
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func TestNewFromSlice(t *testing.T) {
	type testCase[T any] struct {
		name     string
		slice    []T
		expected []T
	}
	testCases := []testCase[int]{
		{"nil slice", nil, []int{}},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			iter := NewFromSlice(tc.slice)
			actual := ToSlice(iter)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNewFromSlice2(t *testing.T) {
	type testCase[V any] struct {
		name     string
		slice    []V
		expected []V
	}
	testCases := []testCase[int]{
		{"nil slice", nil, []int{}},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			iter := NewFromSlice2(tc.slice)
			actual := ToSlice2(iter)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestToMap(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"empty iterator", NewFromSlice2([]int{}), map[int]int{}},
		{"single element", NewFromSlice2([]int{42}), map[int]int{0: 42}},
		{"multiple elements", NewFromSlice2([]int{1, 2, 3}), map[int]int{0: 1, 1: 2, 2: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToMap(tc.iter)
			assert.Equal(t, tc.expected, actual)
		})
	}
	// Test for panic
	assert.PanicsWithValue(t, nilIterPanicMsg, func() {
		ToMap[int, int](nil)
	})
}

func TestToMapN(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		n        uint
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"empty iterator", NewFromSlice2([]int{}), 0, map[int]int{}},
		{"single element", NewFromSlice2([]int{42}), 1, map[int]int{0: 42}},
		{"multiple elements", NewFromSlice2([]int{1, 2, 3}), 2, map[int]int{0: 1, 1: 2}},
		{"n greater than length", NewFromSlice2([]int{1, 2, 3}), 5, map[int]int{0: 1, 1: 2, 2: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToMapN(tc.iter, tc.n)
			assert.Equal(t, tc.expected, actual)
		})
	}
	// Test for panic
	assert.PanicsWithValue(t, nilIterPanicMsg, func() {
		ToMapN[int, int](nil, 0)
	})
}

func TestToSlice(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected []T
	}
	testCases := []testCase[int]{
		{"empty iterator", NewFromSlice([]int{}), []int{}},
		{"single element", NewFromSlice([]int{42}), []int{42}},
		{"multiple elements", NewFromSlice([]int{1, 2, 3}), []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToSlice(tc.iter)
			assert.Equal(t, tc.expected, actual)
		})
	}
	// Test for panic
	assert.PanicsWithValue(t, nilIterPanicMsg, func() {
		ToSlice[int](nil)
	})
}

func TestToSlice2(t *testing.T) {
	type testCase[K types.Integer, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"empty iterator", NewFromSlice2([]int{}), map[int]int{}},
		{"single element", NewFromSlice2([]int{42}), map[int]int{0: 42}},
		{"multiple elements", NewFromSlice2([]int{1, 2, 3}), map[int]int{0: 1, 1: 2, 2: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToMap(tc.iter)
			assert.Equal(t, tc.expected, actual)
		})
	}
	// Test for panic
	assert.PanicsWithValue(t, nilIterPanicMsg, func() {
		ToSlice2[int, int](nil)
	})
}

func TestToSliceN(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		n        uint
		expected []T
	}
	testCases := []testCase[int]{
		{"empty iterator", NewFromSlice([]int{}), 0, []int{}},
		{"single element", NewFromSlice([]int{42}), 1, []int{42}},
		{"multiple elements", NewFromSlice([]int{1, 2, 3}), 2, []int{1, 2}},
		{"n greater than length", NewFromSlice([]int{1, 2, 3}), 5, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ToSliceN(tc.iter, tc.n)
			assert.Equal(t, tc.expected, actual)
		})
	}
	// Test for panic
	assert.PanicsWithValue(t, nilIterPanicMsg, func() {
		ToSliceN[int](nil, 0)
	})
}
