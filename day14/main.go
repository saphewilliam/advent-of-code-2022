package day14

import (
	"aoc-2022/lib"
	"math"
	"strings"
)

type Line struct {
	startX int
	startY int
	endX   int
	endY   int
}

type Cave struct {
	lowBound   int
	leftBound  int
	rightBound int
	blocked    []bool
}

const SPAN = 200
const WIDTH = SPAN * 2
const CENTER = 500

func buildCave(input []string) (c Cave) {
	lines := []Line{}
	for _, line := range input {
		coords := strings.Split(line, " -> ")
		for i, coord := range coords {
			xy := strings.Split(coord, ",")
			x := lib.UnsafeParseInt(xy[0])
			y := lib.UnsafeParseInt(xy[1])

			if i != 0 {
				lines[len(lines)-1].endX = x
				lines[len(lines)-1].endY = y
			}
			if i != len(coords)-1 {
				lines = append(lines, Line{startX: x, startY: y})
			}

			if c.lowBound == 0 || c.lowBound < y {
				c.lowBound = y
			}
			if c.leftBound == 0 || c.leftBound > x {
				c.leftBound = x
			}
			if c.rightBound == 0 || c.rightBound < x {
				c.rightBound = x
			}
		}
	}

	c.blocked = make([]bool, WIDTH*(c.lowBound+2))
	for _, l := range lines {
		if l.startX == l.endX {
			min := int(math.Min(float64(l.startY), float64(l.endY)))
			max := int(math.Max(float64(l.startY), float64(l.endY)))
			for i := min; i <= max; i++ {
				c.blocked[(l.startX+SPAN-CENTER)+i*WIDTH] = true
			}
		}
		if l.startY == l.endY {
			min := int(math.Min(float64(l.startX), float64(l.endX)))
			max := int(math.Max(float64(l.startX), float64(l.endX)))
			for i := min; i <= max; i++ {
				c.blocked[(i+SPAN-CENTER)+l.startY*WIDTH] = true
			}
		}
	}
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	c, isOutOfBounds := buildCave(input), false

	for {
		s := SPAN
		for {
			if s+WIDTH >= WIDTH*(c.lowBound+1) ||
				s%WIDTH < c.leftBound+SPAN-CENTER ||
				s%WIDTH > c.rightBound+SPAN-CENTER {
				isOutOfBounds = true
			}

			// Straight down
			if s+WIDTH < len(c.blocked) && !c.blocked[s+WIDTH] {
				s += WIDTH
				continue
			}

			// Left down
			if s%WIDTH == 0 {
				panic("Cannot go left, increase SPAN constant")
			}
			if s+WIDTH-1 < len(c.blocked) && !c.blocked[s+WIDTH-1] {
				s += WIDTH - 1
				continue
			}

			// Right down
			if s%WIDTH == WIDTH-1 {
				panic("Cannot go right, increase SPAN constant")
			}
			if s+WIDTH+1 < len(c.blocked) && !c.blocked[s+WIDTH+1] {
				s += WIDTH + 1
				continue
			}

			// Tally scores
			c.blocked[s] = true
			if !isOutOfBounds {
				solution1.I++
			}
			solution2.I++
			break
		}
		if s == SPAN {
			break
		}
	}
	return
}
