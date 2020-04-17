package main

type node struct {
	parent *node
	value  string
}

func NodeToChain(node node) *[]string {
	chain := []string {node.value}


	for node.parent != nil {
		chain = append(chain, node.parent.value)
		node = *node.parent
	}

	var reverse []string
	for i := len(chain)-1; i >= 0; i-- {
		reverse = append(reverse, chain[i])
	}

	return &reverse
}
