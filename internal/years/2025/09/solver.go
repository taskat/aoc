package day09

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 9

// init registers the solver for day 09
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 09
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []coordinate.Coordinate2D[int] {
	return slices.Map(lines, parseCoord)
}

// parseCoord parses a coordinate from a string
func parseCoord(s string) coordinate.Coordinate2D[int] {
	var i, j int
	fmt.Sscanf(s, "%d,%d", &i, &j)
	return coordinate.FromIndexes(i, j)
}

// sizeOfRectangle computes the size of the rectangle defined by two coordinates
func sizeOfRectangle(c1, c2 coordinate.Coordinate2D[int]) int {
	width := intutils.Abs(c1.X-c2.X) + 1
	height := intutils.Abs(c1.Y-c2.Y) + 1
	return width * height
}

// findMaxRectangle finds the maximum rectangle from a list of coordinates
func findMaxRectangle(coords []coordinate.Coordinate2D[int]) int {
	maxSize := 0
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			size := sizeOfRectangle(c1, c2)
			if size > maxSize {
				maxSize = size
			}
		}
	}
	return maxSize
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	coords := s.parse(lines)
	maxRectSize := findMaxRectangle(coords)
	return fmt.Sprintf("%d", maxRectSize)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
