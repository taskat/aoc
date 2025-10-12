package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/taskat/aoc/pkg/common"
)

// mode represents the mode of the program
type mode string

const (
	day  mode = "day"
	year mode = "year"
)

// arguments represents the arguments of the program
type arguments struct {
	help bool
	mode mode
	day  int
	year int
}

// String returns a string representation of the arguments
func (args arguments) String() string {
	return fmt.Sprintf("Arguments{help: %t, mode: %s, day: %d, year: %d}", args.help, args.mode, args.day, args.year)
}

// toTemplateValues converts the arguments to template values
func (args arguments) toTemplateValues() templateValues {
	return newTemplateValues(args.year, args.day)
}

func parseDayArgs(arguments *arguments) (*flag.FlagSet, int) {
	arguments.mode = day
	flagStartOffset := 0
	if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
		day, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error parsing day:", err)
			printDayUsage()
			os.Exit(1)
		}
		arguments.day = day
		flagStartOffset++
	} else {
		arguments.day = -1
	}
	fs := flag.NewFlagSet("day", flag.ExitOnError)
	common.AddFlag(fs.BoolVar, &arguments.help, "help", false, "Print this help message")
	common.AddFlag(fs.IntVar, &arguments.year, "year", getDefaultYearForDay(), "Year number")
	fs.Usage = printDayUsage
	return fs, flagStartOffset
}

func parseYearArgs(arguments *arguments) (*flag.FlagSet, int) {
	arguments.mode = year
	flagStartOffset := 0
	if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
		year, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error parsing year:", err)
			printYearUsage()
			os.Exit(1)
		}
		arguments.year = year
		flagStartOffset++
	} else {
		arguments.year = getDefaultYear()
	}
	fs := flag.NewFlagSet("year", flag.ExitOnError)
	common.AddFlag(fs.BoolVar, &arguments.help, "help", false, "Print this help message")
	fs.Usage = printYearUsage
	return fs, flagStartOffset
}

// parseArgs parses the command line arguments
func parseArgs() arguments {
	if len(os.Args) == 1 {
		printGlobalUsage()
		os.Exit(1)
	}
	arguments := arguments{}
	fs := &flag.FlagSet{}
	flagStart, flagStartOffset := 2, 0
	switch os.Args[1] {
	case "day":
		fs, flagStartOffset = parseDayArgs(&arguments)
		flagStart += flagStartOffset
	case "year":
		fs, flagStartOffset = parseYearArgs(&arguments)
		flagStart += flagStartOffset
	case "help":
		printGlobalUsage()
		os.Exit(0)
	default:
		fmt.Println("Invalid mode:", os.Args[1])
		printGlobalUsage()
		os.Exit(1)
	}
	fs.Parse(os.Args[flagStart:])
	if arguments.help {
		fmt.Println("Help requested")
		fs.Usage()
		os.Exit(0)
	}
	if arguments.day == -1 {
		var err error
		arguments.day, err = getDefaultDay(arguments.year)
		common.QuitIfError(err, "Error getting default day:")
	}
	return arguments
}

// getDefaultDay returns the default day number for the given year
// It returns 1 if there are no days in the year
// It returns the next day number if there are days in the year
func getDefaultDay(year int) (int, error) {
	folders := common.ListFolders(fmt.Sprintf("internal/years/%d", year))
	for i := 0; i < len(folders); i++ {
		if !isInt(folders[i]) {
			folders = append(folders[:i], folders[i+1:]...)
			i--
		}
	}
	if len(folders) == 0 {
		return 1, nil
	}
	day, _ := strconv.Atoi(folders[0])
	if day >= 25 {
		return -1, fmt.Errorf("All days are already added for this year")
	}
	return day + 1, nil
}

// getDefaultYear returns the default year number
// It returns the current year if the current month is December
// It returns the first year in the internal/years folder if there are any
// It returns 2015 otherwise
func getDefaultYear() int {
	date := time.Now()
	if date.Month() == time.December {
		return date.Year()
	}
	folders := common.ListFolders("internal/years")
	if len(folders) == 0 {
		return common.FirstYear
	}
	year, err := strconv.Atoi(folders[0])
	if err != nil {
		fmt.Println("Error parsing year:", err)
		os.Exit(1)
	}
	return year + 1
}

// getDefaultYearForDay returns the default year number for the day mode
// It returns the current year if the current month is December
// It returns the latest year in the internal/years folder if there are any
// It prints an error message and exits if there are no years
func getDefaultYearForDay() int {
	date := time.Now()
	if date.Month() == time.December {
		return date.Year()
	}
	folders := common.ListFolders("internal/years")
	if len(folders) == 0 {
		fmt.Println("No years found in internal/years")
		fmt.Println("Please create a year first")
		os.Exit(1)
	}
	year, err := strconv.Atoi(folders[0])
	common.QuitIfError(err, "Error parsing year:")
	return year
}

// isInt returns true if the given string is an integer
func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// printDayUsage prints the usage of the day mode
func printDayUsage() {
	fmt.Println("add day adds a new day to the repository based on templates")
	fmt.Println("Usage: add day [day] [flags]")
	fmt.Println("  Flags:")
	fmt.Println("    -h,  --help\t\tPrint this help message")
	fmt.Println("    -y,  --year\t\tYear number, defaults to the current year in december")
	fmt.Println("               \t\tOtherwise, it defaults to the latest year in the repository")
}

// printGlobalUsage prints the usage of the program
func printGlobalUsage() {
	fmt.Println("add adds a new day or year to the repository based on templates")
	fmt.Println("Usage: add <command>")
	fmt.Println("  Commands:")
	fmt.Println("    help\tPrint this help message")
	fmt.Println("    year\tGenerate a year")
	fmt.Println("    day\t\tGenerate a day")
	fmt.Println("Use 'add <command> --help' for more information about a command")
}

// printYearUsage prints the usage of the year mode
func printYearUsage() {
	fmt.Println("add year adds a new year to the repository based on templates")
	fmt.Println("Usage: add year [year] [flags]")
	fmt.Println("  Flags:")
	fmt.Println("    -h,  --help\t\tPrint this help message")
}
