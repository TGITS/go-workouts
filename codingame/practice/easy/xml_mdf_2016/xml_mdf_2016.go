//Puzzle CG : https://www.codingame.com/training/easy/xml-mdf-2016
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	sequence := scanner.Text()

	fmt.Fprintln(os.Stderr, "Given input sequence ", sequence)
	fmt.Println(findGreatestWeight(computeWeights(sequence)))
}

func computeWeights(s string) map[rune]float64 {
	weights := make(map[rune]float64)
	depth := 0
	up := false
	for _, r := range s {
		if r == '-' {
			depth--
			up = true
		} else if up {
			up = false
		} else {
			depth++
			value, ok := weights[r]
			if ok {
				weights[r] = value + (1.0 / float64(depth))
			} else {
				weights[r] = 1.0 / float64(depth)
			}
		}
	}
	return weights
}

func findGreatestWeight(weights map[rune]float64) string {
	greatestV := -1.0
	greatestK := ""
	for k, v := range weights {
		if greatestV < v {
			greatestV = v
			greatestK, _ = strconv.Unquote(strconv.QuoteRune(k))
		}
	}
	return greatestK
}