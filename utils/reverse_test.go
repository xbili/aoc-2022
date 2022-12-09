package utils

import "testing"

func TestReverseSliceOdd(t *testing.T) {
	assertReversed(t, []string{"A"}, []string{"A"})
	assertReversed(t, []string{"A", "B", "C"}, []string{"C", "B", "A"})
}

func TestReverseSliceEven(t *testing.T) {
	assertReversed(t, []string{"A", "B"}, []string{"B", "A"})
}

func TestReverseSliceEmpty(t *testing.T) {
	assertReversed(t, []string{}, []string{})
}

func assertReversed(t *testing.T, input, expected []string) {
	reversed := ReverseSlice(input)

	if len(reversed) != len(expected) {
		t.Fatalf("Expected length of reversed to be %d, but received %d", len(expected), len(reversed))
	}

	for i, s := range expected {
		if reversed[i] != s {
			t.Fatalf("Expected %s at position %d, but received %s", s, i, reversed[i])
		}
	}
}
