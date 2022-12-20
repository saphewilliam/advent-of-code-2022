package main

import (
	"aoc-2022/day1"
	"aoc-2022/day10"
	"aoc-2022/day11"
	"aoc-2022/day12"
	"aoc-2022/day13"
	"aoc-2022/day14"
	"aoc-2022/day15"
	"aoc-2022/day16"
	"aoc-2022/day17"
	"aoc-2022/day18"
	"aoc-2022/day19"
	"aoc-2022/day2"
	"aoc-2022/day20"
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
		"9":  {Process: day9.Process, TestPart1: "88", TestPart2: "36", Part1: "6503", Part2: "2724"},
		"10": {Process: day10.Process, TestPart1: "13140", TestPart2: "\n##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....", Part1: "15680", Part2: "\n####.####.###..####.#..#..##..#..#.###..\n...#.#....#..#.#....#..#.#..#.#..#.#..#.\n..#..###..###..###..####.#....#..#.#..#.\n.#...#....#..#.#....#..#.#.##.#..#.###..\n#....#....#..#.#....#..#.#..#.#..#.#....\n####.#....###..#....#..#..###..##..#...."},
		"11": {Process: day11.Process, TestPart1: "10605", TestPart2: "2713310158", Part1: "112815", Part2: "25738411485"},
		"12": {Process: day12.Process, TestPart1: "31", TestPart2: "29", Part1: "490", Part2: "488"},
		"13": {Process: day13.Process, TestPart1: "13", TestPart2: "140", Part1: "4809", Part2: "22600"},
		"14": {Process: day14.Process, TestPart1: "24", TestPart2: "93", Part1: "897", Part2: "26683"},
		"15": {Process: day15.Process, TestPart1: "26", TestPart2: "56000011", Part1: "5564017", Part2: "11558423398893"},
		"16": {Process: day16.Process, TestPart1: "1651", TestPart2: "1707", Part1: "1559", Part2: ""},
		"17": {Process: day17.Process, TestPart1: "3068", TestPart2: "1514285714288", Part1: "3130", Part2: "1556521739139"},
		"18": {Process: day18.Process, TestPart1: "64", TestPart2: "58", Part1: "4302", Part2: "2492"},
		"19": {Process: day19.Process, TestPart1: "33", TestPart2: "", Part1: "", Part2: ""},
		"20": {Process: day20.Process, TestPart1: "", TestPart2: "", Part1: "", Part2: ""},
	}
	day, isPresent = days[number]
	return
}
