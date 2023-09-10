package day17

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"slices"

	"github.com/espang/aoc/aoc"
)

var (
	openDoorChars = []byte{'b', 'c', 'd', 'e', 'f'}

	// directions based on the index of the hash - defined by the tasks
	directions = [...]aoc.Direction{aoc.Up, aoc.Down, aoc.Left, aoc.Right}
)

// insideGrid checks if the given coordinate is within the grid. The grid starts
// with 0, 0 at the top left and ends with 3, -3 at the bottem right.
func insideGrid(c aoc.Coordinate) bool {
	return c.X >= 0 && c.X <= 3 && c.Y <= 0 && c.Y >= -3
}

// doorIsOpen checks wether a door is open based on a hex digit.
func doorIsOpen(b byte) bool {
	return slices.Contains(openDoorChars, b)
}

type Hasher struct {
	h   hash.Hash
	buf []byte
}

func NewHasher() *Hasher {
	return &Hasher{
		h:   md5.New(),
		buf: make([]byte, 64),
	}
}

func (h *Hasher) HashOf(state []byte) []byte {
	h.h.Reset()
	h.h.Write(state)
	_ = hex.Encode(h.buf, h.h.Sum(nil))
	return h.buf[:4]
}

func (h *Hasher) DoorStateOf(state []byte) []byte {
	return h.HashOf(state)
}

func OpenDoors(passcode string, state []byte) []aoc.Direction {
	h := NewHasher()
	var dirs []aoc.Direction
	for i, r := range h.HashOf(append([]byte(passcode), state...)) {
		direction := directions[i]
		if doorIsOpen(r) {
			dirs = append(dirs, direction)
		}
	}
	return dirs
}

// followPath follows the given path from a given coordinate.
func followPath(path []byte, start aoc.Coordinate) aoc.Coordinate {
	for _, d := range path {
		switch d {
		case 'U':
			start = aoc.MoveUp(start)
		case 'D':
			start = aoc.MoveDown(start)
		case 'L':
			start = aoc.MoveLeft(start)
		case 'R':
			start = aoc.MoveRight(start)
		default:
			panic("followPath")
		}
	}
	return start
}

// directionToPathElement translates the direction into one of the bytes used
// in the path. One of U|D|L|R.
func directionToPathElement(dir aoc.Direction) byte {
	switch dir {
	case aoc.Up:
		return 'U'
	case aoc.Down:
		return 'D'
	case aoc.Left:
		return 'L'
	case aoc.Right:
		return 'R'
	}
	panic("in directionToPathElement")
}
