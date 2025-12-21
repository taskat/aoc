package day12

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/slices"
)

// day is the day of the solver
const day = 12

// init registers the solver for day 12
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 12
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) ([]present, []region) {
	presents := []present{}
	for !strings.Contains(lines[0], "x") {
		p, consumed := parsePresents(lines)
		presents = append(presents, p)
		lines = lines[consumed:]
	}
	regions := slices.Map(lines, parseRegion)
	return presents, regions
}

// present represents the different present blocks
type present struct {
	id     int
	layout [][]bool
	size   int
}

// parsePresents parses the present from the input lines, until an empty line is found
// The first line expected to be in the format "<id>:", and the following lines represent the layout
// It returns the parsed present and the number of lines consumed
func parsePresents(lines []string) (present, int) {
	var id int
	fmt.Sscanf(lines[0], "%d:", &id)
	layout := [][]bool{}
	size := 0
	linesConsumed := 1
	for _, line := range lines[1:] {
		linesConsumed++
		if line == "" {
			break
		}
		row := []bool{}
		for _, ch := range line {
			if ch == '#' {
				row = append(row, true)
				size++
			} else {
				row = append(row, false)
			}
		}
		layout = append(layout, row)
	}
	return present{
		id:     id,
		layout: layout,
		size:   size,
	}, linesConsumed
}

// dimension returns the dimensions of the present as width and height
func (p present) dimension() (int, int) {
	if len(p.layout) == 0 {
		return 0, 0
	}
	return len(p.layout[0]), len(p.layout)
}

// region represents  region which should be filled with presents
type region struct {
	width    int
	height   int
	presents []int
}

// parseRegion parses the region from the input line
// The line is expected to be in he format "<width>x<height>: <number of presents>"
func parseRegion(line string) region {
	var width, height int
	parts := strings.Split(line, ": ")
	fmt.Sscanf(parts[0], "%dx%d", &width, &height)
	parts = strings.Split(parts[1], " ")
	presents := []int{}
	for _, p := range parts {
		var pid int
		fmt.Sscanf(p, "%d", &pid)
		presents = append(presents, pid)
	}
	return region{
		width:    width,
		height:   height,
		presents: presents,
	}
}

// area computes the area of the region
func (r region) area() int {
	return r.width * r.height
}

// fitByArea checks if the sum of the presents' areas can fit in the region area
func (r region) fitByArea(presents []present) bool {
	presentArea := 0
	for id, present := range presents {
		presentArea += present.size * r.presents[id]
	}
	return presentArea <= r.area()
}

// fitByBlock checks if the presents can fit in the region by block size,
// meaning the enclosing rectangles of the presents fits in the region
func (r region) fitByBlock(presents []present) bool {
	maxWidth := 0
	maxHeight := 0
	for _, present := range presents {
		pw, ph := present.dimension()
		if pw > maxWidth {
			maxWidth = pw
		}
		if ph > maxHeight {
			maxHeight = ph
		}
	}
	sumOfPresents := slices.Sum(r.presents)
	presentsHorizontally := r.width / maxWidth
	presentsVertically := r.height / maxHeight
	return presentsHorizontally*presentsVertically >= sumOfPresents
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	presents, regions := s.parse(lines)
	surelyFits := 0
	surelyNotFits := 0
	unknown := 0
	for _, region := range regions {
		if !region.fitByArea(presents) {
			surelyNotFits++
		} else if region.fitByBlock(presents) {
			surelyFits++
		} else {
			unknown++
		}
	}
	if unknown != 0 {
		panic("There are unknown fit cases")
	}
	return fmt.Sprintf("%d", surelyFits)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
