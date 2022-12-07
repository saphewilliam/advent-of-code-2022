package day7

import (
	"aoc-2022/lib"
	"strconv"
	"strings"
)

type Dir struct {
	content []*Dir
	size    int
}

func getDirSize(dir *Dir) (size int) {
	size = dir.size
	for _, v := range dir.content {
		size += getDirSize(v)
	}
	dir.size = size
	return size
}

func traverse(dir *Dir, s1, s2 *lib.Solution, target int) {
	if dir.size <= 100000 {
		s1.I += dir.size
	}
	if dir.size-target >= 0 && dir.size < s2.I {
		s2.I = dir.size
	}
	for _, v := range dir.content {
		traverse(v, s1, s2, target)
	}
}

func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	root := Dir{}
	s := lib.NewStack(&root)
	for _, v := range input {
		switch {
		case v == "$ ls" || v == "$ cd /" || v[:4] == "dir ":
			continue
		case v == "$ cd ..":
			s.Pop()
		case v[:5] == "$ cd ":
			s.Peek().content = append(s.Peek().content, s.Push(&Dir{}))
		default:
			size, _ := strconv.Atoi(strings.Split(v, " ")[0])
			s.Peek().size += size
		}
	}

	rootSize := getDirSize(&root)
	solution2 = lib.IScore(rootSize)
	traverse(&root, &solution1, &solution2, rootSize-40000000)
	return
}
