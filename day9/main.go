package day9

import (
	"aoc-2022/lib"
	"fmt"
)

type Point struct {
	x int
	y int
}

func getVisitedPoints() int {

}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	points := make(map[Point]bool)
	H, T := Point{0, 0}, Point{0, 0}

	var dir string
	var n int
	for _, v := range input {
		fmt.Sscanf(v, "%s %d", &dir, &n)
		for i := 0; i < n; i++ {
			switch dir {
			case "U":
				H.y++
				if T.y == H.y-2 {
					T.y = H.y - 1
					T.x = H.x
				}
			case "D":
				H.y--
				if T.y == H.y+2 {
					T.y = H.y + 1
					T.x = H.x
				}
			case "L":
				H.x--
				if T.x == H.x+2 {
					T.x = H.x + 1
					T.y = H.y
				}
			case "R":
				H.x++
				if T.x == H.x-2 {
					T.x = H.x - 1
					T.y = H.y
				}
			}
			points[T] = true
		}
	}

	solution1.I = len(points)
	return
}
