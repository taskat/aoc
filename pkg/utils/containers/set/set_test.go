package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	type testCase[T comparable] struct {
		testName    string
		s           []T
		expectedSet Set[T]
	}
	testCases := []testCase[int]{
		{"Empty slice", []int{}, New[int]()},
		{"Non-empty slice", []int{1, 2}, Set[int](map[int]struct{}{1: {}, 2: {}})},
		{"Duplicate elements", []int{1, 1}, Set[int](map[int]struct{}{1: {}})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			s := FromSlice(tc.s)
			assert.Equal(t, tc.expectedSet, s)
		})
	}
}

func TestNew(t *testing.T) {
	s := New[int]()
	assert.Equal(t, Set[int](map[int]struct{}{}), s)
}

func TestAdd(t *testing.T) {
	type testCase[T comparable] struct {
		testName    string
		s           Set[T]
		e           T
		expectedSet Set[T]
	}
	testCases := []testCase[int]{
		{"Empty set", New[int](), 1, Set[int](map[int]struct{}{1: {}})},
		{"Element not present", Set[int](map[int]struct{}{1: {}}), 2, Set[int](map[int]struct{}{1: {}, 2: {}})},
		{"Element present", Set[int](map[int]struct{}{1: {}}), 1, Set[int](map[int]struct{}{1: {}})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.s.Add(tc.e)
			assert.Equal(t, tc.expectedSet, tc.s)
		})
	}
}

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		testName    string
		s           Set[T]
		e           T
		expectedRes bool
	}
	testCases := []testCase[int]{
		{"Empty set", New[int](), 1, false},
		{"Element not present", Set[int](map[int]struct{}{1: {}}), 2, false},
		{"Element present", Set[int](map[int]struct{}{1: {}}), 1, true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			res := tc.s.Contains(tc.e)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestDelete(t *testing.T) {
	type testCase[T comparable] struct {
		testName    string
		s           Set[T]
		e           T
		expectedSet Set[T]
	}
	testCases := []testCase[int]{
		{"Empty set", New[int](), 1, New[int]()},
		{"Element not present", Set[int](map[int]struct{}{1: {}}), 2, Set[int](map[int]struct{}{1: {}})},
		{"Element present", Set[int](map[int]struct{}{1: {}}), 1, New[int]()},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.s.Delete(tc.e)
			assert.Equal(t, tc.expectedSet, tc.s)
		})
	}
}

func TestLength(t *testing.T) {
	type testCase[T comparable] struct {
		testName    string
		s           Set[T]
		expectedLen int
	}
	testCases := []testCase[int]{
		{"Empty set", New[int](), 0},
		{"Non-empty set", Set[int](map[int]struct{}{1: {}, 2: {}}), 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			l := tc.s.Length()
			assert.Equal(t, tc.expectedLen, l)
		})
	}
}

func TestMap(t *testing.T) {
	type testCase[T comparable] struct {
		testName    string
		s           Set[T]
		f           func(T) T
		expectedSet Set[T]
	}
	testCases := []testCase[int]{
		{"Empty set", New[int](), func(e int) int { return e + 1 }, New[int]()},
		{"Non-empty set", Set[int](map[int]struct{}{1: {}, 2: {}}), func(e int) int { return e + 1 }, Set[int](map[int]struct{}{2: {}, 3: {}})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.s = Map(tc.s, tc.f)
			assert.Equal(t, tc.expectedSet, tc.s)
		})
	}
}

func TestMerge(t *testing.T) {
	type testCase[T comparable] struct {
		testName    string
		s1          Set[T]
		s2          Set[T]
		expectedSet Set[T]
	}
	testCases := []testCase[int]{
		{"Both empty sets", New[int](), New[int](), New[int]()},
		{"First set empty", New[int](), Set[int](map[int]struct{}{1: {}}), Set[int](map[int]struct{}{1: {}})},
		{"Second set empty", Set[int](map[int]struct{}{1: {}}), New[int](), Set[int](map[int]struct{}{1: {}})},
		{"Non-empty sets", Set[int](map[int]struct{}{1: {}}), Set[int](map[int]struct{}{2: {}}), Set[int](map[int]struct{}{1: {}, 2: {}})},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			s := tc.s1.Merge(tc.s2)
			assert.Equal(t, tc.expectedSet, s)
		})
	}
}

func TestReduce(t *testing.T) {
	type testCase[T comparable] struct {
		testName      string
		s             Set[T]
		f             func(int, T) int
		initialValue  int
		expectedValue int
	}
	testCases := []testCase[int]{
		{"Empty set", New[int](), func(acc int, e int) int { return acc + e }, 0, 0},
		{"Non-empty set", Set[int](map[int]struct{}{1: {}, 2: {}}), func(acc int, e int) int { return acc + e }, 0, 3},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			value := Reduce(tc.s, tc.f, tc.initialValue)
			assert.Equal(t, tc.expectedValue, value)
		})
	}
}

func TestToSlice(t *testing.T) {
	type testCase[T comparable] struct {
		testName   string
		s          Set[T]
		expectedSl []T
	}
	testCases := []testCase[int]{
		{"Empty set", New[int](), []int{}},
		{"Non-empty set", Set[int](map[int]struct{}{1: {}, 2: {}}), []int{1, 2}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			sl := tc.s.ToSlice()
			assert.ElementsMatch(t, tc.expectedSl, sl)
		})
	}
}
