package day1

import (
	"sort"
	"xbili.com/aoc/utils"
)

// Finds the sum of the total calories carried by the richest K elves.
func FindRichestElves(k int, input utils.Iterator[int]) int {
	calories := []int{}
	for {
		calorie, ok := input.Next()
		if !ok {
			break
		}

		calories = append(calories, calorie)
	}

	// We sort them and sum the biggest K values
	sort.Ints(calories)

	// Lower bound needs to be truncated if there are fewer than 3 elves
	bound := len(calories) - k
	if bound < 0 {
		bound = 0
	}

	res := 0
	for _, calorie := range calories[bound:] {
		res += calorie
	}

	return res
}
