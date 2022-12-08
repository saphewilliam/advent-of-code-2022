package main

import (
	"aoc-2022/lib"
	"bufio"
	"fmt"
	"os"
	"time"
)

func readLines(path string) (lines []string) {
	read, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(read)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return
}

func gradeSolution(name, correct string, test lib.Solution) {
	if correct == "" {
		fmt.Print(" \033[33m?\033[0m  ")
	} else if test.GetValue() == correct {
		fmt.Print(" \033[32m✓\033[0m  ")
	} else {
		fmt.Printf(" \033[31mx\033[0m  (%v) ", correct)
	}
	fmt.Printf("%v: %v\n", name, test.GetValue())
}

// TODO testonly/realonly flag
func main() {
	for _, v := range os.Args[1:] {
		day, dayFound := getDay(v)
		if !dayFound {
			fmt.Println(v + " is not a valid day value!")
			continue
		}
		testInput := readLines("./day" + v + "/testinput.txt")
		input := readLines("./day" + v + "/input.txt")

		timeStart := time.Now()
		testSolution1, testSolution2 := day.Process(testInput)
		solution1, solution2 := day.Process(input)
		timeElapsed := time.Since(timeStart)

		fmt.Println("Day", v)
		gradeSolution("Test solution 1", day.TestPart1, testSolution1)
		gradeSolution("Test solution 2", day.TestPart2, testSolution2)
		gradeSolution("Solution 1", day.Part1, solution1)
		gradeSolution("Solution 2", day.Part2, solution2)
		fmt.Printf("Time elapsed: %dms\n\n", timeElapsed.Milliseconds())
	}
}
