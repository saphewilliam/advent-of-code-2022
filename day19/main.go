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

func canBuildRobot(supplies, cost Materials) bool {
	return supplies.ore >= cost.ore && supplies.clay >= cost.clay && supplies.obsidian >= cost.obsidian && supplies.geode >= supplies.geode
}

func buyRobot(supplies, cost Materials) Materials {
	return Materials{ore: supplies.ore - cost.ore, clay: supplies.clay - cost.clay, obsidian: supplies.obsidian - cost.obsidian, geode: supplies.geode - cost.geode}
}

func greedySimulate(blueprint Blueprint, robots, supplies Materials, time int) int {
	additionalRobots, potentialSupplies := Materials{}, Materials{ore: supplies.ore, clay: supplies.clay, obsidian: supplies.obsidian, geode: supplies.geode}

	for i := 0; i < time; i++ {
		potentialSupplies = Materials{
			ore:      potentialSupplies.ore + robots.ore + additionalRobots.ore,
			clay:     potentialSupplies.clay + robots.clay + additionalRobots.clay,
			obsidian: potentialSupplies.obsidian + robots.obsidian + additionalRobots.obsidian,
			geode:    potentialSupplies.geode + robots.geode + additionalRobots.geode,
		}

		eeeh := func(robot Materials) bool {
			return potentialSupplies.ore >= robot.ore*(additionalRobots.ore+1) && potentialSupplies.clay >= robot.clay*(additionalRobots.ore+1) && potentialSupplies.obsidian >= robot.obsidian*(additionalRobots.ore+1) && potentialSupplies.geode >= robot.geode*(additionalRobots.ore+1)
		}

		if eeeh(blueprint.ore) {
			additionalRobots.ore++
		}
		if eeeh(blueprint.clay) {
			additionalRobots.clay++
		}
		if eeeh(blueprint.obsidian) {
			additionalRobots.obsidian++
		}
		if eeeh(blueprint.geode) {
			additionalRobots.geode++
		}
	}

	return potentialSupplies.geode

	// additional_robots = defaultdict(int)
	// potential_materials = defaultdict(int, materials)
	// for _ in range(time):
	//     for robot in blueprint:
	//         potential_materials[robot] += robots[robot] + additional_robots[robot]
	//     for robot, costs in blueprint.items():
	//         if all(
	//             potential_materials[material] >= cost * (additional_robots[robot] + 1)
	//             for material, cost in costs.items()
	//         ):
	//             additional_robots[robot] += 1
	// return potential_materials["geode"]
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	blueprints := make([]Blueprint, len(input))
	for i, l := range input {
		bp := Blueprint{}
		fmt.Sscanf(l, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &bp.index, &bp.ore.ore, &bp.clay.ore, &bp.obsidian.ore, &bp.obsidian.clay, &bp.geode.ore, &bp.geode.obsidian)
		blueprints[i] = bp
	}

	for _, blueprint := range blueprints {
		robotLimits := Materials{
			ore:      lib.Max(lib.Max(lib.Max(blueprint.ore.ore, blueprint.clay.ore), blueprint.obsidian.ore), blueprint.geode.ore),
			clay:     blueprint.obsidian.clay,
			obsidian: blueprint.geode.obsidian,
			geode:    1000,
		}

		initialState := State{robots: Materials{ore: 1}, time: 24}
		q := lib.NewQueue(initialState)
		maxGeode := 0
		time := 24
		for !q.IsEmpty() {
			s := q.Dequeue()

			maxGeode = lib.Max(maxGeode, s.supplies.geode)
			if time != s.time {
				time = s.time
				fmt.Println(time, maxGeode)
			}

			if greedySimulate(blueprint, s.robots, s.supplies, time) < maxGeode {
				continue
			}

			estimatedGeodes := s.supplies.geode + s.robots.geode*time
			maxGeode = lib.Max(maxGeode, estimatedGeodes)

			if s.time == 0 {
				continue
			}

			nextSupplies := Materials{
				ore:      s.supplies.ore + s.robots.ore,
				clay:     s.supplies.clay + s.robots.clay,
				obsidian: s.supplies.obsidian + s.robots.obsidian,
				geode:    s.supplies.geode + s.robots.geode,
			}

			// Build no robot and continue time
			// TODO test if just passing s.robots to robots has any effect
			moves := []State{{time: s.time - 1, supplies: nextSupplies, robots: Materials{ore: s.robots.ore, clay: s.robots.clay, obsidian: s.robots.obsidian, geode: s.robots.geode}}}

			if s.robots.ore < robotLimits.ore && canBuildRobot(s.supplies, blueprint.ore) {
				newRobots := Materials{ore: s.robots.ore + 1, clay: s.robots.clay, obsidian: s.robots.obsidian, geode: s.robots.geode}
				newSupplies := buyRobot(nextSupplies, blueprint.ore)
				newState := State{time: s.time - 1, supplies: newSupplies, robots: newRobots}
				moves = append(moves, newState)
			}
			if s.robots.clay < robotLimits.clay && canBuildRobot(s.supplies, blueprint.clay) {
				newRobots := Materials{ore: s.robots.ore, clay: s.robots.clay + 1, obsidian: s.robots.obsidian, geode: s.robots.geode}
				newSupplies := buyRobot(nextSupplies, blueprint.clay)
				newState := State{time: s.time - 1, supplies: newSupplies, robots: newRobots}
				moves = append(moves, newState)
			}
			if s.robots.obsidian < robotLimits.obsidian && canBuildRobot(s.supplies, blueprint.obsidian) {
				newRobots := Materials{ore: s.robots.ore, clay: s.robots.clay, obsidian: s.robots.obsidian + 1, geode: s.robots.geode}
				newSupplies := buyRobot(nextSupplies, blueprint.obsidian)
				newState := State{time: s.time - 1, supplies: newSupplies, robots: newRobots}
				moves = append(moves, newState)
			}
			if s.robots.geode < robotLimits.geode && canBuildRobot(s.supplies, blueprint.geode) {
				newRobots := Materials{ore: s.robots.ore, clay: s.robots.clay, obsidian: s.robots.obsidian, geode: s.robots.geode + 1}
				newSupplies := buyRobot(nextSupplies, blueprint.geode)
				newState := State{time: s.time - 1, supplies: newSupplies, robots: newRobots}
				moves = append(moves, newState)
			}

			for _, m := range moves {
				q.Enqueue(m)
			}

		}

		solution1.I += blueprint.index * maxGeode
	}

	return
}
