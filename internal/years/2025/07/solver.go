package day07

import (
	"fmt"
	"slices"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 7

// init registers the solver for day 07
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 07
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) diagram {
	d := diagram{
		height:    len(lines),
		splitters: make(map[int][]coordinate.Coordinate2D[int]),
	}
	for i, line := range lines {
		splitters := make([]coordinate.Coordinate2D[int], 0)
		for j, char := range line {
			switch char {
			case 'S':
				d.start = coordinate.FromIndexes(i, j)
			case '.':
				continue
			case '^':
				splitters = append(splitters, coordinate.FromIndexes(i, j))
			}
		}
		d.splitters[i] = splitters
	}
	return d
}

// diagram represents the tachyon manifold diagram
type diagram struct {
	start     coordinate.Coordinate2D[int]
	height    int
	splitters map[int][]coordinate.Coordinate2D[int]
}

// countSplits counts the number of splits of the tachyon beam
func (d *diagram) countSplits() int {
	splits := 0
	beams := set.New[int]()
	beams.Add(d.start.X)
	for i := range d.height {
		if len(d.splitters[i]) == 0 {
			continue
		}
		nextBeams := set.New[int]()
		for beam := range beams {
			matchX := func(c coordinate.Coordinate2D[int]) bool { return beam == c.X }
			if slices.ContainsFunc(d.splitters[i], matchX) {
				nextBeams.Add(beam - 1)
				nextBeams.Add(beam + 1)
				splits++
			} else {
				nextBeams.Add(beam)
			}
		}
		beams = nextBeams
	}
	return splits
}

// countTimelines counts the different timelines of the tachyon particle
func (d *diagram) countTimelines() int {
	timelines := make(map[int]int)
	timelines[d.start.X] = 1
	for i := range d.height {
		if len(d.splitters[i]) == 0 {
			continue
		}
		nextTimelines := make(map[int]int)
		for timeline, count := range timelines {
			matchX := func(c coordinate.Coordinate2D[int]) bool { return timeline == c.X }
			if slices.ContainsFunc(d.splitters[i], matchX) {
				nextTimelines[timeline-1] += count
				nextTimelines[timeline+1] += count
			} else {
				nextTimelines[timeline] += count
			}
		}
		timelines = nextTimelines
	}
	return maps.Sum(timelines)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	d := s.parse(lines)
	splits := d.countSplits()
	return fmt.Sprintf("%d", splits)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	d := s.parse(lines)
	timelines := d.countTimelines()
	return fmt.Sprintf("%d", timelines)
}
