package iterutils

import (
	"fmt"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	type testCase[V any] struct {
		name     string
		iter     iter.Seq[V]
		value    V
		expected []V
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), 0, []int{0}},
		{"Empty slice", NewFromSlice([]int{}), 0, []int{0}},
		{"Append to slice", NewFromSlice([]int{1, 2, 3}), 0, []int{1, 2, 3, 0}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Append(tc.iter, tc.value)
			actual := ToSlice(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Append(nil, 0)
		})
	})
}

func TestAppend2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		key      K
		value    V
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), 0, 0, map[int]int{0: 0}},
		{"Empty slice", NewFromSlice2([]int{}), 0, 0, map[int]int{0: 0}},
		{"Append to slice", NewFromSlice2([]int{0: 0, 1: 1, 2: 2}), 3, 3, map[int]int{0: 0, 1: 1, 2: 2, 3: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Append2(tc.iter, tc.key, tc.value)
			actual := ToMap(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Append2(nil, 0, 0)
		})
	})
}

func TestConcat(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iters    []iter.Seq[T]
		expected []T
	}
	testCases := []testCase[int]{
		{"Nil slices", []iter.Seq[int]{NewFromSlice[int](nil), NewFromSlice[int](nil)}, []int{}},
		{"Empty slices", []iter.Seq[int]{NewFromSlice([]int{}), NewFromSlice([]int{})}, []int{}},
		{"Concat single slice", []iter.Seq[int]{NewFromSlice([]int{1, 2})}, []int{1, 2}},
		{"Concat slices", []iter.Seq[int]{NewFromSlice([]int{1, 2}), NewFromSlice([]int{3, 4})}, []int{1, 2, 3, 4}},
		{"Concat slices with nil", []iter.Seq[int]{NewFromSlice([]int{1, 2}), NewFromSlice[int](nil)}, []int{1, 2}},
		{"Concat slices from nil", []iter.Seq[int]{NewFromSlice[int](nil), NewFromSlice([]int{3, 4})}, []int{3, 4}},
		{"Concat multiple slices", []iter.Seq[int]{NewFromSlice([]int{1, 2}), NewFromSlice([]int{3, 4}), NewFromSlice([]int{5, 6})}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Concat(tc.iters...)
			actual := ToSlice(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name  string
		iters []iter.Seq[T]
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter1", []iter.Seq[int]{nil, NewFromSlice([]int{})}},
		{"Nil iter2", []iter.Seq[int]{NewFromSlice([]int{}), nil}},
		{"Nil iters", []iter.Seq[int]{nil, nil}},
		{"Nil iter in between", []iter.Seq[int]{NewFromSlice([]int{1, 2}), nil, NewFromSlice([]int{3, 4})}},
	}
	for _, tc := range panicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.PanicsWithValue(t, nilIterPanicMsg, func() {
				Concat(tc.iters...)
			})
		})
	}
}

func TestConcat2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iters    []iter.Seq2[K, V]
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"Nil maps", []iter.Seq2[int, int]{NewFromMap[int, int](nil), NewFromMap[int, int](nil)}, map[int]int{}},
		{"Empty maps", []iter.Seq2[int, int]{NewFromMap(map[int]int{}), NewFromMap(map[int]int{})}, map[int]int{}},
		{"Concat single map", []iter.Seq2[int, int]{NewFromMap(map[int]int{1: 1, 2: 2})}, map[int]int{1: 1, 2: 2}},
		{"Concat maps", []iter.Seq2[int, int]{NewFromMap(map[int]int{1: 1, 2: 2}), NewFromMap(map[int]int{3: 3, 4: 4})}, map[int]int{1: 1, 2: 2, 3: 3, 4: 4}},
		{"Concat maps with nil", []iter.Seq2[int, int]{NewFromMap(map[int]int{1: 1, 2: 2}), NewFromMap[int, int](nil)}, map[int]int{1: 1, 2: 2}},
		{"Concat maps from nil", []iter.Seq2[int, int]{NewFromMap[int, int](nil), NewFromMap(map[int]int{3: 3, 4: 4})}, map[int]int{3: 3, 4: 4}},
		{"Concat multiple maps", []iter.Seq2[int, int]{NewFromMap(map[int]int{1: 1, 2: 2}), NewFromMap(map[int]int{3: 3, 4: 4}), NewFromMap(map[int]int{5: 5, 6: 6})}, map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Concat2(tc.iters...)
			actual := ToMap(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name  string
		iters []iter.Seq2[K, V]
	}
	panicTestCases := []panicTestCase[int, int]{
		{"Nil iter1", []iter.Seq2[int, int]{nil, NewFromSlice2([]int{})}},
		{"Nil iter2", []iter.Seq2[int, int]{NewFromSlice2([]int{}), nil}},
		{"Nil iters", []iter.Seq2[int, int]{nil, nil}},
		{"Nil iter in between", []iter.Seq2[int, int]{NewFromSlice2([]int{0: 0, 1: 1}), nil, NewFromSlice2([]int{2: 2, 3: 3})}},
	}
	for _, tc := range panicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.PanicsWithValue(t, nilIterPanicMsg, func() {
				Concat2(tc.iters...)
			})
		})
	}
}

func TestFilter(t *testing.T) {
	type testCase[T any] struct {
		name      string
		iter      iter.Seq[T]
		predicate Predicate[T]
		expected  []T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) bool { return true }, []int{}},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) bool { return true }, []int{}},
		{"Empty slice", NewFromSlice([]int{}), func(v int) bool { return true }, []int{}},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) bool { return true }, []int{}},
		{"All from slice", NewFromSlice([]int{1, 2, 3}), func(v int) bool { return true }, []int{1, 2, 3}},
		{"All from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) bool { return true }, []int{1, 2, 3}},
		{"None from slice", NewFromSlice([]int{1, 2, 3}), func(v int) bool { return false }, []int{}},
		{"None from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) bool { return false }, []int{}},
		{"Even from slice", NewFromSlice([]int{1, 2, 3}), func(v int) bool { return v%2 == 0 }, []int{2}},
		{"Even from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) bool { return v%2 == 0 }, []int{2}},
		{"Odd from slice", NewFromSlice([]int{1, 2, 3}), func(v int) bool { return v%2 != 0 }, []int{1, 3}},
		{"Odd from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) bool { return v%2 != 0 }, []int{1, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Filter(tc.iter, tc.predicate)
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Filter(nil, func(v int) bool { return true })
		})
	})
}

func TestFilter2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		predicate Predicate2[K, V]
		expected  map[K]V
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), func(k int, v int) bool { return true }, map[int]int{}},
		{"Nil map", NewFromMap[int, int](nil), func(k int, v int) bool { return true }, map[int]int{}},
		{"Empty slice", NewFromSlice2([]int{}), func(k int, v int) bool { return true }, map[int]int{}},
		{"Empty map", NewFromMap(map[int]int{}), func(k int, v int) bool { return true }, map[int]int{}},
		{"All from slice", NewFromSlice2([]int{0, 1, 2, 3}), func(k int, v int) bool { return true }, map[int]int{0: 0, 1: 1, 2: 2, 3: 3}},
		{"All from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), func(k int, v int) bool { return true }, map[int]int{0: 0, 1: 1, 2: 2, 3: 3}},
		{"None from slice", NewFromSlice2([]int{0, 1, 2, 3}), func(k int, v int) bool { return false }, map[int]int{}},
		{"None from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), func(k int, v int) bool { return false }, map[int]int{}},
		{"Even from slice", NewFromSlice2([]int{0, 1, 2, 3}), func(k int, v int) bool { return v%2 == 0 }, map[int]int{0: 0, 2: 2}},
		{"Even from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), func(k int, v int) bool { return v%2 == 0 }, map[int]int{0: 0, 2: 2}},
		{"Odd from slice", NewFromSlice2([]int{0, 1, 2, 3}), func(k int, v int) bool { return v%2 != 0 }, map[int]int{1: 1, 3: 3}},
		{"Odd from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), func(k int, v int) bool { return v%2 != 0 }, map[int]int{1: 1, 3: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Filter2(tc.iter, tc.predicate)
			actual := ToMap(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Filter2(nil, func(k int, v int) bool { return true })
		})
	})
}

func TestKeys(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected []K
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), []int{}},
		{"Nil map", NewFromMap[int, int](nil), []int{}},
		{"Empty slice", NewFromSlice2([]int{}), []int{}},
		{"Empty map", NewFromMap(map[int]int{}), []int{}},
		{"All from slice", NewFromSlice2([]int{0, 1, 2, 3}), []int{0, 1, 2, 3}},
		{"All from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), []int{0, 1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Keys(tc.iter)
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Keys[int, int](nil)
		})
	})
}

func TestMap(t *testing.T) {
	type testCase[T, U any] struct {
		name     string
		iter     iter.Seq[T]
		f        func(T) U
		expected []U
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) string { return "" }, []string{}},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) string { return "" }, []string{}},
		{"Empty slice", NewFromSlice([]int{}), func(v int) string { return "" }, []string{}},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) string { return "" }, []string{}},
		{"All from slice", NewFromSlice([]int{1, 2, 3}), func(v int) string { return fmt.Sprint(v) }, []string{"1", "2", "3"}},
		{"All from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) string { return fmt.Sprint(v) }, []string{"1", "2", "3"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Map(tc.iter, tc.f)
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Map(nil, func(v int) string { return "" })
		})
	})
}

func TestMap12(t *testing.T) {
	type testCase[K comparable, V, U any] struct {
		name     string
		iter     iter.Seq[V]
		f        func(V) (K, U)
		expected map[K]U
	}
	testCases := []testCase[int, int, string]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) (int, string) { return 0, "" }, map[int]string{}},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) (int, string) { return 0, "" }, map[int]string{}},
		{"Empty slice", NewFromSlice([]int{}), func(v int) (int, string) { return 0, "" }, map[int]string{}},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) (int, string) { return 0, "" }, map[int]string{}},
		{"All from slice", NewFromSlice([]int{1, 2, 3}), func(v int) (int, string) { return v, fmt.Sprint(v) }, map[int]string{1: "1", 2: "2", 3: "3"}},
		{"All from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) (int, string) { return v, fmt.Sprint(v) }, map[int]string{1: "1", 2: "2", 3: "3"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Map12(tc.iter, tc.f)
			actual := ToMap(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Map12(nil, func(v int) (int, string) { return 0, "" })
		})
	})
}

func TestMap2(t *testing.T) {
	type testCase[K, L comparable, V, U any] struct {
		name     string
		iter     iter.Seq2[K, V]
		f        func(K, V) (L, U)
		expected map[L]U
	}
	testCases := []testCase[int, int, int, string]{
		{"Nil slice", NewFromSlice2[int](nil), func(k int, v int) (int, string) { return 0, "" }, map[int]string{}},
		{"Nil map", NewFromMap[int, int](nil), func(k int, v int) (int, string) { return 0, "" }, map[int]string{}},
		{"Empty slice", NewFromSlice2([]int{}), func(k int, v int) (int, string) { return 0, "" }, map[int]string{}},
		{"Empty map", NewFromMap(map[int]int{}), func(k int, v int) (int, string) { return 0, "" }, map[int]string{}},
		{"All from slice", NewFromSlice2([]int{0, 1, 2, 3}), func(k int, v int) (int, string) { return k, fmt.Sprint(v) }, map[int]string{0: "0", 1: "1", 2: "2", 3: "3"}},
		{"All from map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), func(k int, v int) (int, string) { return k, fmt.Sprint(v) }, map[int]string{1: "1", 2: "2", 3: "3"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Map2(tc.iter, tc.f)
			actual := ToMap(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Map2(nil, func(k int, v int) (int, string) { return 0, "" })
		})
	})
}

func TestMap21(t *testing.T) {
	type testCase[K comparable, V, U any] struct {
		name     string
		iter     iter.Seq2[K, V]
		f        func(K, V) U
		expected []U
	}
	testCases := []testCase[int, int, string]{
		{"Nil slice", NewFromSlice2[int](nil), func(k int, v int) string { return "" }, []string{}},
		{"Nil map", NewFromMap[int, int](nil), func(k int, v int) string { return "" }, []string{}},
		{"Empty slice", NewFromSlice2([]int{}), func(k int, v int) string { return "" }, []string{}},
		{"Empty map", NewFromMap(map[int]int{}), func(k int, v int) string { return "" }, []string{}},
		{"All from slice", NewFromSlice2([]int{0, 1, 2, 3}), func(k int, v int) string { return fmt.Sprintf("%d:%d", k, v) }, []string{"0:0", "1:1", "2:2", "3:3"}},
		{"All from map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), func(k int, v int) string { return fmt.Sprintf("%d:%d", k, v) }, []string{"1:1", "2:2", "3:3"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Map21(tc.iter, tc.f)
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Map21(nil, func(k int, v int) string { return "" })
		})
	})
}

func TestRemoveKey(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		key      K
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), 0, map[int]int{}},
		{"Nil map", NewFromMap[int, int](nil), 0, map[int]int{}},
		{"Empty slice", NewFromSlice2([]int{}), 0, map[int]int{}},
		{"Empty map", NewFromMap(map[int]int{}), 0, map[int]int{}},
		{"Remove 0 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 0, map[int]int{1: 1, 2: 2, 3: 3}},
		{"Remove 0 from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), 0, map[int]int{1: 1, 2: 2, 3: 3}},
		{"Remove 1 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 1, map[int]int{0: 0, 2: 2, 3: 3}},
		{"Remove 1 from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), 1, map[int]int{0: 0, 2: 2, 3: 3}},
		{"Remove 2 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 2, map[int]int{0: 0, 1: 1, 3: 3}},
		{"Remove 2 from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), 2, map[int]int{0: 0, 1: 1, 3: 3}},
		{"Remove 3 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 3, map[int]int{0: 0, 1: 1, 2: 2}},
		{"Remove 3 from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), 3, map[int]int{0: 0, 1: 1, 2: 2}},
		{"Remove 4 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 4, map[int]int{0: 0, 1: 1, 2: 2, 3: 3}},
		{"Remove 4 from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), 4, map[int]int{0: 0, 1: 1, 2: 2, 3: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RemoveKey(tc.iter, tc.key)
			actual := ToMap(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			RemoveKey[int, int](nil, 0)
		})
	})
}

func TestRemoveNth(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		index    int
		expected []T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), 0, []int{}},
		{"Empty slice", NewFromSlice([]int{}), 0, []int{}},
		{"Remove 0 from slice", NewFromSlice([]int{1, 2, 3}), 0, []int{2, 3}},
		{"Remove 1 from slice", NewFromSlice([]int{1, 2, 3}), 1, []int{1, 3}},
		{"Remove 2 from slice", NewFromSlice([]int{1, 2, 3}), 2, []int{1, 2}},
		{"Remove 3 from slice", NewFromSlice([]int{1, 2, 3}), 3, []int{1, 2, 3}},
		{"Remove 4 from slice", NewFromSlice([]int{1, 2, 3}), 4, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RemoveNth(tc.iter, tc.index)
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			RemoveNth[int](nil, 0)
		})
	})
}

func TestSkip(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		n        int
		expected []T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), 0, []int{}},
		{"Empty slice", NewFromSlice([]int{}), 0, []int{}},
		{"Skip 0 from slice", NewFromSlice([]int{1, 2, 3}), 0, []int{1, 2, 3}},
		{"Skip 1 from slice", NewFromSlice([]int{1, 2, 3}), 1, []int{2, 3}},
		{"Skip 2 from slice", NewFromSlice([]int{1, 2, 3}), 2, []int{3}},
		{"Skip 3 from slice", NewFromSlice([]int{1, 2, 3}), 3, []int{}},
		{"Skip 4 from slice", NewFromSlice([]int{1, 2, 3}), 4, []int{}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToSlice(Skip(tc.iter, tc.n))
			assert.ElementsMatch(t, tc.expected, result)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Skip[int](nil, 0)
		})
	})
}

func TestSkip2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		n        int
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), 0, map[int]int{}},
		{"Empty slice", NewFromSlice2([]int{}), 0, map[int]int{}},
		{"Skip 0 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 0, map[int]int{0: 0, 1: 1, 2: 2, 3: 3}},
		{"Skip 1 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 1, map[int]int{1: 1, 2: 2, 3: 3}},
		{"Skip 2 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 2, map[int]int{2: 2, 3: 3}},
		{"Skip 3 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 3, map[int]int{3: 3}},
		{"Skip 4 from slice", NewFromSlice2([]int{0, 1, 2, 3}), 4, map[int]int{}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToMap(Skip2(tc.iter, tc.n))
			assert.Equal(t, tc.expected, result)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Skip2[int, int](nil, 0)
		})
	})
}

func TestSort(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		less     func(T, T) bool
		expected []T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(a, b int) bool { return a < b }, []int{}},
		{"Empty slice", NewFromSlice([]int{}), func(a, b int) bool { return a < b }, []int{}},
		{"Sort slice", NewFromSlice([]int{3, 2, 1}), func(a, b int) bool { return a < b }, []int{1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sort(tc.iter, tc.less)
			actual := ToSlice(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Sort(nil, func(a, b int) bool { return a < b })
		})
	})
}

func TestSort2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name           string
		iter           iter.Seq2[K, V]
		less           func(K, V, K, V) bool
		expectedKeys   []K
		expectedValues []V
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromMap[int, int](nil), func(a int, b int, c int, d int) bool { return a < c }, []int{}, []int{}},
		{"Empty slice", NewFromMap(map[int]int{}), func(a int, b int, c int, d int) bool { return a < c }, []int{}, []int{}},
		{"Sort slice", NewFromMap(map[int]int{5: 2, 2: 0, 1: 3}), func(a int, b int, c int, d int) bool { return a+b < c+d }, []int{2, 1, 5}, []int{0, 3, 2}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sort2(tc.iter, tc.less)
			actualKeys := ToSlice(Keys(result))
			actualValues := ToSlice(Values(result))
			assert.Equal(t, tc.expectedKeys, actualKeys)
			assert.Equal(t, tc.expectedValues, actualValues)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Sort2(nil, func(a int, b int, c int, d int) bool { return a < c })
		})
	})
}

func TestSplit(t *testing.T) {
	type testCase[T any] struct {
		name      string
		iter      iter.Seq[T]
		predicate func(T) bool
		expected  [][]T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) bool { return true }, [][]int{}},
		{"Empty slice", NewFromSlice([]int{}), func(v int) bool { return true }, [][]int{}},
		{"Split slice", NewFromSlice([]int{1, 2, 3, 4, 5, 6}), func(v int) bool { return v%2 == 0 }, [][]int{{1}, {3}, {5}}},
		{"Split slice with trailing", NewFromSlice([]int{1, 2, 3, 4, 5}), func(v int) bool { return v%2 == 0 }, [][]int{{1}, {3}, {5}}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultIter := Split(tc.iter, tc.predicate)
			result := ToSlice(resultIter)
			actual := make([][]int, len(result))
			for i, iter := range result {
				actual[i] = ToSlice(iter)
			}
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(resultIter, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Split(nil, func(v int) bool { return true })
		})
	})
}

func TestSplit2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		predicate func(K, V) bool
		expected  []map[K]V
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), func(k int, v int) bool { return true }, []map[int]int{}},
		{"Empty slice", NewFromSlice2([]int{}), func(k int, v int) bool { return true }, []map[int]int{}},
		{"Split slice", NewFromSlice2([]int{0, 1, 2, 3, 4, 5, 6}), func(k int, v int) bool { return v%2 == 0 }, []map[int]int{{1: 1}, {3: 3}, {5: 5}}},
		{"Split slice with trailing", NewFromSlice2([]int{0, 1, 2, 3, 4, 5}), func(k int, v int) bool { return v%2 == 0 }, []map[int]int{{1: 1}, {3: 3}, {5: 5}}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultIter := Split2(tc.iter, tc.predicate)
			result := ToSlice(resultIter)
			actual := make([]map[int]int, len(result))
			for i, iter := range result {
				actual[i] = ToMap(iter)
			}
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(resultIter, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Split2(nil, func(k int, v int) bool { return true })
		})
	})
}

func TestSwap(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		i        int
		j        int
		expected []T
	}
	testCases := []testCase[int]{
		{"Swap 0 and 1 in slice", NewFromSlice([]int{1, 2, 3}), 0, 1, []int{2, 1, 3}},
		{"Swap 1 and 2 in slice", NewFromSlice([]int{1, 2, 3}), 1, 2, []int{1, 3, 2}},
		{"Swap 0 and 2 in slice", NewFromSlice([]int{1, 2, 3}), 0, 2, []int{3, 2, 1}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Swap(tc.iter, tc.i, tc.j)
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name        string
		iter        iter.Seq[T]
		i           int
		j           int
		expectedMsg string
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter", nil, 0, 1, nilIterPanicMsg},
		{"Nil slice", NewFromSlice[int](nil), 0, 1, indexOOBPanicMsg},
		{"Empty slice", NewFromSlice([]int{}), 0, 1, indexOOBPanicMsg},
		{"Negative i", NewFromSlice([]int{1, 2, 3}), -1, 1, indexOOBPanicMsg},
		{"Negative j", NewFromSlice([]int{1, 2, 3}), 0, -1, indexOOBPanicMsg},
		{"i out of bounds", NewFromSlice([]int{1, 2, 3}), 3, 1, indexOOBPanicMsg},
		{"j out of bounds", NewFromSlice([]int{1, 2, 3}), 0, 3, indexOOBPanicMsg},
	}
	for _, tc := range panicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.PanicsWithValue(t, tc.expectedMsg, func() {
				Swap(tc.iter, tc.i, tc.j)
			})
		})
	}
}

func TestSwap2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		i        int
		j        int
		expected map[K]V
	}
	testCases := []testCase[int, int]{
		{"Swap 0 and 1 in slice", NewFromSlice2([]int{0, 1, 2, 3}), 0, 1, map[int]int{0: 1, 1: 0, 2: 2, 3: 3}},
		{"Swap 1 and 2 in slice", NewFromSlice2([]int{0, 1, 2, 3}), 1, 2, map[int]int{0: 0, 1: 2, 2: 1, 3: 3}},
		{"Swap 0 and 2 in slice", NewFromSlice2([]int{0, 1, 2, 3}), 0, 2, map[int]int{0: 2, 1: 1, 2: 0, 3: 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Swap2(tc.iter, tc.i, tc.j)
			actual := ToMap(result)
			assert.Equal(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name        string
		iter        iter.Seq2[K, V]
		i           int
		j           int
		expectedMsg string
	}
	panicTestCases := []panicTestCase[int, int]{
		{"Nil iter", nil, 0, 1, nilIterPanicMsg},
		{"Nil slice", NewFromSlice2[int](nil), 0, 1, indexOOBPanicMsg},
		{"Empty slice", NewFromSlice2([]int{}), 0, 1, indexOOBPanicMsg},
		{"Negative i", NewFromSlice2([]int{0, 1, 2, 3}), -1, 1, indexOOBPanicMsg},
		{"Negative j", NewFromSlice2([]int{0, 1, 2, 3}), 0, -1, indexOOBPanicMsg},
		{"i out of bounds", NewFromSlice2([]int{0, 1, 2, 3}), 4, 1, indexOOBPanicMsg},
		{"j out of bounds", NewFromSlice2([]int{0, 1, 2, 3}), 0, 4, indexOOBPanicMsg},
	}
	for _, tc := range panicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.PanicsWithValue(t, tc.expectedMsg, func() {
				Swap2(tc.iter, tc.i, tc.j)
			})
		})
	}
}

func TestSwapByKey(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name           string
		iter           iter.Seq2[K, V]
		key1           K
		key2           K
		expectedKeys   []K
		expectedValues []V
	}
	testCases := []testCase[int, int]{
		{"Swap 0 and 1 in slice", NewFromSlice2([]int{0, 1, 2, 3}), 0, 1, []int{1, 0, 2, 3}, []int{1, 0, 2, 3}},
		{"Swap 1 and 2 in slice", NewFromSlice2([]int{0, 1, 2, 3}), 1, 2, []int{0, 2, 1, 3}, []int{0, 2, 1, 3}},
		{"Swap 0 and 2 in slice", NewFromSlice2([]int{0, 1, 2, 3}), 0, 2, []int{2, 1, 0, 3}, []int{2, 1, 0, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SwapByKey(tc.iter, tc.key1, tc.key2)
			actualKeys := ToSlice(Keys(result))
			actualValues := ToSlice(Values(result))
			assert.Equal(t, tc.expectedKeys, actualKeys)
			assert.Equal(t, tc.expectedValues, actualValues)
			// Coverage for the case where the iterator is used again
			ToMapN(result, 0)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name        string
		iter        iter.Seq2[K, V]
		key1        K
		key2        K
		expectedMsg string
	}
	panicTestCases := []panicTestCase[int, int]{
		{"Nil iter", nil, 0, 1, nilIterPanicMsg},
		{"Nil slice", NewFromSlice2[int](nil), 0, 1, keyNotFoundPanicMsg},
		{"Empty slice", NewFromSlice2([]int{}), 0, 1, keyNotFoundPanicMsg},
		{"Key 1 not found", NewFromSlice2([]int{0, 1, 2, 3}), 4, 1, keyNotFoundPanicMsg},
		{"Key 2 not found", NewFromSlice2([]int{0, 1, 2, 3}), 0, 4, keyNotFoundPanicMsg},
	}
	for _, tc := range panicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.PanicsWithValue(t, tc.expectedMsg, func() {
				SwapByKey(tc.iter, tc.key1, tc.key2)
			})
		})
	}
}

func TestValues(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected []V
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), []int{}},
		{"Nil map", NewFromMap[int, int](nil), []int{}},
		{"Empty slice", NewFromSlice2([]int{}), []int{}},
		{"Empty map", NewFromMap(map[int]int{}), []int{}},
		{"All from slice", NewFromSlice2([]int{0, 1, 2, 3}), []int{0, 1, 2, 3}},
		{"All from map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), []int{0, 1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Values(tc.iter)
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Values[int, int](nil)
		})
	})
}

func TestZip(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter1    iter.Seq[T]
		iter2    iter.Seq[T]
		expected []T
	}
	testCases := []testCase[int]{
		{"Nil slices", NewFromSlice[int](nil), NewFromSlice[int](nil), []int{}},
		{"Empty slices", NewFromSlice([]int{}), NewFromSlice([]int{}), []int{}},
		{"Zip slices", NewFromSlice([]int{1, 2, 3}), NewFromSlice([]int{4, 5, 6}), []int{5, 7, 9}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Zip(tc.iter1, tc.iter2, func(a, b int) int { return a + b })
			actual := ToSlice(result)
			assert.ElementsMatch(t, tc.expected, actual)
			// Coverage for the case where the iterator is used again
			ToSliceN(result, 0)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name        string
		iter1       iter.Seq[T]
		iter2       iter.Seq[T]
		expectedMsg string
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter1", nil, NewFromSlice[int](nil), nilIterPanicMsg},
		{"Nil iter2", NewFromSlice[int](nil), nil, nilIterPanicMsg},
	}
	for _, tc := range panicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.PanicsWithValue(t, tc.expectedMsg, func() {
				Zip(tc.iter1, tc.iter2, func(a, b int) int { return a + b })
			})
		})
	}
}
