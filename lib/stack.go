package lib

type Stack[T any] struct {
	Values []T
}

func NewStack[T any](values ...T) Stack[T] {
	return Stack[T]{values}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Size() int {
	return len(s.Values)
}

func (s *Stack[T]) Push(value T) T {
	s.Values = append(s.Values, value)
	return value
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("Pop failed, stack has size 0")
	}
	top := s.Values[len(s.Values)-1]
	s.Values = s.Values[: len(s.Values)-1 : len(s.Values)-1]
	return top
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("Peek failed, stack has size 0")
	}
	return s.Values[len(s.Values)-1]
}
