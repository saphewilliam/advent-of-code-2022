package day19

import (
	"aoc-2022/lib"
	"fmt"
	"math"
)

type Materials struct {
	ore, clay, obsidian, geode int
}

type Blueprint struct {
	index                      int
	ore, clay, obsidian, geode Materials
}

func buildRobot(supplies, cost, robots Materials, timeDelta, geodes int) Materials {
	return Materials{
		ore:      supplies.ore + timeDelta*robots.ore - cost.ore,
		clay:     supplies.clay + timeDelta*robots.clay - cost.clay,
		obsidian: supplies.obsidian + timeDelta*robots.obsidian - cost.obsidian,
		geode:    supplies.geode + geodes,
	}
}

func canBuildRobot(supplies, cost Materials) bool {
	return supplies.ore >= cost.ore && supplies.clay >= cost.clay && supplies.obsidian >= cost.obsidian && supplies.geode >= cost.geode
}

func dfs(blueprint Blueprint, robotLimits, robots, supplies Materials, time, score int) int {
	if time < 1 || supplies.geode+(time*time)/2 < score {
		return score
	}
	score = lib.Max(score, supplies.geode)

	//Build geode robot
	if robots.obsidian > 0 {
		timeDelta := 1
		if !canBuildRobot(supplies, blueprint.geode) {
			timeDelta += int(math.Max(
				math.Ceil(float64(blueprint.geode.ore-supplies.ore)/float64(robots.ore)),
				math.Ceil(float64(blueprint.geode.obsidian-supplies.obsidian)/float64(robots.obsidian)),
			))
		}

		score = lib.Max(score, dfs(
			blueprint,
			robotLimits,
			robots,
			buildRobot(supplies, blueprint.geode, robots, timeDelta, time-timeDelta),
			time-timeDelta,
			score,
		))

		if canBuildRobot(supplies, blueprint.geode) {
			return score
		}
	}

	//Build obsidian robot
	if robots.clay > 0 {
		timeDelta := 1
		if !canBuildRobot(supplies, blueprint.obsidian) {
			timeDelta += int(math.Max(
				math.Ceil(float64(blueprint.obsidian.ore-supplies.ore)/float64(robots.ore)),
				math.Ceil(float64(blueprint.obsidian.clay-supplies.clay)/float64(robots.clay)),
			))
		}

		if time-timeDelta > 2 {
			score = lib.Max(score, dfs(
				blueprint,
				robotLimits,
				Materials{
					ore:      robots.ore,
					clay:     robots.clay,
					obsidian: robots.obsidian + 1,
					geode:    robots.geode,
				},
				buildRobot(supplies, blueprint.obsidian, robots, timeDelta, 0),
				time-timeDelta,
				score,
			))
		}
	}

	//Build clay robot
	if robots.clay < robotLimits.clay {
		timeDelta := 1
		if !canBuildRobot(supplies, blueprint.clay) {
			timeDelta += int(math.Ceil(float64(blueprint.clay.ore-supplies.ore) / float64(robots.ore)))
		}

		if time-timeDelta > 3 {
			score = lib.Max(score, dfs(
				blueprint,
				robotLimits,
				Materials{
					ore:      robots.ore,
					clay:     robots.clay + 1,
					obsidian: robots.obsidian,
					geode:    robots.geode,
				},
				buildRobot(supplies, blueprint.clay, robots, timeDelta, 0),
				time-timeDelta,
				score,
			))
		}
	}

	//Build ore robot
	if robots.ore < robotLimits.ore {
		timeDelta := 1
		if !canBuildRobot(supplies, blueprint.ore) {
			timeDelta += int(math.Ceil(float64(blueprint.ore.ore-supplies.ore) / float64(robots.ore)))
		}

		if time-timeDelta > 4 {
			score = lib.Max(score, dfs(
				blueprint,
				robotLimits,
				Materials{
					ore:      robots.ore + 1,
					clay:     robots.clay,
					obsidian: robots.obsidian,
					geode:    robots.geode,
				},
				buildRobot(supplies, blueprint.ore, robots, timeDelta, 0),
				time-timeDelta,
				score,
			))
		}
	}
	return score
}

func evaluateBlueprint(blueprint Blueprint, time int) int {
	robotLimits := Materials{
		ore:      lib.Max(blueprint.ore.ore, blueprint.clay.ore, blueprint.obsidian.ore, blueprint.geode.ore),
		clay:     blueprint.obsidian.clay,
		obsidian: blueprint.geode.obsidian,
		geode:    1000,
	}
	return dfs(blueprint, robotLimits, Materials{ore: 1}, Materials{}, time, 0)
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	blueprints := make([]Blueprint, len(input))
	for i, l := range input {
		bp := Blueprint{}
		fmt.Sscanf(l, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &bp.index, &bp.ore.ore, &bp.clay.ore, &bp.obsidian.ore, &bp.obsidian.clay, &bp.geode.ore, &bp.geode.obsidian)
		blueprints[i] = bp
	}

	solution2.I = 1
	for i, blueprint := range blueprints {
		solution1.I += blueprint.index * evaluateBlueprint(blueprint, 24)
		if i < 3 {
			solution2.I *= evaluateBlueprint(blueprint, 32)
		}
	}
	return
}
