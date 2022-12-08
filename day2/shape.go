package day2

// We represent move with an int enum, instead of strings since they can be
// encoded differently for different players.
type Shape int

const (
	InvalidShape Shape = iota
	Rock
	Paper
	Scissors
)

// This is used for Part 1, where we assumed that the second column is the
// move that we should play.
func DecodeSelf(encodedShape string) Shape {
	var shape Shape

	switch encodedShape {
	case "X":
		shape = Rock
	case "Y":
		shape = Paper
	case "Z":
		shape = Scissors
	default:
		panic("Invalid shape")
	}

	return shape
}

func DecodeOpponent(encodedShape string) Shape {
	var shape Shape

	switch encodedShape {
	case "A":
		shape = Rock
	case "B":
		shape = Paper
	case "C":
		shape = Scissors
	default:
		panic("Invalid shape")
	}

	return shape
}
