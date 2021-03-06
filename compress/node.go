package compress

import (
	"fmt"
)

// Code is a map used to store got symbols and respective huffman codes
type Code map[byte]string

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
// Code in string and true if it found the symbol, empty string and false if
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
// a huffman tree and returns the Code table for the sequence of frequencies
func encode(freqs frequencies) Code {
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

// Reverse is used for decode, when you have the string Code and wants the
// byte
func (c Code) Reverse() map[string]byte {
	newTable := make(map[string]byte)
	for k, v := range c {
		newTable[v] = k
	}
	return newTable
}

// Keys is a
func (c Code) Keys() []byte {
	keys := make([]byte, len(c))

	var i int
	for k := range c {
		keys[i] = k
		i++
	}

	return keys
}

// Values TODO
func (c Code) Values() []string {
	keys := make([]string, len(c))

	var i int
	for _, v := range c {
		keys[i] = v
		i++
	}

	return keys
}
