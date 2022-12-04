package day04

import (
	"regexp"
	"strconv"
)

const Path = "./day04/input.txt"

func Calculate(input []string) (overlaps int, fullOverlaps int) {
	for _, v := range input {
		sections := regexp.MustCompile(`-|,`).Split(v, -1)
		min1, _ := strconv.Atoi(sections[0])
		max1, _ := strconv.Atoi(sections[1])
		min2, _ := strconv.Atoi(sections[2])
		max2, _ := strconv.Atoi(sections[3])
		if min1 <= max2 && min1 >= min2 || min2 <= max1 && min2 >= min1 {
			overlaps++
		}
		if min1 <= min2 && max1 >= max2 || min2 <= min1 && max2 >= max1 {
			fullOverlaps++
		}
	}
	return
}
