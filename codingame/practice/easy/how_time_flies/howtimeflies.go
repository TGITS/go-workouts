package main

import (
	"fmt"
	"math"
	"time"

	//"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

//Input:
//01.01.2000
//01.01.2016

//Output:
//16 years, total 5844 days

func main() {
	var BEGIN string
	fmt.Scan(&BEGIN)

	var END string
	fmt.Scan(&END)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(computeDuration(BEGIN, END)) // Write answer to stdout
}

func parseDate(date string) time.Time {
	//example date in string input : 01.01.2000
	slice := strings.Split(date, ".")
	year, _ := strconv.Atoi(slice[2])
	month, _ := strconv.Atoi(slice[1])
	day, _ := strconv.Atoi(slice[0])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func isLeapYear(year int) bool {
	//Leap Years are any year that can be exactly divided by 4 (such as 2012, 2016, etc)
	//except if it can be exactly divided by 100, then it isn't (such as 2100, 2200, etc)
	//except if it can be exactly divided by 400, then it is (such as 2000, 2400)

	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	} else {
		return false
	}
}

func computeDuration(startDate string, endDate string) string {

	daysByMonth := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	beginning := parseDate(startDate)
	end := parseDate(endDate)
	duration := end.Sub(beginning)
	hours := duration.Hours()
	days := int64(math.Round(hours / float64(24)))

	beginningYear, beginningMonth, beginningDay := beginning.Date()
	endYear, endMonth, endDay := end.Date()

	for beginningDay != endDay {
		beginningDay++
		if (beginningMonth == 2 && isLeapYear(beginningYear) && beginningDay > (daysByMonth[int(beginningMonth)]+1)) || (beginningDay > daysByMonth[int(beginningMonth)]) {
			beginningMonth++
			if beginningMonth > 12 {
				beginningYear++
				beginningMonth = 1
			}
			beginningDay = 1
		}
	}

	countMonth := 0
	for beginningMonth != endMonth {
		beginningMonth++
		countMonth++
		if beginningMonth > 12 {
			beginningYear++
			beginningMonth = 1
		}
	}

	countYear := 0
	for beginningYear != endYear {
		beginningYear++
		countYear++
	}

	var strs []string

	if countYear == 1 {
		strs = append(strs, "1 year, ")
	} else if countYear > 1 {
		strs = append(strs, fmt.Sprintf("%d years, ", countYear))
	}

	if countMonth == 1 {
		strs = append(strs, "1 month, ")
	} else if countMonth > 1 {
		strs = append(strs, fmt.Sprintf("%d months, ", countMonth))
	}

	strs = append(strs, fmt.Sprintf("total %d days", days))

	return strings.Join(strs, "")
}
