package day12

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
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
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) garden {
	return parseGarden(lines)
}

type garden map[rune][]region

func parseGarden(lines []string) garden {
	plots := make(map[rune][]coord)
	for i, line := range lines {
		for j, c := range line {
			if !maps.Contains(plots, c) {
				plots[c] = make([]coord, 0)
			}
			plots[c] = append(plots[c], coordinate.NewCoordinate2D(j, i))
		}
	}
	g := make(garden)
	for k, coords := range plots {
		g[k] = separateRegions(coords)
	}
	return g
}

func separateRegions(plots []coord) []region {
	regions := make([]region, 0)
	for i := 0; len(plots) > 0; i = (i + 1) % len(plots) {
		current := plots[i]
		plots = append(plots[:i], plots[i+1:]...)
		i--
		currentRegion := make(region)
		currentRegion.add(current)
		neighborsToVisit := neighbors(current)
		for j := 0; j < len(neighborsToVisit); j++ {
			n := neighborsToVisit[j]
			if maps.Contains(currentRegion, n) {
				continue
			}
			index := slices.FindIndex(plots, func(c coord) bool {
				return c == n
			})
			if index == -1 {
				continue
			}
			currentRegion.add(n)
			neighborsToVisit = append(neighborsToVisit, neighbors(n)...)
			plots = append(plots[:index], plots[index+1:]...)
			if index < i {
				i--
			}
		}
		regions = append(regions, currentRegion)
		if len(plots) == 0 {
			break
		}
	}
	return regions
}

func (g garden) cost(bulk bool) int {
	costs := slices.Map(maps.Values(g), func(regions []region) int {
		return slices.Sum(slices.Map(regions, func(r region) int { return r.cost(bulk) }))
	})
	return slices.Sum(costs)
}

type region map[coord]struct{}

func (r *region) add(c coord) {
	(*r)[c] = struct{}{}
}

func (r region) area() int {
	return len(r)
}

func (r region) cost(bulk bool) int {
	if bulk {
		return r.area() * r.sides()
	}
	return r.area() * r.perimeter()
}

func (r region) perimeter() int {
	isDifferentPlant := func(c coord) bool {
		return !maps.Contains(r, c)
	}
	plotPerimeter := func(c coord) int {
		return slices.Count(neighbors(c), isDifferentPlant)
	}
	perimeters := slices.Map(maps.Keys(r), plotPerimeter)
	return slices.Sum(perimeters)
}

type fence struct {
	c   coord
	dir coordinate.Direction
}

func (f fence) neighbors() []fence {
	if f.dir == coordinate.Up() || f.dir == coordinate.Down() {
		return []fence{
			{f.c.Go(coordinate.Left()), f.dir},
			{f.c.Go(coordinate.Right()), f.dir},
		}
	}
	return []fence{
		{f.c.Go(coordinate.Up()), f.dir},
		{f.c.Go(coordinate.Down()), f.dir},
	}
}

// String returns the string representation of the fence
func (f fence) String() string {
	return fmt.Sprintf("%v %v", f.c, f.dir)
}

func (r region) sides() int {
	fences := make(map[fence]struct{})
	for c := range r {
		for i, n := range neighbors(c) {
			if !maps.Contains(r, n) {
				fences[fence{c, dirs[i]}] = struct{}{}
			}
		}
	}
	count := 0
	for f := range fences {
		delete(fences, f)
		neighborFences := f.neighbors()
		for i := 0; i < len(neighborFences); i++ {
			nf := neighborFences[i]
			_, ok := fences[nf]
			if !ok {
				continue
			}
			delete(fences, nf)
			neighborFences = append(neighborFences, nf.neighbors()...)
		}
		count++
	}
	return count
}

type coord = coordinate.Coordinate2D[int]

var dirs = []coordinate.Direction{coordinate.Up(), coordinate.Right(), coordinate.Down(), coordinate.Left()}

func neighbors(c coord) []coord {
	return c.Neighbors(dirs)
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
