package day15

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 15

// init registers the solver for day 15
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 15
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) warehouse {
	return parseWarehouse(lines)
}

// warehouse contains the state of the warehouse. It stores the walls, boxes,
// robot, the size of the warehouse and the moves, that the robot should make.
type warehouse struct {
	walls set.Set[coord]
	boxes set.Set[coord]
	limit coord
	robot coord
	moves []coordinate.Direction
}

// parseWarehouse parses the warehouse from the input lines. It returns the parsed warehouse.
func parseWarehouse(lines []string) warehouse {
	w := warehouse{
		walls: set.New[coord](),
		boxes: set.New[coord](),
	}
	idx := slices.FindIndex(lines, stringutils.IsEmpty)
	for y, line := range lines[:idx] {
		for x, c := range line {
			coord := coord{X: x, Y: y}
			switch c {
			case '#':
				w.walls.Add(coord)
			case 'O':
				w.boxes.Add(coord)
			case '@':
				w.robot = coord
			}
		}
	}
	limitX := len(lines[0])
	limitY := idx
	w.limit = coord{X: limitX, Y: limitY}
	coordinate.Format = coordinate.Characters
	slices.ForEach(lines[idx+1:], func(line string) {
		for _, c := range line {
			w.moves = append(w.moves, coordinate.MustParse(string(c)))
		}
	})
	return w
}

// moveRobot moves the robot in the warehouse in the given direction one step.
// If the robot hits a wall, it stops. If the robot hits a box, it tries to move the box.
// It can push multiple boxes if there are multiple boxes in a row.
func (w *warehouse) moveRobot(dir coordinate.Direction) {
	robotNext := w.robot.Go(dir)
	next := robotNext
	for {
		switch {
		case w.walls.Contains(next):
			return
		case w.boxes.Contains(next):
			next = next.Go(dir)
		default:
			w.robot = robotNext
			if next != robotNext {
				w.boxes.Delete(w.robot)
				w.boxes.Add(next)
			}
			return
		}
	}
}

// moveRobotInBiggerWarehouse moves the robot in the warehouse in the given direction one step.
// If the robot hits a wall, it stops. If the robot hits a box, it tries to move the box.
// It can push multiple boxes if there are multiple boxes in a row. It also can push multiple boxes
// if they are offset by one cell (the boxes are 2 cell wide in the bigger warehouse).
func (w *warehouse) moveRobotInBiggerWarehouse(dir coordinate.Direction) {
	robotNext := w.robot.Go(dir)
	boxesToPush := make([]coord, 0, len(w.boxes))
	if dir.Horizontal() {
		next := robotNext
		stopLoop := false
		for !stopLoop {
			switch {
			case w.walls.Contains(next):
				return
			case w.boxes.Contains(next):
				boxesToPush = append(boxesToPush, next)
				next = next.GoN(dir, 2)
			case dir == coordinate.Left() && w.boxes.Contains(next.Go(coordinate.Left())):
				boxesToPush = append(boxesToPush, next.Go(coordinate.Left()))
				next = next.GoN(coordinate.Left(), 2)
			default:
				stopLoop = true
			}
		}
	} else {
		nexts := []coord{robotNext}
		for len(nexts) > 0 {
			newNexts := make([]coord, 0, len(nexts))
			for _, n := range nexts {
				switch {
				case w.walls.Contains(n):
					return
				case w.boxes.Contains(n):
					newNexts = append(newNexts, n.Go(dir))
					newNexts = append(newNexts, n.Go(dir).Go(coordinate.Right()))
					boxesToPush = append(boxesToPush, n)
				case w.boxes.Contains(n.Go(coordinate.Left())):
					newNexts = append(newNexts, n.Go(dir))
					newNexts = append(newNexts, n.Go(dir).Go(coordinate.Left()))
					boxesToPush = append(boxesToPush, n.Go(coordinate.Left()))
				}
			}
			nexts = newNexts
		}
	}
	w.robot = robotNext
	for i := len(boxesToPush) - 1; i >= 0; i-- {
		box := boxesToPush[i]
		w.boxes.Delete(box)
		w.boxes.Add(box.Go(dir))
	}
}

// simulate simulates the moves in the warehouse.
func (w *warehouse) simulate(moveFunc func(wh *warehouse, dir coordinate.Direction)) {
	move := func(dir coordinate.Direction) {
		moveFunc(w, dir)
	}
	slices.ForEach(w.moves, move)
}

// String returns the string representation of the warehouse.
// The robot is represented by '@', the walls by '#' and the boxes by 'O'.
func (wh warehouse) String() string {
	grid := make([][]rune, wh.limit.Y)
	for i := range grid {
		grid[i] = slices.Repeat('.', wh.limit.X)
	}
	set := func(r rune) func(c coord) {
		return func(c coord) {
			grid[c.Y][c.X] = r
		}
	}
	slices.ForEach(wh.walls.ToSlice(), set('#'))
	slices.ForEach(wh.boxes.ToSlice(), set('O'))
	grid[wh.robot.Y][wh.robot.X] = '@'
	lines := make([]string, len(grid))
	for i, row := range grid {
		lines[i] = string(row)
	}
	return strings.Join(lines, "\n")
}

// upscale upscales the warehouse by the given factor in horizontal direction.
// It also fills the holes in the wall created by upscaling.
func (wh *warehouse) upscale(n int) {
	wh.limit = upscale(wh.limit, n)
	wh.robot = upscale(wh.robot, n)
	wh.walls = set.FromSlice(slices.Map(wh.walls.ToSlice(), func(c coord) coord { return upscale(c, n) }))
	extraWalls := set.Map(wh.walls, func(c coord) coord { return coord{X: c.X + 1, Y: c.Y} })
	wh.walls = wh.walls.Merge(extraWalls)
	wh.boxes = set.FromSlice(slices.Map(wh.boxes.ToSlice(), func(c coord) coord { return upscale(c, n) }))
}

// coord is a wrapper around coordinate.Coordinate2D[int] to make it easier to use.
type coord = coordinate.Coordinate2D[int]

// gps calculates the GPS coordinate from the given coordinate.
func gps(c coord) int {
	return 100*c.Y + c.X
}

// upscale upscales the coordinate by the given factor in horizontal direction.
func upscale(c coord, n int) coord {
	return coord{X: c.X * n, Y: c.Y}
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	wh := s.parse(lines)
	wh.simulate((*warehouse).moveRobot)
	gpsCoordinates := slices.Map(wh.boxes.ToSlice(), gps)
	sum := slices.Sum(gpsCoordinates)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	wh := s.parse(lines)
	wh.upscale(2)
	wh.simulate((*warehouse).moveRobotInBiggerWarehouse)
	gpsCoordinates := slices.Map(wh.boxes.ToSlice(), gps)
	sum := slices.Sum(gpsCoordinates)
	return fmt.Sprint(sum)
}
