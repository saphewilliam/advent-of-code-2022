package day14

import (
	"aoc-2022/lib"
	"strings"
)

func buildCave(input []string) (cave lib.Set[lib.Point], lowBound int) {
	cave = lib.NewSet[lib.Point]()
	var prevP lib.Point
	for _, line := range input {
		points := strings.Split(line, " -> ")
		for i, point := range points {
			s := strings.Split(point, ",")
			p := lib.Point{X: lib.ParseInt(s[0]), Y: lib.ParseInt(s[1])}
			if i != 0 {
				for i := lib.Min(p.X, prevP.X); i <= lib.Max(p.X, prevP.X); i++ {
					for j := lib.Min(p.Y, prevP.Y); j <= lib.Max(p.Y, prevP.Y); j++ {
						cave.Add(lib.NewPoint(i, j))
					}
				}
			}
			if i != len(points)-1 {
				prevP = p
			}
			if lowBound == 0 || lowBound < p.Y {
				lowBound = p.Y
			}
		}
	}
	return
}

func getSandDestination(cave lib.Set[lib.Point], lowBound int, s lib.Point) lib.Point {
	points := []lib.Point{
		lib.NewPoint(s.X, s.Y+1),
		lib.NewPoint(s.X-1, s.Y+1),
		lib.NewPoint(s.X+1, s.Y+1),
	}
	for _, p := range points {
		if s.Y != lowBound+1 && !cave.Has(p) {
			return getSandDestination(cave, lowBound, p)
		}
	}
	return s
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	c, lowBound := buildCave(input)
	source := lib.NewPoint(500, 0)
	for {
		s := getSandDestination(c, lowBound, source)
		if solution1.I == 0 && s.Y == lowBound+1 {
			solution1.I = solution2.I
		}
		solution2.I++

		c.Add(s)
		if s == source {
			break
		}
	}
	return
}
