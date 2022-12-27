package day25

import (
	"aoc-2022/lib"
	"math"
)

func fromSnafu(in string) (out int) {
	getMultplier := func(r rune) int {
		switch r {
		case '=':
			return -2
		case '-':
			return -1
		default:
			return lib.ParseInt(string(r))
		}
	}
	for i, c := range in {
		out += int(math.Pow(5, float64((len(in)-i-1)))) * getMultplier(c)
	}
	return
}

func toSnafu(in int) (out string) {
	// Convert to raw 5-decimal 1747 => 023442
	q := lib.NewQueue[int]()
	curr := in
	for i := 0; ; i++ {
		n := (in / (int(math.Pow(5, float64(i))))) % 5
		q.Enqueue(n)
		curr /= 5
		if curr == 0 {
			break
		}
	}

	// Convert back to snafu 023442 => 1=-0-2
	carry := 0
	for !q.IsEmpty() {
		v := q.Dequeue()
		switch {
		case v+carry == 5:
			out = "0" + out
			carry = 1
		case v+carry == 4:
			out = "-" + out
			carry = 1
		case v+carry == 3:
			out = "=" + out
			carry = 1
		default:
			out = lib.ParseString(v+carry) + out
			carry = 0
		}
	}
	if carry != 0 {
		out = lib.ParseString(carry) + out
	}
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	for _, l := range input {
		solution1.I += fromSnafu(l)
	}
	solution1.S = toSnafu(solution1.I)
	solution1.I = 0
	return
}
