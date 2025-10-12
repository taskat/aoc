package main

import (
	"strings"
	"text/template"
	"time"
)

const (
	star = "⭐"
)

// collectedStars returns the total number of stars collected across all years
func collectedStars(results FullResult) int {
	total := 0
	for _, year := range results {
		total += year.getStars()
	}
	return total
}

// customFuncMap returns a template.FuncMap with custom functions for templates
func customFuncMap() template.FuncMap {
	return template.FuncMap{
		"collectedStars": collectedStars,
		"dailyStars":     dailyStars,
		"partTime":       partTime,
		"totalStars":     totalStars,
		"starString":     starString,
		"yearlyMaxStars": yearlyMaxStars,
		"yearlyStars":    yearlyStars,
	}
}

// dailyStars returns a string of star emojis representing the number of stars collected for a day
func dailyStars(day DayResult) string {
	return strings.Repeat(star, day.getStars())
}

// partTime returns the time taken for a specific part of a day
func partTime(day DayResult, part int) string {
	times, ok := day[part]
	if !ok {
		return "-"
	}
	return times.String()
}

// starString returns a string of the star emoji
func starString() string {
	return star
}

// totalStars returns the total number of stars available up to the current day in December of the current year
func totalStars() int {
	thisYear := time.Now().Year()
	return (thisYear-2015+1)*50 + yearlyMaxStars(thisYear)
}

// yearlyMaxStars returns the maximum number of stars available in a given year
func yearlyMaxStars(year int) int {
	if year < 2015 {
		return 0
	}
	thisYear := time.Now().Year()
	if year < thisYear {
		return 50
	}
	if year == thisYear {
		if time.Now().Month() < time.December {
			return 0
		}
		today := time.Now().Day()
		if today >= 25 {
			return 50
		}
		return today * 2
	}
	return 0
}

// yearlyStars returns the number of stars collected in a given year
func yearlyStars(year YearResult) int {
	return year.getStars()
}
