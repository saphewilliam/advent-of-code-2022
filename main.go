package main

import (
	day "aoc-2022/day03"
	"bufio"
	"fmt"
	"os"
	"time"
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

	timeStart := time.Now()
	score1, score2 := day.Calculate(lines)
	timeElapsed := time.Since(timeStart)

	fmt.Printf("Score 1: %d\nScore 2: %d\nTime elapsed: %dms", score1, score2, timeElapsed.Milliseconds())
}
