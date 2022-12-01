package day1

import (
	"math"
	"testing"
)

func TestEmpty(t *testing.T) {
	input, k := []int{}, 3
	assert(t, input, k, 0)
}

func TestBasic(t *testing.T) {
	input, k := []int{1, 2, 3, 4, 5, 6}, 3
	assert(t, input, k, 15)
}

func TestBasicUnsorted(t *testing.T) {
	input, k := []int{3, 1, 2, 4, 6, 5}, 3
	assert(t, input, k, 15)
}

func TestFewerThanK(t *testing.T) {
	input, k := []int{1, 2}, 3
	assert(t, input, k, 3)
}

// Helper assert method
func assert(t *testing.T, calories []int, k int, expected int) {
	ecr := NewElvesCaloriesReaderStub(calories)

	res := FindRichestElves(k, ecr)

	if res != expected {
		t.Fatalf("Richest elves expected to be %d, but was %d", expected, res)
	}
}

// Stub implementation for an ElvesCaloriesReader
// TODO: Pretty sure there is a better way to fake interfaces, just have find
// out how.
type ElvesCaloriesReaderStub struct {
	calories []int // The list of calories to return
	index    int   // The next index to read from
}

func NewElvesCaloriesReaderStub(values []int) *ElvesCaloriesReaderStub {
	ecr := new(ElvesCaloriesReaderStub)
	ecr.index = 0

	// Copy values of fake calories
	ecr.calories = make([]int, len(values))
	copy(ecr.calories, values)

	return ecr
}

func (ecr *ElvesCaloriesReaderStub) Next() (value int, ok bool) {
	if ecr.index == len(ecr.calories) {
		return math.MinInt, false
	}

	val := ecr.calories[ecr.index]
	ecr.index++

	return val, true
}
