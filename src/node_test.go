package main

import (
	"testing"
)

func TestNodeToChain(t *testing.T) {
	testData := []struct {
		node   node
		result []string
	}{
		{
			node:   node{value: "word", parent: nil},
			result: []string{"word"},
		},
		{
			node:   node{value: "child", parent: &node{value: "root", parent: nil}},
			result: []string{"root", "child"},
		},
		{
			node:   node{value: "child_lvl_2", parent: &node{value: "child_lvl_1", parent: &node{value: "root", parent: nil}}},
			result: []string{"root", "child_lvl_1", "child_lvl_2"},
		},
	}

	for _, e := range testData {
		actual := NodeToChain(e.node)

		for i, el := range *actual {
			if el != e.result[i] {
				t.Errorf("Error actual = %v, and Expected = %v.", el, e.result[i])
			}
		}
	}

}
