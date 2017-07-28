package main

import "fmt"

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// n: the number of cards for player 1
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		// cardp1: the n cards of player 1
		var cardp1 string
		fmt.Scan(&cardp1)
	}
	// m: the number of cards for player 2
	var m int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		// cardp2: the m cards of player 2
		var cardp2 string
		fmt.Scan(&cardp2)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("PAT") // Write answer to stdout
}
