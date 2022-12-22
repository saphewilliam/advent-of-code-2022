package day22

import (
	"aoc-2022/lib"
	"regexp"
)

type Me struct {
	p      lib.Point2D
	facing int
}

type Move struct {
	forward int
	rotate  func(int) int
}

func parseRotate(rotation string) func(int) int {
	rotate := func(n, r int) int { return (n + r + 4) % 4 }
	switch rotation {
	case "R":
		return func(n int) int { return rotate(n, 1) }
	case "L":
		return func(n int) int { return rotate(n, -1) }
	default:
		return func(n int) int { return n }
	}
}

func wrap3D(edgeRight, edgeBottom, edgeLeft, edgeTop map[int]int) map[Me]lib.Point2D {
	wrap := map[Me]lib.Point2D{}
	// no
	return wrap
}

func wrap2D(edgeRight, edgeBottom, edgeLeft, edgeTop map[int]int) map[Me]lib.Point2D {
	wrap := map[Me]lib.Point2D{}
	for y, x := range edgeRight {
		if x != -1 && edgeLeft[y] != -1 {
			wrap[Me{p: lib.NewP2D(x, y), facing: 0}] = lib.NewP2D(edgeLeft[y], y)
			wrap[Me{p: lib.NewP2D(edgeLeft[y], y), facing: 2}] = lib.NewP2D(x, y)
		}
	}
	for x, y := range edgeBottom {
		if y != -1 && edgeTop[x] != -1 {
			wrap[Me{p: lib.NewP2D(x, y), facing: 1}] = lib.NewP2D(x, edgeTop[x])
			wrap[Me{p: lib.NewP2D(x, edgeTop[x]), facing: 3}] = lib.NewP2D(x, y)
		}
	}
	return wrap
}

func evaluatePuzzle(points lib.Set[lib.Point2D], wrap map[Me]lib.Point2D, me Me, moves lib.Queue[Move]) int {
	for !moves.IsEmpty() {
		m := moves.Dequeue()
		for steps := 0; steps < m.forward; steps++ {
			switch me.facing {
			case 0:
				if points.Has(lib.NewP2D(me.p.X+1, me.p.Y)) {
					me.p.X++
					continue
				}
			case 1:
				if points.Has(lib.NewP2D(me.p.X, me.p.Y+1)) {
					me.p.Y++
					continue
				}
			case 2:
				if points.Has(lib.NewP2D(me.p.X-1, me.p.Y)) {
					me.p.X--
					continue
				}
			case 3:
				if points.Has(lib.NewP2D(me.p.X, me.p.Y-1)) {
					me.p.Y--
					continue
				}
			}
			if _, exists := wrap[me]; exists {
				me.p = wrap[me]
				continue
			}
			break
		}
		me.facing = m.rotate(me.facing)
	}
	return 1000*me.p.Y + 4*me.p.X + me.facing
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	points := lib.NewSet[lib.Point2D]()
	// TODO refactor into map[int]map[int]int{0: map[int]int{}, 1: map[int]int{}, 2: map[int]int{}, 3: map[int]int{}}
	edgeRight := map[int]int{}
	edgeBottom := map[int]int{}
	edgeLeft := map[int]int{}
	edgeTop := map[int]int{}
	var firstX int

	for y := 1; y <= len(input)-2; y++ {
		for x := 1; x <= len(input[y-1]); x++ {
			value := input[y-1][x-1]

			setEdge := func(condition bool, edge map[int]int, a, b int) {
				if condition {
					if value == '#' {
						edge[a] = -1
					} else if value == '.' {
						edge[a] = b
					}
				}
			}

			setEdge(x == len(input[y-1]), edgeRight, y, x)
			setEdge(y == len(input)-1 || len(input[y]) < x || input[y][x-1] == ' ', edgeBottom, x, y)
			setEdge(edgeLeft[y] == 0, edgeLeft, y, x)
			setEdge(edgeTop[x] == 0, edgeTop, x, y)

			if value == '.' {
				points.Add(lib.NewP2D(x, y))
				if firstX == 0 {
					firstX = x
				}
			}
		}
	}

	me := Me{p: lib.NewP2D(firstX, 1)}
	re := regexp.MustCompile(`(\d+)(R|L)?`)
	matches := re.FindAllStringSubmatch(input[len(input)-1], -1)
	moves := lib.NewQueue[Move]()
	for _, match := range matches {
		moves.Enqueue(Move{forward: lib.ParseInt(match[1]), rotate: parseRotate(match[2])})
	}

	solution1.I = evaluatePuzzle(points, wrap2D(edgeRight, edgeBottom, edgeLeft, edgeTop), me, moves)
	solution2.I = evaluatePuzzle(points, wrap3D(edgeRight, edgeBottom, edgeLeft, edgeTop), me, moves)
	return
}
