package main

import (
	day "aoc-2022/day04"
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open(day.Path)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	score1, score2 := day.Calculate(lines)
	fmt.Println("Score 1:", score1)
	fmt.Println("Score 2:", score2)
}
