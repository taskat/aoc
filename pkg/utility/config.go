package utility

import (
	"os"
	"sort"
)

const (
	FirstYear = 2015
)

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
