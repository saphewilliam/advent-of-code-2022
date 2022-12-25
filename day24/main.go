package day24

import "aoc-2022/lib"

type State struct {
	p            lib.Point2D
	blizardIndex int
}

func getLCM(a, b int) int {
	originalA, originalB := a, b
	for a != b {
		if a < b {
			a += originalA
		} else {
			b += originalB
		}
	}
	return a
}

func dfs(in, out lib.Point2D, size lib.Size2D, blizards []lib.Set[lib.Point2D], initialState State) (int, State) {
	seen := map[State]int{initialState: 0}
	q := lib.NewQueue(initialState)
	for !q.IsEmpty() {
		s := q.Dequeue()
		if s.p == out {
			return seen[s], s
		}
		for _, edge := range []lib.Point2D{{X: s.p.X, Y: s.p.Y}, {X: s.p.X + 1, Y: s.p.Y}, {X: s.p.X - 1, Y: s.p.Y}, {X: s.p.X, Y: s.p.Y + 1}, {X: s.p.X, Y: s.p.Y - 1}} {
			bi := (s.blizardIndex + 1) % len(blizards)
			ns := State{p: edge, blizardIndex: bi}
			if _, exists := seen[ns]; !exists && ((ns.p == out || ns.p == in) || (ns.p.X >= size.MinX && ns.p.X < size.MaxX && ns.p.Y >= size.MinY && ns.p.Y < size.MaxY && !blizards[bi].Has(ns.p))) {
				seen[ns] = seen[s] + 1
				q.Enqueue(ns)
			}
		}
	}
	return -1, initialState
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	in, out := lib.Point2D{}, lib.Point2D{}
	size := lib.Size2D{MaxX: len(input[0]) - 2, MaxY: len(input) - 2}
	startBlizards := map[lib.Point2D][]int{}
	for y, l := range input {
		for x, c := range l {
			switch c {
			case '.':
				if y == 0 {
					in = lib.NewP2D(x-1, y-1)
				} else if y == len(input)-1 {
					out = lib.NewP2D(x-1, y-1)
				}
			case '>':
				startBlizards[lib.NewP2D(x-1, y-1)] = append(startBlizards[lib.NewP2D(x-1, y-1)], 0)
			case 'v':
				startBlizards[lib.NewP2D(x-1, y-1)] = append(startBlizards[lib.NewP2D(x-1, y-1)], 1)
			case '<':
				startBlizards[lib.NewP2D(x-1, y-1)] = append(startBlizards[lib.NewP2D(x-1, y-1)], 2)
			case '^':
				startBlizards[lib.NewP2D(x-1, y-1)] = append(startBlizards[lib.NewP2D(x-1, y-1)], 3)
			}
		}
	}

	// Simulate blizard pattern
	lcm := getLCM(size.MaxX, size.MaxY)
	blizards := make([]lib.Set[lib.Point2D], lcm)
	currBlizards := startBlizards
	for i := 0; i < lcm; i++ {
		blizards[i] = lib.NewSet[lib.Point2D]()
		newBlizards := map[lib.Point2D][]int{}
		for p, bs := range currBlizards {
			blizards[i].Add(p)
			for _, b := range bs {
				newP := lib.Point2D{}
				switch b {
				case 0:
					newP = lib.NewP2D((p.X+1)%size.MaxX, p.Y)
				case 1:
					newP = lib.NewP2D(p.X, (p.Y+1)%size.MaxY)
				case 2:
					newP = lib.NewP2D((p.X-1+size.MaxX)%size.MaxX, p.Y)
				case 3:
					newP = lib.NewP2D(p.X, (p.Y-1+size.MaxY)%size.MaxY)
				}
				newBlizards[newP] = append(newBlizards[newP], b)
			}
		}
		currBlizards = newBlizards
	}

	initialState := State{p: in}
	trip1Time, trip1State := dfs(in, out, size, blizards, initialState)
	trip2Time, trip2State := dfs(out, in, size, blizards, trip1State)
	trip3Time, _ := dfs(in, out, size, blizards, trip2State)

	solution1.I = trip1Time
	solution2.I = trip1Time + trip2Time + trip3Time
	return
}
