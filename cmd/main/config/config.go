package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Config is a struct that contains the configuration of the solver
// for a specific part of a specific day of a specific year
type Config struct {
	year        int
	day         int
	part        int
	inputType   InputType
	hyperParams []any
}

// NewConfig creates a new Config struct with the given parameters
func NewConfig(year, day, part int, inputType InputType, hyperParams ...any) *Config {
	return &Config{
		year:        year,
		day:         day,
		part:        part,
		inputType:   inputType,
		hyperParams: hyperParams,
	}
}

// GetYear returns the year of the Config
func (c *Config) GetYear() int {
	return c.year
}

// GetDay returns the day of the Config
func (c *Config) GetDay() int {
	return c.day
}

// GetHyperParams returns the hyper parameters of the Config
func (c *Config) GetHyperParams() []any {
	return c.hyperParams
}

// GetInputFileName returns the input file name corresponding to the Config
func (c *Config) GetInputFileName() string {
	return fmt.Sprintf("/workspaces/aoc/internal/years/%d/%.2d/inputs/%s.txt", c.year, c.day, c.inputType.String())
}

// GetInputType returns the input type of the Config
func (c *Config) GetInputType() InputType {
	return c.inputType
}

// GetPart returns the part of the Config
func (c *Config) GetPart() int {
	return c.part
}

// ReadInputFile reads the input file defined in the Config and returns the content
// of the file as an array of lines
func (c *Config) ReadInputFile() []string {
	return ReadInputLines(c.GetInputFileName())
}

// ReadInputLines reads the input file and returns the lines of the file as a slice of strings
func ReadInputLines(filename string) []string {
	filename, _ = filepath.Abs(filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// String is a method that returns a string representation of Config
func (c *Config) String() string {
	return fmt.Sprintf("Config{Year: %d, Day: %d, Part: %d, InputType: %s, HyperParams: %v}",
		c.year, c.day, c.part, c.inputType, c.hyperParams)
}
