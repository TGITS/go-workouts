package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type palindrometestpair struct {
	input  string
	output bool
}

func TestIsPalindrome(t *testing.T) {

	var tests = []palindrometestpair{
		{"madam", true},
		{"été", true},
		{"ClearlyNotAPalindrome", false},
		{"Palindrome", false},
		{"a", true},
		{"aa", true},
		{"abb", false},
		{"aba", true},
		{"1111111111", true},
		{"1111111112", false},
	}

	for _, s := range tests {

		if isPalindrome(s.input) != s.output {
			t.Error("Not the expected result for : ", s)
		}
	}
}

type allpalindrometestpair struct {
	input  string
	output []string
}

func TestFindPalindromes(t *testing.T) {

	var tests = []allpalindrometestpair{
		{"madam", []string{"madam", "ada"}},
		{"ete", []string{"ete"}},
		{"1111", []string{"1111", "111", "11", "111", "11", "11"}},
		{"1112", []string{"111", "11", "11"}},
	}

	for _, s := range tests {
		answers := findPalindromes(s.input)
		fmt.Fprintf(os.Stderr, "original string : %s\n", s.input)
		fmt.Fprintf(os.Stderr, "expected results : %s\n", s.output)
		fmt.Fprintf(os.Stderr, "results : %v\n", answers)
		if !reflect.DeepEqual(s.output, answers) {
			t.Error("Not all the palindromes were found : ", answers)
		}
	}
}

func TestFindLongestPalindromes(t *testing.T) {

	var tests = []allpalindrometestpair{
		{"madam", []string{"madam"}},
		{"ete", []string{"ete"}},
		{"1111", []string{"1111"}},
		{"111222", []string{"111", "222"}},
		{"111222333111222333111222333111222333", []string{"111", "222", "333", "111", "222", "333", "111", "222", "333", "111", "222", "333"}},
	}

	for _, s := range tests {
		answers := findLongest(findPalindromes(s.input))
		fmt.Fprintf(os.Stderr, "original string : %s\n", s.input)
		fmt.Fprintf(os.Stderr, "expected results : %s\n", s.output)
		fmt.Fprintf(os.Stderr, "results : %v\n", answers)
		if !reflect.DeepEqual(s.output, answers) {
			t.Error("Not the correct longest palindromes were found : ", answers)
		}
	}
}
