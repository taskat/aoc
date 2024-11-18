package common

import (
	"os"
	"sort"
)

const (
	FirstYear = 2015
)

// FlagAddFunc is a function type that adds a flag to the flag set
type FlagAddFunc[T any] func(varRef *T, name string, defaultValue T, usage string)

// AddFlag adds a flag to the flag set, it adds the flag with the full name and the first letter of the name
// as a shorthand
func AddFlag[T any](addFunc FlagAddFunc[T], varRef *T, name string, defaultValue T, usage string) {
	addFunc(varRef, name, defaultValue, usage)
	addFunc(varRef, name[:1], defaultValue, usage)
}

// ListFolders returns a list of folders in the given path
// It also sorts the list in reverse order
func ListFolders(path string) []string {
	entries, err := os.ReadDir(path)
	QuitIfError(err, "Error reading directory:")
	var folders []string
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(folders)))
	return folders
}
