package day4

import (
	"aoc-2022/lib"
	"fmt"
)

func Process(input []string) (overlaps lib.Solution, fullOverlaps lib.Solution) {
	for _, v := range input {
		var min1, max1, min2, max2 int
		fmt.Sscanf(v, "%d-%d,%d-%d", &min1, &max1, &min2, &max2)
		if min1 <= max2 && min2 <= max1 {
			overlaps.I++
		}
		if min1 <= min2 && max1 >= max2 || min2 <= min1 && max2 >= max1 {
			fullOverlaps.I++
		}
	}
	return
}
