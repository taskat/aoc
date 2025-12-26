package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPaths(t *testing.T) {
	type testCase struct {
		testName      string
		startNodeId   int
		goal          func(int) bool
		expectedPaths []Path[int]
	}
	node1 := &BaseNode[int]{id: 1, neighbors: make(map[int]int)}
	node2 := &BaseNode[int]{id: 2, neighbors: make(map[int]int)}
	node3 := &BaseNode[int]{id: 3, neighbors: make(map[int]int)}
	node4 := &BaseNode[int]{id: 4, neighbors: make(map[int]int)}
	node5 := &BaseNode[int]{id: 5, neighbors: make(map[int]int)}
	node6 := &BaseNode[int]{id: 6, neighbors: make(map[int]int)}
	node1.neighbors[2] = 1
	node1.neighbors[3] = 20
	node1.neighbors[4] = 7
	node1.neighbors[5] = 1
	node2.neighbors[3] = 3
	node2.neighbors[4] = 5
	node3.neighbors[4] = 2
	graph := New[int]()
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	graph.AddNode(node5)
	graph.AddNode(node6)
	testCases := []testCase{
		{
			"Single path",
			1,
			func(n int) bool { return n == 5 },
			[]Path[int]{
				{nodes: []int{1, 5}, cost: 1},
			},
		},
		{
			"Multiple paths",
			1,
			func(n int) bool { return n == 4 },
			[]Path[int]{
				{nodes: []int{1, 2, 4}, cost: 6},
				{nodes: []int{1, 2, 3, 4}, cost: 6},
				{nodes: []int{1, 4}, cost: 7},
				{nodes: []int{1, 3, 4}, cost: 22},
			},
		},
		{
			"No path",
			1,
			func(n int) bool { return n == 6 },
			[]Path[int]{},
		},
		{
			"Same node",
			1,
			func(n int) bool { return n == 1 },
			[]Path[int]{
				{nodes: []int{1}, cost: 0},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			paths := graph.AllPaths(graph.GetNode(tc.startNodeId), tc.goal)
			assert.Equal(t, len(tc.expectedPaths), len(paths))
			for _, expectedPath := range tc.expectedPaths {
				found := false
				for _, path := range paths {
					if expectedPath.Equals(path) {
						found = true
						break
					}
				}
				assert.True(t, found, "Expected path not found: %v", expectedPath)
			}
		})
	}
}

func TestNodesBestPaths(t *testing.T) {
	type testCase struct {
		testName      string
		goal          func(int) bool
		expectedPaths []int
	}
	node1 := &BaseNode[int]{id: 1, neighbors: make(map[int]int)}
	node2 := &BaseNode[int]{id: 2, neighbors: make(map[int]int)}
	node3 := &BaseNode[int]{id: 3, neighbors: make(map[int]int)}
	node4 := &BaseNode[int]{id: 4, neighbors: make(map[int]int)}
	node5 := &BaseNode[int]{id: 5, neighbors: make(map[int]int)}
	node1.neighbors[2] = 1
	node1.neighbors[3] = 20
	node1.neighbors[4] = 7
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
		{"Single route", func(n int) bool { return n == 2 }, []int{1, 2}},
		{"Multiple routes", func(n int) bool { return n == 4 }, []int{1, 2, 3, 4}},
		{"No path", func(n int) bool { return n == 5 }, []int{}},
		{"Same node", func(n int) bool { return n == 1 }, []int{1}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			paths := graph.NodesOfBestPaths(node1, tc.goal)
			assert.ElementsMatch(t, tc.expectedPaths, paths)
		})
	}
}

func TestDijkstra(t *testing.T) {
	type testCase struct {
		testName     string
		goal         func(int) bool
		expectedPath Path[int]
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
	node2.neighbors[4] = 6
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
			expectedPath: Path[int]{nodes: []int{1, 2}, cost: 1},
		},
		{
			testName:     "Complex test",
			goal:         func(n int) bool { return n == 4 },
			expectedPath: Path[int]{nodes: []int{1, 2, 3, 4}, cost: 6},
		},
		{
			testName:     "No path",
			goal:         func(n int) bool { return n == 5 },
			expectedPath: NoPath[int](),
		},
		{
			testName:     "Same node",
			goal:         func(n int) bool { return n == 1 },
			expectedPath: Path[int]{nodes: []int{1}, cost: 0},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			path := graph.Dijkstra(node1, tc.goal)
			assert.Equal(t, tc.expectedPath, path)
		})
	}
}

func TestHasDirectedCycle(t *testing.T) {
	node1 := &BaseNode[int]{id: 1, neighbors: make(map[int]int)}
	node2 := &BaseNode[int]{id: 2, neighbors: make(map[int]int)}
	node3 := &BaseNode[int]{id: 3, neighbors: make(map[int]int)}
	node4 := &BaseNode[int]{id: 4, neighbors: make(map[int]int)}
	node1.neighbors[2] = 1
	node2.neighbors[3] = 1
	node3.neighbors[4] = 1
	graph := New[int]()
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	assert.False(t, graph.HasDirectedCycle(node1), "Graph A should not have a directed cycle")
	node4.neighbors[2] = 1
	assert.True(t, graph.HasDirectedCycle(node1), "Graph B should have a directed cycle")
}

func TestStartNodes(t *testing.T) {
	node1 := &BaseNode[int]{id: 1, neighbors: make(map[int]int)}
	node2 := &BaseNode[int]{id: 2, neighbors: make(map[int]int)}
	node3 := &BaseNode[int]{id: 3, neighbors: make(map[int]int)}
	node4 := &BaseNode[int]{id: 4, neighbors: make(map[int]int)}
	node1.neighbors[2] = 1
	node1.neighbors[3] = 1
	node2.neighbors[3] = 1
	node2.neighbors[4] = 1
	node3.neighbors[4] = 1
	graph := New[int]()
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	startNodes := graph.StartNodes()
	expectedStartNodes := []int{1}
	assert.ElementsMatch(t, startNodes, expectedStartNodes)
}

func TestTopologicalOrder(t *testing.T) {
	node1 := &BaseNode[int]{id: 1, neighbors: make(map[int]int)}
	node2 := &BaseNode[int]{id: 2, neighbors: make(map[int]int)}
	node3 := &BaseNode[int]{id: 3, neighbors: make(map[int]int)}
	node4 := &BaseNode[int]{id: 4, neighbors: make(map[int]int)}
	node1.neighbors[2] = 1
	node1.neighbors[3] = 1
	node2.neighbors[3] = 1
	node2.neighbors[4] = 1
	node3.neighbors[4] = 1
	graph := New[int]()
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	order := graph.TopologicalOrder()
	expectedOrder := []int{1, 2, 3, 4}
	assert.Equal(t, expectedOrder, order)
}
