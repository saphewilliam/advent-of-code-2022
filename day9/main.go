package day9

import (
	"aoc-2022/lib"
	"fmt"
	"math"
)

type Knot struct {
	x int
	y int
}

func getVisitedPoints(input []string, ropeSize int) int {
	rope := make([]Knot, ropeSize)
	visitedPoints := make(map[Knot]bool)

	var dir string
	var n int
	for _, v := range input {
		fmt.Sscanf(v, "%s %d", &dir, &n)
		for i := 0; i < n; i++ {
			switch dir {
			case "U":
				rope[0].y--
			case "D":
				rope[0].y++
			case "R":
				rope[0].x--
			case "L":
				rope[0].x++
			}
			for j := 0; j < len(rope)-1; j++ {
				dx, dy := (rope[j+1].x - rope[j].x), (rope[j+1].y - rope[j].y)
				if math.Abs(float64(dx)) != 2 && math.Abs(float64(dy)) != 2 {
					break
				}
				rope[j+1].y = rope[j].y + dy/2
				rope[j+1].x = rope[j].x + dx/2
			}
			visitedPoints[rope[len(rope)-1]] = true
		}
	}
	return len(visitedPoints)
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	solution1.I = getVisitedPoints(input, 2)
	solution2.I = getVisitedPoints(input, 10)
	return
}
