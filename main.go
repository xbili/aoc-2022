package main

import "xbili.com/aoc/day1"
import "xbili.com/aoc/day2"
import "xbili.com/aoc/day3"
import "xbili.com/aoc/day4"
import "xbili.com/aoc/day5"
import "xbili.com/aoc/day6"
import "xbili.com/aoc/day7"
import "xbili.com/aoc/day8"

func main() {
	day1.Run("inputs/day1.txt")
	day2.RunPart1("inputs/day2.txt")
	day2.RunPart2("inputs/day2.txt")
	day3.RunPart1("inputs/day3.txt")
	day3.RunPart2("inputs/day3.txt")
	day4.Run("inputs/day4.txt")
	day5.RunPart1("inputs/day5.txt")
	day5.RunPart2("inputs/day5.txt")
	day6.Run("inputs/day6.txt", 4)
	day6.Run("inputs/day6.txt", 14)
	day7.Run("inputs/day7.txt")
	day8.Run("inputs/day8.txt")
}
