package day08

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
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

type antennaMap struct {
	width    int
	height   int
	antennas map[rune][]coordinate
}

func parseAntennaMap(lines []string) antennaMap {
	am := antennaMap{
		width:    len(lines[0]),
		height:   len(lines),
		antennas: make(map[rune][]coordinate),
	}
	for i, line := range lines {
		for j, r := range line {
			if r != '.' {
				am.antennas[r] = append(am.antennas[r], coordinate{x: j, y: i})
			}
		}
	}
	return am
}

func (am antennaMap) getAntinodes() map[coordinate][]rune {
	antinodes := make(map[coordinate][]rune)
	for antennaType, antennas := range am.antennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				dist1 := antennas[i].distanceTo(antennas[j])
				dist2 := antennas[j].distanceTo(antennas[i])
				if am.inBounds(antennas[i].add(dist1)) {
					antinodes[antennas[i].add(dist1)] = append(antinodes[antennas[i].add(dist1)], antennaType)
				}
				if am.inBounds(antennas[j].add(dist2)) {
					antinodes[antennas[j].add(dist2)] = append(antinodes[antennas[j].add(dist2)], antennaType)
				}
			}
		}
	}
	return antinodes
}

func (am antennaMap) getAntinodesWithHarmonics() map[coordinate][]rune {
	antinodes := make(map[coordinate][]rune)
	for antennaType, antennas := range am.antennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				dist1 := antennas[i].distanceTo(antennas[j])
				dist2 := antennas[j].distanceTo(antennas[i])
				for point := antennas[i]; am.inBounds(point); point = point.add(dist1) {
					antinodes[point] = append(antinodes[point], antennaType)
				}
				for point := antennas[j]; am.inBounds(point); point = point.add(dist2) {
					antinodes[point] = append(antinodes[point], antennaType)
				}
			}
		}
	}
	return antinodes
}

func (am antennaMap) inBounds(c coordinate) bool {
	return c.x >= 0 && c.x < am.width && c.y >= 0 && c.y < am.height
}

type coordinate struct {
	x, y int
}

func (c coordinate) add(other coordinate) coordinate {
	return coordinate{x: c.x + other.x, y: c.y + other.y}
}

func (c coordinate) distanceTo(other coordinate) coordinate {
	return coordinate{x: c.x - other.x, y: c.y - other.y}
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	antennaMap := s.parse(lines)
	antinodes := antennaMap.getAntinodes()
	uniqueCount := len(antinodes)
	return fmt.Sprint(uniqueCount)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	antennaMap := s.parse(lines)
	antinodes := antennaMap.getAntinodesWithHarmonics()
	uniqueCount := len(antinodes)
	return fmt.Sprint(uniqueCount)
}
