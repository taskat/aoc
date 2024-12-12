package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
