package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/taskat/aoc/pkg/common"
)

// templateValues represents the values used in the templates
type templateValues struct {
	Day   int
	Year  int
	Parts []int
}

// newTemplateValues creates a new template values
// with the given year and day. The parts are always 1 and 2
func newTemplateValues(year, day int) templateValues {
	return templateValues{
		Day:   day,
		Year:  year,
		Parts: []int{1, 2},
	}
}

// addDay creates the folder structure for a new day
// and instantiates the template files. It also adds the
// import to the imports.go file
func addDay(args arguments) {
	fmt.Printf("Adding day %d of year %d\n", args.day, args.year)
	src := "cmd/add/templates/day"
	dest := fmt.Sprintf("internal/years/%d/%02d", args.year, args.day)
	instantiateFolder(src, dest, args.toTemplateValues())
	importFile := fmt.Sprintf("internal/years/%d/imports/imports.go", args.year)
	line := fmt.Sprintf("\t_ \"github.com/taskat/aoc/internal/years/%d/%02d\"", args.year, args.day)
	addLineToFile(importFile, line)
}

// addLineToFile adds a line to a file. It adds the line
// before the line of the closing parenthesis
func addLineToFile(file, line string) {
	f, err := os.Open(file)
	common.QuitIfError(err, "Error opening file:")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		nextLine := scanner.Text()
		if nextLine == ")" {
			lines = append(lines, line)
		}
		lines = append(lines, nextLine)
	}
	output := strings.Join(lines, "\n") + "\n"
	err = os.WriteFile(f.Name(), []byte(output), 0644)
	common.QuitIfError(err, "Error writing file:")
}

// addYear creates the folder structure for a new year
// and instantiates the template files. It also adds the
// import to the imports.go file
func addYear(args arguments) {
	fmt.Println("Adding year", args.year)
	src := "cmd/add/templates/year"
	dest := fmt.Sprintf("internal/years/%d", args.year)
	instantiateFolder(src, dest, args.toTemplateValues())
	importFile := "cmd/main/imports/imports.go"
	line := fmt.Sprintf("\t_ \"github.com/taskat/aoc/internal/years/%d\"", args.year)
	addLineToFile(importFile, line)
}

// instantiateFile creates a file from a template file
func instantiateFile(src, dest string, values templateValues) {
	fmt.Println("Creating file", dest, "from", src)
	f, err := os.Create(dest)
	common.QuitIfError(err, "Error creating file:")
	defer f.Close()
	if strings.HasSuffix(f.Name(), ".txt") {
		return
	}
	err = os.Rename(f.Name(), strings.Replace(f.Name(), ".tmpl", "", 1))
	common.QuitIfError(err, "Error renaming file:")
	templateText, err := os.ReadFile(src)
	common.QuitIfError(err, "Error reading template file:")
	t := template.Must(template.New("file").Parse(string(templateText)))
	err = t.Execute(f, values)
	common.QuitIfError(err, "Error executing template:")
}

// instantiateFolder creates a folder from a template folder
// and recursively instantiates the folders and files inside
func instantiateFolder(src, dest string, values templateValues) {
	fmt.Println("Creating folder", dest, "from", src)
	err := os.Mkdir(dest, 0755)
	common.QuitIfError(err, "Error creating folder:")
	entries, err := os.ReadDir(src)
	common.QuitIfError(err, "Error reading directory:")
	for _, entry := range entries {
		newSrc := filepath.Join(src, entry.Name())
		newDest := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			instantiateFolder(newSrc, newDest, values)
		} else {
			instantiateFile(newSrc, newDest, values)
		}
	}
}

func main() {
	args := parseArgs()
	switch args.mode {
	case day:
		addDay(args)
	case year:
		addYear(args)
	}
}
