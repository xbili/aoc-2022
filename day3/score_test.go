package day3

import "testing"

func TestCalculateScore(t *testing.T) {
	assertScore(t, "a", 1)
	assertScore(t, "z", 26)
	assertScore(t, "A", 27)
	assertScore(t, "Z", 52)
}

func assertScore(t *testing.T, input string, expected int) {
	score := CalculateScore(input)

	if score != expected {
		t.Fatalf("Expected score to be %d, but received %d", expected, score)
	}
}
