package day8

import (
	"fmt"
	"os"
)

func Run(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	tmr := NewTreeMapReader(f)
	grid := tmr.ReadMap()
	tm := NewTreeMap(grid)

	visible := make([][]bool, len(grid))
	for i, row := range grid {
		visible[i] = make([]bool, len(row))
		for j := range row {
			for _, direction := range []Direction{Left, Right, Up, Down} {
				if tm.IsTreeVisibleFromDirection(i, j, direction) {
					visible[i][j] = true
				}
			}
		}
	}

	numVisible := 0
	for _, row := range visible {
		for _, isVisible := range row {
			if isVisible {
				numVisible++
			}
		}
	}

	fmt.Printf("Visible: %d\n", numVisible)

	score := 0
	for i, row := range grid {
		for j := range row {
			localScore := 1
			for _, direction := range []Direction{Left, Right, Up, Down} {
				visibleFromPos := tm.GetNumberOfVisibleTrees(i, j, direction)
				localScore *= visibleFromPos
			}

			if localScore > score {
				score = localScore
			}
		}
	}

	fmt.Printf("Scenic score: %d\n", score)
}
