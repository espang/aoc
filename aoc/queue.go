package aoc

import "container/list"

type Queue[V any] struct {
	l *list.List
}

func NewQueue[V any](vs ...V) *Queue[V] {
	q := &Queue[V]{
		l: list.New(),
	}
	for _, v := range vs {
		q.Push(v)
	}
	return q
}

func (q *Queue[V]) NotEmpty() bool {
	return q.l.Len() > 0
}

func (q *Queue[V]) Push(v V) {
	q.l.PushBack(v)
}

func (q *Queue[V]) Pop() V {
	return q.l.Remove(q.l.Front()).(V)
}
