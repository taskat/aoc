package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/taskat/aoc/pkg/common"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/stringutils"
)

// getMaxDay returns the maximum day that has been added to the repository
func getMaxDay(year int) int {
	folders := common.ListFolders(fmt.Sprintf("internal/years/%d", year))
	if len(folders) == 0 {
		return 0
	}
	folders = slices.Filter(folders, stringutils.IsInteger)
	day, err := strconv.Atoi(folders[0])
	common.QuitIfError(err, "Error parsing day:")
	return day
}

// getMaxYear returns the maximum year that has been added to the repository
func getMaxYear() int {
	folders := common.ListFolders("internal/years")
	if len(folders) == 0 {
		return common.FirstYear
	}
	year, err := strconv.Atoi(folders[0])
	common.QuitIfError(err, "Error parsing year:")
	return year
}

// parseArguments parses the command line arguments and returns the year, day, part, input type
// and the hyper parameters
func parseArguments() (year, day, part int, inputType string, hyperParams []string) {
	common.AddFlag(flag.IntVar, &year, "year", getMaxYear(), "the year of the puzzle")
	common.AddFlag(flag.IntVar, &day, "day", -1, "the day of the puzzle")
	common.AddFlag(flag.IntVar, &part, "part", 1, "the part of the puzzle")
	common.AddFlag(flag.StringVar, &inputType, "input", "real", "the input type to use")
	flag.Parse()
	args := flag.Args()
	hyperParams = make([]string, len(args))
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
	validate := func(value int, values []int, name string) {
		if !slices.Contains(values, value) {
			fmt.Println(fmt.Sprintf("%s: %d is not in possible values: %v", name, value, values))
			os.Exit(1)
		}
	}
	yearFolders := common.ListFolders("internal/years")
	years := slices.Map(yearFolders, stringutils.Atoi)
	validate(year, years, "Year")
	dayFolders := common.ListFolders(fmt.Sprintf("internal/years/%d", year))
	dayFolders = slices.Filter(dayFolders, stringutils.IsInteger)
	days := slices.Map(dayFolders, stringutils.Atoi)
	validate(day, days, "Day")
	validate(part, []int{1, 2}, "Part")
	inputType := parseInputType(inputType_)
	if inputType == nil {
		fmt.Println("Input must be either 'real' or a 'test-n' where n is a number")
		os.Exit(1)
	}
	return inputType
}
