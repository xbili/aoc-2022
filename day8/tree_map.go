package day8

// Enum to keep track of directions
type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

// Map of trees' height
type TreeMap struct {
	height [][]int // height of trees in each position
}

func NewTreeMap(grid [][]int) *TreeMap {
	rows, cols := getDimensions(grid)

	tm := new(TreeMap)
	tm.height = make([][]int, rows)
	for i := 0; i < rows; i++ {
		tm.height[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			tm.height[i][j] = grid[i][j]
		}
	}

	return tm
}

// Returns true if the tree at position is visible from a certain direction
func (tm *TreeMap) IsTreeVisibleFromDirection(row, col int, direction Direction) bool {
	isVisible := true

	height := tm.height[row][col]
	iterate(tm.height, row, col, direction, func(i, j int) bool {
		if i == row && j == col {
			return false
		}

		otherHeight := tm.height[i][j]
		if height <= otherHeight {
			isVisible = false
			return true
		}

		return false
	})

	return isVisible
}

// Returns the number of visible trees from a single position
func (tm *TreeMap) GetNumberOfVisibleTrees(row, col int, direction Direction) int {
	numVisible := 0
	height := tm.height[row][col]
	iterate(tm.height, row, col, direction, func(i, j int) bool {
		if i == row && j == col {
			return false
		}

		otherHeight := tm.height[i][j]
		numVisible++

		if height <= otherHeight {
			return true
		}

		return false
	})

	return numVisible
}

func getDimensions(grid [][]int) (rows, cols int) {
	return len(grid), len(grid[0])
}

func createEmptyGrid(rows, cols int) [][][]int {
	empty := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		empty[i] = make([][]int, cols)
		for j := 0; j < cols; j++ {
			empty[i][j] = make([]int, 2)
		}
	}

	return empty
}

// Perform f, and start iterating from (row,col) towards a certain direction
func iterate(grid [][]int, row, col int, direction Direction, f func(i, j int) bool) {
	delta := map[Direction][2]int{
		Left:  {0, -1},
		Right: {0, 1},
		Up:    {-1, 0},
		Down:  {1, 0},
	}

	terminate := false
	i, j := row, col
	for isValid(grid, i, j) && !terminate {
		terminate = f(i, j)

		i += delta[direction][0]
		j += delta[direction][1]
	}
}

func isValid(grid [][]int, row, col int) bool {
	rows, cols := getDimensions(grid)
	if row < 0 || col < 0 {
		return false
	}

	if row >= rows || col >= cols {
		return false
	}

	return true
}
