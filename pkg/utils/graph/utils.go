package graph

import (
	"slices"

	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// FromSliceOfCoordinates creates a graph from a slice of coordinates
func FromSliceOfCoordinates(coords []coordinate.Coordinate2D[int], neighborDirs []coordinate.Direction) *Graph[coordinate.Coordinate2D[int]] {
	g := New[coordinate.Coordinate2D[int]]()
	for _, c := range coords {
		possibleNeighbors := c.Neighbors(neighborDirs)
		neighbors := maps.Filter(possibleNeighbors, func(dir coordinate.Direction, c coordinate.Coordinate2D[int]) bool { return slices.Contains(coords, c) })
		for _, neighbor := range neighbors {
			g.AddNode(NewBaseNode(c))
			g.AddNode(NewBaseNode(neighbor))
			g.AddEdge(c, neighbor, 1)
		}
	}
	return g
}
