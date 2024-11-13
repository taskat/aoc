package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type templateValues struct {
	Day  int
	Year int
}

func addDay(args arguments) {
	fmt.Println(fmt.Sprintf("Adding day %d of year %d", args.day, args.year))
}

func addYear(args arguments) {
	fmt.Println("Adding year", args.year)
	src := "cmd/add/templates/year"
	dest := fmt.Sprintf("internal/years/%d", args.year)
	instantiateFolder(src, dest, args.toTemplateValues())
}

func instantiateFile(src, dest string, values templateValues) {
	f, err := os.Create(dest)
	quitIfError(err, "Error creating file:")
	if strings.HasSuffix(f.Name(), ".txt") {
		return
	}
	err = os.Rename(f.Name(), strings.Replace(f.Name(), ".tmpl", "", 1))
	quitIfError(err, "Error renaming file:")
	templateText, err := os.ReadFile(src)
	quitIfError(err, "Error reading template file:")
	t := template.Must(template.New("file").Parse(string(templateText)))
	err = t.Execute(f, values)
	quitIfError(err, "Error executing template:")
}

func instantiateFolder(src, dest string, values templateValues) {
	fmt.Println("Creating folder", dest, "from", src)
	err := os.Mkdir(dest, 0755)
	quitIfError(err, "Error creating folder:")
	entries, err := os.ReadDir(src)
	quitIfError(err, "Error reading directory:")
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

func quitIfError(err error, message string) {
	if err == nil {
		return
	}
	fmt.Println(message, err)
	os.Exit(1)
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
