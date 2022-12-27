package day23

import (
	"aoc-2022/lib"
)

type Direction struct {
	left, mid, right lib.Point2D
}

func getDirections(origin lib.Point2D) []Direction {
	d := getPoints(origin)
	return []Direction{
		// North
		{left: d[0], mid: d[1], right: d[2]},
		// South
		{left: d[8], mid: d[7], right: d[6]},
		// West
		{left: d[6], mid: d[3], right: d[0]},
		// East
		{left: d[2], mid: d[5], right: d[8]},
	}
}

func getPoints(origin lib.Point2D) []lib.Point2D {
	points := []lib.Point2D{
		{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
		{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0},
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
	}

	return lib.Map(points, func(p lib.Point2D) lib.Point2D {
		return lib.Point2D{X: p.X + origin.X, Y: p.Y + origin.Y}
	})
}

func getElvesSize(elves lib.Set[lib.Point2D]) lib.Size2D {
	size := lib.Size2D{MinX: 1000, MaxX: -1000, MinY: 1000, MaxY: -1000}
	for _, e := range elves.Elements() {
		if e.X < size.MinX {
			size.MinX = e.X
		}
		if e.X > size.MaxX {
			size.MaxX = e.X
		}
		if e.Y < size.MinY {
			size.MinY = e.Y
		}
		if e.Y > size.MaxY {
			size.MaxY = e.Y
		}
	}
	return size
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	elves := lib.NewSet[lib.Point2D]()
	for y, l := range input {
		for x, c := range l {
			if c == '#' {
				elves.Add(lib.NewP2D(x, y))
			}
		}
	}

	oldElves := lib.NewSet(elves.Elements()...)
	for i := 0; ; i++ {
		propositions := map[lib.Point2D][]lib.Point2D{}

		for _, elf := range elves.Elements() {
			points := getPoints(elf)
			shouldPropose := false

			for _, p := range points {
				if elves.Has(p) && p != elf {
					shouldPropose = true
				}
			}

			if shouldPropose {
				directions := getDirections(elf)
				for j := 0; j < len(directions); j++ {
					d := directions[(j+i)%len(directions)]
					if !elves.Has(d.left) && !elves.Has(d.mid) && !elves.Has(d.right) {
						propositions[d.mid] = append(propositions[d.mid], elf)
						break
					}
				}
			}
		}

		for dest, e := range propositions {
			if len(e) == 1 {
				elves.Remove(e[0])
				elves.Add(dest)
			}
		}

		if i == 10 {
			size := getElvesSize(elves)
			solution1.I = (size.MaxX-size.MinX+1)*(size.MaxY-size.MinY+1) - elves.Size()
		}

		diff := lib.SetDifference(elves, oldElves)
		if diff.Size() == 0 {
			solution2.I = i + 1
			break
		} else {
			oldElves = lib.NewSet(elves.Elements()...)
		}
	}

	return
}
