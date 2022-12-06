package day1

import (
	"aoc-2022/lib"
	"sort"
	"strconv"
)

func Process(input []string) (topCalories lib.Solution, top3Calories lib.Solution) {
	elves := []int{0}
	for _, v := range input {
		val, err := strconv.Atoi(v)
		if err != nil {
			elves = append(elves, 0)
		} else {
			elves[len(elves)-1] += val
		}
	}
	sort.Ints(elves)

	return lib.IScore(elves[len(elves)-1]), lib.IScore(elves[len(elves)-3] + elves[len(elves)-2] + elves[len(elves)-1])
}
