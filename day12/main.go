package day12

import "aoc-2022/lib"

func getDistance(points []rune, width, start, goal int) (dist int) {
	q, parents := lib.NewQueue(start), make(map[int]int)
	parents[start] = -1
	for !q.IsEmpty() {
		p := q.Dequeue()
		if p == goal {
			break
		}

		evaluateDirection := func(pDir int, cond bool) {
			if _, explored := parents[pDir]; cond && !explored && points[pDir] <= points[p]+1 {
				parents[pDir] = p
				q.Enqueue(pDir)
			}
		}
		evaluateDirection(p-width, p >= width)            // up
		evaluateDirection(p+width, p+width < len(points)) // down
		evaluateDirection(p-1, p%width != 0)              // left
		evaluateDirection(p+1, p%width != width-1)        // right
	}
	for _, found := parents[goal]; found && parents[goal] != -1; dist++ {
		goal = parents[goal]
	}
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	points, starts, goal := []rune{}, []int{}, -1
	for _, row := range input {
		for _, v := range row {
			switch v {
			case 'a':
				starts = append(starts, len(points))
			case 'S':
				v = 'a'
				starts = append([]int{len(points)}, starts...)
			case 'E':
				v = 'z'
				goal = len(points)
			}
			points = append(points, v)
		}
	}

	for i, s := range starts {
		dist := getDistance(points, len(input[0]), s, goal)
		if i == 0 {
			solution1.I = dist
			solution2.I = dist
		}
		if dist > 0 && dist < solution2.I {
			solution2.I = dist
		}
	}
	return
}
