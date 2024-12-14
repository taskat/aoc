package day14

import (
	"fmt"

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
func (s *Solver) parse(lines []string) []robot {
	return slices.Map(lines, parseRobot)
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

func (s Solver) elapseSeconds(n int, robots []robot) {
	for j := range robots {
		robots[j].move(n, s.width, s.height)
	}
}

func (s Solver) robotsInQuadrants(robots []robot) []int {
	quadrants := make([]int, 4)
	for _, robot := range robots {
		quadrantIndex := 0
		if robot.pos.X == s.width/2 || robot.pos.Y == s.height/2 {
			continue
		}
		if robot.pos.Y > s.height/2 {
			quadrantIndex += 2
		}
		if robot.pos.X > s.width/2 {
			quadrantIndex++
		}
		quadrants[quadrantIndex]++
	}
	return quadrants
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	robots := s.parse(lines)
	s.elapseSeconds(100, robots)
	quadrants := s.robotsInQuadrants(robots)
	safetyfactor := slices.Product(quadrants)
	fmt.Println(s.width, s.height)
	return fmt.Sprint(safetyfactor)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
