package day5

import "testing"

func TestMoveReversed(t *testing.T) {
	stacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	cs := NewCrateStacks(stacks)
	cs.Move(1, 2, 1, false)
	cs.Move(3, 1, 3, false)
	cs.Move(2, 2, 1, false)
	cs.Move(1, 1, 2, false)

	top := cs.GetTopCrates()

	assertTopCrates(t, top, []string{"C", "M", "Z"})
}

func TestMoveOrdered(t *testing.T) {
	stacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	cs := NewCrateStacks(stacks)
	cs.Move(1, 2, 1, true)
	cs.Move(3, 1, 3, true)
	cs.Move(2, 2, 1, true)
	cs.Move(1, 1, 2, true)

	top := cs.GetTopCrates()

	assertTopCrates(t, top, []string{"M", "C", "D"})
}

func assertTopCrates(t *testing.T, actual, expected []string) {
	if len(actual) != len(expected) {
		t.Fatalf("Expected length of top to be %d, but received %d", len(expected), len(actual))
	}

	for i, crate := range expected {
		if actual[i] != crate {
			t.Fatalf("Expected stack %d to have crate %s at top, but received %s", i, crate, actual[i])
		}
	}
}
