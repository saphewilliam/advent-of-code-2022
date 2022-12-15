package day15

import (
	"aoc-2022/lib"
	"fmt"
)

func manhattanDistance(source, dest lib.Point) int {
	return lib.Abs(source.X-dest.X) + lib.Abs(source.Y-dest.Y)
}

func searchSpace(xLow, xUp, yLow, yUp int, sensors map[lib.Point]int, beacons lib.Set[lib.Point]) (points lib.Set[lib.Point]) {
	points = lib.NewSet[lib.Point]()
	for x := xLow; x <= xUp; x++ {
		for y := yLow; y <= yUp; y++ {
			p := lib.NewPoint(x, y)
			for sensor, dist := range sensors {
				if beacons.Has(p) {
					break
				} else if manhattanDistance(sensor, p) <= dist {
					points.Add(p)
					break
				}
			}
		}
		if x%100000 == 0 {
			fmt.Println(x, "/", xUp)
		}
	}
	return
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	sensors := map[lib.Point]int{}
	beacons := lib.NewSet[lib.Point]()
	for _, line := range input {
		sensor, beacon := lib.Point{}, lib.Point{}
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
		sensors[sensor] = manhattanDistance(beacon, sensor)
		beacons.Add(beacon)
	}

	xLow, xUp, y, cLow, cUp := -10000000, 10000000, 2000000, 0, 4000000
	if len(input) == 14 {
		xLow, xUp, y, cLow, cUp = -5, 30, 10, 0, 20
	}

	s1 := searchSpace(xLow, xUp, y, y, sensors, beacons)
	s2 := searchSpace(cLow, cUp, cLow, cUp, sensors, beacons)
	solution1.I = s1.Size()

	u := lib.SetUnion(s2, beacons)
	// fmt.Println(cLow, cUp, len(s), s)

	for x := cLow; x <= cUp; x++ {
		for y := cLow; y <= cUp; y++ {
			found := false
			// u.has
			for _, p := range u.Elements() {
				if p == lib.NewPoint(x, y) {
					found = true
					break
				}
			}
			if !found {
				fmt.Println(x, y)
			}
		}
	}

	// solution2.I = len(searchSpace(cLow, cUp, cLow, cUp, sensors, beacons))

	return
}
