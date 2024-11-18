package years

import (
	"github.com/taskat/aoc/cmd/main/collection"
)

// years is a map that contains all the yearCollections for all the years
var years = make(collection.Collection[collection.Year])

// AddYear adds a new year to the years map
func AddYear(year int, d collection.Year) {
	years[year] = d
}

// GetYear returns the year for a specific year
func GetYear(year int) collection.Year {
	return years[year]
}
