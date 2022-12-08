package day7

import (
	"aoc-2022/lib"
	"strconv"
	"strings"
)

type Dir struct {
	parent *Dir
	dirs   []*Dir
	size   int
}

func addFile(dir *Dir, size int) {
	dir.size += size
	if dir.parent != nil {
		addFile(dir.parent, size)
	}
}

func traverse(dir *Dir, s1, s2 *lib.Solution, target int) {
	if dir.size <= 100000 {
		s1.I += dir.size
	}
	if dir.size-target >= 0 && dir.size < s2.I {
		s2.I = dir.size
	}
	for _, v := range dir.dirs {
		traverse(v, s1, s2, target)
	}
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	root := &Dir{}
	currentDir := root
	for _, v := range input {
		switch {
		case v == "$ ls" || v == "$ cd /" || v[:4] == "dir ":
			continue
		case v == "$ cd ..":
			currentDir = currentDir.parent
		case v[:5] == "$ cd ":
			newDir := &Dir{parent: currentDir}
			currentDir.dirs = append(currentDir.dirs, newDir)
			currentDir = newDir
		default:
			size, _ := strconv.Atoi(strings.Split(v, " ")[0])
			addFile(currentDir, size)
		}
	}

	solution2 = lib.IScore(root.size)
	traverse(root, &solution1, &solution2, root.size-40000000)
	return
}
