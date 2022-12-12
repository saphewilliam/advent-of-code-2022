package day10

import (
	"aoc-2022/lib"
	"fmt"
	"math"
)

type Command struct {
	isAddx bool
	value  int
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	q := lib.NewQueue[Command]()
	for _, v := range input {
		if v == "noop" {
			q.Enqueue(Command{})
		} else {
			var value int
			fmt.Sscanf(v, "addx %d", &value)
			q.Enqueue(Command{isAddx: true, value: value})
		}
	}

	X, isProcessing, value := 1, false, 0
	for cycle := 0; !q.IsEmpty(); cycle++ {
		if (cycle-19)%40 == 0 {
			solution1.I += (cycle + 1) * X
		}
		if cycle%40 == 0 {
			solution2.S += "\n"
		}
		if math.Abs(float64((cycle%40)-X)) < 2 {
			solution2.S += "#"
		} else {
			solution2.S += "."
		}

		if !isProcessing {
			command := q.Dequeue()
			isProcessing = command.isAddx
			value = command.value
		} else {
			X += value
			isProcessing = false
		}
	}

	return
}
