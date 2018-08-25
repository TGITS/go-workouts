package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

type testpair struct {
	GivenInput     string
	ExpectedOutput time.Time
}

func TestDateParsing(t *testing.T) {

	tests := []testpair{
		{"01.01.2000", time.Date(2000, time.Month(1), 1, 0, 0, 0, 0, time.UTC)},
		{"01.01.2016", time.Date(2016, time.Month(1), 1, 0, 0, 0, 0, time.UTC)},
		{"15.03.2014", time.Date(2014, time.Month(3), 15, 0, 0, 0, 0, time.UTC)},
		{"31.03.2061", time.Date(2061, time.Month(3), 31, 0, 0, 0, 0, time.UTC)},
		{"30.11.1997", time.Date(1997, time.Month(11), 30, 0, 0, 0, 0, time.UTC)},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Input string :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected time :", pair.ExpectedOutput)
		date := parseDate(pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Parse date :", date)
		if !reflect.DeepEqual(date, pair.ExpectedOutput) {
			t.Error("Not the expected result for : ", pair)
		}
	}
}

type testpair1 struct {
	GivenInput     []string
	ExpectedOutput string
}

func TestComputedDurationFormating(t *testing.T) {
	tests := []testpair1{
		{[]string{"01.01.2000", "01.01.2016"}, "16 years, total 5844 days"},
		{[]string{"15.12.2014", "14.02.2016"}, "1 year, 1 month, total 426 days"},
		{[]string{"01.01.2016", "18.08.2016"}, "7 months, total 230 days"},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Input string :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected time :", pair.ExpectedOutput)
		duration := computeDuration(pair.GivenInput[0], pair.GivenInput[1])
		fmt.Fprintln(os.Stderr, "Computed duration :", duration)
		if duration != pair.ExpectedOutput {
			t.Error("Not the expected result for : ", pair)
		}
	}
}

type testpair2 struct {
	GivenInput     int
	ExpectedOutput bool
}

//Leap Years are any year that can be exactly divided by 4 (such as 2012, 2016, etc)
//except if it can be exactly divided by 100, then it isn't (such as 2100, 2200, etc)
//except if it can be exactly divided by 400, then it is (such as 2000, 2400)
func TestIsLeapYear(t *testing.T) {
	tests := []testpair2{
		{2012, true},
		{2016, true},
		{2100, false},
		{2200, false},
		{2000, true},
		{2400, true},
		{1885, false},
		{1975, false},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Input string :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected time :", pair.ExpectedOutput)
		test := isLeapYear(pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Is leap year ? :", test)
		if test != pair.ExpectedOutput {
			t.Error("Not the expected result for : ", pair)
		}
	}
}
