package day5

import (
	"xbili.com/aoc/utils"
)

type CrateStacks struct {
	// Each stack represented as a slice, ordered from bottom to top
	// i.e. the top item is the last index
	stacks [][]string
}

func NewCrateStacks(stacks [][]string) *CrateStacks {
	cs := new(CrateStacks)
	cs.stacks = stacks

	return cs
}

// Moves a certain quantity of crates from src stack to dst stack
func (cs *CrateStacks) Move(qty, src, dst int, keepOrder bool) {
	// src and dst are 1-indexed, we need to decrement to map to array indexes
	srcStack, dstStack := cs.stacks[src-1], cs.stacks[dst-1]

	// Since crates are moved one at a time, we will reverse the order when we
	// append them into the new stack.
	toMove := srcStack[len(srcStack)-qty:]
	if !keepOrder {
		toMove = utils.ReverseSlice(toMove)
	}

	cs.stacks[dst-1] = append(dstStack, toMove...)
	cs.stacks[src-1] = srcStack[0 : len(srcStack)-qty]
}

func (cs *CrateStacks) GetTopCrates() []string {
	top := []string{}
	for _, stack := range cs.stacks {
		top = append(top, stack[len(stack)-1])
	}

	return top
}
