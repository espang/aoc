package aoc

func reduce[A any, B any](f func(acc A, val B) A, initial A, vs []B) A {
	ret := initial
	for _, v := range vs {
		ret = f(ret, v)
	}
	return ret
}

func Map[A any, B any](f func(input A) B, vs []A) []B {
	ret := make([]B, 0, len(vs))
	for _, v := range vs {
		ret = append(ret, f(v))
	}
	return ret
}

func plus(v1, v2 int) int { return v1 + v2 }

func Sum(vs []int) int {
	return reduce(plus, 0, vs)
}
