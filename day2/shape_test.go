package day2

import "testing"

func TestDecodeSelf(t *testing.T) {
	f := DecodeSelf

	assertDecode(t, "X", Rock, f)
	assertDecode(t, "Y", Paper, f)
	assertDecode(t, "Z", Scissors, f)
}

func TestDecodeOpponent(t *testing.T) {
	f := DecodeOpponent

	assertDecode(t, "A", Rock, f)
	assertDecode(t, "B", Paper, f)
	assertDecode(t, "C", Scissors, f)
}

func assertDecode(t *testing.T, input string, expected Shape, decodeFunc func(string) Shape) {
	shape := decodeFunc(input)

	if shape != expected {
		t.Fatalf("Expected shape %d, but was %d", expected, shape)
	}
}
