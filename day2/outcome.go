package day2

type Outcome int

const (
	InvalidOutcome Outcome = iota
	Lose
	Draw
	Win
)

func DecodeOutcome(encodedOutcome string) Outcome {
	switch encodedOutcome {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		return InvalidOutcome
	}
}

// Determine what shape to play in order to achieve the desired outcome given
// the opponent's known shape.
func FindShapeForDesiredOutcome(opponentShape Shape, desiredOutcome Outcome) Shape {
	if desiredOutcome == Draw {
		// To draw we just need to play the same shape as our opponent
		return opponentShape
	}

	if desiredOutcome == Win {
		// We want opponent to lose
		return LoseRules[opponentShape]
	}

	if desiredOutcome == Lose {
		// We want opponent to win
		return WinRules[opponentShape]
	}

	panic("Invalid outcome")
}
