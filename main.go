package main

import (
	"errors"
	"log"
	"strings"
)

func findShortestChain(start string, end string) ([]string, error) {
	if len(start) != len(end) {
		return nil, errors.New("not possible to find a chain, start and end words should have same size")
	}

	if strings.Compare(start, end) == 0 {
		log.Print("Same word for start and end, not much of an algorithm needed..")
		return []string{start}, nil
	}

	//todo implement algorithm
	return []string{}, nil
}
