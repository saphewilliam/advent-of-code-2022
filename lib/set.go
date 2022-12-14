package lib

type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable](elements ...T) Set[T] {
	s := Set[T]{}
	s.values = make(map[T]struct{})
	s.Add(elements...)
	return s
}

func SetUnion[T comparable](s, t Set[T]) Set[T] {
	unionSet := NewSet(s.Elements()...)
	unionSet.Add(t.Elements()...)
	return unionSet
}

func SetIntersection[T comparable](s, t Set[T]) Set[T] {
	intersectionSet := NewSet[T]()
	for _, element := range s.Elements() {
		if t.Has(element) {
			intersectionSet.Add(element)
		}
	}
	return intersectionSet
}

func SetDifference[T comparable](s, t Set[T]) Set[T] {
	differenceSet := NewSet[T]()
	for _, element := range s.Elements() {
		if !t.Has(element) {
			differenceSet.Add(element)
		}
	}
	for _, element := range t.Elements() {
		if !s.Has(element) {
			differenceSet.Add(element)
		}
	}
	return differenceSet
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, len(s.values))
	i := 0
	for k := range s.values {
		elements[i] = k
		i++
	}
	return elements
}

func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s.values[element] = struct{}{}
	}
}

func (s *Set[T]) Remove(elements ...T) {
	for _, element := range elements {
		delete(s.values, element)
	}
}

func (s *Set[T]) Has(element T) bool {
	_, ok := s.values[element]
	return ok
}

func (s *Set[T]) HasSubset(t Set[T]) bool {
	for _, element := range t.Elements() {
		if !s.Has(element) {
			return false
		}
	}
	return true
}
