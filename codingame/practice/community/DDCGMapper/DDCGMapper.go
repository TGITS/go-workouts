package main

import "fmt"

//import "os"

/*
 * https://www.codingame.com/training/community/ddcg-mapper
 * Your program must automate the creation of maps for the new game "Dance Dance CodinGame". In this game, the player must push 4 keys in rhythm.
 * One map is composed of lines of 4 characters each: a zero (0) indicates that the corresponding key is released, a cross (X) indicates that the corresponding key is pushed.
 *
 * You are given the patterns of lines Pattern and their tempo Tempo. You must reproduce the pattern every Tempo lines.
 * If one line has no pattern, it is composed of 4 zeros: 0000.
 * If one line has multiple patterns, you must accumulate cross. For example, XX00 and X0X0 becomes XXX0.
 *
 *  Warning: the map starts from the bottom to the top!
 * Input
 *  Line 1: The length L of the map.
 *  Line 2: The number N of pair Pattern Tempo.
 *  N next lines: A string Pattern and a number Tempo.
 * Output
 *  L lines representing the map.
 * Constraints
 *   0 < L < 100
 *   0 < N < 10
 *   0 < Tempo < 100
 */

//The Element that will allow to build the map
type Element struct {
	Pattern           string
	Tempo             int
	ApplicablePattern [4]bool
}

func NewElement(pattern string, tempo int) *Element {
	element := new(Element)
	element.Pattern = pattern
	element.Tempo = tempo
	for i, v := range pattern {
		if c := string(v); c == "0" {
			element.ApplicablePattern[i] = false
		} else {
			element.ApplicablePattern[i] = true
		}
	}
	return element
}

func CombinePattern(elements ...*Element) [4]bool {
	result := [4]bool{false, false, false, false}
	for _, e := range elements {
		for i, v := range e.ApplicablePattern {
			result[i] = result[i] || v
		}
	}
	return result
}

func ToPattern(applicablePattern [4]bool) string {
	s := ""
	for _, b := range applicablePattern {
		if b {
			s = s + "X"
		} else {
			s = s + "0"
		}
	}
	return s
}

func main() {
	var L int // Number of lines of the map
	fmt.Scan(&L)
	gameMap := make([]string, L, L)

	var N int // Number of pair Pattern/Tempo
	fmt.Scan(&N)
	elements := make([]*Element, N, N)

	for i := 0; i < N; i++ {
		var pattern string
		var tempo int
		fmt.Scan(&pattern, &tempo)
		elements[i] = NewElement(pattern, tempo)
	}

	for i := 0; i < L; i++ {
		applicableElements := make([]*Element, 0, 0)
		for _, v := range elements {
			if (i+1)%v.Tempo == 0 {
				applicableElements = append(applicableElements, v)
			}
		}
		gameMap[i] = ToPattern(CombinePattern(applicableElements...))
	}
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	for i := L - 1; i >= 0; i-- {
		fmt.Println(gameMap[i])
	}
	//fmt.Println("answer") // Write answer to stdout
}
