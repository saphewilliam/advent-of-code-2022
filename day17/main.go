package day17

import (
	"aoc-2022/lib"
)

type RockCacheKey struct {
	rockIndex, jetIndex, col0, col1, col2, col3, col4, col5, col6 int
}

type RockCacheValue struct {
	maxHeight, index int
}

func getRock(index int, origin lib.Point) (lib.Set[lib.Point], int) {
	rocks := [][]lib.Point{
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
	}
	translate := func(r lib.Point) lib.Point { return lib.Point{X: r.X + origin.X, Y: r.Y + origin.Y} }
	id := index % len(rocks)
	return lib.NewSet(lib.Map(rocks[id], translate)...), id
}

func execute(jets []int, target int) (maxHeight int) {
	cache := map[RockCacheKey]RockCacheValue{}
	rocks := lib.NewSet[lib.Point]()
	jetIndex := 0
	top0, top1, top2, top3, top4, top5, top6 := 0, 0, 0, 0, 0, 0, 0
	for i := 0; i < target; i++ {
		origin := lib.Point{X: 2, Y: maxHeight + 4}
		rock, rockIndex := getRock(i, origin)

		// fmt.Println(solution1.I, jetIndex, rockIndex, top0, top1, top2, top3, top4, top5, top6)
		ck := RockCacheKey{jetIndex: jetIndex, rockIndex: rockIndex, col0: top0, col1: top1, col2: top2, col3: top3, col4: top4, col5: top5, col6: top6}
		cvNew := RockCacheValue{maxHeight: maxHeight, index: i}
		if cv, exists := cache[ck]; exists {
			restTarget := target - i
			segmentSize := cvNew.index - cv.index
			segmentHeight := cvNew.maxHeight - cv.maxHeight
			segmentsLeft := restTarget / segmentSize

			i += segmentsLeft * segmentSize
			maxHeight += segmentsLeft * segmentHeight

			newRocks := lib.NewSet[lib.Point]()
			for _, r := range rocks.Elements() {
				newRocks.Add(lib.Point{X: r.X, Y: r.Y + segmentsLeft*segmentHeight})
			}
			rocks = newRocks

			translate := func(r lib.Point) lib.Point { return lib.Point{X: r.X + origin.X, Y: r.Y + segmentsLeft*segmentHeight} }
			rock = lib.NewSet(lib.Map(rock.Elements(), translate)...)

			// fmt.Println("Cache hit:", i, solution1.I, segmentsLeft)
		} else {
			cache[ck] = cvNew
		}

		for {
			isSettled := false

			// Move by jet pulse
			jetRock := lib.NewSet[lib.Point]()
			for _, r := range rock.Elements() {
				newR := lib.Point{X: r.X + jets[jetIndex], Y: r.Y}
				if newR.X == 7 || newR.X == -1 || rocks.Has(newR) {
					jetRock = rock
					break
				} else {
					jetRock.Add(newR)
				}
			}
			jetIndex = (jetIndex + 1) % len(jets)
			rock = jetRock

			// Move by gravity
			gravityRock := lib.NewSet[lib.Point]()
			for _, r := range rock.Elements() {
				newR := lib.Point{X: r.X, Y: r.Y - 1}
				if newR.Y == 0 || rocks.Has(newR) {
					gravityRock = rock
					isSettled = true
					break
				} else {
					gravityRock.Add(newR)
				}
			}
			rock = gravityRock

			// Settle
			if isSettled {
				top0 += 4
				top1 += 4
				top2 += 4
				top3 += 4
				top4 += 4
				top5 += 4
				top6 += 4
				for _, r := range rock.Elements() {
					maxHeight = lib.Max(maxHeight, r.Y)
					switch r.X {
					case 0:
						top0 = lib.Min(top0, origin.Y-r.Y)
					case 1:
						top1 = lib.Min(top1, origin.Y-r.Y)
					case 2:
						top2 = lib.Min(top2, origin.Y-r.Y)
					case 3:
						top3 = lib.Min(top3, origin.Y-r.Y)
					case 4:
						top4 = lib.Min(top4, origin.Y-r.Y)
					case 5:
						top5 = lib.Min(top5, origin.Y-r.Y)
					case 6:
						top6 = lib.Min(top6, origin.Y-r.Y)
					}
				}
				rocks = lib.SetUnion(rocks, rock)
				break
			}
		}
	}
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	jets := make([]int, len(input[0]))
	for i, l := range input[0] {
		if l == '<' {
			jets[i] = -1
		} else {
			jets[i] = 1
		}
	}

	solution1.I = execute(jets, 2022)
	solution2.I = execute(jets, 1000000000000)
	return
}

// 1556521739090 incorrect
