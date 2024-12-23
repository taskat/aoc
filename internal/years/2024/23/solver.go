package day23

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/graph"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 23

// init registers the solver for day 23
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 23
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) *graph.Graph[string] {
	g := graph.New[string]()
	for _, line := range lines {
		computers := strings.Split(line, "-")
		if !g.HasNode(computers[0]) {
			g.AddNode(graph.NewBaseNode(computers[0]))
		}
		if !g.HasNode(computers[1]) {
			g.AddNode(graph.NewBaseNode(computers[1]))
		}
		g.AddEdge(computers[0], computers[1], 0)
	}
	return g
}

type setOf3 [3]string

func (s setOf3) hasT() bool {
	return slices.Any(s[:], func(computer string) bool { return strings.HasPrefix(computer, "t") })
}

func get3Sets(g *graph.Graph[string]) []setOf3 {
	sets := make([]setOf3, 0)
	nodes := maps.Values(g.GetNodes())
	for i := 0; i < len(nodes); i++ {
		iId := nodes[i].Id()
		for j := i + 1; j < len(nodes); j++ {
			jId := nodes[j].Id()
			for k := j + 1; k < len(nodes); k++ {
				kId := nodes[k].Id()
				iNeighbors := nodes[i].GetNeighbors()
				jNeighbors := nodes[j].GetNeighbors()
				if maps.Contains(iNeighbors, jId) && maps.Contains(iNeighbors, kId) && maps.Contains(jNeighbors, kId) {
					sets = append(sets, setOf3{iId, jId, kId})
				}
			}
		}
	}
	return sets
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	g := s.parse(lines)
	sets := get3Sets(g)
	sets = slices.Filter(sets, setOf3.hasT)
	return fmt.Sprint(len(sets))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
