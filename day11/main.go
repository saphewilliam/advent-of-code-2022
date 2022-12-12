package day11

import (
	"aoc-2022/lib"
	"strconv"
	"strings"
)

type Monkey struct {
	inspects  int
	items     lib.Queue[int]
	plusVal   int
	multVal   int
	divisor   int
	destTrue  int
	destFalse int
}

func monkeyBusiness(monkeys []*Monkey, rounds, divisor int) int {
	modulus, top1, top2 := 1, 0, 0
	for _, m := range monkeys {
		modulus *= m.divisor
	}

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.inspects += monkey.items.Size()
			for !monkey.items.IsEmpty() {
				worry := monkey.items.Dequeue() + monkey.plusVal
				if monkey.multVal == -1 {
					worry *= worry
				} else {
					worry *= monkey.multVal
				}

				worry = worry % modulus / divisor
				if worry%monkey.divisor == 0 {
					monkeys[monkey.destTrue].items.Enqueue(worry)
				} else {
					monkeys[monkey.destFalse].items.Enqueue(worry)
				}
			}
		}
	}

	for _, monkey := range monkeys {
		if monkey.inspects > top1 {
			top2 = top1
			top1 = monkey.inspects
		} else if monkey.inspects > top2 {
			top2 = monkey.inspects
		}
	}
	return top1 * top2
}

func parseMonkeys(input []string) (monkeys []*Monkey) {
	for i := 1; i < len(input); i += 7 {
		m := Monkey{
			items:     lib.NewQueue(lib.Map(strings.Split(input[i][18:], ", "), lib.UnsafeParseInt)...),
			divisor:   lib.UnsafeParseInt(input[i+2][21:]),
			destTrue:  lib.UnsafeParseInt(input[i+3][29:]),
			destFalse: lib.UnsafeParseInt(input[i+4][30:]),
		}

		switch input[i+1][23] {
		case '+':
			m.multVal = 1
			m.plusVal = lib.UnsafeParseInt(input[i+1][25:])
		case '*':
			multVal, err := strconv.Atoi(input[i+1][25:])
			if err != nil {
				multVal = -1
			}
			m.multVal = multVal
		}
		monkeys = append(monkeys, &m)
	}
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	solution1.I = monkeyBusiness(parseMonkeys(input), 20, 3)
	solution2.I = monkeyBusiness(parseMonkeys(input), 10000, 1)
	return
}
