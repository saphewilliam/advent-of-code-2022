package day04

import "fmt"

const Path = "./day04/input.txt"

func Calculate(input []string) (overlaps int, fullOverlaps int) {
	for _, v := range input {
		var min1, max1, min2, max2 int
		fmt.Sscanf(v, "%d-%d,%d-%d", &min1, &max1, &min2, &max2)
		if min1 <= max2 && min2 <= max1 {
			overlaps++
		}
		if min1 <= min2 && max1 >= max2 || min2 <= min1 && max2 >= max1 {
			fullOverlaps++
		}
	}
	return
}
