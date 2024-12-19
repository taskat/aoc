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
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 18
type Solver struct {
	mapLimit  int
	byteLimit int
}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {
	if len(params) < 2 {
		s.mapLimit = 70
		s.byteLimit = 1024
		return
	}
	s.mapLimit = stringutils.Atoi(params[0].(string))
	s.byteLimit = stringutils.Atoi(params[1].(string))
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []coord {
	return slices.Map(lines, parseCoord)

}

type coord = coordinate.Coordinate2D[int]

func parseCoord(line string) coord {
	var x, y int
	_, _ = fmt.Sscanf(line, "%d,%d", &x, &y)
	return coordinate.NewCoordinate2D(x, y)
}

func (s *Solver) generateGraph(coords []coord) *graph.Graph[coord] {
	g := graph.New[coord]()
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
		for _, neighbor := range node.Id().Neighbors(coordinate.Straights()) {
			if g.HasNode(neighbor) {
				g.AddEdge(node.Id(), neighbor, 1)
			}
		}
	}
	return g
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	coords := s.parse(lines)
	if len(coords) > s.byteLimit {
		coords = coords[:s.byteLimit]
	}
	graph := s.generateGraph(coords)
	startNode := graph.GetNode(coordinate.NewCoordinate2D(0, 0))
	goalCoord := coordinate.NewCoordinate2D(s.mapLimit, s.mapLimit)
	path := graph.Dijkstra(startNode, goalCoord.Equal)
	return fmt.Sprint(path.Cost())
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
