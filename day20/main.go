package day20

import (
	"aoc-2022/lib"
)

type Item struct {
	value, originalIndex int
}

func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

func mix(numbers []int, repeat int) (items []Item) {
	items = make([]Item, len(numbers))
	for i, n := range numbers {
		items[i] = Item{value: n, originalIndex: i}
	}

	for i := 0; i < repeat; i++ {
		for j := 0; j < len(items); j++ {
			var itemIndex int
			for k := 0; k < len(items); k++ {
				if items[k].originalIndex == j {
					itemIndex = k
					break
				}
			}
			itemValue := items[itemIndex].value
			newItems := append(items[:itemIndex], items[itemIndex+1:]...)
			insertIndex := (itemIndex + (itemValue % len(newItems)) + len(newItems)) % len(newItems)
			items = concatMultipleSlices([][]Item{newItems[:insertIndex], {{originalIndex: j, value: itemValue}}, newItems[insertIndex:]})
		}
	}

	return items
}

func getCoordinates(items []Item) int {
	var itemIndex int
	for i := 0; i < len(items); i++ {
		if items[i].value == 0 {
			itemIndex = i
			break
		}
	}
	getValue := func(i int) int { return items[(itemIndex+i*1000)%len(items)].value }
	return getValue(1) + getValue(2) + getValue(3)
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	numbers := lib.Map(input, lib.ParseInt)
	decryptedNumbers := lib.Map(numbers, func(n int) int { return n * 811589153 })

	solution1.I = getCoordinates(mix(numbers, 1))
	solution2.I = getCoordinates(mix(decryptedNumbers, 10))

	return
}
