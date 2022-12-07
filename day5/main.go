package day5

import (
	"aoc-2022/lib"
	"fmt"
)

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	var start, move, from, to int
	for ; input[start] != ""; start++ {
	}

	stacks := make([]Stacks, (len(input[start-1])+1)/4)
	for i := start - 2; i >= 0; i-- {
		for j := 0; j*4 < len(input[i]); j++ {
			stacks[j].Init(rune(input[i][j*4+1]))
		}
	}

	for i := start + 1; i < len(input); i++ {
		fmt.Sscanf(input[i], "move %d from %d to %d", &move, &from, &to)
		stacks[to-1].PushMany(stacks[from-1].PopMany(move))
		for j := 0; j < move; j++ {
			stacks[to-1].s1.Push(stacks[from-1].s1.Pop())
		}
	}

	for i := 0; i < len(stacks); i++ {
		solution1.S += string(stacks[i].s1.Pop())
		solution2.S += string(stacks[i].s2.Pop())
	}
	return
}

type Stacks struct {
	s1 lib.Stack[rune]
	s2 lib.Stack[rune]
}

func (s *Stacks) Init(r rune) {
	if r != 32 {
		s.s1.Push(r)
		s.s2.Push(r)
	}
}

func (s *Stacks) PushMany(values []rune) {
	s.s2.Values = append(s.s2.Values, values...)
}

func (s *Stacks) PopMany(n int) []rune {
	top := s.s2.Values[s.s2.Size()-n : s.s2.Size()]
	s.s2.Values = s.s2.Values[:s.s2.Size()-n]
	return top
}
