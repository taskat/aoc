package day11

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/graph"
)

// day is the day of the solver
const day = 11

// init registers the solver for day 11
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 11
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) *graph.Graph[string] {
	g := graph.New[string]()
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		nodeId := parts[0]
		if !g.HasNode(nodeId) {
			g.AddNode(graph.NewBaseNode(nodeId))
		}
		parts = strings.Split(parts[1], " ")
		for _, neighborId := range parts {
			if !g.HasNode(neighborId) {
				g.AddNode(graph.NewBaseNode(neighborId))
			}
			g.AddDirectedEdge(nodeId, neighborId, 0)
		}
	}
	return g
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	g := s.parse(lines)
	startNode := g.GetNode("you")
	isGoal := func(id string) bool { return id == "out" }
	paths := g.AllPaths(startNode, isGoal)
	return fmt.Sprint(len(paths))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
