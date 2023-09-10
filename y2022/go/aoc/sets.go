package aoc

type Set[K comparable] map[K]struct{}

func (s *Set[K]) Add(element K) Set[K] {
	(*s)[element] = struct{}{}
	return *s
}

func (s Set[K]) IsMember(element K) bool {
	_, ok := s[element]
	return ok
}

func (s Set[K]) Union(s2 Set[K]) Set[K] {
	for e := range s2 {
		s = s.Add(e)
	}
	return s
}

func (s Set[K]) Remove(element K) (bool, Set[K]) {
	if s.IsMember(element) {
		delete(s, element)
		return true, s
	}
	return false, s
}
