package day04

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/iterutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 4

// init registers the solver for day 04
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 04
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) map[rune][]coordinate.Coordinate2D[int] {
	coordinates := make(map[rune][]coordinate.Coordinate2D[int])
	for i, line := range lines {
		for j, c := range line {
			coordinates[c] = append(coordinates[c], coordinate.FromIndexes(i, j))
		}
	}
	return coordinates
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	coordinates := s.parse(lines)
	count := 0
	for _, c := range coordinates['X'] {
		count += countwords(coordinates, c, "XMAS")
	}
	return fmt.Sprint(count)
}

// countwords counts how many times the word can be found starting from the given coordinate
func countwords(coordinates map[rune][]coordinate.Coordinate2D[int], start coordinate.Coordinate2D[int], word string) int {
	directions := coordinate.Directions()
	counter := func(d coordinate.Direction) bool {
		return containsWordInDirection(coordinates, start, word, d)
	}
	return iterutils.Count(iterutils.NewFromSlice(directions), counter)
}

// containsWordInDirection checks if the word can be found starting from the given coordinate in the given direction
func containsWordInDirection(coordinates map[rune][]coordinate.Coordinate2D[int], start coordinate.Coordinate2D[int], word string, d coordinate.Direction) bool {
	c := start
	for _, letter := range word {
		if !slices.Contains(coordinates[letter], c) {
			return false
		}
		c.Move(d)
	}
	return true
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	coordinates := s.parse(lines)
	isX := func(c coordinate.Coordinate2D[int]) bool {
		return isX(coordinates, c)
	}
	count := iterutils.Count(iterutils.NewFromSlice(coordinates['A']), isX)
	return fmt.Sprint(count)
}

// isX checks if the coordinate is the center of an X shape
func isX(coordinates map[rune][]coordinate.Coordinate2D[int], start coordinate.Coordinate2D[int]) bool {
	line1 := []coordinate.Coordinate2D[int]{start.Go(coordinate.UpLeft()), start.Go(coordinate.DownRight())}
	line2 := []coordinate.Coordinate2D[int]{start.Go(coordinate.UpRight()), start.Go(coordinate.DownLeft())}
	lines := [][]coordinate.Coordinate2D[int]{line1, line2}
	for _, line := range lines {
		if !(isPartOfX(coordinates, line[0]) && isPartOfX(coordinates, line[1]) && !isSame(coordinates, line[0], line[1])) {
			return false
		}
	}
	return true
}

// isPartOfX checks if the coordinate is part of the X shape (either M or S)
func isPartOfX(coordinates map[rune][]coordinate.Coordinate2D[int], c coordinate.Coordinate2D[int]) bool {
	return slices.Contains(coordinates['M'], c) || slices.Contains(coordinates['S'], c)
}

// isSame checks if both coordinates are of the same type (both M or both S)
func isSame(coordinates map[rune][]coordinate.Coordinate2D[int], c1, c2 coordinate.Coordinate2D[int]) bool {
	return (slices.Contains(coordinates['M'], c1) && slices.Contains(coordinates['M'], c2)) ||
		(slices.Contains(coordinates['S'], c1) && slices.Contains(coordinates['S'], c2))
}
