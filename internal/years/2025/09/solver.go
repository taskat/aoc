package day09

import (
	"fmt"
	"sort"

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
	var x, y int
	fmt.Sscanf(s, "%d,%d", &x, &y)
	return coordinate.NewCoordinate2D(x, y)
}

// findMaxRectangle finds the maximum rectangle from a list of coordinates
func findMaxRectangle(coords []coordinate.Coordinate2D[int]) int {
	maxSize := 0
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {
			r := rectangle{c1: c1, c2: c2}
			size := r.getSize()
			if size > maxSize {
				maxSize = size
			}
		}
	}
	return maxSize
}

// rectangle represents a rectangle defined by two opposite corners
type rectangle struct {
	c1, c2 coordinate.Coordinate2D[int]
	size   int
}

// getSize computes the getSize of the rectangle
func (r *rectangle) getSize() int {
	if r.size != 0 {
		return r.size
	}
	width := intutils.Abs(r.c1.X-r.c2.X) + 1
	height := intutils.Abs(r.c1.Y-r.c2.Y) + 1
	r.size = width * height
	return r.size
}

// createRectangles creates all rectangles from a list of coordinates
func createRectangles(coords []coordinate.Coordinate2D[int], anomalies []coordinate.Coordinate2D[int]) []rectangle {
	var rectangles []rectangle
	if len(anomalies) == 0 {
		for i, c1 := range coords {
			for _, c2 := range coords[i+1:] {
				rectangles = append(rectangles, rectangle{c1: c1, c2: c2})
			}
		}
		return rectangles
	}
	for _, anomaly := range anomalies {
		for _, c := range coords {
			if anomaly != c {
				rectangles = append(rectangles, rectangle{c1: anomaly, c2: c})
			}
		}
	}
	return rectangles
}

type line struct {
	from, to       int
	otherDimension int
}

// hasIntersection checks if two lines intersect
func (l1 line) hasIntersection(l2 line) bool {
	return l1.from < l2.otherDimension && l1.to > l2.otherDimension &&
		l2.from < l1.otherDimension && l2.to > l1.otherDimension
}

type edges struct {
	horizontal []line
	vertical   []line
}

// newEdges creates a new edges instance
func newEdges() edges {
	return edges{
		horizontal: []line{},
		vertical:   []line{},
	}
}

// addEdge adds an edge coordinate to the edges
func (e *edges) addEdge(c1, c2 coordinate.Coordinate2D[int]) {
	if c1.X == c2.X {
		fromY, toY := c1.Y, c2.Y
		if fromY > toY {
			fromY, toY = toY, fromY
		}
		e.vertical = append(e.vertical, line{from: fromY, to: toY, otherDimension: c1.X})
		return
	} else if c1.Y == c2.Y {
		fromX, toX := c1.X, c2.X
		if fromX > toX {
			fromX, toX = toX, fromX
		}
		e.horizontal = append(e.horizontal, line{from: fromX, to: toX, otherDimension: c1.Y})
		return
	}
	panic("coordinates are not aligned")
}

// getEdges returns the edge coordinates of the given coordinate list
func getEdges(coords []coordinate.Coordinate2D[int]) edges {
	edges := newEdges()
	for i, c1 := range coords {
		nextIndex := (i + 1) % len(coords)
		c2 := coords[nextIndex]
		edges.addEdge(c1, c2)
	}
	return edges
}

// hasIntersection checks if a line has an intersection with the edges
func (e *edges) hasIntersection(c1, c2 coordinate.Coordinate2D[int]) bool {
	if c1.X == c2.X {
		l := line{}
		fromY, toY := c1.Y, c2.Y
		if fromY > toY {
			fromY, toY = toY, fromY
		}
		l.from = fromY
		l.to = toY
		l.otherDimension = c1.X
		for _, h := range e.horizontal {
			if h.hasIntersection(l) {
				return true
			}
		}
		return false
	}
	if c1.Y == c2.Y {
		fromX, toX := c1.X, c2.X
		if fromX > toX {
			fromX, toX = toX, fromX
		}
		l := line{from: fromX, to: toX, otherDimension: c1.Y}
		for _, v := range e.vertical {
			if v.hasIntersection(l) {
				return true
			}
		}
		return false
	}
	return false
}

// isRectangleInside checks if a rectangle is entirely inside the lines
func isRectangleInside(r rectangle, edges edges) bool {
	fromX, toX := r.c1.X, r.c2.X
	if fromX > toX {
		fromX, toX = toX, fromX
	}
	for x := fromX; x <= toX; x++ {
		c1 := coordinate.NewCoordinate2D(x, r.c1.Y)
		c2 := coordinate.NewCoordinate2D(x, r.c2.Y)
		if edges.hasIntersection(c1, c2) {
			return false
		}
	}
	fromY, toY := r.c1.Y, r.c2.Y
	if fromY > toY {
		fromY, toY = toY, fromY
	}
	for y := fromY; y <= toY; y++ {
		c1 := coordinate.NewCoordinate2D(r.c1.X, y)
		c2 := coordinate.NewCoordinate2D(r.c2.X, y)
		if edges.hasIntersection(c1, c2) {
			return false
		}
	}
	return true
}

// getMaxRectangleInside finds the maximum rectangle entirely inside the lines
// It assumes that the rectangles are sorted by size in descending order
func getMaxRectangleInside(rectangles []rectangle, edges edges) rectangle {
	for _, r := range rectangles {
		if isRectangleInside(r, edges) {
			return r
		}
	}
	panic("no rectangle found inside the lines")
}

// getAnomalyCoordinates finds the coordinates of anomalies
func getAnomalyCoordinates(coords []coordinate.Coordinate2D[int]) []coordinate.Coordinate2D[int] {
	maxLength := 0
	maxIdx := -1
	for i, c1 := range coords {
		nextIdx := (i + 1) % len(coords)
		c2 := coords[nextIdx]
		length := (&rectangle{c1: c1, c2: c2}).getSize()
		if length > maxLength {
			maxLength = length
			maxIdx = i
		}
	}
	line1 := rectangle{c1: coords[maxIdx], c2: coords[(maxIdx+1)%len(coords)]}
	line2 := rectangle{c1: coords[(maxIdx+2)%len(coords)], c2: coords[(maxIdx+3)%len(coords)]}
	if line1.getSize() > 1000 && line2.getSize() > 1000 {
		return []coordinate.Coordinate2D[int]{line1.c2, line2.c1}
	}
	return nil
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	coords := s.parse(lines)
	maxRectSize := findMaxRectangle(coords)
	return fmt.Sprintf("%d", maxRectSize)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	coords := s.parse(lines)
	edges := getEdges(coords)
	anomalies := getAnomalyCoordinates(coords)
	rectangles := createRectangles(coords, anomalies)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].getSize() > rectangles[j].getSize()
	})
	r := getMaxRectangleInside(rectangles, edges)
	return fmt.Sprintf("%d", r.getSize())
}
