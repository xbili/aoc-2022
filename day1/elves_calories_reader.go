package day1

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// A ElvesCaloriesReader reads chunks of calories from file system, and sums
// them up grouped by elf.
type ElvesCaloriesReader struct {
	scanner *bufio.Scanner
}

// Creates a new ElvesCaloriesReader object
func NewElvesCaloriesReader(f *os.File) *ElvesCaloriesReader {
	ecr := new(ElvesCaloriesReader)
	ecr.scanner = bufio.NewScanner(f)

	return ecr
}

// Reads the sum of next chunk of numbers
func (ecr *ElvesCaloriesReader) Next() (value int, ok bool) {
	currentSum := 0
	for ecr.scanner.Scan() {
		line := ecr.scanner.Text()

		if line == "" {
			return currentSum, true
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentSum += calories
		}
	}

	if err := ecr.scanner.Err(); err != nil {
		panic(err)
	}

	return math.MinInt, false
}
