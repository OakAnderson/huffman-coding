package compress

import (
	"fmt"
)

// code is a not importes type that implements a map[byte]String
// it is used to store huffman tree codes
type code map[byte]string

// node is a struct that stores huffman tree nodes information
type node struct {
	b     byte  // b stores the symbol byte
	freq  int   // freq is the frequency that this symbol appears
	left  *node // the left child of this symbol on the huffman tree
	right *node // the left child of this symbol on the huffman tree
}

// String is used to show the byte character and to avoid showing the childs
// pointers
func (n node) String() string {
	return fmt.Sprintf("{'%s': %d}", string(n.b), n.freq)
}

// find is a recursive method that find a symbol in the tree and returns its
// code in string and true if it found the symbol, empty string and false if
// not
func (n *node) find(search node) (string, bool) {
	if n.b == search.b {
		return "", true
	}
	if n.left != nil && n.freq-search.freq > 0 {
		if s, found := n.left.find(search); found {
			return "0" + s, true
		}
	}
	if n.right != nil && n.freq-search.freq > 0 {
		if s, found := n.right.find(search); found {
			return "1" + s, true
		}
	}
	return "", false
}

// encode is the core of huffman's algorithm. Here the symbols is putted on
// a huffman tree and returns the code table for the sequence of frequencies
func encode(freqs frequencies) code {
	tree := make(frequencies, len(freqs))
	copy(tree, freqs)

	for len(tree) > 1 {
		node0 := &tree[0]
		node1 := &tree[1]
		var newNode node
		if node0.freq == node1.freq && node1.b == 0 {
			node0, node1 = node1, node0
		}
		newNode.left, newNode.right = node1, node0
		newNode.freq = node0.freq + node1.freq
		tree = append(tree[2:], newNode)
		tree.sort()
	}
	return freqs.getCodes(tree[0])
}

// reverseCode is used for decode, when you got the string code and wants the
// byte
func reverseCode(table code) map[string]byte {
	newTable := make(map[string]byte)
	for k, v := range table {
		newTable[v] = k
	}
	return newTable
}
