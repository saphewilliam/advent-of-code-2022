package day8

import "aoc-2022/lib"

func scanDirection(input []string, dir, axis, row, col, size int) (isHidden bool, score int) {
	for k := (((axis+1)%2)*row + axis*col) + dir; (dir == -1 && k >= 0) || (dir == 1 && k < size); k += dir {
		score++
		if (axis == 0 && input[k][col] >= input[row][col]) || (axis == 1 && input[row][k] >= input[row][col]) {
			isHidden = true
			break
		}
	}
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	for i, line := range input {
		for j := range line {
			if i == 0 || j == 0 || i == len(line)-1 || j == len(line)-1 {
				solution1.I++
				continue
			}

			upHidden, upScore := scanDirection(input, -1, 0, i, j, len(line))
			downHidden, downScore := scanDirection(input, 1, 0, i, j, len(line))
			leftHidden, leftScore := scanDirection(input, -1, 1, i, j, len(line))
			rightHidden, rightScore := scanDirection(input, 1, 1, i, j, len(line))

			score := upScore * downScore * leftScore * rightScore
			if score > solution2.I {
				solution2.I = score
			}
			if !(upHidden && downHidden && leftHidden && rightHidden) {
				solution1.I++
			}
		}
	}
	return
}
