package aoc

func SliceMap[K, V any](ks []K, f func(K) V) []V {
	if len(ks) == 0 {
		return nil
	}
	vs := make([]V, 0, len(ks))
	for _, k := range ks {
		vs = append(vs, f(k))
	}
	return vs
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Do calls the given function n times.
func Do(n int, f func()) {
	for i := 0; i < n; i++ {
		f()
	}
}

type Set[V comparable] map[V]struct{}

func NewSet[V comparable]() *Set[V] {
	m := map[V]struct{}{}
	s := Set[V](m)
	return &s
}

func (s *Set[V]) Add(v V) {
	(*s)[v] = struct{}{}
}

func (s *Set[V]) Contains(v V) bool {
	_, ok := (*s)[v]
	return ok
}
