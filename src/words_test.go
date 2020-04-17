package main

import (
	"testing"
)

func TestIsWordPresent(t *testing.T) {
	testData := []struct {
		words  []string
		word   string
		result bool
	}{
		{
			words:  []string{"abandons", "abandum", "abanet", "abanga"},
			word:   "abannition",
			result: false,
		},
		{
			words:  []string{"abandons", "abandum", "abanet", "abanga"},
			word:   "abandons",
			result: true,
		},
	}

	for _, e := range testData {
		actual := IsWordPresent(e.word, &e.words)
		if actual != e.result {
			t.Errorf("Error actual = %v, and Expected = %v.", actual, e.result)
		}
	}
}

func TestGetWordsWithSize(t *testing.T) {
	testData := []struct {
		words  []string
		size   int
		result []string
	}{
		{
			words:  []string{"a", "aa", "aa", "aaa"},
			size:   2,
			result: []string{"aa", "aa"},
		},
		{
			words:  []string{"a", "aa", "aa", "aaa"},
			size:   5,
			result: []string{},
		},
	}

	for _, e := range testData {
		actual := GetWordsWithSize(&e.words, e.size)
		for i, el := range *actual {
			if el != e.result[i] {
				t.Errorf("Error actual = %v, and Expected = %v.", el, e.result[i])
			}
		}
	}
}

func TestGetNeighbors(t *testing.T) {
	testData := []struct {
		words  []string
		word   string
		result []string
	}{
		{
			words:  []string{"rube", "ruby", "rubs", "ruly", "code", "robe", "rode"},
			word:   "ruby",
			result: []string{"rube", "rubs", "ruly"},
		},
	}

	for _, e := range testData {
		actual := GetNeighbors(e.word, &e.words)

		for i, el := range *actual {
			if el != e.result[i] {
				t.Errorf("Error actual = %v, and Expected = %v.", el, e.result[i])
			}
		}
	}
}

func TestIsNeighbor(t *testing.T) {
	testData := []struct {
		word1  string
		word2  string
		result bool
	}{
		{
			word1:  "ruby",
			word2:  "rube",
			result: true,
		},
		{
			word1:  "rube",
			word2:  "code",
			result: false,
		},
		{
			word1:  "ruby",
			word2:  "ruby",
			result: false,
		},
	}

	for _, e := range testData {
		actual := isNeighbor(e.word1, e.word2)

		if actual != e.result {
			t.Errorf("Error actual = %v, and Expected = %v.", actual, e.result)
		}
	}
}
