package day15

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 15

// init registers the solver for day 15
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 15
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) warehouse {
	return parseWarehouse(lines)
}

type warehouse struct {
	walls set.Set[coord]
	boxes set.Set[coord]
	limit coord
	robot coord
	moves []coordinate.Direction
}

func parseWarehouse(lines []string) warehouse {
	w := warehouse{
		walls: set.New[coord](),
		boxes: set.New[coord](),
	}
	idx := slices.FindIndex(lines, func(s string) bool { return s == "" })
	for y, line := range lines {
		if y == idx {
			break
		}
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
			w.limit = coord
		}
	}
	w.limit.X++
	w.limit.Y++
	coordinate.Format = coordinate.Characters
	for i := idx + 1; i < len(lines); i++ {
		for _, c := range lines[i] {
			newMove, err := coordinate.Parse(string(c))
			if err != nil {
				panic(err)
			}
			w.moves = append(w.moves, newMove)
		}
	}
	return w
}

func (w *warehouse) moveRobot(dir coordinate.Direction) {
	next := w.robot.Go(dir)
	for {
		switch {
		case w.walls.Contains(next):
			return
		case w.boxes.Contains(next):
			next = next.Go(dir)
		default:
			w.robot = w.robot.Go(dir)
			if next != w.robot {
				w.boxes.Delete(w.robot)
				w.boxes.Add(next)
			}
			return
		}
	}
}

func (w *warehouse) moveRobotInBiggerWarehouse(dir coordinate.Direction) {
	nexts := []coord{w.robot.Go(dir)}
	boxes := make([]coord, 0, len(w.boxes))
	if dir.Horizontal() {
		for len(nexts) > 0 {
			next := nexts[0]
			switch {
			case w.walls.Contains(next):
				return
			case w.boxes.Contains(next):
				nexts = []coord{next.Go(dir).Go(dir)}
				boxes = append(boxes, next)
			case dir == coordinate.Left() && w.boxes.Contains(next.Go(coordinate.Left())):
				nexts = []coord{next.Go(coordinate.Left()).Go(coordinate.Left())}
				boxes = append(boxes, next.Go(coordinate.Left()))
			default:
				nexts = []coord{}
			}
		}
	} else {
		for len(nexts) > 0 {
			newNexts := make([]coord, 0, len(nexts))
			for _, n := range nexts {
				switch {
				case w.walls.Contains(n):
					return
				case w.boxes.Contains(n):
					newNexts = append(newNexts, n.Go(dir))
					newNexts = append(newNexts, n.Go(dir).Go(coordinate.Right()))
					boxes = append(boxes, n)
				case w.boxes.Contains(n.Go(coordinate.Left())):
					newNexts = append(newNexts, n.Go(dir))
					newNexts = append(newNexts, n.Go(dir).Go(coordinate.Left()))
					boxes = append(boxes, n.Go(coordinate.Left()))
				}
			}
			nexts = newNexts
		}
	}
	w.robot = w.robot.Go(dir)
	for i := len(boxes) - 1; i >= 0; i-- {
		box := boxes[i]
		w.boxes.Delete(box)
		w.boxes.Add(box.Go(dir))
	}
}

func (w *warehouse) simulate(moveFunc func(wh *warehouse, dir coordinate.Direction)) {
	for _, move := range w.moves {
		moveFunc(w, move)
	}
}

func (wh warehouse) String() string {
	grid := make([][]rune, wh.limit.Y)
	for i := range grid {
		grid[i] = slices.Repeat('.', wh.limit.X)
	}
	for wall := range wh.walls {
		grid[wall.Y][wall.X] = '#'
	}
	for box := range wh.boxes {
		grid[box.Y][box.X] = 'O'
	}
	grid[wh.robot.Y][wh.robot.X] = '@'
	lines := make([]string, len(grid))
	for i, row := range grid {
		lines[i] = string(row)
	}
	return strings.Join(lines, "\n")
}

func (wh *warehouse) upscale(n int) {
	wh.limit = upscale(wh.limit, n)
	wh.robot = upscale(wh.robot, n)
	wh.walls = set.FromSlice(slices.Map(wh.walls.ToSlice(), func(c coord) coord { return upscale(c, n) }))
	extraWalls := wh.walls.Map(func(c coord) coord { return coord{X: c.X + 1, Y: c.Y} })
	wh.walls = wh.walls.Merge(extraWalls)
	wh.boxes = set.FromSlice(slices.Map(wh.boxes.ToSlice(), func(c coord) coord { return upscale(c, n) }))
}

type coord = coordinate.Coordinate2D[int]

func gps(c coord) int {
	return 100*c.Y + c.X
}

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
