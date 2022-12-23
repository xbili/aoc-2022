package day9

import "testing"

func TestMoveSampleInput(t *testing.T) {
	bridge := NewBridge([2]int{0, 0})

	// Taken from sample input
	bridge.Move(Right, 4)
	bridge.Move(Up, 4)
	bridge.Move(Left, 3)
	bridge.Move(Down, 1)
	bridge.Move(Right, 4)
	bridge.Move(Down, 1)
	bridge.Move(Left, 5)
	bridge.Move(Right, 2)

	assertTotalVisited(t, bridge, 13)
}

func assertTotalVisited(t *testing.T, bridge *Bridge, expected int) {
	visited := bridge.GetTotalVisited()

	if visited != expected {
		t.Fatalf("Expected total visited to be %d, but got %d", expected, visited)
	}
}
