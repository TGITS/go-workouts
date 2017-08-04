package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int
	fmt.Scan(&N)
	results := make(map[int]string)
	temporary := make([]int, N, N)
	fmt.Fprintf(os.Stderr, "temporary before sort after init: %v\n", temporary)
	for i := 0; i < N; i++ {
		var t string
		fmt.Scan(&t)
		timeInInt, _ := strconv.Atoi(strings.Replace(t, ":", "", 3))
		fmt.Fprintf(os.Stderr, "timeInInt : %v\n", timeInInt)
		fmt.Fprintf(os.Stderr, "time : %v\n", t)
		temporary[i] = timeInInt
		results[timeInInt] = t
	}

	fmt.Fprintf(os.Stderr, "temporary before sort: %v\n", temporary)
	sort.Ints(temporary)
	fmt.Fprintf(os.Stderr, "temporary : %v\n", temporary)
	fmt.Fprintf(os.Stderr, "results : %v\n", results)
	fmt.Println(results[temporary[0]])
}
