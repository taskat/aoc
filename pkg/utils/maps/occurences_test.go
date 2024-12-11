package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOccurrences(t *testing.T) {
	testCases := []struct {
		testName       string
		s              []int
		expectedOccMap map[int]int
	}{
		{"Empty slice", []int{}, map[int]int{}},
		{"Slice with elements", []int{1, 2, 1, 3, 2}, map[int]int{1: 2, 2: 2, 3: 1}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			occMap := Occurrences(tc.s)
			assert.Equal(t, tc.expectedOccMap, occMap)
		})
	}
}

func TestAddOccurence(t *testing.T) {
	testCases := []struct {
		testName       string
		m              map[int]int
		e              int
		count          int
		expectedOccMap map[int]int
	}{
		{
			"Empty map",
			map[int]int{},
			1,
			1,
			map[int]int{1: 1},
		},
		{
			"Element not present",
			map[int]int{1: 1},
			2,
			1,
			map[int]int{1: 1, 2: 1},
		},
		{
			"Element present",
			map[int]int{1: 1},
			1,
			1,
			map[int]int{1: 2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			AddOccurence(tc.m, tc.e, tc.count)
			assert.Equal(t, tc.expectedOccMap, tc.m)
		})
	}
}
