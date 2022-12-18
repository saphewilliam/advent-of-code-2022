package day15

import (
	"aoc-2022/lib"
	"fmt"
)

type Range struct {
	min int
	max int
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	y, yMax := 2000000, 4000000
	if len(input) == 14 {
		y, yMax = 10, 20
	}

	sensors := make([][]Range, yMax)
	beacons := lib.NewSet[lib.Point2D]()
	for _, line := range input {
		sensor, beacon := lib.Point2D{}, lib.Point2D{}
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
		beacons.Add(beacon)
		dist := lib.Abs(beacon.X-sensor.X) + lib.Abs(beacon.Y-sensor.Y)
		for y := lib.Max(0, sensor.Y-dist); y <= lib.Min(yMax-1, sensor.Y+dist); y++ {
			rNew := Range{sensor.X - (dist - lib.Abs(y-sensor.Y)), sensor.X + (dist - lib.Abs(y-sensor.Y))}

			// fmt.Println(rNew, sensors[y])

			if len(sensors[y]) == 0 {
				sensors[y] = append(sensors[y], rNew)
			}

			for i := 0; i < len(sensors[y]); i++ {
				if i == 0 && rNew.max < sensors[y][i].min-1 {
					sensors[y] = append(sensors[y], rNew)
					copy(sensors[y][i+1:], sensors[y][i:])
					sensors[y][i] = rNew
				}

				if rNew.min > sensors[y][i].max+1 && (i == len(sensors[y])-1 || rNew.max < sensors[y][i+1].max) {
					sensors[y] = append(sensors[y], rNew)
					copy(sensors[y][i+2:], sensors[y][i+1:])
					sensors[y][i+1] = rNew
				}

				if sensors[y][i].min <= rNew.max && rNew.min <= sensors[y][i].max+1 {
					sensors[y] = append(sensors[y], rNew)
					copy(sensors[y][i+2:], sensors[y][i+1:])
					sensors[y][i+1] = rNew
					for i < len(sensors[y])-1 && sensors[y][i].min <= sensors[y][i+1].max && sensors[y][i+1].min <= sensors[y][i].max+1 {
						sensors[y][i] = Range{lib.Min(sensors[y][i+1].min, sensors[y][i].min), lib.Max(sensors[y][i+1].max, sensors[y][i].max)}
						sensors[y] = append(sensors[y][:i+1], sensors[y][i+2:]...)
					}
				}
			}

			// fmt.Println(sensors[y])
		}
	}

	for _, r := range sensors[y] {
		solution1.I += r.max - r.min
	}

	// TODO somehow two possible ranges are found
	count := 0
	for _, s := range sensors {
		if len(s) == 2 {
			count++
		}
	}
	fmt.Println("count:", count)

	for i, s := range sensors {
		if len(s) == 2 {
			solution2.I = (s[0].max+1)*4000000 + i
			break
		}
	}
	return
}
