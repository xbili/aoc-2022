package day6

import (
	"bufio"
	"os"
)

type StreamReader struct {
	scanner *bufio.Scanner
}

func NewStreamReader(f *os.File) *StreamReader {
	sr := new(StreamReader)
	sr.scanner = bufio.NewScanner(f)

	return sr
}

func (sr *StreamReader) Read() (chunk string, ok bool) {
	sr.scanner.Scan()
	chunk = sr.scanner.Text()

	return chunk, true
}
