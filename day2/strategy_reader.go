package day2

import (
	"bufio"
	"os"
	"strings"
)

type StrategyReader struct {
	scanner *bufio.Scanner
}

func NewStrategyReader(f *os.File) *StrategyReader {
	sr := new(StrategyReader)
	sr.scanner = bufio.NewScanner(f)

	return sr
}

func (sr *StrategyReader) Next() (row []string, ok bool) {
	sr.scanner.Scan()
	line := sr.scanner.Text()

	if line == "" {
		return []string{}, false
	}

	moves := strings.Split(line, " ")

	if err := sr.scanner.Err(); err != nil {
		panic(err)
	}

	return moves, true
}
