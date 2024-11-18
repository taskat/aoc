package day01

import (
	"testing"

	"github.com/taskat/aoc/cmd/main/config"
	"github.com/taskat/aoc/internal/years/2023/days"

	"github.com/stretchr/testify/assert"
)

var (
	// result stores the result of the solver
	// this is used to prevent the compiler from optimizing the code
	result any
)

// Test_Day_01_Part1 tests the day 01 solver for part 1
func Test_Day_01_Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.InputType
		expectedValue string
		hyperParams   []any
	}{
		{"Test 1", config.TestInput(1), "142", nil},
		{"Real", config.RealInput{}, "54927", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := config.NewConfig(days.Year, day, 1, tc.input, tc.hyperParams...)
			solver := &Solver{}
			solver.AddHyperParams(tc.hyperParams...)
			solver.ParsePart1(cfg.ReadInputFile())
			solution := solver.SolvePart1()
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

// Benchmark_Day_01_Part1 benchmarks the day 01 solver for part 1
func Benchmark_Day_01_Part1(b *testing.B) {
	cfg := config.NewConfig(days.Year, day, 1, config.RealInput{}, nil)
	solver := &Solver{}
	solver.AddHyperParams(cfg.GetHyperParams()...)
	solver.ParsePart1(cfg.ReadInputFile())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart1()
	}
}

// Test_Day_01_Part2 tests the day 01 solver for part 2
func Test_Day_01_Part2(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.InputType
		expectedValue string
		hyperParams   []any
	}{
		{"Test 1", config.TestInput(1), "", nil},
		{"Real", config.RealInput{}, "", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := config.NewConfig(days.Year, day, 1, tc.input, tc.hyperParams...)
			solver := &Solver{}
			solver.AddHyperParams(tc.hyperParams...)
			solver.ParsePart2(cfg.ReadInputFile())
			solution := solver.SolvePart2()
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

// Benchmark_Day_01_Part2 benchmarks the day 01 solver for part 2
func Benchmark_Day_01_Part2(b *testing.B) {
	cfg := config.NewConfig(days.Year, day, 1, config.RealInput{}, nil)
	solver := &Solver{}
	solver.AddHyperParams(cfg.GetHyperParams()...)
	solver.ParsePart2(cfg.ReadInputFile())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart2()
	}
}
