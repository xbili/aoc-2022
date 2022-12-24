package day9

import (
	"fmt"
	"os"
)

func Run(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	mr := NewMovementReader(f)
	bridge1 := NewBridge([2]int{0, 0}, 2)
	bridge2 := NewBridge([2]int{0, 0}, 10)

	for {
		direction, distance, ok := mr.Next()
		if !ok {
			break
		}

		bridge1.Move(direction, distance)
		bridge2.Move(direction, distance)
	}

	fmt.Printf("Part 1 tail visited: %d\n", bridge1.GetTotalVisited())
	fmt.Printf("Part 2 tails visited: %d\n", bridge2.GetTotalVisited())
}
