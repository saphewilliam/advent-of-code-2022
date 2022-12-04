package day03

import (
	"math"
	"strings"
)

const Path = "./day03/input.txt"
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Calculate(input []string) (individualSum int, groupSum int) {
	var groupContent, leftContent, rightContent int64 = 0, 0, 0
	for i := 0; i < len(input); i++ {
		content := input[i]
		leftContent, rightContent = 0, 0
		for _, v := range content[:len(content)/2] {
			leftContent |= 1 << strings.IndexRune(alphabet, v)
		}
		for _, v := range content[len(content)/2:] {
			rightContent |= 1 << strings.IndexRune(alphabet, v)
		}

		individualSum += 1 + int(math.Logb(float64(leftContent&rightContent)))
		if groupContent == 0 {
			groupContent = (leftContent | rightContent)
		} else {
			groupContent &= (leftContent | rightContent)
		}
		if i%3 == 2 {
			groupSum += 1 + int(math.Logb(float64(groupContent)))
			groupContent = 0
		}
	}
	return
}
