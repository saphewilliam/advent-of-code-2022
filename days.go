package main

import (
	"aoc-2022/day1"
	"aoc-2022/day10"
	"aoc-2022/day2"
	"aoc-2022/day3"
	"aoc-2022/day4"
	"aoc-2022/day5"
	"aoc-2022/day6"
	"aoc-2022/day7"
	"aoc-2022/day8"
	"aoc-2022/day9"
	"aoc-2022/lib"
)

type Day struct {
	Process                            func(input []string) (lib.Solution, lib.Solution)
	Part1, Part2, TestPart1, TestPart2 string
}

func getDay(number string) (day Day, isPresent bool) {
	days := map[string]Day{
		"1":  {Process: day1.Process, TestPart1: "24000", TestPart2: "45000", Part1: "72511", Part2: "212117"},
		"2":  {Process: day2.Process, TestPart1: "15", TestPart2: "12", Part1: "12740", Part2: "11980"},
		"3":  {Process: day3.Process, TestPart1: "157", TestPart2: "70", Part1: "7917", Part2: "2585"},
		"4":  {Process: day4.Process, TestPart1: "4", TestPart2: "2", Part1: "888", Part2: "471"},
		"5":  {Process: day5.Process, TestPart1: "CMZ", TestPart2: "MCD", Part1: "QNNTGTPFN", Part2: "GGNPJBTTR"},
		"6":  {Process: day6.Process, TestPart1: "10", TestPart2: "29", Part1: "1833", Part2: "3425"},
		"7":  {Process: day7.Process, TestPart1: "95437", TestPart2: "24933642", Part1: "1334506", Part2: "7421137"},
		"8":  {Process: day8.Process, TestPart1: "21", TestPart2: "8", Part1: "1825", Part2: "235200"},
		"9":  {Process: day9.Process, TestPart1: "", TestPart2: "", Part1: "", Part2: ""},
		"10": {Process: day10.Process, TestPart1: "", TestPart2: "", Part1: "", Part2: ""},
	}
	day, isPresent = days[number]
	return
}
