package day3

import (
	"bufio"
	"os"
)

type RucksackGroupReader struct {
	scanner *bufio.Scanner
}

func NewRucksackGroupReader(f *os.File) *RucksackGroupReader {
	rr := new(RucksackGroupReader)
	rr.scanner = bufio.NewScanner(f)

	return rr
}

func (sr *RucksackGroupReader) Next(groupSize int) (group []string, ok bool) {
	group = make([]string, groupSize)
	for i := 0; i < groupSize; i++ {
		sr.scanner.Scan()
		line := sr.scanner.Text()

		if line == "" {
			return group, false
		}

		if err := sr.scanner.Err(); err != nil {
			panic(err)
		}

		group[i] = line
	}

	return group, true
}
