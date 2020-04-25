package compress

import (
	"fmt"
	"strings"
)

// frequencies is a not imported type that implements methods that help to work
// with byte frequencies.
type frequencies []node

// frequencies.String is a method that returns a formatted string for fmt
// package. I use
func (freqs frequencies) String() string {
	var values []string
	for _, v := range freqs {
		values = append(values, fmt.Sprintf("%v", v))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

// frequencies.sort is a method that uses the insertion sort algorithm
func (freqs frequencies) sort() {
	for i := 0; i < len(freqs); i++ {
		cursor := freqs[i]
		pos := i

		for pos > 0 && freqs[pos-1].freq >= cursor.freq {
			freqs[pos] = freqs[pos-1]
			pos--
		}
		freqs[pos] = cursor
	}
}

// frequencies.contains check if the passed frequencies slice has already an element
func (freqs frequencies) contains(b byte) (bool, int) {
	for i, v := range freqs {
		if v.b == b {
			return true, i
		}
	}
	return false, -1
}

// getNodes receives a root node and for every byte in frequencies slice it
// will save into a map[byte]string type, where the value is the code for
// respective key or byte
func (freqs frequencies) getCodes(root node) (c code) {
	c = make(code)
	for _, n := range freqs {
		c[n.b], _ = root.find(n)
	}
	return
}
