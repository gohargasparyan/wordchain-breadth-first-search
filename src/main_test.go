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

func TestWordIsNotInDictionary(t *testing.T) {
	expected := "start or end word is not in the dictionary"
	_, actual := findShortestChain("aaaaaa", "aaaaab")

	if actual == nil || actual.Error() != expected {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
	}
}

func TestWordChain(t *testing.T) {
	actual, err := findShortestChain("cat", "dog")

	if err != nil {
		t.Error(err)
	}

	if len(actual) != 4 {
		t.Errorf("From cat to dog shortest chain gotta be 4")
	}
}

func Benchmark(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		findShortestChain("cat", "dog")
	}
}

func benchmark(start, end string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		findShortestChain(start, end)
	}
}

func BenchmarkForwards(b *testing.B)  { benchmark("lead", "gold", b) }

func BenchmarkBackwards(b *testing.B)  {  benchmark("gold", "lead", b) }