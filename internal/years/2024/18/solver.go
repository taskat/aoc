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
	switch len(params) {
	case 0:
		s.mapLimit = 70
		s.byteLimit = 1024
	case 1:
		s.mapLimit = stringutils.Atoi(params[0].(string))
		s.byteLimit = 1024
	default:
		s.mapLimit = stringutils.Atoi(params[0].(string))
		s.byteLimit = stringutils.Atoi(params[1].(string))
	}
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []coord {
	return slices.Map(lines, parseCoord)

}

type coord struct {
	coordinate.Coordinate2D[int]
}

func newCoord(i, j int) coord {
	return coord{coordinate.FromIndexes(i, j)}
}

func parseCoord(line string) coord {
	var x, y int
	_, _ = fmt.Sscanf(line, "%d,%d", &x, &y)
	return newCoord(y, x)
}

func (c coord) Equal(other coord) bool {
	return c.Coordinate2D.Equal(other.Coordinate2D)
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (s *Solver) generateGraph(coords []coord) *graph.Graph[coord] {
	g := graph.New[coord]()
	fallenBytes := set.FromSlice(coords)
	for i := 0; i <= s.mapLimit; i++ {
		for j := 0; j <= s.mapLimit; j++ {
			c := newCoord(i, j)
			if fallenBytes.Contains(c) {
				continue
			}
			g.AddNode(graph.NewBaseNode(c))
		}
	}
	for _, node := range g.GetNodes() {
		for _, neighborCoord := range node.Id().Neighbors(coordinate.Straights()) {
			neighbor := newCoord(neighborCoord.Y, neighborCoord.X)
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
	coords = coords[:s.byteLimit]
	graph := s.generateGraph(coords)
	startNode := graph.GetNode(newCoord(0, 0))
	goalCoord := newCoord(s.mapLimit, s.mapLimit)
	path := graph.Dijkstra(startNode, goalCoord.Equal)
	return fmt.Sprint(path.Cost())
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	coords := s.parse(lines)
	g := s.generateGraph(coords[:s.byteLimit])
	startNode := g.GetNode(newCoord(0, 0))
	goalCoord := newCoord(s.mapLimit, s.mapLimit)
	var closingIndex int
	for closingIndex = s.byteLimit; closingIndex < len(coords); closingIndex++ {
		fmt.Println("Closing index", closingIndex)
		g.RemoveNode(coords[closingIndex])
		path := g.Dijkstra(startNode, goalCoord.Equal)
		if !path.IsValid() {
			break
		}
	}
	return coords[closingIndex].String()
}
