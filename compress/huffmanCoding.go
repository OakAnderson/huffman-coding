package compress

/*
HuffmanEncode compresses the passed data with the Huffman Coding algorithm. It
returns a byte slice with encoded data and the table of in codes.
example:

compress.HuffmanEncode([]byte("AAAABBBCCD")) returns:

[ '1' '1' '1' '1' '0' '0' '0' '0' '0' '0'
	'0' '1' '0' '0' '1' '0' '0' '1' '1' ]
map['A': "1", 'B': "00", 'C': "010", 'D', "011"]

The map can be used for decode the encoded byte sequence
*/
func HuffmanEncode(data []byte) ([]byte, Code) {
	freqs := symbolsFrequency(data)
	codeTable := encode(freqs)

	var encoded []byte
	for _, b := range data {
		encoded = append(encoded, []byte(codeTable[b])...)
	}
	return encoded, codeTable
}

// HuffmanDecode decodes an encoded string by HuffmanEncode function.
//
// It receives a byte sequence and a map[byte]string Code table. This
// parameters is returned by HuffmanEncode function. It also returns a byte
// sequence for the passed encoded byte sequence
func HuffmanDecode(data []byte, codeTable Code) []byte {
	var decoded []byte
	reversedTable := codeTable.Reverse()

	var word string
	for _, v := range data {
		word += string(v)
		if b, ok := reversedTable[word]; ok {
			decoded = append(decoded, b)
			word = ""
		}
	}
	return decoded
}

// symbolsFrequency is a function that returns a slice of frequencies. That
// means that this function counts every symbol in passed data and return a
// []node slice with the symbol byte and counted frequency
func symbolsFrequency(data []byte) (freqs frequencies) {
	for _, b := range data {
		if ok, indx := freqs.contains(b); ok {
			freqs[indx].freq++
		} else {
			freqs = append(freqs, node{b: b, freq: 1, left: nil, right: nil})
		}
	}
	freqs.sort()
	return
}
