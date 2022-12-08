package day2

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

	sr := NewStrategyReader(f)

	total := 0
	for {
		row, ok := sr.Next()
		if !ok {
			break
		}

		opponentShape, selfShape := DecodeOpponent(row[0]), DecodeSelf(row[1])
		total += CalculateScore(selfShape, opponentShape)
	}

	fmt.Printf("Part 1: %d\n", total)
}

func RunPart2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sr := NewStrategyReader(f)

	total := 0
	for {
		row, ok := sr.Next()
		if !ok {
			break
		}

		opponentShape, desiredOutcome := DecodeOpponent(row[0]), DecodeOutcome(row[1])
		selfShape := FindShapeForDesiredOutcome(opponentShape, desiredOutcome)
		total += CalculateScore(selfShape, opponentShape)
	}

	fmt.Printf("Part 2: %d\n", total)
}
