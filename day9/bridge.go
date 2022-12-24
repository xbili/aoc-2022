package day9

import (
	"fmt"
)

type Bridge struct {
	seen  map[string]bool // visited set of positions by tail in the bridge
	knots []*Knot         // all knots in the bridge
}

type Knot struct {
	position [2]int // current position of the knot
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
	Stationary
)

var Diff = map[Direction][2]int{
	Stationary: {0, 0},
	Left:       {0, -1},
	Right:      {0, 1},
	Up:         {-1, 0},
	Down:       {1, 0},
}

func NewBridge(start [2]int, numKnots int) *Bridge {
	knots := []*Knot{}
	for i := 0; i < numKnots; i++ {
		// All knots start at the same point
		knots = append(knots, &Knot{
			position: start,
		})
	}

	b := &Bridge{
		knots: knots,
		seen:  map[string]bool{},
	}

	// Start position will always be visited
	b.seen[formatPositionKey(start)] = true

	return b
}

// Moves the head in a specific direction for a fixed number of units
func (bridge *Bridge) Move(direction Direction, distance int) {
	for i := 0; i < distance; i++ {
		// Move head
		bridge.knots[0].position[0] += Diff[direction][0]
		bridge.knots[0].position[1] += Diff[direction][1]

		// Follow tails
		for j := 1; j < len(bridge.knots); j++ {
			leader, follower := bridge.knots[j-1].position, bridge.knots[j].position
			drow, dcol := leader[0]-follower[0], leader[1]-follower[1]

			// No need to move if the knots are already within range
			if abs(drow) <= 1 && abs(dcol) <= 1 {
				continue
			}

			if drow > 0 {
				bridge.knots[j].position[0] += 1
			} else if drow < 0 {
				bridge.knots[j].position[0] -= 1
			}

			if dcol > 0 {
				bridge.knots[j].position[1] += 1
			} else if dcol < 0 {
				bridge.knots[j].position[1] -= 1
			}

			if j == len(bridge.knots)-1 {
				bridge.seen[formatPositionKey(bridge.knots[j].position)] = true
			}
		}
	}
}

func (bridge *Bridge) GetTotalVisited() int {
	return len(bridge.seen)
}

func (bridge *Bridge) PrintSeen(rows, cols int) {
	for i := -rows / 2; i < rows/2; i++ {
		for j := -cols / 2; j < cols/2; j++ {
			if val, ok := bridge.seen[formatPositionKey([2]int{i, j})]; ok && val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n\n")
}

func (bridge *Bridge) PrintState(rows, cols int) {
	for i := -rows / 2; i < rows/2; i++ {
		for j := -cols / 2; j < cols/2; j++ {
			found := false
			for k, tail := range bridge.knots {
				if i == tail.position[0] && j == tail.position[1] {
					fmt.Print(k)
					found = true
					break
				}
			}

			if !found {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n\n")
}

func formatPositionKey(pos [2]int) string {
	return fmt.Sprintf("%d_%d", pos[0], pos[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x

}
