package day3

const LowercaseOffset int = 96
const UppercaseOffset int = 64

// Returns the priority score for a single item
func CalculateScore(input string) int {
	// Turns out this is the way to perform this conversion in Go, not 100%
	// certain that this is really correct.
	charCode := int(input[0])

	// Case: a-z
	if charCode > LowercaseOffset && charCode <= LowercaseOffset+26 {
		return int(input[0]) - LowercaseOffset
	}

	// Case: A-Z
	if charCode > UppercaseOffset && charCode <= UppercaseOffset+26 {
		return int(input[0]) - UppercaseOffset + 26
	}

	panic("Invalid character provided")
}
