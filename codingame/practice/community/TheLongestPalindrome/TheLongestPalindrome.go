package main

import "fmt"

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var s string
	fmt.Scan(&s)
}

func isPalindrome(s string) bool {
	runes := toRuneSlice(s)
	reversed := make([]rune, len(runes))
	copy(reversed, runes)
	reverseRuneSlice(reversed)
	return string(runes) == string(reversed)
}

func toRuneSlice(s string) []rune {
	n := 0
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}
	return runes[0:n]
}

func reverseRuneSlice(runes []rune) {
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

}

// Adapted from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
// and Russ Cox, on the golang-nuts mailing list (https://groups.google.com/forum/#!topic/golang-nuts/oPuBaYJ17t4)
// Also see : https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverseString(input string) string {
	runes := toRuneSlice(input)
	reverseRuneSlice(runes)
	// Convert back to UTF-8.
	//output := string(rune)
	//fmt.Println("output : " + output)
	return string(runes)
}
