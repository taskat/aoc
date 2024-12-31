package day12

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 12

// init registers the solver for day 12
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 12
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) garden {
	return parseGarden(lines)
}

// garden represents the plots of the garden
type garden map[rune][]region

// parseGarden parses the garden from the lines
func parseGarden(lines []string) garden {
	plots := make(map[rune][]coord)
	for i, line := range lines {
		for j, c := range line {
			maps.Append(plots, c, coordinate.FromIndexes(i, j))
		}
	}
	return maps.MapValues(plots, separateRegions)
}

// separateRegions separates the plots with the same plant into regions
// based on neighbors
func separateRegions(plots []coord) []region {
	regions := make([]region, 0)
	for len(plots) > 0 {
		current := plots[0]
		plots = slices.RemoveNth(plots, 0)
		currentRegion := newRegion()
		currentRegion.Add(current)
		filter := func(c coord) bool {
			return slices.Contains(plots, c) && !currentRegion.Contains(c)
		}
		getNeighbors := func(c coord) []coord {
			return slices.Filter(maps.Values(neighbors(c)), filter)
		}
		neighborsToVisit := getNeighbors(current)
		for len(neighborsToVisit) > 0 {
			n := neighborsToVisit[0]
			neighborsToVisit = slices.RemoveNth(neighborsToVisit, 0)
			currentRegion.Add(n)
			neighborsToVisit = append(neighborsToVisit, getNeighbors(n)...)
			neighborsToVisit = slices.Filter(neighborsToVisit, filter)
			index := slices.FindIndex(plots, n.Equal)
			plots = slices.RemoveNth(plots, index)
		}
		regions = append(regions, currentRegion)
	}
	return regions
}

// cost returns the cost of the garden
func (g garden) cost(bulk bool) int {
	costPerRegions := maps.MapValues(g, func(regions []region) []int {
		return slices.Map(regions, func(r region) int { return r.cost(bulk) })
	})
	costPerType := maps.MapValues(costPerRegions, slices.Sum)
	return slices.Sum(maps.Values(costPerType))
}

// region represents a region of the garden with the same plant
type region struct {
	set.Set[coord]
}

// newRegion creates a new region
func newRegion() region {
	return region{set.New[coord]()}
}

// area returns the area of the region
func (r region) area() int {
	return len(r.Set)
}

// cost returns the cost of the region. If bulk is true, the cost is calculated
// based on the number of sides of the region, otherwise it is calculated based
// on the perimeter of the region
func (r region) cost(bulk bool) int {
	perimeter := r.perimeter()
	if bulk {
		perimeter = r.sides()
	}
	return r.area() * perimeter
}

// fences returns the fences of the region
func (r region) fences() set.Set[fence] {
	fences := set.New[fence]()
	for c := range r.Set {
		for dir, n := range neighbors(c) {
			if !r.Contains(n) {
				fences.Add(newFence(c, dir))
			}
		}
	}
	return fences
}

// perimeter returns the perimeter of the region
func (r region) perimeter() int {
	isPerimeter := func(other coord) bool {
		return !r.Contains(other)
	}
	perimeterLength := func(c coord) int {
		return slices.Count(maps.Values(neighbors(c)), isPerimeter)
	}
	perimeters := slices.Map(r.ToSlice(), perimeterLength)
	return slices.Sum(perimeters)
}

// sides returns the number of sides of the region
func (r region) sides() int {
	fences := r.fences()
	count := 0
	for f := range fences {
		delete(fences, f)
		neighborFences := f.neighbors()
		for i := 0; i < len(neighborFences); i++ {
			nf := neighborFences[i]
			if !fences.Contains(nf) {
				continue
			}
			delete(fences, nf)
			neighborFences = append(neighborFences, nf.neighbors()...)
		}
		count++
	}
	return count
}

// fence represents a fenc on the dir side of the c ccordinate
type fence struct {
	c   coord
	dir coordinate.Direction
}

// newFence creates a new fence
func newFence(c coord, dir coordinate.Direction) fence {
	return fence{c, dir}
}

// neighbors returns the continuation of the fence
func (f fence) neighbors() []fence {
	var dirs []coordinate.Direction
	if f.dir.Vertical() {
		dirs = coordinate.Horizontals()
	} else {
		dirs = coordinate.Verticals()
	}
	coords := slices.Map(dirs, f.c.Go)
	return slices.Map(coords, func(c coord) fence { return newFence(c, f.dir) })
}

// String returns the string representation of the fence
func (f fence) String() string {
	return fmt.Sprintf("%v %v", f.c, f.dir)
}

type coord = coordinate.Coordinate2D[int]

func neighbors(c coord) map[coordinate.Direction]coord {
	return c.Neighbors(coordinate.Straights())
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	g := s.parse(lines)
	cost := g.cost(false)
	return fmt.Sprint(cost)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	g := s.parse(lines)
	cost := g.cost(true)
	return fmt.Sprint(cost)
}
