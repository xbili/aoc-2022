package day7

import (
	"fmt"
	"math"
	"os"
)

func Run(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tr := NewTerminalReader(f)
	context := NewContext()
	ok := context.Build(tr)
	if !ok {
		panic("Context not built")
	}

	sizes := context.ComputeDirectoriesSizes()

	// Part 1
	part1 := 0
	for _, size := range sizes {
		if size <= 100000 {
			part1 += size
		}
	}
	fmt.Println("Part 1:", part1)

	// Part 2
	unused := 70000000 - sizes[context.root]
	target := 30000000 - unused
	part2 := math.MaxInt
	for _, size := range sizes {
		if size < part2 && size > target {
			part2 = size
		}
	}

	fmt.Println("Part 2:", part2)
}
