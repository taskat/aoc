package day08

import (
	"fmt"
	"sort"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 8

// init registers the solver for day 08
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 08
type Solver struct {
	numberOfConnections int
}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {
	if len(params) > 0 {
		fmt.Sscanf(params[0], "%d", &s.numberOfConnections)
	}
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []coordinate.Coordinate3D[int] {
	return slices.Map(lines, parseCoordinate)
}

// parseCoordinate parses a coordinate from a string
func parseCoordinate(line string) coordinate.Coordinate3D[int] {
	var x, y, z int
	fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
	return coordinate.NewCoordinate3D(x, y, z)
}

// circuit represents a circuit of junction boxes.
// It contains the id of the junction boxes
type circuit []int

// hasBox checks if the circuit contains the given junction box
func (c circuit) hasBox(box int) bool {
	return slices.Contains(c, box)
}

// length returns the length of the circuit
func (c circuit) length() int {
	return len(c)
}

// merge merges two circuits into one
func (c circuit) merge(other circuit) circuit {
	return append(c, other...)
}

// distance represents a pair of junction boxes and the distance between them
type distance struct {
	from     int
	to       int
	distance float64
}

// distances calculates the distances between all pairs of junction boxes
func distances(coords []coordinate.Coordinate3D[int]) []distance {
	distances := make([]distance, 0)
	for i, c1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]
			distances = append(distances, distance{from: i, to: j, distance: c1.Distance(c2)})
		}
	}
	return distances
}

// circuits generates all the circuits that contain a single junction box
func circuits(numberOfBoxes int) []circuit {
	circuits := make([]circuit, numberOfBoxes)
	for i := range numberOfBoxes {
		circuits[i] = circuit{i}
	}
	return circuits
}

// connectBoxes connects the junction boxes based on the given distances
func connectBoxes(circuits []circuit, distances []distance) []circuit {
	for _, d := range distances {
		fromCircuitIndex := -1
		toCircuitIndex := -1
		for i, c := range circuits {
			if c.hasBox(d.from) {
				fromCircuitIndex = i
			}
			if c.hasBox(d.to) {
				toCircuitIndex = i
			}
			if fromCircuitIndex != -1 && toCircuitIndex != -1 {
				break
			}
		}
		if fromCircuitIndex == toCircuitIndex {
			continue
		}
		circuits[fromCircuitIndex] = circuits[fromCircuitIndex].merge(circuits[toCircuitIndex])
		circuits = append(circuits[:toCircuitIndex], circuits[toCircuitIndex+1:]...)
	}
	return circuits
}

// findFinishingConnection finds the finishing connection in the distances
func findFinishingConnection(circuits []circuit, distances []distance) distance {
	for _, d := range distances {
		fromCircuitIndex := -1
		toCircuitIndex := -1
		for i, c := range circuits {
			if c.hasBox(d.from) {
				fromCircuitIndex = i
			}
			if c.hasBox(d.to) {
				toCircuitIndex = i
			}
			if fromCircuitIndex != -1 && toCircuitIndex != -1 {
				break
			}
		}
		if fromCircuitIndex == toCircuitIndex {
			continue
		}
		circuits[fromCircuitIndex] = circuits[fromCircuitIndex].merge(circuits[toCircuitIndex])
		circuits = append(circuits[:toCircuitIndex], circuits[toCircuitIndex+1:]...)
		if len(circuits) == 1 {
			return d
		}
	}
	panic("no finishing connection found")
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	coords := s.parse(lines)
	distances := distances(coords)
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
	distances = distances[:s.numberOfConnections]
	circuits := circuits(len(coords))
	circuits = connectBoxes(circuits, distances)
	lengths := slices.Map(circuits, circuit.length)
	sort.Ints(lengths)
	product := slices.Product(lengths[len(lengths)-3:])
	return fmt.Sprintf("%d", product)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	coords := s.parse(lines)
	distances := distances(coords)
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
	circuits := circuits(len(coords))
	finishingDistance := findFinishingConnection(circuits, distances)
	fromCoord := coords[finishingDistance.from]
	toCoord := coords[finishingDistance.to]
	product := fromCoord.X * toCoord.X
	return fmt.Sprintf("%d", product)
}
