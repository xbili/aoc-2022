package day2

var ShapeScore = map[Shape]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var OutcomeBonus = map[Outcome]int{
	Win:  6,
	Draw: 3,
	Lose: 0,
}

func CalculateScore(selfShape Shape, opponentShape Shape) int {
	// Determine how much score to get for your own shape
	shapeScore := ShapeScore[selfShape]

	// Outcome: draw
	if selfShape == opponentShape {
		return shapeScore + OutcomeBonus[Draw]
	}

	selfWon := WinRules[selfShape] == opponentShape

	if selfWon {
		// Outcome: self won
		return shapeScore + OutcomeBonus[Win]
	} else {
		// Outcome: self lost
		return shapeScore + OutcomeBonus[Lose]
	}
}
