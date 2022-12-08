package day3

import (
	"fmt"
	"os"
)

func RunPart1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rr := NewRucksackReader(f)

	total := 0
	for {
		row, ok := rr.Next()
		if !ok {
			break
		}

		rucksack := NewRucksack(row)
		duplicatedItem, ok := rucksack.FindDuplicatedItem()
		if !ok {
			panic("No duplicated item found")
		}

		total += CalculateScore(duplicatedItem)
	}

	fmt.Printf("Part 1: %d\n", total)
}

func RunPart2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rr := NewRucksackGroupReader(f)

	total := 0
	for {
		group, ok := rr.Next(3)
		if !ok {
			break
		}

		badge, ok := FindBadge(group)
		if !ok {
			panic("Cannot find badge")
		}

		total += CalculateScore(badge)
	}

	fmt.Printf("Part 2: %d\n", total)
}
