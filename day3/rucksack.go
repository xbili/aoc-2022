package day3

import "strings"
import "xbili.com/aoc/utils"

type Rucksack struct {
	// Items contained in each compartment
	compartments [][]string
}

// Builds a new rucksack using the encoded string.
func NewRucksack(encoded string) *Rucksack {
	compartmentSize := len(encoded) / 2
	first := strings.Split(encoded[0:compartmentSize], "")
	second := strings.Split(encoded[compartmentSize:], "")

	var rucksack = Rucksack{
		compartments: [][]string{first, second},
	}

	return &rucksack
}

func (rucksack *Rucksack) FindDuplicatedItem() (item string, ok bool) {
	sets := make([]*utils.Set[string], len(rucksack.compartments))

	// Build up set of items for all compartments
	for i, compartment := range rucksack.compartments {
		sets[i] = utils.NewSet(compartment)
	}

	// Find common item across all compartments
	var intersection = utils.IntersectAll(sets)
	for _, item := range intersection.GetItems() {
		return item, true
	}

	return "", false
}

func (rucksack *Rucksack) CreateItemSet() *utils.Set[string] {
	items := []string{}
	for _, compartment := range rucksack.compartments {
		items = append(items, compartment...)
	}

	return utils.NewSet(items)
}
