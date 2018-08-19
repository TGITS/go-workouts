package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type testpair1 struct {
	GivenInput     string
	ExpectedOutput map[rune]float64
}

func TestComputeRawWeight(t *testing.T) {

	tests := []testpair1{
		{"a-a", map[rune]float64{'a': 1.0}},
		{"a-abab-b-a-b", map[rune]float64{'a': 1.5, 'b': 1.3333333333333333}},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Input string :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected weight table :", pair.ExpectedOutput)
		weights := computeWeights(pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Computed weight for input data :", weights)
		if !reflect.DeepEqual(weights, pair.ExpectedOutput) {
			t.Error("Not the expected result for : ", pair)
		}
	}
}

type testpair2 struct {
	GivenInput     string
	ExpectedOutput string
}

func TestComputeGreatestWeight(t *testing.T) {
	tests := []testpair2{
		{"a-a", "a"},
		{"a-abab-b-a-b", "a"},
	}

	for _, pair := range tests {
		fmt.Fprintln(os.Stderr, "Test pair :", pair)
		fmt.Fprintln(os.Stderr, "Input string :", pair.GivenInput)
		fmt.Fprintln(os.Stderr, "Expected weight table :", pair.ExpectedOutput)
		weights := computeWeights(pair.GivenInput)
		greatest := findGreatestWeight(weights)
		fmt.Fprintln(os.Stderr, "Greatest weights for input data :", greatest)
		if greatest != pair.ExpectedOutput {
			t.Error("Not the expected result for : ", pair)
		}
	}
}
