package day16

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/graph"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 16

// init registers the solver for day 16
func init() {
	days.AddDay(day, &Solver{})
}

type coord = coordinate.Coordinate2D[int]

type gate struct {
	c coord
	d coordinate.Direction
}

// String returns a string representation of the gate
func (g gate) String() string {
	return fmt.Sprintf("%v %v", g.c, g.d)
}

// Solver implements the puzzle solver for day 16
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) (*graph.Graph[gate], gate, coord) {
	nodes := set.New[coord]()
	var start, end coord
	for i, line := range lines {
		for j, cell := range line {
			if cell == '#' {
				continue
			}
			c := coordinate.FromIndexes(i, j)
			switch cell {
			case 'S':
				start = c
			case 'E':
				end = c
			}
			nodes.Add(c)
		}
	}
	startGate := gate{start, coordinate.Right()}
	g := graph.New[gate]()
	g.AddNode(graph.NewBaseNode(startGate))
	for _, c := range nodes.ToSlice() {
		possibleNeighbors := c.Neighbors(coordinate.Straights())
		neighbors := maps.Filter(possibleNeighbors, func(dir coordinate.Direction, c coord) bool { return nodes.Contains(c) })
		for dir, neighbor := range neighbors {
			selfGate := gate{c, dir}
			otherGate := gate{neighbor, dir.Opposite()}
			g.AddNode(graph.NewBaseNode(selfGate))
			g.AddNode(graph.NewBaseNode(otherGate))
			g.AddEdge(selfGate, otherGate, 1)
		}
		directions := maps.Keys(neighbors)
		for i := 0; i < len(directions); i++ {
			for j := i + 1; j < len(directions); j++ {
				selfGate1 := gate{c, directions[i]}
				selfGate2 := gate{c, directions[j]}
				cost := 0
				if directions[i].Opposite() != directions[j] {
					cost = 1000
				}
				g.AddEdge(selfGate1, selfGate2, cost)
			}
		}
		if c == start {
			for dir := range neighbors {
				if dir == startGate.d {
					continue
				}
				selfGate := gate{c, dir}
				cost := 0
				if dir.Opposite() != startGate.d {
					cost = 1000
				}
				g.AddEdge(startGate, selfGate, cost)
			}
		}
	}
	return g, startGate, end
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	graph, start, end := s.parse(lines)
	path := graph.Dijkstra(graph.GetNode(start), func(g gate) bool { return g.c.Equal(end) })
	return fmt.Sprint(path.Cost())
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	g, start, end := s.parse(lines)
	nodes := g.NodesOfBestPaths(g.GetNode(start), func(g gate) bool { return g.c.Equal(end) })
	coords := slices.Map(nodes, func(n gate) coord { return n.c })
	coords = set.FromSlice(coords).ToSlice()
	return fmt.Sprint(len(coords))
}
