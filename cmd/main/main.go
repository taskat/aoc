package main

import (
	"fmt"

	"github.com/taskat/aoc/cmd/main/config"
	_ "github.com/taskat/aoc/cmd/main/imports"
	"github.com/taskat/aoc/cmd/main/years"
)

func main() {
	fmt.Println("Welcome to taskat's Advent of Code solutions!")
	cfg := config.ParseConfig()
	msg := "Solving year %d, day %d, part %d with %s input:"
	fmt.Println(fmt.Sprintf(msg, cfg.GetYear(), cfg.GetDay(), cfg.GetPart(), cfg.GetInputType()))
	year := years.GetYear(cfg.GetYear())
	if year == nil {
		fmt.Println("Year not found")
		return
	}
	day := year.Get(cfg.GetDay())
	if day == nil {
		fmt.Println("Day not found")
		return
	}
	input := cfg.ReadInputFile()
	if input == nil {
		fmt.Println("Input file not found")
		return
	}
	if cfg.GetPart() == 1 {
		day.ParsePart1(input)
		fmt.Println("Solution for part 1:", day.SolvePart1())
	} else {
		day.ParsePart2(input)
		fmt.Println("Solution for part 2:", day.SolvePart2())
	}

}
