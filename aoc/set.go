package aoc

type Set[V comparable] map[V]struct{}

func NewSet[V comparable](vs ...V) *Set[V] {
	m := map[V]struct{}{}
	s := Set[V](m)
	for _, v := range vs {
		s.Add(v)
	}
	return &s
}

func (s *Set[V]) Add(v V) {
	(*s)[v] = struct{}{}
}

func (s *Set[V]) Contains(v V) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *Set[V]) Len() int {
	return len(*s)
}
