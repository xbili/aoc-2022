package day3

import "testing"

func TestFindBadge(t *testing.T) {
	assertFindBadge(t, []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}, "r")
	assertFindBadge(t, []string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}, "Z")
}

func assertFindBadge(t *testing.T, input []string, expected string) {
	badge, ok := FindBadge(input)

	if !ok {
		t.Fatal("Invalid result returned by FindBadge")
	}

	if badge != expected {
		t.Fatalf("Expected badge to be %s, but received %s instead", expected, badge)
	}
}
