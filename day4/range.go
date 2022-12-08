package day4

import (
	"strconv"
	"strings"
)

type SectionRange struct {
	start int
	end   int
}

func NewSectionRange(input string) *SectionRange {
	sectionRange := new(SectionRange)

	split := strings.Split(input, "-")

	start, err := strconv.Atoi(split[0])
	if err != nil {
		panic("Invalid start")
	}
	sectionRange.start = start

	end, err := strconv.Atoi(split[1])
	if err != nil {
		panic("Invalid end")
	}
	sectionRange.end = end

	return sectionRange
}

func (sr *SectionRange) FullyContains(other *SectionRange) bool {
	return sr.start <= other.start && sr.end >= other.end
}

func (sr *SectionRange) Contains(other *SectionRange) bool {
	return sr.start <= other.end && sr.end >= other.start
}
