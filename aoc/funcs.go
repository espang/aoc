package aoc

type Number interface {
	~int | ~float64
}

func SliceDrop[K any](n int, vs []K) []K {
	if n >= len(vs) {
		return nil
	}
	return vs[n:]
}

func SliceSum[N Number](vs []N) N {
	var total N
	for _, v := range vs {
		total += v
	}
	return total
}

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

func Map[A any, B any](f func(input A) B, vs []A) []B {
	ret := make([]B, 0, len(vs))
	for _, v := range vs {
		ret = append(ret, f(v))
	}
	return ret
}

// Run executes 'f' for each element of 'vs' for the side effects of f.
func Run[A any](f func(A), vs []A) {
	for _, v := range vs {
		f(v)
	}
}

func MapTranspose[K, V comparable](m map[K]V) map[V]K {
	result := map[V]K{}
	for k, v := range m {
		result[v] = k
	}
	return result
}

func MapValues[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
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
