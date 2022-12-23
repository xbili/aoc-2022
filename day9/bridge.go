package day9

import (
	"fmt"
)

type Bridge struct {
	seen map[string]bool // visited set of positions in the bridge
	head [2]int          // current position of head
	tail [2]int          // current position of tail
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

var Diff = map[Direction][2]int{
	Left:  {0, -1},
	Right: {0, 1},
	Up:    {-1, 0},
	Down:  {1, 0},
}

func NewBridge(start [2]int) *Bridge {
	b := &Bridge{
		seen: map[string]bool{},
		head: start,
		tail: start,
	}

	// Tail needs to see start position
	b.seen[formatPositionKey(start)] = true

	return b
}

// Moves the head in a specific direction for a fixed number of units
func (bridge *Bridge) Move(direction Direction, distance int) {

	fmt.Println(direction, distance)

	visited := map[string]bool{}

	for i := 0; i < distance; i++ {
		// Move head
		bridge.head[0] += Diff[direction][0]
		bridge.head[1] += Diff[direction][1]

		// Follow tail
		drow, dcol := bridge.head[0]-bridge.tail[0], bridge.head[1]-bridge.tail[1]

		// Nothing needs to change, tail is still within range of head
		if abs(drow) <= 1 && abs(dcol) <= 1 {
			continue
		}

		bridge.tail[0] += Diff[direction][0]
		bridge.tail[1] += Diff[direction][1]

		// Diagonal, need to catch up with one more move
		if abs(drow)+abs(dcol) > 2 {
			var catchup Direction
			if abs(drow) > abs(dcol) {
				if dcol > 0 {
					catchup = Right
				} else {
					catchup = Left
				}
			} else {
				if drow > 0 {
					catchup = Down
				} else {
					catchup = Up
				}
			}

			bridge.tail[0] += Diff[catchup][0]
			bridge.tail[1] += Diff[catchup][1]
		}

		visited[formatPositionKey(bridge.tail)] = true
	}

	// Add visited into seen map
	for k := range visited {
		bridge.seen[k] = true
	}
}

func (bridge *Bridge) GetTotalVisited() int {
	return len(bridge.seen)
}

func formatPositionKey(pos [2]int) string {
	return fmt.Sprintf("%d-%d", pos[0], pos[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x

}
