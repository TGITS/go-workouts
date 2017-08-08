package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestCreateElement(t *testing.T) {
	pattern := "X000"
	tempo := 2
	expected := [4]bool{true, false, false, false}

	element := NewElement(pattern, tempo)

	fmt.Fprintf(os.Stderr, "element : %v\n", element)
	fmt.Fprintf(os.Stderr, "expected : %v\n", expected)
	if !reflect.DeepEqual(element.ApplicablePattern, expected) {
		t.Error("Element creation not correct : ", element)
	}
}

func TestCombinePattern(t *testing.T) {
	element1 := NewElement("X000", 2)
	expected := [4]bool{true, false, false, false}
	result := CombinePattern(element1, element1)
	fmt.Fprintf(os.Stderr, "element1 : %v\n", element1)
	fmt.Fprintf(os.Stderr, "expected : %v\n", expected)

	if !reflect.DeepEqual(result, expected) {
		t.Error("Pattern combination incorrect : ", result)
	}

	element2 := NewElement("0X00", 2)
	expected = [4]bool{true, true, false, false}

	result = CombinePattern(element1, element2)
	fmt.Fprintf(os.Stderr, "element_1 : %v\n", element1)
	fmt.Fprintf(os.Stderr, "element_2 : %v\n", element2)
	fmt.Fprintf(os.Stderr, "expected : %v\n", expected)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Pattern combination incorrect : ", result)
	}
}
