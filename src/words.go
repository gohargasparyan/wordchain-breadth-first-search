package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const wordsPath string = "./english_words_alpha.txt"

func ReadWords() (*[]string, error) {
	file, err := os.Open(wordsPath)
	readingError := errors.New("can not read words from: " + wordsPath)

	if err != nil {
		return nil, readingError
	}

	defer file.Close()
	var words []string

	s := bufio.NewScanner(file)

	for s.Scan() {
		words = append(words, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, readingError
	}

	return &words, nil
}

func GetWordsWithSize(dictionary *[]string , size int) *[]string {
	var result []string
	for _, word := range *dictionary {
		if len(word) == size {
			result = append(result, word)
		}
	}

	return &result
}

func IsWordPresent(word string, words *[]string) bool {
	for _, w := range *words {
		if strings.Compare(word, w) == 0 {
			return true
		}
	}
	return false
}

/*
 	Neighbors to the given word are all those words that have same size and differ with one letter
 */
func GetNeighbors(word string, words *[]string) *[]string {
	var result []string
	for _, w := range *words {
		if isNeighbor(w, word) {
			result = append(result, w)
		}
	}

	return &result
}


func isNeighbor(word1, word2 string) bool {
	var diff int
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}

	return diff == 1
}
