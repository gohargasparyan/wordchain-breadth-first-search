package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	ruby := "ruby"
	code := "code"
	res, err := findShortestChain(ruby, code)
	if err !=nil {
		log.Fatal(err)
	}

	fmt.Printf("From '%s' to '%s' I can get you in %d steps, here it is: %v \n",ruby, code, len(res), res)
	fmt.Println("Lets try another one")

	cat := "cat"
	dog := "dog"
	res, err = findShortestChain(cat, dog)
	if err !=nil {
		log.Fatal(err)
	}

	fmt.Printf("From '%s' to '%s' I can get you in %d steps, here it is: %v",cat, dog, len(res), res)
}

func findShortestChain(start string, end string) ([]string, error) {
	if len(start) != len(end) {
		return nil, errors.New("not possible to find a chain, start and end words should have same size")
	}

	if strings.Compare(start, end) == 0 {
		log.Print("Same word for start and end, not much of an algorithm needed..")
		return []string{start}, nil
	}

	words, err := ReadWords()
	if err != nil {
		return nil, err
	}

	return findShortestChainInDictionary(start, end, words)
}

func findShortestChainInDictionary(start string, end string, words *[]string) ([]string, error) {
	//filter to words that have same size as start and end words
	words = GetWordsWithSize(words, len(start))

	if !IsWordPresent(start, words) || !IsWordPresent(end, words) {
		return nil, errors.New("start or end word is not in the dictionary")
	}

	//root node
	toVisit := []node{
		{
			parent: nil,
			value:  start,
		},
	}

	var visited []string
	var neighbors *[]string
	var endNode node

	for i := 0; i < len(toVisit); i++ {

		currentWord := toVisit[i].value
		if strings.Compare(currentWord, end) == 0 {
			endNode = toVisit[i]
			return *NodeToChain(endNode), nil
		}

		neighbors = GetNeighbors(currentWord, words)

		for _, neighbor := range *neighbors {
			if !Contains(neighbor, &visited) {
				visited = append(visited, neighbor)

				node := node {
					value:  neighbor,
					parent: &toVisit[i],
				}

				toVisit = append(toVisit, node)
			}
		}

		neighbors = nil
	}

	return nil, errors.New("looked up all dictionary and couldn't find a chain")
}
