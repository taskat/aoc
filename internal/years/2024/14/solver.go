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
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 14
type Solver struct {
	width  int
	height int
}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {
	if len(params) != 2 {
		s.width = 101
		s.height = 103
		return
	}
	s.width = stringutils.Atoi(params[0])
	s.height = stringutils.Atoi(params[1])
}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) lobby {
	robots := slices.Map(lines, parseRobot)
	return newLobby(s.width, s.height, robots)
}

// lobby holds the robots and the dimensions of the lobby
type lobby struct {
	width  int
	height int
	robots []robot
}

// newLobby creates a new lobby
func newLobby(width, height int, robots []robot) lobby {
	return lobby{width: width, height: height, robots: robots}
}

// elapseSeconds moves the robots n seconds into the future
func (l *lobby) elapseSeconds(n int) {
	slices.ForEach_m(l.robots, func(r *robot) { r.move(n, l.width, l.height) })
}

// robotsInQuadrants returns the number of robots in each quadrant. It ignores
// robots that are on the axes
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

// safetyFactor returns the product of the number of robots in each quadrant,
// which is a measure of how close the robots are to each other
func (l lobby) safetyFactor() int {
	quadrants := l.robotsInQuadrants()
	return slices.Product(quadrants)
}

// String returns a string representation of the lobby, where every empty space
// is represented by a dot. If there are multiple robots in the same position,
// they are represented by a number
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

// robot represents a robot in the lobby, with a position and a velocity
type robot struct {
	pos      coord
	velocity coord
}

// parseRobot parses a robot from a string. The format is "p=x,y v=x,y"
func parseRobot(line string) robot {
	posX, posY, velX, velY := 0, 0, 0, 0
	fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posX, &posY, &velX, &velY)
	return robot{
		pos:      coord{X: posX, Y: posY},
		velocity: coord{X: velX, Y: velY},
	}
}

// move moves the robot n seconds into the future, wrapping around the lobby
func (r *robot) move(seconds, width, height int) {
	newPos := func(oldPos, velocity, limit int) int {
		newPos := oldPos + velocity*seconds
		newPos %= limit
		if newPos < 0 {
			newPos += limit
		}
		return newPos
	}
	r.pos.X = newPos(r.pos.X, r.velocity.X, width)
	r.pos.Y = newPos(r.pos.Y, r.velocity.Y, height)
}

// coord is a wrapper around a 2D coordinate
type coord = coordinate.Coordinate2D[int]

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
	safetyFactors := slices.For(s.width*s.height, func(_ int) int {
		lobby.elapseSeconds(1)
		return lobby.safetyFactor()
	})
	_, minIndex := slices.Min_i(safetyFactors)
	minIndex++
	// prints out the resulting image
	// lobby = s.parse(lines)
	// lobby.elapseSeconds(minIndex)
	// fmt.Println(lobby)
	return fmt.Sprint(minIndex)
}
