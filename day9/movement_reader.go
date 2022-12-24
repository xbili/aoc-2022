package day9

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MovementReader struct {
	scanner *bufio.Scanner
}

func NewMovementReader(f *os.File) *MovementReader {
	return &MovementReader{
		scanner: bufio.NewScanner(f),
	}
}

func (mr *MovementReader) Next() (direction Direction, distance int, ok bool) {
	for {
		mr.scanner.Scan()
		line := mr.scanner.Text()

		if line == "" {
			break
		}

		tokens := strings.Split(line, " ")
		direction := parseDirection(tokens[0])
		distance := parseDistance(tokens[1])

		return direction, distance, true
	}

	return Stationary, 0, false
}

func parseDirection(token string) Direction {
	switch token {
	case "L":
		return Left
	case "R":
		return Right
	case "U":
		return Up
	case "D":
		return Down
	}

	return Stationary
}

func parseDistance(token string) int {
	val, err := strconv.Atoi(token)
	if err != nil {
		panic(err)
	}

	return val
}
