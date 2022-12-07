package lib

type Queue[T any] struct {
	Values []T
}

func NewQueue[T any](values ...T) Queue[T] {
	return Queue[T]{values}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Size() int {
	return len(q.Values)
}

func (q *Queue[T]) Enqueue(value T) {
	q.Values = append(q.Values, value)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("Dequeue failed, queue has size 0")
	}
	top := q.Values[0]
	q.Values = q.Values[1:]
	return top
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		panic("Peek failed, queue has size 0")
	}
	return q.Values[0]
}
