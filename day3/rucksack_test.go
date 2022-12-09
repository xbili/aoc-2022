package day3

import "testing"

func TestFindDuplicatedItem(t *testing.T) {
	assertDuplicatedItem(t, "AA", "A")
	assertDuplicatedItem(t, "vJrwpWtwJgWrhcsFMMfFFhFp", "p")
	assertDuplicatedItem(t, "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "L")
	assertDuplicatedItem(t, "PmmdzqPrVvPwwTWBwg", "P")
	assertDuplicatedItem(t, "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "v")
	assertDuplicatedItem(t, "ttgJtRGJQctTZtZT", "t")
	assertDuplicatedItem(t, "CrZsJsPPZsGzwwsLwLmpwMDw", "s")
}

func assertDuplicatedItem(t *testing.T, input string, expected string) {
	rucksack := NewRucksack(input)
	duplicatedItem, ok := rucksack.FindDuplicatedItem()

	if !ok {
		t.Fatal("Find duplicated item failed")
	}

	if duplicatedItem != expected {
		t.Fatalf("Expected duplicated item to be %s, but received %s", expected, duplicatedItem)
	}
}
