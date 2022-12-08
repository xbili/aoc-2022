package day2

// Key Shape wins Value Shape
var WinRules = map[Shape]Shape{
	Rock:     Scissors,
	Scissors: Paper,
	Paper:    Rock,
}

// Key Shape loses to Value Shape
var LoseRules = map[Shape]Shape{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}
