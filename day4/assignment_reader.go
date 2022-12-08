package day4

import (
	"bufio"
	"os"
	"strings"
)

type AssignmentReader struct {
	scanner *bufio.Scanner
}

func NewAssignmentReader(f *os.File) *AssignmentReader {
	rr := new(AssignmentReader)
	rr.scanner = bufio.NewScanner(f)

	return rr
}

func (ar *AssignmentReader) Next() (range1 string, range2 string, ok bool) {
	ar.scanner.Scan()
	line := ar.scanner.Text()

	if line == "" {
		return "", "", false
	}

	if err := ar.scanner.Err(); err != nil {
		panic(err)
	}

	split := strings.Split(line, ",")

	return split[0], split[1], true
}
