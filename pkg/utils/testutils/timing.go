package testutils

import (
	"fmt"
	"time"

	"github.com/taskat/aoc/pkg/utils/slices"
)

// TimeRepeatedCalls is a struct that stores the time taken for each call
type TimeRepeatedCalls struct {
	calls []time.Duration
}

// Call stores the time taken for a single call
// It should be used in a defer statement, like:
// defer t.Call(Start())
func (t *TimeRepeatedCalls) Call(startTime time.Time) {
	t.calls = append(t.calls, TimeSingleCall(startTime))
}

// Average returns the average time taken for each call
func (t TimeRepeatedCalls) Average() time.Duration {
	var sum time.Duration
	for _, call := range t.calls {
		sum += call
	}
	return sum / time.Duration(len(t.calls))
}

// Max returns the maximum time taken for a call
func (t TimeRepeatedCalls) Max() time.Duration {
	return slices.Max(t.calls)
}

// Min returns the minimum time taken for a call
func (t TimeRepeatedCalls) Min() time.Duration {
	return slices.Min(t.calls)
}

// Sum returns the sum of all the time taken for each call
func (t TimeRepeatedCalls) Sum() time.Duration {
	return slices.Sum(t.calls)
}

// PrintStats prints the average, max, min and sum time taken for each call
func (t TimeRepeatedCalls) PrintStats() {
	fmt.Println("Average time taken: ", t.Average())
	fmt.Println("Max time taken: ", t.Max())
	fmt.Println("Min time taken: ", t.Min())
	fmt.Println("Sum time taken: ", t.Sum())
}

// PrintSingleCall prints the time taken for a single call
// It should be used in a defer statement, like:
// defer PrintSingleCall(Start())
func PrintSingleCall(startTime time.Time) {
	fmt.Println("Time taken for single call: ", TimeSingleCall(startTime))
}

// Start returns the current time. It is a hlper function, used to single line measurements
func Start() time.Time {
	return time.Now()
}

// TimeSingleCall returns the time taken for a single call
// It should be used in a defer statement, like:
// defer TimeSingleCall(Start())
func TimeSingleCall(startTime time.Time) time.Duration {
	return time.Since(startTime)
}
