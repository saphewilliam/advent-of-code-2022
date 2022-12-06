package day5

import (
	"aoc-2022/lib"
	"fmt"
)

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	var start, move, from, to int
	for ; input[start] != ""; start++ {
	}

	stacks := make([]Stacks[rune], (len(input[start-1])+1)/4)
	for i := start - 2; i >= 0; i-- {
		for j := 0; j*4 < len(input[i]); j++ {
			r := rune(input[i][j*4+1])
			if r != 32 {
				stacks[j].s1.Push(r)
				stacks[j].s2.Push(r)
			}
		}
	}

	for i := start + 1; i < len(input); i++ {
		fmt.Sscanf(input[i], "move %d from %d to %d", &move, &from, &to)
		var stack Stack[rune]
		for j := 0; j < move; j++ {
			stacks[to-1].s1.Push(stacks[from-1].s1.Pop())
			stack.Push(stacks[from-1].s2.Pop())
		}
		for i := 0; i < move; i++ {
			stacks[to-1].s2.Push(stack.Pop())
		}
	}

	for i := 0; i < len(stacks); i++ {
		solution1.S += string(stacks[i].s1.Pop())
		solution2.S += string(stacks[i].s2.Pop())
	}
	return
}

type Stacks[T any] struct {
	s1 Stack[T]
	s2 Stack[T]
}

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() T {
	top := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return top
}
