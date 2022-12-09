package day5

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type CratesReader struct {
	scanner *bufio.Scanner
	rx      *regexp.Regexp
}

func NewCratesReader(f *os.File) *CratesReader {
	cr := new(CratesReader)
	cr.scanner = bufio.NewScanner(f)

	rx, err := regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")
	if err != nil {
		panic("Invalid regexp")
	}
	cr.rx = rx

	return cr
}

// Reads the initial crate configuration from the input.
func (cr *CratesReader) ReadCrateStacks() (cs *CrateStacks, ok bool) {
	// Read to the last line to determine the total number of columns
	lines := []string{}
	for {
		cr.scanner.Scan()
		line := cr.scanner.Text()
		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	// Use the last line to build the index lookup for all columns
	columnsIndexes := []int{}
	for i, c := range []rune(lines[len(lines)-1]) {
		// We assume that we will only have single character digit as a stack
		// identifier.
		if unicode.IsNumber(c) {
			columnsIndexes = append(columnsIndexes, i)
		}
	}
	lines = lines[0 : len(lines)-1]

	// Build up the stacks from the bottom up
	stacks := [][]string{}
	for range columnsIndexes {
		stacks = append(stacks, []string{})
	}

	for i := range lines {
		// We want to read in reversed order to build the stack from bottom up
		i = len(lines) - 1 - i

		// Convert all to runes for easier character checks
		line := []rune(lines[i])
		for stackIndex, lineIndex := range columnsIndexes {
			// Since we know that the indexes are sorted in ascending order, we can
			// be sure that other stacks will not have any crates
			if lineIndex >= len(line) {
				break
			}

			// If a stack does not have a crate, but neighboring stacks do, a space
			// character may appear. These should not be considered as crates, but
			// empty spaces.
			if !unicode.IsLetter(line[lineIndex]) {
				continue
			}

			stacks[stackIndex] = append(stacks[stackIndex], string(line[lineIndex]))
		}
	}

	return NewCrateStacks(stacks), true
}

// Reads the moves one line at a time.
func (cr *CratesReader) ReadNextMove() (qty, src, dst int, ok bool) {
	cr.scanner.Scan()
	line := cr.scanner.Text()
	if line == "" {
		return 0, 0, 0, false
	}

	if !cr.rx.MatchString(line) {
		panic("Invalid move line")
	}

	p := cr.rx.FindStringSubmatch(line)

	if len(p) < 4 {
		panic("Invalid subgroups parsed")
	}

	qty, err := strconv.Atoi(p[1])
	if err != nil {
		panic("Invalid qty")
	}
	src, err = strconv.Atoi(p[2])
	if err != nil {
		panic("Invalid src")
	}
	dst, err = strconv.Atoi(p[3])
	if err != nil {
		panic("Invalid dst")
	}

	return qty, src, dst, true
}
