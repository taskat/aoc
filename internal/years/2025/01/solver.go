package day01

import (
	"github.com/taskat/aoc/internal/years/2025/days"
	"github.com/taskat/aoc/pkg/utils/intutils"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 1

// init registers the solver for day 01
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 01
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []move {
	return slices.Map(lines, newMove)
}

// dial represents a dial in the puzzle
type dial int

// newDial creates a new dial and initializes it to 50
func newDial() dial {
	return dial(50)
}

// countValues counts the number of times the dial is left at the given value,
// given a series of moves
func (d *dial) countValues(moves []move, value dial) int {
	count := 0
	current := *d
	for _, m := range moves {
		current.rotate(m)
		if current == value {
			count++
		}
	}
	return count
}

// countClicksAtValue counts the number of clicks the dial makes at the given value,
// given a series of moves
func (d *dial) countClicksAtZero(moves []move) int {
	count := 0
	for _, m := range moves {
		count += m.fullTurns()
		current := d.toInt()
		next := (m.toInt() % 100) + current
		d.rotate(m)
		if current == 0 {
			continue
		}
		if next < 0 || next >= 100 {
			count++
		}
		if d.toInt() == 0 && m.toInt() < 0 {
			count++
		}
	}
	return count
}

// rotate rotates the dial by the given move
func (d *dial) rotate(m move) {
	*d += dial(m)
	*d %= 100
	if *d < 0 {
		*d += 100
	}
}

// toInt converts the dial to an integer
func (d *dial) toInt() int {
	return int(*d)
}

// move represents a move in the puzzle
type move int

// newMove creates a new move from a string representation
// If the string starts with 'R', it creates a right move
// If the string starts with 'L', it creates a left move
func newMove(s string) move {
	amount := stringutils.Atoi(s[1:])
	switch s[0] {
	case 'R':
		return rightMove(amount)
	case 'L':
		return leftMove(amount)
	default:
		panic("invalid move string: " + s)
	}
}

// rightMove creates a move to the right by the given amount
// It returns a positive move
func rightMove(amount int) move {
	return move(amount)
}

// leftMove creates a move to the left by the given amount
// It returns a negative move
func leftMove(amount int) move {
	return move(-amount)
}

// fullTurns returns the number of full turns in the move
func (m move) fullTurns() int {
	return intutils.Abs(m.toInt()) / 100
}

// toInt converts the move to an integer
func (m move) toInt() int {
	return int(m)
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	dial := newDial()
	moves := s.parse(lines)
	zeroCount := dial.countValues(moves, 0)
	return stringutils.Itoa(zeroCount)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	dial := newDial()
	moves := s.parse(lines)
	clickCount := dial.countClicksAtZero(moves)
	return stringutils.Itoa(clickCount)
}
