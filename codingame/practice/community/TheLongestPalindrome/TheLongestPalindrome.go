package main

import (
	"fmt"
	// "os"
	"sort"
)

/*
 * Longest Palindrome
 * A contribution by Coni
 * https://www.codingame.com/ide/puzzle/longest-palindrome
 * A palindrome is a sequence of letters which reads the same backward as forward, like “madam” for example.
 * For a given input string S, you have to return the longest palindrome found within.
 * If multiple substrings qualify, print them all in the same order as they appear in S.
 *
 * I have excluded from the results of findPalindromes the trivial 1 character palindrome
 */

func main() {
	var s string
	fmt.Scan(&s)
	// fmt.Fprintf(os.Stderr, "input string : %s\n", s)
	results := findLongestPalindromes(s)
	for _, s := range results {
		fmt.Println(s)
	}
}

//This solution with slices of strings to obtain substrings will not function correctly
//with strings that have characters encoded on more than one byte
func findPalindromes(s string) []string {
	size := len(s)
	palindromes := make([]string, 0)
	for i := 0; i < size; i++ {
		// j > i+1 to exclude the possibility of the trivial 1 character palindrome
		for j := size; j > i+1; j-- {
			if isPalindrome(s[i:j]) {
				palindromes = append(palindromes, s[i:j])
			}
		}
	}
	return palindromes
}

func findLongest(slice []string) []string {
	sizes := make([]int, 0)
	stringBySize := make(map[int][]string)
	for _, s := range slice {
		size := len(s)
		sizes = append(sizes, size)
		if v, ok := stringBySize[size]; !ok {
			v = make([]string, 0)
			v = append(v, s)
			stringBySize[size] = v
		} else {
			v = append(v, s)
			stringBySize[size] = v
		}
	}
	sort.Ints(sizes)
	return stringBySize[sizes[len(sizes)-1]]
}

//This solution with slices of strings to obtain substrings will not function correctly
//with strings that have characters encoded on more than one byte
func findLongestPalindromes(s string) []string {
	size := len(s)
	palindromeMaxSize := 0
	palindromes := make([]string, 0)
	for i := 0; i < size && (size-i) >= palindromeMaxSize; i++ {
		// j > i+1 to exclude the possibility of the trivial 1 character palindrome
		for j := size; j > i+1; j-- {
			current := s[i:j]
			if currentSize := len(current); currentSize >= palindromeMaxSize {
				if isPalindromeOptimized(current) {
					if currentSize == palindromeMaxSize {
						palindromes = append(palindromes, current)
					} else if currentSize > palindromeMaxSize {
						palindromeMaxSize = currentSize
						palindromes = []string{current}
					}
				}
			}
		}
	}
	return palindromes
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

//Does not work with strings that have characters encoded on more than one byte
//But for the CodinGame Puzzle with runes, it wasn't fast enough
func isPalindromeOptimized(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
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
