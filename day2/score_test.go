package day2

import "testing"

func TestCalculateScoreDraws(t *testing.T) {
	assertCalculateScore(t, Rock, Rock, 4)
	assertCalculateScore(t, Paper, Paper, 5)
	assertCalculateScore(t, Scissors, Scissors, 6)
}

func TestCalculateScoreWins(t *testing.T) {
	assertCalculateScore(t, Rock, Scissors, 7)
	assertCalculateScore(t, Scissors, Paper, 9)
	assertCalculateScore(t, Paper, Rock, 8)
}

func TestCalculateScoreLose(t *testing.T) {
	assertCalculateScore(t, Scissors, Rock, 3)
	assertCalculateScore(t, Rock, Paper, 1)
	assertCalculateScore(t, Paper, Scissors, 2)
}

func assertCalculateScore(t *testing.T, selfShape Shape, opponentShape Shape, expected int) {
	score := CalculateScore(selfShape, opponentShape)

	if score != expected {
		t.Fatalf("Expected %d, but got %d", expected, score)
	}
}
