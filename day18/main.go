package day18

import (
	"aoc-2022/lib"
	"strings"
)

func getEdges(p lib.Point3D) []lib.Point3D {
	return []lib.Point3D{
		{X: p.X + 1, Y: p.Y, Z: p.Z},
		{X: p.X - 1, Y: p.Y, Z: p.Z},
		{X: p.X, Y: p.Y + 1, Z: p.Z},
		{X: p.X, Y: p.Y - 1, Z: p.Z},
		{X: p.X, Y: p.Y, Z: p.Z + 1},
		{X: p.X, Y: p.Y, Z: p.Z - 1},
	}
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	lava := lib.NewSet[lib.Point3D]()
	size := lib.NewS3D(1000, 0, 1000, 0, 1000, 0)
	for _, l := range input {
		c := lib.Map(strings.Split(l, ","), lib.ParseInt)
		p := lib.NewP3D(c[0], c[1], c[2])
		lava.Add(p)
		size.SetAll(p)
		edges := getEdges(p)

		solution1.I += 6
		for _, e := range edges {
			if lava.Has(e) {
				solution1.I -= 2
			}
		}
	}

	initialPoint := lib.NewP3D(size.MinX, size.MinY, size.MinZ)
	q := lib.NewQueue(initialPoint)
	explored := lib.NewSet(initialPoint)
	for !q.IsEmpty() {
		p := q.Dequeue()
		edges := getEdges(p)
		for _, e := range edges {
			if lava.Has(e) {
				solution2.I++
				continue
			}
			if !explored.Has(e) &&
				e.X >= size.MinX-1 &&
				e.X <= size.MaxX+1 &&
				e.Y >= size.MinY-1 &&
				e.Y <= size.MaxY+1 &&
				e.Z >= size.MinZ-1 &&
				e.Z <= size.MaxZ+1 {
				explored.Add(e)
				q.Enqueue(e)
			}
		}
	}

	return
}
