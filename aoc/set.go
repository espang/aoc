package aoc

type Set[V comparable] map[V]struct{}

func EmptySet[V comparable]() *Set[V] {
	return NewSet[V]()
}

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

func (s *Set[V]) ToList() []V {
	if s.Len() == 0 {
		return nil
	}
	vs := make([]V, 0, s.Len())
	for k := range *s {
		vs = append(vs, k)
	}
	return vs
}
