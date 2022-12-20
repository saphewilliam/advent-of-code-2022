package day19

import (
	"aoc-2022/lib"
	"fmt"
)

type Materials struct {
	ore, clay, obsidian, geode int
}

type Blueprint struct {
	index                      int
	ore, clay, obsidian, geode Materials
}

type State struct {
	time             int
	supplies, robots Materials
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	blueprints := make([]Blueprint, len(input))
	for i, l := range input {
		bp := Blueprint{}
		fmt.Sscanf(l, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &bp.index, &bp.ore.ore, &bp.clay.ore, &bp.obsidian.ore, &bp.obsidian.clay, &bp.geode.ore, &bp.geode.obsidian)
		blueprints[i] = bp
	}

	// TODO
	// Less geode robots is always worse
	for _, blueprint := range blueprints {
		initialState := State{robots: Materials{ore: 1}}
		seen := lib.NewSet(initialState)
		q := lib.NewQueue(initialState)
		blueprintMax := 0

		time := 0
		for !q.IsEmpty() {
			s := q.Dequeue()

			if time != s.time {
				time = s.time
				fmt.Println(time, "/ 24")
			}

			if s.time == 24 {
				blueprintMax = lib.Max(blueprintMax, s.supplies.geode)
				continue
			}

			nextSupplies := Materials{
				ore:      s.supplies.ore + s.robots.ore,
				clay:     s.supplies.clay + s.robots.clay,
				obsidian: s.supplies.obsidian + s.robots.obsidian,
				geode:    s.supplies.geode + s.robots.geode,
			}

			// Build no robot and continue time
			moves := []State{{time: s.time + 1, supplies: nextSupplies, robots: Materials{ore: s.robots.ore, clay: s.robots.clay, obsidian: s.robots.obsidian, geode: s.robots.geode}}}

			// Build geode robot
			if s.supplies.ore >= blueprint.geode.ore && s.supplies.obsidian >= blueprint.geode.obsidian {
				newRobots := Materials{ore: s.robots.ore, clay: s.robots.clay, obsidian: s.robots.obsidian, geode: s.robots.geode + 1}
				newSupplies := Materials{ore: nextSupplies.ore - blueprint.geode.ore, clay: nextSupplies.clay - blueprint.geode.clay, obsidian: nextSupplies.obsidian - blueprint.geode.obsidian, geode: nextSupplies.geode - blueprint.obsidian.geode}
				newState := State{time: s.time + 1, supplies: newSupplies, robots: newRobots}
				moves = append(moves, newState)
			} else {
				// Build obsidian robot
				if s.supplies.ore >= blueprint.obsidian.ore && s.supplies.clay >= blueprint.obsidian.clay {
					newRobots := Materials{ore: s.robots.ore, clay: s.robots.clay, obsidian: s.robots.obsidian + 1, geode: s.robots.geode}
					newSupplies := Materials{ore: nextSupplies.ore - blueprint.obsidian.ore, clay: nextSupplies.clay - blueprint.obsidian.clay, obsidian: nextSupplies.obsidian - blueprint.obsidian.obsidian, geode: nextSupplies.geode - blueprint.obsidian.geode}
					newState := State{time: s.time + 1, supplies: newSupplies, robots: newRobots}
					moves = append(moves, newState)
				}

				// Build clay robot
				if s.supplies.ore >= blueprint.clay.ore {
					newRobots := Materials{ore: s.robots.ore, clay: s.robots.clay + 1, obsidian: s.robots.obsidian, geode: s.robots.geode}
					newSupplies := Materials{ore: nextSupplies.ore - blueprint.clay.ore, clay: nextSupplies.clay - blueprint.clay.clay, obsidian: nextSupplies.obsidian - blueprint.clay.obsidian, geode: nextSupplies.geode - blueprint.obsidian.geode}
					newState := State{time: s.time + 1, supplies: newSupplies, robots: newRobots}
					moves = append(moves, newState)
				}

				// Build ore robot
				if s.supplies.ore >= blueprint.ore.ore {
					newRobots := Materials{ore: s.robots.ore + 1, clay: s.robots.clay, obsidian: s.robots.obsidian, geode: s.robots.geode}
					newSupplies := Materials{ore: nextSupplies.ore - blueprint.ore.ore, clay: nextSupplies.clay - blueprint.ore.clay, obsidian: nextSupplies.obsidian - blueprint.ore.obsidian, geode: nextSupplies.geode - blueprint.obsidian.geode}
					newState := State{time: s.time + 1, supplies: newSupplies, robots: newRobots}
					moves = append(moves, newState)
				}
			}

			// fmt.Println(s, moves)
			for _, m := range moves {
				if !seen.Has(m) {
					seen.Add(m)
					q.Enqueue(m)
				}
			}

		}

		solution1.I += blueprint.index * blueprintMax
	}

	return
}
