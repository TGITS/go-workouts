package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "madam"
	answer := isPalindrome(s)
	if !answer {
		t.Error("Is not a palindrome : ", s)
	}
}

func TestIsNotPalindrome(t *testing.T) {
	s := "ClearlyNotAPalindrome"
	answer := isPalindrome(s)
	if answer {
		t.Error("Is a palindrome : ", s)
	}
}
