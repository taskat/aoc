package day18

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/graph"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 18

// init registers the solver for day 18
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 18
type Solver struct {
	mapLimit  int
	byteLimit int
}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {
	switch len(params) {
	case 0:
		s.mapLimit = 70
		s.byteLimit = 1024
	case 1:
		s.mapLimit = stringutils.Atoi(params[0])
		s.byteLimit = 1024
	default:
		s.mapLimit = stringutils.Atoi(params[0])
		s.byteLimit = stringutils.Atoi(params[1])
	}
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []coordinate.Coordinate2D[int] {
	return slices.Map(lines, parseCoord)
}

// parseCoord parses a coordinate from a line
func parseCoord(line string) coordinate.Coordinate2D[int] {
	var i, j int
	_, _ = fmt.Sscanf(line, "%d,%d", &i, &j)
	return coordinate.FromIndexes(i, j)
}

// generateGraph generates the graph from the fallen byte coordinates
func (s *Solver) generateGraph(coords []coordinate.Coordinate2D[int]) *graph.Graph[coordinate.Coordinate2D[int]] {
	g := graph.New[coordinate.Coordinate2D[int]]()
	fallenBytes := set.FromSlice(coords)
	for i := 0; i <= s.mapLimit; i++ {
		for j := 0; j <= s.mapLimit; j++ {
			c := coordinate.FromIndexes(i, j)
			if fallenBytes.Contains(c) {
				continue
			}
			g.AddNode(graph.NewBaseNode(c))
		}
	}
	for _, node := range g.GetNodes() {
		for _, neighborCoord := range node.Id().Neighbors(coordinate.Straights()) {
			if g.HasNode(neighborCoord) {
				g.AddEdge(node.Id(), neighborCoord, 1)
			}
		}
	}
	return g
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	coords := s.parse(lines)
	coords = coords[:s.byteLimit]
	graph := s.generateGraph(coords)
	startNode := graph.GetNode(coordinate.FromIndexes(0, 0))
	goalCoord := coordinate.FromIndexes(s.mapLimit, s.mapLimit)
	path := graph.Dijkstra(startNode, goalCoord.Equal)
	return fmt.Sprint(path.Cost())
}

// removeNodes removes nodes from the graph between from (inclusive) and to (exclusive)
func removeNodes(g *graph.Graph[coordinate.Coordinate2D[int]], coords []coordinate.Coordinate2D[int], from, to int) {
	for i := from; i < to; i++ {
		g.RemoveNode(coords[i])
	}
}

// addNodes adds nodes to the graph between from (inclusive) and to (exclusive)
func addNodes(g *graph.Graph[coordinate.Coordinate2D[int]], coords []coordinate.Coordinate2D[int], from, to int) {
	for i := from; i < to; i++ {
		g.AddNode(graph.NewBaseNode(coords[i]))
		for _, neighborCoord := range coords[i].Neighbors(coordinate.Straights()) {
			neighbor := coordinate.NewCoordinate2D(neighborCoord.X, neighborCoord.Y)
			if g.HasNode(neighbor) {
				g.AddEdge(coords[i], neighbor, 1)
			}
		}
	}
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	coords := s.parse(lines)
	g := s.generateGraph(coords[:s.byteLimit])
	startNode := g.GetNode(coordinate.FromIndexes(0, 0))
	goalCoord := coordinate.FromIndexes(s.mapLimit, s.mapLimit)
	from := s.byteLimit
	to := len(coords)
	prev := s.byteLimit
	for from != to {
		halfway := (from + to) / 2
		if halfway == from {
			break
		}
		if prev < halfway {
			removeNodes(g, coords, prev, halfway)
		} else {
			addNodes(g, coords, halfway, prev)
		}
		path := g.Dijkstra(startNode, goalCoord.Equal)
		if path.IsValid() {
			from = halfway
		} else {
			to = halfway
		}
		prev = halfway
	}
	return fmt.Sprintf("%d,%d", coords[from].I(), coords[from].J())
}
