package aoc

// SetBit takes v and sets the bit at position pos to 1.
// SetBit(0, 0) == 1,
// SetBit(0, 1) == 2,
// SetBit(1, 1) == 3,
// ...
func SetBit(v int, pos int) int {
	v |= 1 << pos
	return v
}
