package day03

import (
	"testing"

	"github.com/taskat/aoc/cmd/main/config"
	"github.com/taskat/aoc/internal/years/2024/days"

	"github.com/stretchr/testify/assert"
)

var (
	// result stores the result of the solver
	// this is used to prevent the compiler from optimizing the code
	_result any
)

// Test_Day_03_Part1 tests the day 03 solver for part 1
func Test_Day_03_Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.InputType
		expectedValue string
		hyperParams   []any
	}{
		{"Test 1", config.TestInput(1), "161", nil},
		{"Real", config.RealInput{}, "166905464", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := config.NewConfig(days.Year, day, 1, tc.input, tc.hyperParams...)
			solver := &Solver{}
			solver.AddHyperParams(tc.hyperParams...)
			solution := solver.SolvePart1(cfg.ReadInputFile())
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

// Benchmark_Day_03_Part1 benchmarks the day 03 solver for part 1
func Benchmark_Day_03_Part1(b *testing.B) {
	cfg := config.NewConfig(days.Year, day, 1, config.RealInput{}, nil)
	solver := &Solver{}
	solver.AddHyperParams(cfg.GetHyperParams()...)
	input := cfg.ReadInputFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_result = solver.SolvePart1(input)
	}
}

// Test_Day_03_Part2 tests the day 03 solver for part 2
func Test_Day_03_Part2(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.InputType
		expectedValue string
		hyperParams   []any
	}{
		{"Test 1", config.TestInput(2), "48", nil},
		{"Real", config.RealInput{}, "72948684", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := config.NewConfig(days.Year, day, 2, tc.input, tc.hyperParams...)
			solver := &Solver{}
			solver.AddHyperParams(tc.hyperParams...)
			solution := solver.SolvePart2(cfg.ReadInputFile())
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

// Benchmark_Day_03_Part2 benchmarks the day 03 solver for part 2
func Benchmark_Day_03_Part2(b *testing.B) {
	cfg := config.NewConfig(days.Year, day, 2, config.RealInput{}, nil)
	solver := &Solver{}
	solver.AddHyperParams(cfg.GetHyperParams()...)
	input := cfg.ReadInputFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_result = solver.SolvePart2(input)
	}
}
