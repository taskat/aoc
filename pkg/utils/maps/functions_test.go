package maps

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	testCases := []struct {
		testName    string
		m           map[int][]string
		k           int
		v           string
		expectedMap map[int][]string
	}{
		{"Empty map", map[int][]string{}, 1, "a", map[int][]string{1: {"a"}}},
		{"Key not present", map[int][]string{1: {"a"}}, 2, "b", map[int][]string{1: {"a"}, 2: {"b"}}},
		{"Key present", map[int][]string{1: {"a"}}, 1, "b", map[int][]string{1: {"a", "b"}}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			Append(tc.m, tc.k, tc.v)
			assert.Equal(t, tc.expectedMap, tc.m)
		})
	}
	// Test for panic
	assert.Panics(t, func() {
		Append(nil, 1, "a")
	})
}

func TestContains(t *testing.T) {
	testCases := []struct {
		testName    string
		m           map[int]string
		k           int
		expectedRes bool
	}{
		{
			"Empty map",
			map[int]string{},
			1,
			false,
		},
		{
			"Key not present",
			map[int]string{1: "a", 2: "b"},
			3,
			false,
		},
		{
			"Key present",
			map[int]string{1: "a", 2: "b"},
			2,
			true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			res := Contains(tc.m, tc.k)
			if res != tc.expectedRes {
				t.Errorf("Expected %t, got %t", tc.expectedRes, res)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		testName    string
		m           map[int]string
		expectedMap map[int]string
	}{
		{"Nil map", nil, map[int]string{}},
		{"Empty map", map[int]string{}, map[int]string{}},
		{"Map with elements", map[int]string{1: "a", 2: "b"}, map[int]string{1: "a"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			f := func(k int, v string) bool {
				return k == 1
			}
			result := Filter(tc.m, f)
			assert.Equal(t, tc.expectedMap, result)
		})
	}
}

func TestForEach(t *testing.T) {
	testCases := []struct {
		testName    string
		m           map[int][]string
		expectedMap map[int][]string
	}{
		{"Nil map", nil, nil},
		{"Empty map", map[int][]string{}, map[int][]string{}},
		{"Map with elements", map[int][]string{1: {"a"}, 2: {"b", "c"}}, map[int][]string{1: {"A"}, 2: {"B"}}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			f := func(k int, v []string) {
				v[0] = strings.ToUpper(v[0])
			}
			ForEach(tc.m, f)
		})
	}
}

func TestKeys(t *testing.T) {
	type testCase[T comparable] struct {
		testName     string
		m            map[T]string
		expectedKeys []T
	}
	testCases := []testCase[int]{
		{"Nil map", nil, []int{}},
		{"Empty map", map[int]string{}, []int{}},
		{"Map with elements", map[int]string{1: "a", 2: "b"}, []int{1, 2}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			keys := Keys(tc.m)
			assert.Equal(t, tc.expectedKeys, keys)
		})
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		testName    string
		m           map[int]string
		expectedMap map[string]int
	}{
		{"Nil map", nil, map[string]int{}},
		{"Empty map", map[int]string{}, map[string]int{}},
		{"Map with elements", map[int]string{1: "a", 2: "b"}, map[string]int{"a": 1, "b": 2}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			f := func(k int, v string) (string, int) {
				return v, k
			}
			result := Map(tc.m, f)
			assert.Equal(t, tc.expectedMap, result)
		})
	}
}

func TestMapValues(t *testing.T) {
	testCases := []struct {
		testName    string
		m           map[int]string
		expectedMap map[int]int
	}{
		{"Nil map", nil, map[int]int{}},
		{"Empty map", map[int]string{}, map[int]int{}},
		{"Map with elements", map[int]string{1: "2", 2: "32"}, map[int]int{1: 1, 2: 2}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			f := func(v string) int {
				return len(v)
			}
			result := MapValues(tc.m, f)
			assert.Equal(t, tc.expectedMap, result)
		})
	}
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		testName    string
		map1        map[int]string
		map2        map[int]string
		expectedMap map[int]string
	}{
		{"Nil maps", nil, nil, map[int]string{}},
		{"Empty maps", map[int]string{}, map[int]string{}, map[int]string{}},
		{"Map1 empty", map[int]string{}, map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"}},
		{"Map2 empty", map[int]string{1: "a", 2: "b"}, map[int]string{}, map[int]string{1: "a", 2: "b"}},
		{"Maps with elements", map[int]string{1: "a", 2: "b"}, map[int]string{2: "c", 3: "d"}, map[int]string{1: "a", 2: "c", 3: "d"}},
		{"Overlapping keys", map[int]string{1: "a", 2: "b"}, map[int]string{2: "c", 3: "d"}, map[int]string{1: "a", 2: "c", 3: "d"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := Merge(tc.map1, tc.map2)
			for k, v := range tc.expectedMap {
				if result[k] != v {
					t.Errorf("Expected value %s for key %d, got %s", v, k, result[k])
				}
			}
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		testName    string
		m           map[int]int
		expectedSum int
	}{
		{"Nil map", nil, 0},
		{"Empty map", map[int]int{}, 0},
		{"Map with elements", map[int]int{1: 2, 2: 3}, 5},
		{"Map with negative elements", map[int]int{1: -2, 2: 3}, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			sum := Sum(tc.m)
			assert.Equal(t, tc.expectedSum, sum)
		})
	}
}

func TestToSlice(t *testing.T) {
	testCases := []struct {
		testName      string
		m             map[int]string
		expectedSlice []string
	}{
		{"Nil map", nil, []string{}},
		{"Empty map", map[int]string{}, []string{}},
		{"Map with elements", map[int]string{1: "a", 2: "b"}, []string{"1: a", "2: b"}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			f := func(k int, v string) string {
				return strings.Join([]string{fmt.Sprint(k), v}, ": ")
			}
			result := ToSlice(tc.m, f)
			assert.ElementsMatch(t, tc.expectedSlice, result)
		})
	}
}

func TestValues(t *testing.T) {
	type testCase[T comparable] struct {
		testName        string
		m               map[T]string
		expectedValues  []string
		expectedLength  int
		expectedCapcity int
	}
	testCases := []testCase[int]{
		{"Nil map", nil, []string{}, 0, 0},
		{"Empty map", map[int]string{}, []string{}, 0, 0},
		{"Map with elements", map[int]string{1: "a", 2: "b"}, []string{"a", "b"}, 2, 2},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			values := Values(tc.m)
			assert.Equal(t, tc.expectedValues, values)
			assert.Equal(t, tc.expectedLength, len(values))
			assert.Equal(t, tc.expectedCapcity, cap(values))
		})
	}
}
