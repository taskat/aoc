package day05

import (
	"testing"

	"github.com/taskat/aoc/cmd/main/config"
	"github.com/taskat/aoc/internal/years/2024/days"

	"github.com/stretchr/testify/assert"
)

// Test_Day_05_Part1 tests the day 05 solver for part 1
func Test_Day_05_Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.InputType
		expectedValue string
		hyperParams   []string
	}{
		{"Test 1", config.TestInput(1), "143", nil},
		{"Real", config.RealInput{}, "5762", nil},
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

// Benchmark_Day_05_Part1 benchmarks the day 05 solver for part 1
func Benchmark_Day_05_Part1(b *testing.B) {
	cfg := config.NewConfig(days.Year, day, 1, config.RealInput{})
	solver := &Solver{}
	solver.AddHyperParams(cfg.GetHyperParams()...)
	input := cfg.ReadInputFile()
	for b.Loop() {
		solver.SolvePart1(input)
	}
}

// Test_Day_05_Part2 tests the day 05 solver for part 2
func Test_Day_05_Part2(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.InputType
		expectedValue string
		hyperParams   []string
	}{
		{"Test 1", config.TestInput(1), "123", nil},
		{"Real", config.RealInput{}, "4130", nil},
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

// Benchmark_Day_05_Part2 benchmarks the day 05 solver for part 2
func Benchmark_Day_05_Part2(b *testing.B) {
	cfg := config.NewConfig(days.Year, day, 2, config.RealInput{})
	solver := &Solver{}
	solver.AddHyperParams(cfg.GetHyperParams()...)
	input := cfg.ReadInputFile()
	for b.Loop() {
		solver.SolvePart2(input)
	}
}
