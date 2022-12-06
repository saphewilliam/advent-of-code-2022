package day6

import "aoc-2022/lib"

func findMarker(input string, length int) int {
	for i := length; i <= len(input); i++ {
		m1 := map[byte]bool{}
		for j := i - 1; j >= i-length; j-- {
			m1[input[j]] = true
		}
		if len(m1) == length {
			return i
		}
	}
	return -1
}

func Process(input []string) (marker1 lib.Solution, marker2 lib.Solution) {
	marker1.I = findMarker(input[0], 4)
	marker2.I = findMarker(input[0], 14)
	return
}
