package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	type testCase struct {
		testName     string
		goal         func(int) bool
		expectedPath []int
		expectedCost int
	}
	node1 := &BaseNode[int]{id: 1, neighbors: make(map[int]int)}
	node2 := &BaseNode[int]{id: 2, neighbors: make(map[int]int)}
	node3 := &BaseNode[int]{id: 3, neighbors: make(map[int]int)}
	node4 := &BaseNode[int]{id: 4, neighbors: make(map[int]int)}
	node5 := &BaseNode[int]{id: 5, neighbors: make(map[int]int)}
	node1.neighbors[2] = 1
	node1.neighbors[3] = 20
	node2.neighbors[3] = 3
	node2.neighbors[4] = 5
	node3.neighbors[4] = 2
	graph := New[int]()
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	graph.AddNode(node5)
	testCases := []testCase{
		{
			testName:     "Simple test",
			goal:         func(n int) bool { return n == 2 },
			expectedPath: []int{1, 2},
			expectedCost: 1,
		},
		{
			testName:     "Complex test",
			goal:         func(n int) bool { return n == 4 },
			expectedPath: []int{1, 2, 3, 4},
			expectedCost: 6,
		},
		{
			testName:     "No path",
			goal:         func(n int) bool { return n == 5 },
			expectedPath: nil,
			expectedCost: -1,
		},
		{
			testName:     "Same node",
			goal:         func(n int) bool { return n == 1 },
			expectedPath: []int{1},
			expectedCost: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			path, cost := graph.Dijkstra(node1, tc.goal)
			assert.Equal(t, tc.expectedPath, path)
			assert.Equal(t, tc.expectedCost, cost)
		})
	}
}
