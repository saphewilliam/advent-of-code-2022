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

func main() {
	args, testOnly := os.Args[1:], false
	if args[0] == "0" {
		args, testOnly = args[1:], true
	}
	for _, v := range args {
		day, dayFound := getDay(v)
		if !dayFound {
			fmt.Println(v + " is not a valid day value!")
			continue
		}
		testInput := readLines("./day" + v + "/testinput.txt")
		input := readLines("./day" + v + "/input.txt")

		fmt.Println("Day", v)
		timeStart := time.Now()
		testSolution1, testSolution2 := day.Process(testInput)
		gradeSolution("Test solution 1", day.TestPart1, testSolution1)
		gradeSolution("Test solution 2", day.TestPart2, testSolution2)
		if !testOnly {
			solution1, solution2 := day.Process(input)
			gradeSolution("Solution 1", day.Part1, solution1)
			gradeSolution("Solution 2", day.Part2, solution2)
		}
		timeElapsed := time.Since(timeStart)
		fmt.Printf("Time elapsed: %dms\n\n", timeElapsed.Milliseconds())
	}
}
