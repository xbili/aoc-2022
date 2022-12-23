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
	bridge := NewBridge([2]int{0, 0})

	for {
		direction, distance, ok := mr.Next()
		if !ok {
			break
		}

		bridge.Move(direction, distance)
	}

	fmt.Printf("Tail visited: %d\n", bridge.GetTotalVisited())
}
