package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/taskat/aoc/pkg/common"
)

// templateValues represents the values used in the templates
type templateValues struct {
	Date string
}

// newTemplateValues creates a new template values
func newTemplateValues() templateValues {
	return templateValues{
		Date: time.Now().Format("2006-01-02 15:04:05"),
	}
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

func main() {
	fmt.Println("Generating README.md")
	values := newTemplateValues()
	root := "/workspaces/aoc"
	templatePath := filepath.Join(root, "cmd/readme-generator/templates/readme.md.tmpl")
	destinationPath := filepath.Join(root, "README.md")
	instantiateFile(templatePath, destinationPath, values)
	fmt.Println("Done")
}
