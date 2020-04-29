package file

import (
	"io/ioutil"
	"os"

	"github.com/OakAnderson/huffman-coding/compress"
	"github.com/icza/bitio"
)

// Decode is a function that decodes a .hff file, econded by
// github.com/OakAnderson/huffman-coding/file.Encode
func Decode(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := bitio.NewReader(file)
	defer file.Close()

	keys, sizes, err := readSymSizes(r)
	if err != nil {
		return nil, err
	}

	table, err := readCode(keys, sizes, r)
	if err != nil {
		return nil, err
	}

	decoded, err := readData(r)
	if err != nil {
		return nil, err
	}

	decoded = compress.HuffmanDecode(decoded, table)
	return decoded, nil
}

// Encode is a function that use huffman coding to
// compress the passed srcfile and saves the result into the
// dstfile
func Encode(srcfile, dstfile string) error {
	sourcebytes, _ := ioutil.ReadFile(srcfile)
	encoded, table := compress.HuffmanEncode(sourcebytes)

	f, w, err := loadFileWriter(dstfile)
	if err != nil {
		return err
	}
	defer f.Close()
	defer w.Close()

	writeCodes(w, table)

	for i := 0; i < len(encoded) && w.TryError == nil; i++ {
		w.TryWriteBool(encoded[i] == '1')
	}

	return w.TryError
}

// loadFileWriter returns a os.file and an writer from the passed file
func loadFileWriter(filename string) (*os.File, *bitio.Writer, error) {
	file, err := os.Create(filename)
	w := bitio.NewWriter(file)

	return file, w, err
}

// writeCodes write the symbols and its codes to the passed writer
func writeCodes(w *bitio.Writer, table compress.Code) error {
	keys := table.Keys()
	for _, k := range keys {
		size := len(table[k])
		w.TryWrite([]byte{k, byte(size)})
	}
	w.TryWrite([]byte{'\u0000', '\u0000'})

	for _, k := range keys {
		for _, bit := range table[k] {
			w.TryWriteBool(bit == '1')
		}
	}

	return w.TryError
}

// This function read pairs from passed bitio.Reader and store it into
// two slices, one of read keys and another of code sizes
func readSymSizes(r *bitio.Reader) (keys []byte, sizes []int, err error) {
	ks := []byte{'0', '0'}
	for r.TryRead(ks); r.TryError == nil && ks[0] != '\u0000'; r.TryRead(ks) {
		keys = append(keys, ks[0])
		sizes = append(sizes, int(ks[1]))
	}

	return nil, nil, r.TryError
}

// readCode read a block of the file to get the codes from the passed symbols
func readCode(keys []byte, sizes []int, r *bitio.Reader) (compress.Code, error) {
	table := make(compress.Code)
	var val string
	for idx, k := range keys {
		val = ""
		for i := 0; i < sizes[idx]; i++ {
			if r.TryReadBool() {
				val += "1"
			} else {
				val += "0"
			}
		}
		table[k] = val
	}

	return table, r.TryError
}

// readData reads the reserved file block to compressed data
func readData(r *bitio.Reader) ([]byte, error) {
	var data []byte
	for r.TryError == nil {
		if r.TryReadBool() {
			data = append(data, '1')
		} else {
			data = append(data, '0')
		}
	}
	return data, r.TryError
}
