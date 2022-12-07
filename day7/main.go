package day7

import (
	"aoc-2022/lib"
	"fmt"
)

type Node struct {
	content map[string]*Node
	size    int
}

func traverse(n *Node, s *lib.Solution) (size int) {
	size = n.size
	for _, v := range n.content {
		nodeSize := traverse(v, s)
		size += nodeSize
	}
	if size <= 100000 {
		s.I += size
	}
	return size
}

// TODO solution2 rename
func Process(input []string) (solution1 lib.Solution, solution2 lib.Solution) {
	root := Node{content: make(map[string]*Node)}
	s := Stack[*Node]{values: []*Node{&root}}
	for _, v := range input {
		switch {
		case v[:4] == "dir ":
			s.Peek().content[v[4:]] = &Node{content: make(map[string]*Node)}
		case v == "$ ls" || v == "$ cd /":
			continue
		case v == "$ cd ..":
			s.Pop()
		case v[:5] == "$ cd ":
			s.Push(s.Peek().content[v[5:]])
		default:
			var size int
			var name string
			fmt.Sscanf(v, "%d %s", &size, &name)
			s.Peek().size += size
		}
	}

	traverse(&root, &solution1)
	return
}

// TODO make stack util
type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() T {
	top := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return top
}

func (s *Stack[T]) Peek() T {
	return s.values[len(s.values)-1]
}
