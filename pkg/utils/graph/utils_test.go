package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

func TestFromSliceOfIndexes(t *testing.T) {
	type testCase struct {
		testName      string
		coords        []coordinate.Coordinate2D[int]
		neighborDirs  []coordinate.Direction
		expectedEdges map[coordinate.Coordinate2D[int]]map[coordinate.Coordinate2D[int]]int
		expectedNodes []coordinate.Coordinate2D[int]
	}
	testCases := []testCase{
		{
			testName: "Simple square",
			coords: []coordinate.Coordinate2D[int]{
				coordinate.FromIndexes(0, 0),
				coordinate.FromIndexes(0, 1),
				coordinate.FromIndexes(1, 0),
				coordinate.FromIndexes(1, 1),
			},
			neighborDirs: coordinate.Straights(),
			expectedEdges: map[coordinate.Coordinate2D[int]]map[coordinate.Coordinate2D[int]]int{
				coordinate.FromIndexes(0, 0): {
					coordinate.FromIndexes(0, 1): 1,
					coordinate.FromIndexes(1, 0): 1,
				},
				coordinate.FromIndexes(0, 1): {
					coordinate.FromIndexes(0, 0): 1,
					coordinate.FromIndexes(1, 1): 1,
				},
				coordinate.FromIndexes(1, 0): {
					coordinate.FromIndexes(0, 0): 1,
					coordinate.FromIndexes(1, 1): 1,
				},
				coordinate.FromIndexes(1, 1): {
					coordinate.FromIndexes(0, 1): 1,
					coordinate.FromIndexes(1, 0): 1,
				},
			},
			expectedNodes: []coordinate.Coordinate2D[int]{
				coordinate.FromIndexes(0, 0),
				coordinate.FromIndexes(0, 1),
				coordinate.FromIndexes(1, 0),
				coordinate.FromIndexes(1, 1),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			g := FromSliceOfCoordinates(tc.coords, tc.neighborDirs)
			for nodeId, expectedNeighbors := range tc.expectedEdges {
				node := g.GetNode(nodeId)
				for neighborId, expectedCost := range expectedNeighbors {
					cost, exists := node.GetNeighbors()[neighborId]
					assert.True(t, exists, "Neighbor %v of node %v should exist", neighborId, nodeId)
					assert.Equal(t, expectedCost, cost, "Cost to neighbor %v of node %v should be %d", neighborId, nodeId, expectedCost)
				}
			}
			actualNodes := slices.Map(maps.Values(g.GetNodes()), Node[coordinate.Coordinate2D[int]].Id)
			assert.ElementsMatch(t, tc.expectedNodes, actualNodes, "Graph nodes should match expected nodes")
		})
	}
}
