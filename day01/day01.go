package day01

import (
	"sort"
	"strconv"
)

const Path = "./day01/input.txt"

func Calculate(input []string) (topCalories int, top3Calories int) {
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
	return elves[len(elves)-1], elves[len(elves)-3] + elves[len(elves)-2] + elves[len(elves)-1]
}
