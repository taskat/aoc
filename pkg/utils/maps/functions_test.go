package maps

import "testing"

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

func TestMerge(t *testing.T) {
	testCases := []struct {
		testName    string
		map1        map[int]string
		map2        map[int]string
		expectedMap map[int]string
	}{
		{
			"Nil maps",
			nil,
			nil,
			map[int]string{},
		},
		{
			"Empty maps",
			map[int]string{},
			map[int]string{},
			map[int]string{},
		},
		{
			"First map empty",
			map[int]string{},
			map[int]string{1: "a", 2: "b"},
			map[int]string{1: "a", 2: "b"},
		},
		{
			"Second map empty",
			map[int]string{1: "a", 2: "b"},
			map[int]string{},
			map[int]string{1: "a", 2: "b"},
		},
		{
			"Both maps have elements",
			map[int]string{1: "a", 2: "b"},
			map[int]string{3: "c", 4: "d"},
			map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
		},
		{
			"Overlapping keys",
			map[int]string{1: "a", 2: "b"},
			map[int]string{2: "c", 3: "d"},
			map[int]string{1: "a", 2: "c", 3: "d"},
		},
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
