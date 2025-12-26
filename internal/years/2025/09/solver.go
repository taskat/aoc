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
			r := newRectangle(c1, c2)
			if r.size > maxSize {
				maxSize = r.size
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

// newRectangle creates a new rectangle
func newRectangle(c1, c2 coordinate.Coordinate2D[int]) rectangle {
	r := rectangle{c1: c1, c2: c2}
	width := intutils.Abs(r.c1.X-r.c2.X) + 1
	height := intutils.Abs(r.c1.Y-r.c2.Y) + 1
	r.size = width * height
	return r
}

// compress compresses the rectangle coordinates
func (r *rectangle) compress(compression *coordinate.Compression2D[int]) {
	r.c1 = compression.Compress(r.c1)
	r.c2 = compression.Compress(r.c2)
}

// createRectangles creates all rectangles from a list of coordinates
func createRectangles(coords []coordinate.Coordinate2D[int], anomalies []coordinate.Coordinate2D[int], compression *coordinate.Compression2D[int]) []rectangle {
	var rectangles []rectangle
	if len(anomalies) == 0 {
		for i, c1 := range coords {
			for _, c2 := range coords[i+1:] {
				r := newRectangle(c1, c2)
				r.compress(compression)
				rectangles = append(rectangles, r)
			}
		}
		return rectangles
	}
	for _, anomaly := range anomalies {
		for _, c := range coords {
			if anomaly != c {
				r := newRectangle(anomaly, c)
				r.compress(compression)
				rectangles = append(rectangles, r)
			}
		}
	}
	return rectangles
}

// line represents a vertical or horizontal line segment
type line struct {
	from, to       int
	otherDimension int
}

// hasIntersection checks if two lines intersect
func (l1 line) hasIntersection(l2 line) bool {
	return l1.from <= l2.otherDimension && l1.to >= l2.otherDimension &&
		l2.from <= l1.otherDimension && l2.to >= l1.otherDimension
}

// edges represents a collection of horizontal and vertical edges
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
		fromY, toY := increasing(c1.Y, c2.Y)
		e.addVerticalEdge(c1.X, fromY, toY)
		return
	} else if c1.Y == c2.Y {
		fromX, toX := increasing(c1.X, c2.X)
		e.addHorizontalEdge(c1.Y, fromX, toX)
		return
	}
	panic("coordinates are not aligned")
}

// addVerticalEdge adds a vertical edge
func (e *edges) addVerticalEdge(x, fromY, toY int) {
	e.vertical = append(e.vertical, line{from: fromY, to: toY, otherDimension: x})
}

// addHorizontalEdge adds a horizontal edge
func (e *edges) addHorizontalEdge(y, fromX, toX int) {
	e.horizontal = append(e.horizontal, line{from: fromX, to: toX, otherDimension: y})
}

// isPointOnEdge checks if a point is on any edge
func (e *edges) isPointOnEdge(c coordinate.Coordinate2D[int]) bool {
	for _, h := range e.horizontal {
		if c.Y == h.otherDimension && c.X >= h.from && c.X <= h.to {
			return true
		}
	}
	for _, v := range e.vertical {
		if c.X == v.otherDimension && c.Y >= v.from && c.Y <= v.to {
			return true
		}
	}
	return false
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

// isPointInside checks if a point is inside the lines
func (e edges) isPointInside(c coordinate.Coordinate2D[int]) bool {
	if e.isPointOnEdge(c) {
		return true
	}
	intersections := 0
	l := line{from: 0, to: c.X, otherDimension: c.Y}
	for _, v := range e.vertical {
		if v.hasIntersection(l) {
			intersections++
		}
	}
	if intersections%2 == 1 {
		return true
	}
	extraLine := e.getPossibleIntersectorLine(c)
	if extraLine == nil {
		return false
	}
	return extraLine.hasIntersection(l)
}

// getPossibleIntersectorLine returns a new line that is possibly an intersector
// It is used when the intersecion counting could fail due to the point is aligned with an edge
// but not on the edge. If no such line can be found, it returns nil
func (e edges) getPossibleIntersectorLine(c coordinate.Coordinate2D[int]) *line {
	x1, x2 := -1, -1
	for _, h := range e.horizontal {
		if h.otherDimension == c.Y && h.from < c.X {
			x1 = h.from
			x2 = h.to
			break
		}
	}
	if x1 == -1 && x2 == -1 {
		return nil
	}
	y1, y2 := -1, -1
	for _, v := range e.vertical {
		if v.otherDimension == x1 {
			if v.from == c.Y {
				y1 = v.to
			} else {
				y1 = v.from
			}
		}
		if v.otherDimension == x2 {
			if v.from == c.Y {
				y2 = v.to
			} else {
				y2 = v.from
			}
		}
	}
	if y1 == -1 || y2 == -1 {
		panic("invalid edge configuration")
	}
	y_from, y_to := increasing(y1, y2)
	extraLine := line{from: y_from, to: y_to, otherDimension: x1}
	return &extraLine
}

// isPointInsideWithCache checks if a point is inside the lines with caching
func (e edges) isPointInsideWithCache(c coordinate.Coordinate2D[int], cache map[coordinate.Coordinate2D[int]]bool) bool {
	val, ok := cache[c]
	if ok {
		return val
	}
	val = e.isPointInside(c)
	cache[c] = val
	return val
}

// isRectangleInside checks if a rectangle is entirely inside the lines
// It checks all four edges of the rectangle
func (e edges) isRectangleInside(r rectangle, cache map[coordinate.Coordinate2D[int]]bool) bool {
	xFrom, xTo := increasing(r.c1.X, r.c2.X)
	yFrom, yTo := increasing(r.c1.Y, r.c2.Y)
	for x := xFrom; x <= xTo; x++ {
		if !e.isPointInsideWithCache(coordinate.NewCoordinate2D(x, yFrom), cache) ||
			!e.isPointInsideWithCache(coordinate.NewCoordinate2D(x, yTo), cache) {
			return false
		}
	}
	for y := yFrom; y <= yTo; y++ {
		if !e.isPointInsideWithCache(coordinate.NewCoordinate2D(xFrom, y), cache) ||
			!e.isPointInsideWithCache(coordinate.NewCoordinate2D(xTo, y), cache) {
			return false
		}
	}
	return true
}

// increasing returns the two integers in increasing order
func increasing(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// getAnomalyCoordinates finds the coordinates of anomalies
func getAnomalyCoordinates(coords []coordinate.Coordinate2D[int]) []coordinate.Coordinate2D[int] {
	maxLength := 0
	maxIdx := -1
	for i, c1 := range coords {
		nextIdx := (i + 1) % len(coords)
		c2 := coords[nextIdx]
		length := newRectangle(c1, c2).size
		if length > maxLength {
			maxLength = length
			maxIdx = i
		}
	}
	line1 := newRectangle(coords[maxIdx], coords[(maxIdx+1)%len(coords)])
	line2 := newRectangle(coords[(maxIdx+2)%len(coords)], coords[(maxIdx+3)%len(coords)])
	if line1.size > 1000 && line2.size > 1000 {
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
	compressedCoords, compression := coordinate.Compress2D(coords)
	edges := getEdges(compressedCoords)
	anomalies := getAnomalyCoordinates(coords)
	rectangles := createRectangles(coords, anomalies, compression)
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i].size > rectangles[j].size })
	cache := make(map[coordinate.Coordinate2D[int]]bool)
	rectangles = slices.Filter(rectangles, func(r rectangle) bool {
		return edges.isRectangleInside(r, cache)
	})
	return fmt.Sprintf("%d", rectangles[0].size)
}
