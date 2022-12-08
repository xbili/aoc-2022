package day2

import "testing"

func TestOutcomeDecode(t *testing.T) {
	assertDecodeOutcome(t, "W", InvalidOutcome)
	assertDecodeOutcome(t, "X", Lose)
	assertDecodeOutcome(t, "Y", Draw)
	assertDecodeOutcome(t, "Z", Win)
}

func TestFindShapeForDesiredOutcomeDraws(t *testing.T) {
	assertShapeForDesiredOutcome(t, Rock, Draw, Rock)
	assertShapeForDesiredOutcome(t, Paper, Draw, Paper)
	assertShapeForDesiredOutcome(t, Scissors, Draw, Scissors)
}

func TestFindShapeForDesiredOutcomeWins(t *testing.T) {
	assertShapeForDesiredOutcome(t, Rock, Win, Paper)
	assertShapeForDesiredOutcome(t, Paper, Win, Scissors)
	assertShapeForDesiredOutcome(t, Scissors, Win, Rock)
}

func TestFindShapeForDesiredOutcomeLose(t *testing.T) {
	assertShapeForDesiredOutcome(t, Rock, Lose, Scissors)
	assertShapeForDesiredOutcome(t, Paper, Lose, Rock)
	assertShapeForDesiredOutcome(t, Scissors, Lose, Paper)
}

func assertDecodeOutcome(t *testing.T, input string, expected Outcome) {
	outcome := DecodeOutcome(input)

	if outcome != expected {
		t.Fatalf("Expected outcome %d, but was %d", expected, outcome)
	}
}

func assertShapeForDesiredOutcome(t *testing.T, opponentShape Shape, desiredOutcome Outcome, expected Shape) {
	shape := FindShapeForDesiredOutcome(opponentShape, desiredOutcome)

	if shape != expected {
		t.Fatalf("Expected shape for desired outcome %d, but was %d", expected, shape)
	}
}
