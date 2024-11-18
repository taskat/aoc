package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/taskat/aoc/pkg/utility"
)

// getMaxDay returns the maximum day that has been added to the repository
func getMaxDay(year int) int {
	folders := utility.ListFolders(fmt.Sprintf("internal/years/%d", year))
	if len(folders) == 0 {
		return 0
	}
	day, err := strconv.Atoi(folders[0])
	utility.QuitIfError(err, "Error parsing day:")
	return day
}

// getMaxYear returns the maximum year that has been added to the repository
func getMaxYear() int {
	folders := utility.ListFolders("internal/years")
	if len(folders) == 0 {
		return utility.FirstYear
	}
	year, err := strconv.Atoi(folders[0])
	utility.QuitIfError(err, "Error parsing year:")
	return year
}

// parseArguments parses the command line arguments and returns the year, day, part, input type
// and the hyper parameters
func parseArguments() (year, day, part int, inputType string, hyperParams []any) {
	utility.AddFlag(flag.IntVar, &year, "year", getMaxYear(), "the year of the puzzle")
	utility.AddFlag(flag.IntVar, &day, "day", -1, "the day of the puzzle")
	utility.AddFlag(flag.IntVar, &part, "part", 1, "the part of the puzzle")
	utility.AddFlag(flag.StringVar, &inputType, "input", "real", "the input type to use")
	flag.Parse()
	args := flag.Args()
	hyperParams = make([]any, len(args))
	for i, arg := range args {
		hyperParams[i] = arg
	}
	if day == -1 {
		day = getMaxDay(year)
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
	validate(year, utility.FirstYear, getMaxYear(), "Year")
	validate(day, 1, getMaxDay(year), "Day")
	validate(part, 1, 2, "Part")
	inputType := parseInputType(inputType_)
	if inputType == nil {
		fmt.Println("Input must be either 'real' or a 'test-n' where n is a number")
		os.Exit(1)
	}
	return inputType
}
