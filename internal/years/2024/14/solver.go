package day14

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
	"github.com/taskat/aoc/pkg/utils/types/coordinate"
)

// day is the day of the solver
const day = 14

// init registers the solver for day 14
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 14
type Solver struct {
	width  int
	height int
}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {
	if len(params) != 2 {
		s.width = 101
		s.height = 103
		return
	}
	s.width = stringutils.Atoi(params[0].(string))
	s.height = stringutils.Atoi(params[1].(string))
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) lobby {
	robots := slices.Map(lines, parseRobot)
	return lobby{
		width:  s.width,
		height: s.height,
		robots: robots,
	}
}

type robot struct {
	pos      coord
	velocity coord
}

func parseRobot(line string) robot {
	posX, posY, velX, velY := 0, 0, 0, 0
	fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posX, &posY, &velX, &velY)
	return robot{
		pos:      coord{X: posX, Y: posY},
		velocity: coord{X: velX, Y: velY},
	}
}

func (r *robot) move(seconds, width, height int) {
	r.pos.X += r.velocity.X * seconds
	r.pos.X %= width
	if r.pos.X < 0 {
		r.pos.X += width
	}
	r.pos.Y += r.velocity.Y * seconds
	r.pos.Y %= height
	if r.pos.Y < 0 {
		r.pos.Y += height
	}
}

type coord = coordinate.Coordinate2D[int]

type lobby struct {
	width  int
	height int
	robots []robot
}

func (l *lobby) elapseSeconds(n int) {
	for j := range l.robots {
		l.robots[j].move(n, l.width, l.height)
	}
}

func (l lobby) robotsInQuadrants() []int {
	quadrants := make([]int, 4)
	for _, robot := range l.robots {
		quadrantIndex := 0
		if robot.pos.X == l.width/2 || robot.pos.Y == l.height/2 {
			continue
		}
		if robot.pos.Y > l.height/2 {
			quadrantIndex += 2
		}
		if robot.pos.X > l.width/2 {
			quadrantIndex++
		}
		quadrants[quadrantIndex]++
	}
	return quadrants
}

func (l lobby) safetyFactor() int {
	quadrants := l.robotsInQuadrants()
	return slices.Product(quadrants)
}

func (l lobby) String() string {
	line := slices.Repeat('.', l.width)
	grid := make([][]rune, l.height)
	for i := range grid {
		grid[i] = slices.Copy(line)
	}
	robots := make(map[coord]int)
	for _, robot := range l.robots {
		robots[robot.pos]++
	}
	for pos, count := range robots {
		grid[pos.Y][pos.X] = rune('0' + count)
	}
	return strings.Join(slices.Map(grid, func(line []rune) string { return string(line) }), "\n")
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	lobby := s.parse(lines)
	lobby.elapseSeconds(100)
	safetyFactor := lobby.safetyFactor()
	return fmt.Sprint(safetyFactor)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	lobby := s.parse(lines)
	minSafetyFactors := lobby.safetyFactor()
	minIndex := 0
	for i := 1; i < s.height*s.width; i++ {
		lobby.elapseSeconds(1)
		newSafetyFactor := lobby.safetyFactor()
		if newSafetyFactor < minSafetyFactors {
			minSafetyFactors = newSafetyFactor
			minIndex = i
		}
	}
	lobby = s.parse(lines)
	lobby.elapseSeconds(minIndex)
	fmt.Println(lobby)
	return fmt.Sprint(minIndex)
}
