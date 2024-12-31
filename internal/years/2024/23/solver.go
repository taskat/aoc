package day23

import (
	"fmt"
	"sort"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/graph"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 23

// init registers the solver for day 23
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 23
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

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

func newSet(values ...string) setOf3 {
	if len(values) != 3 {
		panic("setOf3 must have 3 values")
	}
	sort.Strings(values)
	return setOf3{values[0], values[1], values[2]}
}

func get3Sets(g *graph.Graph[string]) []setOf3 {
	sets := set.New[setOf3]()
	nodes := maps.Values(g.GetNodes())
	ts := slices.Filter(nodes, func(n graph.Node[string]) bool { return strings.HasPrefix(n.Id(), "t") })
	for _, t := range ts {
		tId := t.Id()
		tNeighbors := maps.Keys(t.GetNeighbors())
		for i := 0; i < len(tNeighbors); i++ {
			iId := tNeighbors[i]
			iNeighbors := g.GetNode(iId).GetNeighbors()
			for j := i + 1; j < len(tNeighbors); j++ {
				jId := tNeighbors[j]
				if maps.Contains(iNeighbors, jId) {
					sets.Add(newSet(tId, iId, jId))
				}
			}
		}
	}
	return sets.ToSlice()
}

func canAdd(g *graph.Graph[string], newNeighbor string, neighbors []string) bool {
	for _, neighbor := range neighbors {
		if !g.GetNode(neighbor).HasNeighbor(newNeighbor) {
			return false
		}
	}
	return true
}

func getBiggestSetIn(g *graph.Graph[string], currentNeighbors, possibleNeighbors []string, maxLength int) []string {
	if len(currentNeighbors)+len(possibleNeighbors) < maxLength {
		return nil
	}
	if len(possibleNeighbors) == 0 {
		if len(currentNeighbors) <= maxLength {
			return nil
		}
		return currentNeighbors
	}
	neighbor := possibleNeighbors[0]
	possibleNeighbors = possibleNeighbors[1:]
	biggestSet := slices.Copy(currentNeighbors)
	if canAdd(g, neighbor, currentNeighbors) {
		withNeighbor := getBiggestSetIn(g, append(slices.Copy(currentNeighbors), neighbor), possibleNeighbors, maxLength)
		if len(withNeighbor) > maxLength {
			maxLength = len(withNeighbor)
			biggestSet = withNeighbor
		}
	}
	withoutNeighbor := getBiggestSetIn(g, currentNeighbors, possibleNeighbors, maxLength)
	if len(withoutNeighbor) > maxLength {
		maxLength = len(withoutNeighbor)
		biggestSet = withoutNeighbor
	}
	return biggestSet
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	g := s.parse(lines)
	sets := get3Sets(g)
	return fmt.Sprint(len(sets))
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	g := s.parse(lines)
	nodes := maps.Values(g.GetNodes())
	sets := slices.Map(nodes, func(n graph.Node[string]) []string {
		neighbors := maps.Keys(n.GetNeighbors())
		return getBiggestSetIn(g, []string{n.Id()}, neighbors, 0)
	})
	lengths := slices.Map(sets, func(s []string) int { return len(s) })
	_, i := slices.Max_i(lengths)
	maxSet := sets[i]
	// maxSet = set.FromSlice(maxSet).ToSlice()
	sort.Strings(maxSet)
	return strings.Join(maxSet, ",")
}
