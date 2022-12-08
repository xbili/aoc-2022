package day4

import (
	"fmt"
	"os"
)

func Run(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ar := NewAssignmentReader(f)
	fullyContains := 0
	contains := 0

	for {
		range1, range2, ok := ar.Next()
		if !ok {
			break
		}

		sectionRange1 := NewSectionRange(range1)
		sectionRange2 := NewSectionRange(range2)

		if sectionRange1.FullyContains(sectionRange2) ||
			sectionRange2.FullyContains(sectionRange1) {
			fullyContains++
		}

		// Don't need to check the other direction since this applies for both
		// directions
		if sectionRange1.Contains(sectionRange2) {
			contains++
		}
	}

	fmt.Printf("Part 1: %d\n", fullyContains)
	fmt.Printf("Part 2: %d\n", contains)
}
