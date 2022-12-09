package day5

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

	cr := NewCratesReader(f)

	// Read initial crate configuration
	cs, ok := cr.ReadCrateStacks()
	if !ok {
		panic("Failed to read crate stacks")
	}

	// Read and process moves one line at a time
	for {
		qty, src, dst, ok := cr.ReadNextMove()
		if !ok {
			break
		}

		// Part 1 does not keep crates in order
		cs.Move(qty, src, dst, false)
	}

	fmt.Println("Part 1:", cs.GetTopCrates())
}

func RunPart2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cr := NewCratesReader(f)

	// Read initial crate configuration
	cs, ok := cr.ReadCrateStacks()
	if !ok {
		panic("Failed to read crate stacks")
	}

	// Read and process moves one line at a time
	for {
		qty, src, dst, ok := cr.ReadNextMove()
		if !ok {
			break
		}

		// Part 2 keeps crates in order
		cs.Move(qty, src, dst, true)
	}

	fmt.Println("Part 2:", cs.GetTopCrates())
}
