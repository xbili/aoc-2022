package day3

import "xbili.com/aoc/utils"

// Finds a badge within a group
func FindBadge(group []string) (badge string, ok bool) {
	// Get all item sets for all rucksacks in the group
	sets := make([]*utils.Set[string], len(group))
	for i, input := range group {
		rucksack := NewRucksack(input)
		sets[i] = rucksack.CreateItemSet()
	}

	intersection := utils.IntersectAll(sets)

	for _, item := range intersection.GetItems() {
		return item, true
	}

	return "", false
}
