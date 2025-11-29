package day08

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"

	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 8

// init registers the solver for day 08
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 08
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) antennaMap {
	return parseAntennaMap(lines)
}

// antennaMap represents the map of antennas
type antennaMap struct {
	width    int
	height   int
	antennas map[rune][]coordinate.Coordinate2D[int]
}

// parseAntennaMap parses the antenna map from the input lines.
// It does not store the whole map as an array, only the positions of the antennas
func parseAntennaMap(lines []string) antennaMap {
	am := antennaMap{
		width:    len(lines[0]),
		height:   len(lines),
		antennas: make(map[rune][]coordinate.Coordinate2D[int]),
	}
	for i, line := range lines {
		for j, r := range line {
			if r != '.' {
				am.antennas[r] = append(am.antennas[r], coordinate.FromIndexes(i, j))
			}
		}
	}
	return am
}

// getAntinodes returns the antinodes in the antenna map
// An antinode is a point that is on the line between two antennas of the same type,
// and the distance ratio between the two antennas is 2:1. If withHarmonics is true,
// it also includes the harmonics of the antinodes (points that are multiples of the distance
// from the antennas).
func (am antennaMap) getAntinodes(withHarmonics bool) map[coordinate.Coordinate2D[int]][]rune {
	antinodes := make(map[coordinate.Coordinate2D[int]][]rune)
	for antennaType, antennas := range am.antennas {
		for i := range len(antennas) {
			for j := i + 1; j < len(antennas); j++ {
				dist1 := antennas[i].Sub(antennas[j])
				dist2 := antennas[j].Sub(antennas[i])
				if withHarmonics {
					for point := antennas[i]; am.inBounds(point); point = point.Add(dist1) {
						antinodes[point] = append(antinodes[point], antennaType)
					}
					for point := antennas[j]; am.inBounds(point); point = point.Add(dist2) {
						antinodes[point] = append(antinodes[point], antennaType)
					}
				} else {
					if am.inBounds(antennas[i].Add(dist1)) {
						antinodes[antennas[i].Add(dist1)] = append(antinodes[antennas[i].Add(dist1)], antennaType)
					}
					if am.inBounds(antennas[j].Add(dist2)) {
						antinodes[antennas[j].Add(dist2)] = append(antinodes[antennas[j].Add(dist2)], antennaType)
					}
				}
			}
		}
	}
	return antinodes
}

// inBounds checks if the coordinate is within the bounds of the antenna map
func (am antennaMap) inBounds(c coordinate.Coordinate2D[int]) bool {
	return c.X >= 0 && c.X < am.width && c.Y <= 0 && c.Y > -am.height
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	antennaMap := s.parse(lines)
	antinodes := antennaMap.getAntinodes(false)
	uniqueCount := len(antinodes)
	return fmt.Sprint(uniqueCount)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	antennaMap := s.parse(lines)
	antinodes := antennaMap.getAntinodes(true)
	uniqueCount := len(antinodes)
	return fmt.Sprint(uniqueCount)
}
