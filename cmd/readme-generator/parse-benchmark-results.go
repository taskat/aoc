package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"strings"

	"github.com/taskat/aoc/pkg/common"
	"github.com/taskat/aoc/pkg/utils/iterutils"
)

// FullResult represents the full benchmark results
// It is a map of year to yearResult
type FullResult map[int]YearResult

// FromFile reads the benchmark results from the given file
func FromFile(file string) FullResult {
	bytes, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return make(FullResult)
		}
		common.QuitIfError(err, "Error reading "+file+":")
	}
	var existing FullResult
	err = json.Unmarshal(bytes, &existing)
	common.QuitIfError(err, "Error unmarshalling "+file+":")
	return existing
}

// getYear gets the yearResult for the given year, creating it if it doesn't exist
func (f FullResult) getYear(year int) YearResult {
	if f[year] == nil {
		f[year] = make(YearResult)
	}
	return f[year]
}

// MarshalJSON marshals the fullResult as a JSON object
// It also appends a comment at the top
func (f FullResult) MarshalJSON() ([]byte, error) {
	type Alias FullResult
	return json.Marshal(&struct {
		Comment string `json:"_comment"`
		Result  *Alias
	}{
		Comment: "This file is generated. Do not edit manually.",
		Result:  (*Alias)(&f),
	})
}

// merge merges another fullResult into this one
func (f FullResult) merge(other FullResult) {
	for year, yr := range other {
		f.getYear(year).merge(yr)
	}
}

// UnmarshalJSON unmarshals a JSON object into a fullResult
// It ignores the comment at the top
func (f *FullResult) UnmarshalJSON(data []byte) error {
	type Alias FullResult
	aux := &struct {
		Result Alias
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	*f = FullResult(aux.Result)
	return nil
}

// updateResults updates the results from the content of the given file.
// The file is expected to be in the format produced by `go test -bench . -benchmem -run=^$ -count=3 ./... > benchmark.txt`
func (f FullResult) updateResults(file string) {
	parsedResults := parse(file)
	f.merge(parsedResults)
}

// saveToFile saves the fullResult to the given file
func (f FullResult) saveToFile(file string) {
	bytes, err := json.MarshalIndent(f, "", "  ")
	common.QuitIfError(err, "Error marshalling results:")
	err = os.WriteFile(file, bytes, 0644)
	common.QuitIfError(err, "Error writing "+file+":")
}

// YearResult represents the benchmark results for a year
// It is a map of day to dayResult
type YearResult map[int]DayResult

// getDay gets the dayResult for the given day, creating it if it doesn't exist
func (y YearResult) getDay(day int) DayResult {
	if y[day] == nil {
		y[day] = make(DayResult)
	}
	return y[day]
}

// getStars returns the number of stars collected in this year
func (y YearResult) getStars() int {
	count := 0
	for _, dr := range y {
		count += dr.getStars()
	}
	return count
}

// merge merges another yearResult into this one
func (y YearResult) merge(other YearResult) {
	for day, dr := range other {
		y.getDay(day).merge(dr)
	}
}

// DayResult represents the benchmark results for a day
// It is a map of part to times
type DayResult map[int]Times

// addBenchmarkResult adds a benchmark result for the given part
// It appends the time to the list of times for that part
// If the part doesn't exist, it creates it
func (d DayResult) addBenchmarkResult(part int, t int) {
	if d[part] == nil {
		d[part] = make(Times, 0, 3)
	}
	d[part] = append(d[part], t)
}

// getStars returns the number of stars collected for this day
func (d DayResult) getStars() int {
	count := 0
	if _, ok := d[1]; ok {
		count++
	}
	if _, ok := d[2]; ok {
		count++
	}
	return count
}

// merge overrides the times for each part with the times from the other dayResult
func (d DayResult) merge(other DayResult) {
	maps.Copy(d, other)
}

// Times represents a list of benchmark Times
// It is a slice of ints, where each int is a time in nanoseconds
type Times []int

// average returns the average time in nanoseconds
// If there are no times, it returns 0
func (t Times) average() int {
	if len(t) == 0 {
		return 0
	}
	sum := iterutils.Sum(iterutils.NewFromSlice(t))
	return sum / len(t)
}

// MarshalJSON marshals the average time as a JSON number
func (t Times) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "%d", t.average()), nil
}

// String returns a string representation of the times
func (t Times) String() string {
	if len(t) == 0 {
		return "-"
	}
	return fmt.Sprintf("%d ms", t.average()/1000)
}

// UnmarshalJSON unmarshals a JSON number into a times slice with a single element
func (t *Times) UnmarshalJSON(data []byte) error {
	var avg int
	if err := json.Unmarshal(data, &avg); err != nil {
		return err
	}
	*t = append(*t, avg)
	return nil
}

// parse parses the benchmark results from the given file
func parse(file string) FullResult {
	data, err := os.ReadFile(file)
	common.QuitIfError(err, "Error reading file:")
	lines := strings.Split(string(data), "\n")
	allResults := make(FullResult)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "pkg: ") {
			var currentYear, currentDay int
			fmt.Sscanf(line, "pkg: github.com/taskat/aoc/internal/years/%d/%d", &currentYear, &currentDay)
			yearResult := allResults.getYear(currentYear)
			dayResult := yearResult.getDay(currentDay)
			offset := parsePackageResult(lines[i+1:], dayResult)
			i += offset
		}
	}
	return allResults
}

// parsePackageResult parses the benchmark results for a single package
// It returns the number of lines parsed
func parsePackageResult(lines []string, current DayResult) int {
	for i, line := range lines {
		if strings.HasPrefix(line, "PASS") {
			return i
		}
		if strings.HasPrefix(line, "pkg: ") {
			fmt.Println("Failed to parse package result, got new package line")
			os.Exit(1)
		}
		if !strings.HasPrefix(line, "Benchmark") {
			continue
		}
		if !strings.Contains(line, " ") {
			continue
		}
		var _day, part, _runs, nsPerOp int
		fmt.Sscanf(line, "Benchmark_Day_%d_Part%d\t%d\t%d ns/op\t%d", &_day, &part, &_runs, &nsPerOp)
		current.addBenchmarkResult(part, nsPerOp)
	}
	return len(lines)
}
