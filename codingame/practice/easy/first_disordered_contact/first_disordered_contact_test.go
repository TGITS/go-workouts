package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type testpair struct {
	GivenInput     string
	ExpectedOutput string
}

func TestEncode(t *testing.T) {

	tests := []testpair{
		{"hello world", "worlelhlo d"},
		{"abcdefghi", "ghibcadef"},
		{"abcdefghij", "ghijbcadef"},
		{"abcdefghijkmnop", "ghijbcadefkmnop"},
		{"abcdefghijkmno", "ghijbcadefkmno"},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Given Input  :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected output :", pair.ExpectedOutput)
		encoded := Encode(pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Encoded : ", encoded)
		if encoded != pair.ExpectedOutput {
			t.Error("Not the expected result for : ", pair)
		}
	}
}

func TestEncode2(t *testing.T) {

	tests := []testpair{
		{"hello world", "worlelhlo d"},
		{"abcdefghi", "ghibcadef"},
		{"abcdefghij", "ghijbcadef"},
		{"abcdefghijkmnop", "ghijbcadefkmnop"},
		{"abcdefghijkmno", "ghijbcadefkmno"},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Given Input  :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected output :", pair.ExpectedOutput)
		encoded := Encode2(pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Encoded : ", encoded)
		if encoded != pair.ExpectedOutput {
			t.Error("Not the expected result for : ", pair)
		}
	}
}

type testtriplet struct {
	GivenInput           string
	FirstExpectedOutput  []string
	SecondExpectedOutput []int
}

func TestRegroupBySize(t *testing.T) {

	tests := []testtriplet{
		{"hello world", []string{"h", "el", "lo ", "worl", "d"}, []int{1, 2, 3, 4, 1}},
		{"abcdefghi", []string{"a", "bc", "def", "ghi"}, []int{1, 2, 3, 3}},
		{"abcdefghij", []string{"a", "bc", "def", "ghij"}, []int{1, 2, 3, 4}},
		{"abcdefghijkmnop", []string{"a", "bc", "def", "ghij", "kmnop"}, []int{1, 2, 3, 4, 5}},
		{"abcdefghijkmno", []string{"a", "bc", "def", "ghij", "kmno"}, []int{1, 2, 3, 4, 4}},
	}

	for _, triplet := range tests {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Test triplet : %v", triplet))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Given Input : %v", triplet.GivenInput))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("First Expected output : %v", triplet.FirstExpectedOutput))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Second Expected output : %v", triplet.SecondExpectedOutput))
		stringGroup, sizeGroup := RegroupBySize(triplet.GivenInput)
		fmt.Fprintln(os.Stderr, fmt.Sprintf("String group : %v", stringGroup))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Size group : %v", sizeGroup))
		if !reflect.DeepEqual(stringGroup, triplet.FirstExpectedOutput) {
			t.Error("Not the expected result for : ", triplet)
		}
		if !reflect.DeepEqual(sizeGroup, triplet.SecondExpectedOutput) {
			t.Error("Not the expected result for : ", triplet)
		}
	}
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expectedReversedSlice := []int{5, 4, 3, 2, 1}

	if !reflect.DeepEqual(Reverse(slice), expectedReversedSlice) {
		t.Error("Not the expected result for : ", expectedReversedSlice)
	}
}

func TestDecode(t *testing.T) {

	tests := []testpair{
		{"worlelhlo d", "hello world"},
		{"ghibcadef", "abcdefghi"},
		{"ghijbcadef", "abcdefghij"},
		{"ghijbcadefkmnop", "abcdefghijkmnop"},
		{"ghijbcadefkmno", "abcdefghijkmno"},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Given Input  :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected output :", pair.ExpectedOutput)
		decoded := Decode(pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Decoded : ", decoded)
		if decoded != pair.ExpectedOutput {
			t.Error("Not the expected result for : ", pair)
		}
	}
}
