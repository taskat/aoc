package iterutils

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	type testCase[T any] struct {
		name      string
		iter      iter.Seq[T]
		predicate Predicate[T]
		expected  bool
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) bool { return v > 0 }, true},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) bool { return v > 0 }, true},
		{"Empty slice", NewFromSlice([]int{}), func(v int) bool { return v > 0 }, true},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) bool { return v > 0 }, true},
		{"True with slice", NewFromSlice([]int{1, 2, 3}), func(v int) bool { return v > 0 }, true},
		{"True with map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) bool { return v > 0 }, true},
		{"False with slice", NewFromSlice([]int{1, -2, 3}), func(v int) bool { return v > 0 }, false},
		{"False with map", Values(NewFromMap(map[int]int{1: 1, -2: -2, 3: 3})), func(v int) bool { return v > 0 }, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := All(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			All(nil, func(int) bool { return true })
		})
	})
}

func TestAll2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		predicate Predicate2[K, V]
		expected  bool
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice2[string](nil), func(k int, v string) bool { return k > 0 }, true},
		{"Nil map", NewFromMap[int, string](nil), func(k int, v string) bool { return k > 0 }, true},
		{"Empty slice", NewFromSlice2([]string{}), func(k int, v string) bool { return k > 0 }, true},
		{"Empty map", NewFromMap(map[int]string{}), func(k int, v string) bool { return k > 0 }, true},
		{"True with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) > k }, true},
		{"True with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) > k }, true},
		{"False with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) < k }, false},
		{"False with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) < k }, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := All2(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			All2(nil, func(int, string) bool { return true })
		})
	})
}

func TestAny(t *testing.T) {
	type testCase[T any] struct {
		name      string
		iter      iter.Seq[T]
		predicate Predicate[T]
		expected  bool
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) bool { return v > 0 }, false},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) bool { return v > 0 }, false},
		{"Empty slice", NewFromSlice([]int{}), func(v int) bool { return v > 0 }, false},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) bool { return v > 0 }, false},
		{"True with slice", NewFromSlice([]int{1, -2, 3}), func(v int) bool { return v > 0 }, true},
		{"True with map", Values(NewFromMap(map[int]int{1: 1, -2: -2, 3: 3})), func(v int) bool { return v > 0 }, true},
		{"False with slice", NewFromSlice([]int{-1, -2, -3}), func(v int) bool { return v > 0 }, false},
		{"False with map", Values(NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3})), func(v int) bool { return v > 0 }, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Any(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Any(nil, func(int) bool { return true })
		})
	})
}

func TestAny2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		predicate Predicate2[K, V]
		expected  bool
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice2[string](nil), func(k int, v string) bool { return k > 0 }, false},
		{"Nil map", NewFromMap[int, string](nil), func(k int, v string) bool { return k > 0 }, false},
		{"Empty slice", NewFromSlice2([]string{}), func(k int, v string) bool { return k > 0 }, false},
		{"Empty map", NewFromMap(map[int]string{}), func(k int, v string) bool { return k > 0 }, false},
		{"True with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) > k }, true},
		{"True with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) > k }, true},
		{"False with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) < k }, false},
		{"False with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) < k }, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Any2(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Any2(nil, func(int, string) bool { return true })
		})
	})
}

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		iter     iter.Seq[T]
		item     T
		expected bool
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), 1, false},
		{"Nil map", Values(NewFromMap[int, int](nil)), 1, false},
		{"Empty slice", NewFromSlice([]int{}), 1, false},
		{"Empty map", Values(NewFromMap(map[int]int{})), 1, false},
		{"True with slice", NewFromSlice([]int{1, 2, 3}), 2, true},
		{"True with map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), 2, true},
		{"False with slice", NewFromSlice([]int{1, 2, 3}), 4, false},
		{"False with map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), 4, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Contains(testCase.iter, testCase.item)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Contains(nil, 1)
		})
	})
}

func TestCount(t *testing.T) {
	type testCase[T any] struct {
		name      string
		iter      iter.Seq[T]
		predicate Predicate[T]
		expected  int
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) bool { return v > 0 }, 0},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) bool { return v > 0 }, 0},
		{"Empty slice", NewFromSlice([]int{}), func(v int) bool { return v > 0 }, 0},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) bool { return v > 0 }, 0},
		{"True with slice", NewFromSlice([]int{1, -2, 3}), func(v int) bool { return v > 0 }, 2},
		{"True with map", Values(NewFromMap(map[int]int{1: 1, -2: -2, 3: 3})), func(v int) bool { return v > 0 }, 2},
		{"False with slice", NewFromSlice([]int{-1, -2, -3}), func(v int) bool { return v > 0 }, 0},
		{"False with map", Values(NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3})), func(v int) bool { return v > 0 }, 0},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Count(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Count(nil, func(int) bool { return true })
		})
	})
}

func TestCount2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		predicate Predicate2[K, V]
		expected  int
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice2[string](nil), func(k int, v string) bool { return k > 0 }, 0},
		{"Nil map", NewFromMap[int, string](nil), func(k int, v string) bool { return k > 0 }, 0},
		{"Empty slice", NewFromSlice2([]string{}), func(k int, v string) bool { return k > 0 }, 0},
		{"Empty map", NewFromMap(map[int]string{}), func(k int, v string) bool { return k > 0 }, 0},
		{"True with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) > k }, 3},
		{"True with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) > k }, 3},
		{"False with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) < k }, 0},
		{"False with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) < k }, 0},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Count2(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Count2(nil, func(int, string) bool { return true })
		})
	})
}

func TestEqual(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		iter1    iter.Seq[T]
		iter2    iter.Seq[T]
		expected bool
	}
	testCases := []testCase[int]{
		{"Empty slices", NewFromSlice([]int{}), NewFromSlice([]int{}), true},
		{"Different lengths", NewFromSlice([]int{1, 2, 3}), NewFromSlice([]int{1, 2}), false},
		{"Different values", NewFromSlice([]int{1, 2, 3}), NewFromSlice([]int{1, 3, 2}), false},
		{"Same values", NewFromSlice([]int{1, 2, 3}), NewFromSlice([]int{1, 2, 3}), true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Equal(testCase.iter1, testCase.iter2)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	testCases = []testCase[int]{
		{"Nil iter1", nil, NewFromSlice([]int{1, 2, 3}), false},
		{"Nil iter2", NewFromSlice([]int{1, 2, 3}), nil, false},
		{"Both nil", nil, nil, true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, nilIterPanicMsg, func() {
				Equal(testCase.iter1, testCase.iter2)
			})
		})
	}
}

func TestEqual2(t *testing.T) {
	type testCase[K comparable, V comparable] struct {
		name     string
		iter1    iter.Seq2[K, V]
		iter2    iter.Seq2[K, V]
		expected bool
	}
	testCases := []testCase[int, string]{
		{"Empty slices", NewFromSlice2([]string{}), NewFromSlice2([]string{}), true},
		{"Different lengths", NewFromSlice2([]string{"a", "bb", "ccc"}), NewFromSlice2([]string{"a", "bb"}), false},
		{"Different values", NewFromSlice2([]string{"a", "bb", "ccc"}), NewFromSlice2([]string{"a", "ccc", "bb"}), false},
		{"Same values", NewFromSlice2([]string{"a", "bb", "ccc"}), NewFromSlice2([]string{"a", "bb", "ccc"}), true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Equal2(testCase.iter1, testCase.iter2)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	testCases = []testCase[int, string]{
		{"Nil iter1", nil, NewFromSlice2([]string{"a", "bb", "ccc"}), false},
		{"Nil iter2", NewFromSlice2([]string{"a", "bb", "ccc"}), nil, false},
		{"Both nil", nil, nil, true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, nilIterPanicMsg, func() {
				Equal2(testCase.iter1, testCase.iter2)
			})
		})
	}
}

func TestFind(t *testing.T) {
	type testCase[T any] struct {
		name      string
		iter      iter.Seq[T]
		predicate Predicate[T]
		expected  T
		found     bool
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) bool { return v > 0 }, 0, false},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) bool { return v > 0 }, 0, false},
		{"Empty slice", NewFromSlice([]int{}), func(v int) bool { return v > 0 }, 0, false},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) bool { return v > 0 }, 0, false},
		{"Found with slice", NewFromSlice([]int{1, -2, 3}), func(v int) bool { return v > 0 }, 1, true},
		{"Found with map", Values(NewFromMap(map[int]int{1: 1, -2: -2, 3: -3})), func(v int) bool { return v > 0 }, 1, true},
		{"Not found with slice", NewFromSlice([]int{-1, -2, -3}), func(v int) bool { return v > 0 }, 0, false},
		{"Not found with map", Values(NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3})), func(v int) bool { return v > 0 }, 0, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, found := Find(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
			assert.Equal(t, testCase.found, found)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Find(nil, func(int) bool { return true })
		})
	})
}

func TestFind2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		predicate Predicate2[K, V]
		expectedK K
		expectedV V
		found     bool
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice2[string](nil), func(k int, v string) bool { return k > 0 }, 0, "", false},
		{"Nil map", NewFromMap[int, string](nil), func(k int, v string) bool { return k > 0 }, 0, "", false},
		{"Empty slice", NewFromSlice2([]string{}), func(k int, v string) bool { return k > 0 }, 0, "", false},
		{"Empty map", NewFromMap(map[int]string{}), func(k int, v string) bool { return k > 0 }, 0, "", false},
		{"Found with slice", NewFromSlice2([]string{"a", "b", "ccc"}), func(k int, v string) bool { return len(v) == k }, 1, "b", true},
		{"Found with map", NewFromMap(map[int]string{0: "a", 1: "b", 2: "ccc"}), func(k int, v string) bool { return len(v) == k }, 1, "b", true},
		{"Not found with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) < k }, 0, "", false},
		{"Not found with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) < k }, 0, "", false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualK, actualV, found := Find2(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expectedK, actualK)
			assert.Equal(t, testCase.expectedV, actualV)
			assert.Equal(t, testCase.found, found)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Find2(nil, func(int, string) bool { return true })
		})
	})
}

func TestFirst(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected T
	}
	testCases := []testCase[int]{
		{"First element", NewFromSlice([]int{1, 2, 3}), 1},
		{"Only element", NewFromSlice([]int{1}), 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := First(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name             string
		iter             iter.Seq[T]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				First(testCase.iter)
			})
		})
	}
}

func TestFirst2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		expectedK K
		expectedV V
	}
	testCases := []testCase[int, string]{
		{"First element", NewFromSlice2([]string{"a", "bb", "ccc"}), 0, "a"},
		{"Only element", NewFromSlice2([]string{"a"}), 0, "a"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualK, actualV := First2(testCase.iter)
			assert.Equal(t, testCase.expectedK, actualK)
			assert.Equal(t, testCase.expectedV, actualV)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name             string
		iter             iter.Seq2[K, V]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int, string]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice2[string](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice2([]string{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				First2(testCase.iter)
			})
		})
	}
}

func TestForEach(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected []T
	}
	testCases := []testCase[[]int]{
		{"Empty slice", NewFromSlice([][]int{}), [][]int{}},
		{"Non-empty slice", NewFromSlice([][]int{{1}, {2}, {3}}), [][]int{{2}, {3}, {4}}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ForEach(testCase.iter, func(slice []int) { slice[0]++ })
			assert.Equal(t, testCase.expected, ToSlice(testCase.iter))
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			ForEach(nil, func(int) {})
		})
	})
}

func TestForEach2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected map[K]V
	}
	testCases := []testCase[int, []int]{
		{"Empty slice", NewFromSlice2([][]int{}), map[int][]int{}},
		{"Non-empty slice", NewFromSlice2([][]int{{1}, {2}, {3}}), map[int][]int{0: {1}, 1: {3}, 2: {5}}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ForEach2(testCase.iter, func(k int, slice []int) { slice[0] += k })
			assert.Equal(t, testCase.expected, ToMap(testCase.iter))
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			ForEach2(nil, func(int, int) {})
		})
	})
}

func TestIsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected bool
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), true},
		{"Nil map", Values(NewFromMap[int, int](nil)), true},
		{"Empty slice", NewFromSlice([]int{}), true},
		{"Empty map", Values(NewFromMap(map[int]int{})), true},
		{"Non-empty slice", NewFromSlice([]int{1, 2, 3}), false},
		{"Non-empty map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsEmpty(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			IsEmpty[int](nil)
		})
	})
}

func TestIsEmpty2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected bool
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice2[string](nil), true},
		{"Nil map", NewFromMap[int, string](nil), true},
		{"Empty slice", NewFromSlice2([]string{}), true},
		{"Empty map", NewFromMap(map[int]string{}), true},
		{"Non-empty slice", NewFromSlice2([]string{"a", "bb", "ccc"}), false},
		{"Non-empty map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsEmpty2(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			IsEmpty2[int, string](nil)
		})
	})
}

func TestIsValidKey(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		key      K
		expected bool
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2[int](nil), 1, false},
		{"Nil map", NewFromMap[int, int](nil), 1, false},
		{"Empty slice", NewFromSlice2([]int{}), 1, false},
		{"Empty map", NewFromMap(map[int]int{}), 1, false},
		{"True with slice", NewFromSlice2([]int{1, 2, 3}), 2, true},
		{"True with map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), 2, true},
		{"False with slice", NewFromSlice2([]int{1, 2, 3}), 4, false},
		{"False with map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), 4, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsValidKey(testCase.iter, testCase.key)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			IsValidKey[int, int](nil, 1)
		})
	})
}

func TestLast(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Last element", NewFromSlice([]int{1, 2, 3}), 3},
		{"Only element", NewFromSlice([]int{1}), 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Last(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name             string
		iter             iter.Seq[T]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Last(testCase.iter)
			})
		})
	}
}

func TestLast2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		expectedK K
		expectedV V
	}
	testCases := []testCase[int, string]{
		{"Last element", NewFromSlice2([]string{"a", "bb", "ccc"}), 2, "ccc"},
		{"Only element", NewFromSlice2([]string{"a"}), 0, "a"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualK, actualV := Last2(testCase.iter)
			assert.Equal(t, testCase.expectedK, actualK)
			assert.Equal(t, testCase.expectedV, actualV)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name             string
		iter             iter.Seq2[K, V]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int, string]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice2[string](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice2([]string{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Last2(testCase.iter)
			})
		})
	}
}

func TestLen(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected int
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), 0},
		{"Nil map", Values(NewFromMap[int, int](nil)), 0},
		{"Empty slice", NewFromSlice([]int{}), 0},
		{"Empty map", Values(NewFromMap(map[int]int{})), 0},
		{"Non-empty slice", NewFromSlice([]int{1, 2, 3}), 3},
		{"Non-empty map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), 3},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Len(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Len[int](nil)
		})
	})
}

func TestLen2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected int
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice2[string](nil), 0},
		{"Nil map", NewFromMap[int, string](nil), 0},
		{"Empty slice", NewFromSlice2([]string{}), 0},
		{"Empty map", NewFromMap(map[int]string{}), 0},
		{"Non-empty slice", NewFromSlice2([]string{"a", "bb", "ccc"}), 3},
		{"Non-empty map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), 3},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Len2(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Len2[int, string](nil)
		})
	})
}

func TestMax(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Max element from slice", NewFromSlice([]int{1, 2, 3}), 3},
		{"Max element from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), 3},
		{"Only element from slice", NewFromSlice([]int{1}), 1},
		{"Only element from map", Values(NewFromMap(map[int]int{1: 1})), 1},
		{"Negative max from slice", NewFromSlice([]int{-1, -2, -3}), -1},
		{"Negative max from map", Values(NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3})), -1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Max(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name             string
		iter             iter.Seq[T]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Max(testCase.iter)
			})
		})
	}
}

func TestMax2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		expectedK K
		expectedV V
	}
	testCases := []testCase[int, int]{
		{"Max element from slice", NewFromSlice2([]int{1, 2, 3}), 2, 3},
		{"Max element from map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), 3, 3},
		{"Only element from slice", NewFromSlice2([]int{1}), 0, 1},
		{"Only element from map", NewFromMap(map[int]int{1: 1}), 1, 1},
		{"Negative max from slice", NewFromSlice2([]int{-1, -2, -3}), 0, -1},
		{"Negative max from map", NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3}), -1, -1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualK, actualV := Max2(testCase.iter)
			assert.Equal(t, testCase.expectedK, actualK)
			assert.Equal(t, testCase.expectedV, actualV)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name             string
		iter             iter.Seq2[K, V]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int, int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice2[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice2([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Max2(testCase.iter)
			})
		})
	}
}

func TestMiddle(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Middle element from ood slice", NewFromSlice([]int{1, 2, 3}), 2},
		{"Middle element from even slice", NewFromSlice([]int{1, 2, 3, 4}), 3},
		{"Only element from slice", NewFromSlice([]int{1}), 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Middle(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name             string
		iter             iter.Seq[T]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Middle(testCase.iter)
			})
		})
	}
}

func TestMiddle2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		expectedK K
		expectedV V
	}
	testCases := []testCase[int, int]{
		{"Middle element from odd slice", NewFromSlice2([]int{1, 2, 3}), 1, 2},
		{"Middle element from even slice", NewFromSlice2([]int{1, 2, 3, 4}), 2, 3},
		{"Only element from slice", NewFromSlice2([]int{1}), 0, 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualK, actualV := Middle2(testCase.iter)
			assert.Equal(t, testCase.expectedK, actualK)
			assert.Equal(t, testCase.expectedV, actualV)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name             string
		iter             iter.Seq2[K, V]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int, int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice2[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice2([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Middle2(testCase.iter)
			})
		})
	}
}

func TestMin(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Min element from slice", NewFromSlice([]int{1, 2, 3}), 1},
		{"Min element from map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), 1},
		{"Only element from slice", NewFromSlice([]int{1}), 1},
		{"Only element from map", Values(NewFromMap(map[int]int{1: 1})), 1},
		{"Negative min from slice", NewFromSlice([]int{-1, -2, -3}), -3},
		{"Negative min from map", Values(NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3})), -3},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Min(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	type panicTestCase[T any] struct {
		name             string
		iter             iter.Seq[T]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Min(testCase.iter)
			})
		})
	}
}

func TestMin2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		expectedK K
		expectedV V
	}
	testCases := []testCase[int, int]{
		{"Min element from slice", NewFromSlice2([]int{1, 2, 3}), 0, 1},
		{"Min element from map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), 1, 1},
		{"Only element from slice", NewFromSlice2([]int{1}), 0, 1},
		{"Only element from map", NewFromMap(map[int]int{1: 1}), 1, 1},
		{"Negative min from slice", NewFromSlice2([]int{-1, -2, -3}), 2, -3},
		{"Negative min from map", NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3}), -3, -3},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualK, actualV := Min2(testCase.iter)
			assert.Equal(t, testCase.expectedK, actualK)
			assert.Equal(t, testCase.expectedV, actualV)
		})
	}
	// Test for panic
	type panicTestCase[K comparable, V any] struct {
		name             string
		iter             iter.Seq2[K, V]
		expectedPanicMsg string
	}
	panicTestCases := []panicTestCase[int, int]{
		{"Nil iter", nil, nilIterPanicMsg},
		{"Nil slice", NewFromSlice2[int](nil), emptyIterPanicMsg},
		{"Empty slice", NewFromSlice2([]int{}), emptyIterPanicMsg},
	}
	for _, testCase := range panicTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedPanicMsg, func() {
				Min2(testCase.iter)
			})
		})
	}
}

func TestNone(t *testing.T) {
	type testCase[T any] struct {
		name      string
		iter      iter.Seq[T]
		predicate Predicate[T]
		expected  bool
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(v int) bool { return v > 0 }, true},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(v int) bool { return v > 0 }, true},
		{"Empty slice", NewFromSlice([]int{}), func(v int) bool { return v > 0 }, true},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(v int) bool { return v > 0 }, true},
		{"True with slice", NewFromSlice([]int{1, -2, 3}), func(v int) bool { return v < 0 }, false},
		{"True with map", Values(NewFromMap(map[int]int{1: 1, -2: -2, 3: 3})), func(v int) bool { return v < 0 }, false},
		{"False with slice", NewFromSlice([]int{1, 2, 3}), func(v int) bool { return v < 0 }, true},
		{"False with map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(v int) bool { return v < 0 }, true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := None(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			None(nil, func(int) bool { return true })
		})
	})
}

func TestNone2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name      string
		iter      iter.Seq2[K, V]
		predicate Predicate2[K, V]
		expected  bool
	}
	testCases := []testCase[int, string]{
		{"Nil slice", NewFromSlice2[string](nil), func(k int, v string) bool { return k > 0 }, true},
		{"Nil map", NewFromMap[int, string](nil), func(k int, v string) bool { return k > 0 }, true},
		{"Empty slice", NewFromSlice2([]string{}), func(k int, v string) bool { return k > 0 }, true},
		{"Empty map", NewFromMap(map[int]string{}), func(k int, v string) bool { return k > 0 }, true},
		{"True with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) == k }, true},
		{"True with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) == k }, true},
		{"False with slice", NewFromSlice2([]string{"a", "bb", "ccc"}), func(k int, v string) bool { return len(v) == k+1 }, false},
		{"False with map", NewFromMap(map[int]string{0: "a", 1: "bb", 2: "ccc"}), func(k int, v string) bool { return len(v) == k+1 }, false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := None2(testCase.iter, testCase.predicate)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			None2(nil, func(int, string) bool { return true })
		})
	})
}

func TestProduct(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), 1},
		{"Nil map", Values(NewFromMap[int, int](nil)), 1},
		{"Empty slice", NewFromSlice([]int{}), 1},
		{"Empty map", Values(NewFromMap(map[int]int{})), 1},
		{"Product of slice", NewFromSlice([]int{1, 2, 3}), 6},
		{"Product of map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), 6},
		{"Only element from slice", NewFromSlice([]int{1}), 1},
		{"Only element from map", Values(NewFromMap(map[int]int{1: 1})), 1},
		{"Negative product from slice", NewFromSlice([]int{-1, -2, -3}), -6},
		{"Negative product from map", Values(NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3})), -6},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Product(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Product[int](nil)
		})
	})
}

func TestProduct2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected int
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2([]int{}), 1},
		{"Nil map", NewFromMap(map[int]int{}), 1},
		{"Empty slice", NewFromSlice2([]int{}), 1},
		{"Empty map", NewFromMap(map[int]int{}), 1},
		{"Product of slice", NewFromSlice2([]int{1, 2, 3}), 6},
		{"Product of map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), 6},
		{"Only element from slice", NewFromSlice2([]int{1}), 1},
		{"Only element from map", NewFromMap(map[int]int{1: 1}), 1},
		{"Negative product from slice", NewFromSlice2([]int{-1, -2, -3}), -6},
		{"Negative product from map", NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3}), -6},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Product2(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Product2[int, int](nil)
		})
	})
}

func TestReduce(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		f        func(a, b int) int
		initial  int
		expected T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), func(a, b int) int { return a + b }, 0, 0},
		{"Nil map", Values(NewFromMap[int, int](nil)), func(a, b int) int { return a + b }, 0, 0},
		{"Empty slice", NewFromSlice([]int{}), func(a, b int) int { return a + b }, 0, 0},
		{"Empty map", Values(NewFromMap(map[int]int{})), func(a, b int) int { return a + b }, 0, 0},
		{"Sum of slice", NewFromSlice([]int{1, 2, 3}), func(a, b int) int { return a + b }, 0, 6},
		{"Sum of map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(a, b int) int { return a + b }, 0, 6},
		{"Product of slice", NewFromSlice([]int{1, 2, 3}), func(a, b int) int { return a * b }, 1, 6},
		{"Product of map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), func(a, b int) int { return a * b }, 1, 6},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Reduce(testCase.iter, testCase.f, testCase.initial)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Reduce(nil, func(int, int) int { return 0 }, 0)
		})
	})
}

func TestReduce2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		f        func(a, b, c int) int
		initial  int
		expected int
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2([]int{}), func(a, b, c int) int { return a + b + c }, 0, 0},
		{"Nil map", NewFromMap(map[int]int{}), func(a, b, c int) int { return a + b + c }, 0, 0},
		{"Empty slice", NewFromSlice2([]int{}), func(a, b, c int) int { return a + b + c }, 0, 0},
		{"Empty map", NewFromMap(map[int]int{}), func(a, b, c int) int { return a + b + c }, 0, 0},
		{"Sum of slice", NewFromSlice2([]int{1, 2, 3}), func(a, b, c int) int { return a + b + c }, 0, 9},
		{"Sum of map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), func(a, b, c int) int { return a + b + c }, 0, 12},
		{"Product of slice", NewFromSlice2([]int{1, 2, 3}), func(a, b, c int) int { return a * b * c }, 1, 0},
		{"Product of map", NewFromMap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}), func(a, b, c int) int { return a * b * c }, 1, 0},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Reduce2(testCase.iter, testCase.f, testCase.initial)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Reduce2(nil, func(int, int, int) int { return 0 }, 0)
		})
	})
}

func TestSum(t *testing.T) {
	type testCase[T any] struct {
		name     string
		iter     iter.Seq[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Nil slice", NewFromSlice[int](nil), 0},
		{"Nil map", Values(NewFromMap[int, int](nil)), 0},
		{"Empty slice", NewFromSlice([]int{}), 0},
		{"Empty map", Values(NewFromMap(map[int]int{})), 0},
		{"Sum of slice", NewFromSlice([]int{1, 2, 3}), 6},
		{"Sum of map", Values(NewFromMap(map[int]int{1: 1, 2: 2, 3: 3})), 6},
		{"Only element from slice", NewFromSlice([]int{1}), 1},
		{"Only element from map", Values(NewFromMap(map[int]int{1: 1})), 1},
		{"Negative sum from slice", NewFromSlice([]int{-1, -2, -3}), -6},
		{"Negative sum from map", Values(NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3})), -6},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Sum(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Sum[int](nil)
		})
	})
}

func TestSum2(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		iter     iter.Seq2[K, V]
		expected int
	}
	testCases := []testCase[int, int]{
		{"Nil slice", NewFromSlice2([]int{}), 0},
		{"Nil map", NewFromMap(map[int]int{}), 0},
		{"Empty slice", NewFromSlice2([]int{}), 0},
		{"Empty map", NewFromMap(map[int]int{}), 0},
		{"Sum of slice", NewFromSlice2([]int{1, 2, 3}), 6},
		{"Sum of map", NewFromMap(map[int]int{1: 1, 2: 2, 3: 3}), 6},
		{"Only element from slice", NewFromSlice2([]int{1}), 1},
		{"Only element from map", NewFromMap(map[int]int{1: 1}), 1},
		{"Negative sum from slice", NewFromSlice2([]int{-1, -2, -3}), -6},
		{"Negative sum from map", NewFromMap(map[int]int{-1: -1, -2: -2, -3: -3}), -6},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Sum2(testCase.iter)
			assert.Equal(t, testCase.expected, actual)
		})
	}
	// Test for panic
	t.Run("Nil iter", func(t *testing.T) {
		assert.PanicsWithValue(t, nilIterPanicMsg, func() {
			Sum2[int, int](nil)
		})
	})
}
