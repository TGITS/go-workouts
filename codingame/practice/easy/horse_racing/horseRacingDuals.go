package main

import (
	"fmt"
	"sort"
	"strconv"
	//"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int
	fmt.Scan(&N)
	powers := make([]int, N, N)
	for i := 0; i < N; i++ {
		var Pi int
		fmt.Scan(&Pi)
		powers[i] = Pi
	}
	sort.Ints(powers)
	length := len(powers) - 1
	differences := make([]int, length, length)
	for i := 0; i < length; i++ {
		differences[i] = powers[i+1] - powers[i]
	}
	sort.Ints(differences)
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(strconv.Itoa(differences[0])) // Write answer to stdout
}
