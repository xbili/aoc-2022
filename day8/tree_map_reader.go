package day8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type TreeMapReader struct {
	scanner *bufio.Scanner
}

func NewTreeMapReader(f *os.File) *TreeMapReader {
	tr := new(TreeMapReader)
	tr.scanner = bufio.NewScanner(f)

	return tr
}

func (tr *TreeMapReader) ReadMap() [][]int {
	grid := [][]int{}

	for {
		tr.scanner.Scan()
		line := tr.scanner.Text()

		if line == "" {
			break
		}

		tokens := strings.Split(line, "")
		row := make([]int, len(tokens))
		for i, c := range tokens {
			val, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}

			row[i] = val
		}

		grid = append(grid, row)
	}

	return grid
}
