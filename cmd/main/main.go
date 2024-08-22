package main

import (
	"fmt"

	"taskat/aoc/config"
)

func main() {
	fmt.Println("Welcome to taskat's Advent of Code solutions!")
	config := config.ParseConfig()
	fmt.Println("Config:", config.String())
}
