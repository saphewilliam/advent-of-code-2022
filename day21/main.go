package day21

import (
	"aoc-2022/lib"
	"fmt"
)

type Monkey struct {
	value      int
	m1, op, m2 string
	me         string
}

func getValue(monkeys map[string]*Monkey, monkey string) int {
	m := monkeys[monkey]
	if monkey == "humn" {
		monkeys[monkey] = &Monkey{me: "me", value: m.value}
	}
	if m.op == "" {
		return m.value
	}

	v1, v2 := getValue(monkeys, m.m1), getValue(monkeys, m.m2)
	if monkeys[m.m1].me != "" {
		monkeys[monkey] = &Monkey{m1: m.m1, op: m.op, m2: m.m2, me: m.m1}
	}
	if monkeys[m.m2].me != "" {
		monkeys[monkey] = &Monkey{m1: m.m1, op: m.op, m2: m.m2, me: m.m2}
	}

	switch m.op {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	}
	return 0
}

func reverseOperations(monkeys map[string]*Monkey, monkey string, target int) int {
	m := monkeys[monkey]

	if m.me == "me" {
		return target
	}

	if monkeys[m.m1].me != "" {
		val := getValue(monkeys, m.m2)
		switch m.op {
		case "+":
			return reverseOperations(monkeys, m.m1, target-val)
		case "-":
			return reverseOperations(monkeys, m.m1, val+target)
		case "*":
			return reverseOperations(monkeys, m.m1, target/val)
		case "/":
			return reverseOperations(monkeys, m.m1, val*target)
		}
	}
	if monkeys[m.m2].me != "" {
		val := getValue(monkeys, m.m1)
		switch m.op {
		case "+":
			return reverseOperations(monkeys, m.m2, target-val)
		case "-":
			return reverseOperations(monkeys, m.m2, val-target)
		case "*":
			return reverseOperations(monkeys, m.m2, target/val)
		case "/":
			return reverseOperations(monkeys, m.m2, val/target)
		}
	}

	return 0
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	monkeys := map[string]*Monkey{}
	for _, l := range input {
		var name string
		m := Monkey{}
		_, err := fmt.Sscanf(l, "%s %d\n", &name, &m.value)
		if err != nil {
			fmt.Sscanf(l, "%s %s %s %s", &name, &m.m1, &m.op, &m.m2)
		}
		monkeys[name[:len(name)-1]] = &m
	}

	solution1.I = getValue(monkeys, "root")
	if monkeys["root"].me != monkeys["root"].m1 {
		value := getValue(monkeys, monkeys["root"].m1)
		solution2.I = reverseOperations(monkeys, monkeys["root"].m2, value)
	} else {
		value := getValue(monkeys, monkeys["root"].m2)
		solution2.I = reverseOperations(monkeys, monkeys["root"].m1, value)
	}

	return
}
