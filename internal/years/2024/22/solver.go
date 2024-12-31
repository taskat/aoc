package day22

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/maps"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
	"github.com/taskat/aoc/pkg/utils/testutils"
)

// day is the day of the solver
const day = 22

// init registers the solver for day 22
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 22
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []*secret {
	return slices.Map(lines, newSecret)
}

type secret struct {
	initial int
	current int
	prices  []int
	changes []int
}

func newSecret(s string) *secret {
	return &secret{
		initial: stringutils.Atoi(s),
		current: stringutils.Atoi(s),
		prices:  make([]int, 0),
		changes: make([]int, 0),
	}
}

func (s *secret) calculate() {
	s.prices = make([]int, 2001)
	s.changes = make([]int, len(s.prices)-1)
	for i := 0; i < 2001; i++ {
		s.prices[i] = s.current % 10
		if i > 0 {
			s.changes[i-1] = s.prices[i] - s.prices[i-1]
		}
		s.next()
	}
}

func (s *secret) isMatchAt(seq sequence, at int) bool {
	// defer timer.Call(testutils.Start())
	for i := 0; i < 4; i++ {
		if s.prices[at-(3-i)]-s.prices[at-(4-i)] != seq[i] {
			return false
		}
	}
	return true
}

func (s *secret) sellAfterSequence(seq sequence) int {
	for i := 4; i < len(s.prices); i++ {
		if s.isMatchAt(seq, i) {
			return s.prices[i]
		}
	}
	return 0
}

func (s *secret) next() {
	s.current = mix(s.current*64, s.current)
	s.current = prune(s.current)
	s.current = mix(s.current/32, s.current)
	s.current = prune(s.current)
	s.current = mix(s.current*2048, s.current)
	s.current = prune(s.current)
}

var timer testutils.TimeRepeatedCalls

// func (s *secret) sequences() map[sequence]int {
// 	defer timer.Call(testutils.Start())
// 	sequences := make(map[sequence]int, 0)
// 	for i := 3; i < 2000; i++ {
// 		newSequence := sequence{s.changes[i-3], s.changes[i-2], s.changes[i-1], s.changes[i]}
// 		if !maps.Contains(sequences, newSequence) {
// 			sequences[newSequence] = s.prices[i+1]
// 		}
// 	}
// 	return sequences
// }

func (s *secret) sequences(sequences map[sequence]int) map[sequence]int {
	defer timer.Call(testutils.Start())
	current := set.New[sequence]()
	for i := 3; i < 2000; i++ {
		newSequence := sequence{s.changes[i-3], s.changes[i-2], s.changes[i-1], s.changes[i]}
		if !current.Contains(newSequence) {
			sequences[newSequence] += s.prices[i+1]
			current.Add(newSequence)
		}
	}
	return sequences
}

func mix(a, b int) int {
	return a ^ b
}

func mixPrune(a, b int) int {
	return prune(mix(a, b))
}

func prune(a int) int {
	return a % 16777216
}

type sequence [4]int

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	secrets := s.parse(lines)
	for i := 0; i < 2000; i++ {
		slices.ForEach(secrets, (*secret).next)
	}
	finalSecrets := slices.Map(secrets, func(s *secret) int { return s.current })
	sum := slices.Sum(finalSecrets)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	defer testutils.PrintSingleCall(testutils.Start())
	secrets := s.parse(lines)
	slices.ForEach(secrets, (*secret).calculate)
	sequences := make(map[sequence]int, 0)
	slices.ForEach(secrets, func(s *secret) {
		sequences = s.sequences(sequences)
	})
	max := slices.Max(maps.Values(sequences))
	timer.PrintStats()
	return fmt.Sprint(max)
}
