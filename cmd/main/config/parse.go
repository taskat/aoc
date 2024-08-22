package config

import (
	"flag"
	"fmt"
	"os"
)

// addVar is a helper function that adds a variable to the flag package
// It adds the variable with the given name, usage, and default value
// It also adds the variable with the first letter of the name as the name
func addVar[T any](variable *T, name, usage string, defaultValue T, add func(*T, string, T, string)) {
	add(variable, name, defaultValue, usage)
	add(variable, name[:1], defaultValue, usage)
}

// parseArguments parses the command line arguments and returns the year, day, part, input type
// and the hyper parameters
func parseArguments() (year, day, part int, inputType string, hyperParams []any) {
	addVar(&year, "year", "the year of the puzzle", 2023, flag.IntVar)
	addVar(&day, "day", "the day of the puzzle", 1, flag.IntVar)
	addVar(&part, "part", "the part of the puzzle", 1, flag.IntVar)
	addVar(&inputType, "input", "the input type to use", "real", flag.StringVar)
	flag.Parse()
	args := flag.Args()
	hyperParams = make([]any, len(args))
	for i, arg := range args {
		hyperParams[i] = arg
	}
	return year, day, part, inputType, hyperParams
}

// ParseConfig parses the command line arguments and returns a Config struct
// If the arguments are not valid, it prints an error message and exits
func ParseConfig() *Config {
	year, day, part, inputType_, hyperParams := parseArguments()
	inputType := validateArguments(year, day, part, inputType_)
	return NewConfig(year, day, part, inputType, hyperParams...)
}

// validateArguments validates the year, day, part, and input type
// It returns the input type if it is valid, otherwise it prints an error message and exits
func validateArguments(year, day, part int, inputType_ string) InputType {
	validate := func(value, min, max int, name string) {
		if value < min || value > max {
			fmt.Println(fmt.Sprintf("%s: %d is not between %d and %d", name, value, min, max))
			os.Exit(1)
		}
	}
	validate(year, 2015, 2023, "Year")
	validate(day, 1, 25, "Day")
	validate(part, 1, 2, "Part")
	inputType := parseInputType(inputType_)
	if inputType == nil {
		fmt.Println("Input must be either 'real' or a 'test-n' where n is a number")
		os.Exit(1)
	}
	return inputType
}
