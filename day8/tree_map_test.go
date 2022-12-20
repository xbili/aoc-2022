package day8

import (
	"testing"
)

func TestSingleRowAlwaysVisible(t *testing.T) {
	heights := [][]int{
		{1, 2, 3, 4},
	}

	tm := NewTreeMap(heights)

	for j := 0; j < 4; j++ {
		assertTreeVisible(t, tm, 0, j, Up, true)
		assertTreeVisible(t, tm, 0, j, Down, true)
	}
}

func TestSingleColAlwaysVisible(t *testing.T) {
	heights := [][]int{
		{1},
		{2},
		{3},
		{4},
	}

	tm := NewTreeMap(heights)

	for i := 0; i < 4; i++ {
		assertTreeVisible(t, tm, i, 0, Left, true)
		assertTreeVisible(t, tm, i, 0, Right, true)
	}
}

func TestIsTreeVisibleFromDirectionSampleMap(t *testing.T) {
	heights := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	tm := NewTreeMap(heights)

	// Top right 5
	assertTreeVisible(t, tm, 1, 1, Left, true)
	assertTreeVisible(t, tm, 1, 1, Up, true)

	// Middle 5
	assertTreeVisible(t, tm, 1, 2, Right, true)
	assertTreeVisible(t, tm, 1, 2, Up, true)

	// Top right 1
	assertTreeVisible(t, tm, 1, 3, Left, false)
	assertTreeVisible(t, tm, 1, 3, Right, false)
	assertTreeVisible(t, tm, 1, 3, Up, false)
	assertTreeVisible(t, tm, 1, 3, Down, false)

	// Left middle 5
	assertTreeVisible(t, tm, 2, 1, Right, true)

	// Center 3
	assertTreeVisible(t, tm, 2, 2, Left, false)
	assertTreeVisible(t, tm, 2, 2, Right, false)
	assertTreeVisible(t, tm, 2, 2, Up, false)
	assertTreeVisible(t, tm, 2, 2, Down, false)

	// Right middle 3
	assertTreeVisible(t, tm, 2, 3, Left, false)
	assertTreeVisible(t, tm, 2, 3, Right, true)
	assertTreeVisible(t, tm, 2, 3, Up, false)
	assertTreeVisible(t, tm, 2, 3, Down, false)

	// Bottom row middle 5
	assertTreeVisible(t, tm, 3, 2, Left, true)
	assertTreeVisible(t, tm, 3, 2, Right, false)
	assertTreeVisible(t, tm, 3, 2, Up, false)
	assertTreeVisible(t, tm, 3, 2, Down, true)
}

func TestGetNumberOfVisibleTreesSampleMap(t *testing.T) {
	heights := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	tm := NewTreeMap(heights)

	assertNumberTreesVisible(t, tm, 1, 2, Up, 1)
	assertNumberTreesVisible(t, tm, 1, 2, Left, 1)
	assertNumberTreesVisible(t, tm, 1, 2, Right, 2)
	assertNumberTreesVisible(t, tm, 1, 2, Down, 2)

	assertNumberTreesVisible(t, tm, 3, 2, Up, 2)
	assertNumberTreesVisible(t, tm, 3, 2, Left, 2)
	assertNumberTreesVisible(t, tm, 3, 2, Down, 1)
	assertNumberTreesVisible(t, tm, 3, 2, Right, 2)
}

func assertTreeVisible(t *testing.T, tm *TreeMap, row, col int, direction Direction, visible bool) {
	isVisible := tm.IsTreeVisibleFromDirection(row, col, direction)

	if isVisible != visible {
		t.Fatalf("Expect tree in %d, %d to be %t from %d", row, col, visible, direction)
	}
}

func assertNumberTreesVisible(t *testing.T, tm *TreeMap, row, col int, direction Direction, expected int) {
	numVisible := tm.GetNumberOfVisibleTrees(row, col, direction)

	if numVisible != expected {
		t.Fatalf("Expect number of visible trees from %d, %d facing %d to be %d, but got %d", row, col, direction, expected, numVisible)
	}
}
