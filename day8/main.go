package day8

import "aoc-2022/lib"

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	for i, line := range input {
		for j, v := range line {
			if i == 0 || j == 0 || i == len(line)-1 || j == len(line)-1 {
				continue
			}
			up, down, left, right := 0, 0, 0, 0

			// Up
			for k := i - 1; k >= 0; k-- {
				up++
				if rune(input[k][j]) >= v {
					break
				}
			}

			// Down
			for k := i + 1; k < len(line); k++ {
				down++
				if rune(input[k][j]) >= v {
					break
				}
			}

			// Left
			for k := j - 1; k >= 0; k-- {
				left++
				if rune(input[i][k]) >= v {
					break
				}
			}

			// Right
			for k := j + 1; k < len(line); k++ {
				right++
				if rune(input[i][k]) >= v {
					break
				}
			}

			score := up * down * left * right
			if score > solution2.I {
				solution2.I = score
			}
		}
	}

	for i, line := range input {
		for j, v := range line {
			if i == 0 || j == 0 || i == len(line)-1 || j == len(line)-1 {
				solution1.I++
				continue
			}

			isVisible := true
			// Up
			for k := i - 1; k >= 0; k-- {
				if rune(input[k][j]) >= v {
					isVisible = false
					break
				}
			}

			if isVisible {
				solution1.I++
				continue
			}

			isVisible = true

			// Down
			for k := i + 1; k < len(line); k++ {
				if rune(input[k][j]) >= v {
					isVisible = false
					break
				}
			}

			if isVisible {
				solution1.I++
				continue
			}

			isVisible = true

			// Left
			for k := j - 1; k >= 0; k-- {
				if rune(input[i][k]) >= v {
					isVisible = false
					break
				}
			}

			if isVisible {
				solution1.I++
				continue
			}

			isVisible = true

			// Right
			for k := j + 1; k < len(line); k++ {
				if rune(input[i][k]) >= v {
					isVisible = false
					break
				}
			}

			if isVisible {
				solution1.I++
				continue
			}

			isVisible = true
		}
	}
	return
}
