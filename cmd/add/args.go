package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/taskat/aoc/pkg/utility"
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

// flagAddFunc is a function type that adds a flag to the flag set
type flagAddFunc[T any] func(varRef *T, name string, defaultValue T, usage string)

// addFlag adds a flag to the flag set, it adds the flag with the full name and the first letter of the name
// as a shorthand
func addFlag[T any](addFunc flagAddFunc[T], varRef *T, name string, defaultValue T, usage string) {
	addFunc(varRef, name, defaultValue, usage)
	addFunc(varRef, name[:1], defaultValue, usage)
}

// parseArgs parses the command line arguments
func parseArgs() arguments {
	if len(os.Args) == 1 {
		printGlobalUsage()
		os.Exit(1)
	}
	arguments := arguments{}
	fs := &flag.FlagSet{}
	flagStart := 2
	switch os.Args[1] {
	case "day":
		arguments.mode = day
		if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
			day, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Error parsing day:", err)
				printYearUsage()
				os.Exit(1)
			}
			arguments.day = day
			flagStart++
		} else {
			arguments.day = -1
		}
		fs = flag.NewFlagSet("day", flag.ExitOnError)
		addFlag(fs.BoolVar, &arguments.help, "help", false, "Print this help message")
		addFlag(fs.IntVar, &arguments.year, "year", getDefaultYearForDay(), "Year number")
		fs.Usage = printDayUsage
	case "year":
		arguments.mode = year
		if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
			year, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Error parsing year:", err)
				printYearUsage()
				os.Exit(1)
			}
			arguments.year = year
			flagStart++
		} else {
			arguments.year = getDefaultYear()
		}
		fs = flag.NewFlagSet("year", flag.ExitOnError)
		addFlag(fs.BoolVar, &arguments.help, "help", false, "Print this help message")
		fs.Usage = printYearUsage
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
		fs.Usage()
		os.Exit(0)
	}
	if arguments.day == -1 {
		arguments.day = getDefaultDay(arguments.year)
	}
	return arguments
}

// getDefaultDay returns the default day number for the given year
// It returns 1 if there are no days in the year
// It returns the next day number if there are days in the year
func getDefaultDay(year int) int {
	folders := utility.ListFolders(fmt.Sprintf("internal/years/%d", year))
	for i := 0; i < len(folders); i++ {
		if !isInt(folders[i]) {
			folders = append(folders[:i], folders[i+1:]...)
			i--
		}
	}
	if len(folders) == 0 {
		return 1
	}
	day, _ := strconv.Atoi(folders[0])
	return day + 1
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
	folders := utility.listFolders("internal/years")
	if len(folders) == 0 {
		return 2015
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
	folders := utility.listFolders("internal/years")
	if len(folders) == 0 {
		fmt.Println("No years found in internal/years")
		fmt.Println("Please create a year first")
		os.Exit(1)
	}
	year, err := strconv.Atoi(folders[0])
	quitIfError(err, "Error parsing year:")
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
	fmt.Println("    help\t\tPrint this help message")
	fmt.Println("    year\t\tGenerate a year")
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
