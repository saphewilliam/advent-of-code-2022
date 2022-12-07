package day2

import "aoc-2022/lib"

func Process(input []string) (score1 lib.Solution, score2 lib.Solution) {
	for _, v := range input {
		a, b := int(v[0])-65, int(v[2])-84
		score1.I += b + 3*((b-a)%3-1)
		score2.I += (a+b-2)%3 + b*3 - 11
	}
	return
}
