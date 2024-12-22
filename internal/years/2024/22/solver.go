package day22

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// day is the day of the solver
const day = 22

// init registers the solver for day 22
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 22
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) []int {
	return slices.Map(lines, stringutils.Atoi)
}

func nextSecret(secret int) int {
	secret = mix(secret*64, secret)
	secret = prune(secret)
	secret = mix(secret/32, secret)
	secret = prune(secret)
	secret = mix(secret*2048, secret)
	secret = prune(secret)
	return secret
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

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	secrets := s.parse(lines)
	for i := 0; i < 2000; i++ {
		secrets = slices.Map(secrets, nextSecret)
	}
	sum := slices.Sum(secrets)
	return fmt.Sprint(sum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	return ""
}
