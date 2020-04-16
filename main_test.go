package main

import (
	"testing"
)

func TestStartEndHaveDifferentSizes(t *testing.T) {
	expected := "not possible to find a chain, start and end words should have same size"
	_, actual := findShortestChain("a", "ab")

	if actual == nil || actual.Error() != expected {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
	}
}

func TestStartEndAreSameWord(t *testing.T) {
	expected := []string{"same"}
	actual, _ := findShortestChain("same", "same")

	if len(actual) != len(expected) {
		t.Errorf("Error actual length = %d, and Expected = 1%d", len(actual), len(expected))
	}

	if actual[0] != expected[0] {
		t.Errorf("Error actual = %v, and Expected = %v.", actual[0], expected[0])
	}
}
