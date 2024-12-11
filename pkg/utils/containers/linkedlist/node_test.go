package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFirst(t *testing.T) {
	type testCase[T any] struct {
		testName string
		node     *node[T]
		expected bool
	}
	testCases := []testCase[int]{
		{"Test first", &node[int]{value: 1, prev: nil}, true},
		{"Test not first", &node[int]{value: 2, prev: &node[int]{value: 1}}, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := tc.node.IsFirst()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsLast(t *testing.T) {
	type testCase[T any] struct {
		testName string
		node     *node[T]
		expected bool
	}
	testCases := []testCase[int]{
		{"Test last", &node[int]{value: 1, next: nil}, true},
		{"Test not last", &node[int]{value: 2, next: &node[int]{value: 3}}, false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := tc.node.IsLast()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSetValue(t *testing.T) {
	type testCase[T any] struct {
		testName string
		node     *node[T]
		value    T
	}
	testCases := []testCase[int]{
		{"Test 1", &node[int]{value: 1}, 2},
		{"Test 2", &node[int]{value: 2}, 3},
		{"Test negative", &node[int]{value: -1}, -2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.node.SetValue(tc.value)
			assert.Equal(t, tc.value, tc.node.value)
		})
	}
}

func TestNodeString(t *testing.T) {
	type testCase[T any] struct {
		testName string
		node     *node[T]
		expected string
	}
	testCases := []testCase[int]{
		{"Test nil", nil, ""},
		{"Test 1", &node[int]{value: 1}, "1"},
		{"Test 2", &node[int]{value: 2}, "2"},
		{"Test negative", &node[int]{value: -1}, "-1"},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := tc.node.String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestValue(t *testing.T) {
	type testCase[T any] struct {
		testName string
		node     *node[T]
		expected T
	}
	testCases := []testCase[int]{
		{"Test 1", &node[int]{value: 1}, 1},
		{"Test 2", &node[int]{value: 2}, 2},
		{"Test negative", &node[int]{value: -1}, -1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := tc.node.Value()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
