package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
