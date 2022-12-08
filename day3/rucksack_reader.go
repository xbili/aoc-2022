package day3

import (
	"bufio"
	"os"
)

type RucksackReader struct {
	scanner *bufio.Scanner
}

func NewRucksackReader(f *os.File) *RucksackReader {
	rr := new(RucksackReader)
	rr.scanner = bufio.NewScanner(f)

	return rr
}

func (sr *RucksackReader) Next() (row string, ok bool) {
	sr.scanner.Scan()
	line := sr.scanner.Text()

	if line == "" {
		return "", false
	}

	if err := sr.scanner.Err(); err != nil {
		panic(err)
	}

	return line, true
}
