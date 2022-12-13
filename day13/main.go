package day13

import (
	"aoc-2022/lib"
	"encoding/json"
	"math"
	"sort"
)

func comparePackets(left, right []any) int {
	max := int(math.Max(float64(len(left)), float64(len(right))))
	for i := 0; i < max; i++ {
		if i == len(left) {
			return 1
		}
		if i == len(right) {
			return -1
		}

		l, leftIsNumber := left[i].(float64)
		r, rightIsNumber := right[i].(float64)
		if leftIsNumber && rightIsNumber {
			if l < r {
				return 1
			} else if r < l {
				return -1
			}
			continue
		}
		if leftIsNumber {
			result := comparePackets([]any{l}, right[i].([]any))
			if result == 0 {
				continue
			}
			return result
		}
		if rightIsNumber {
			result := comparePackets(left[i].([]any), []any{r})
			if result == 0 {
				continue
			}
			return result
		}

		result := comparePackets(left[i].([]any), right[i].([]any))
		if result == 0 {
			continue
		}
		return result
	}
	return 0
}

func parseArray(value string) (arr []any) {
	json.Unmarshal([]byte(value), &arr)
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	p1, p2 := parseArray("[[2]]"), parseArray("[[6]]")
	packets := [][]any{p1, p2}

	for i := 0; i < len(input); i += 3 {
		left, right := parseArray(input[i]), parseArray(input[i+1])
		if comparePackets(left, right) == 1 {
			solution1.I += i/3 + 1
		}
		packets = append(packets, left, right)
	}

	sort.Slice(packets, func(l, r int) bool {
		return comparePackets(packets[l], packets[r]) == 1
	})

	solution2.I = 1
	for i, p := range packets {
		if comparePackets(p, p1) == 0 || comparePackets(p, p2) == 0 {
			solution2.I *= i + 1
		}
	}

	return
}
