package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput = `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`

func TestDay1(t *testing.T) {
	t.Run("Part1 returns the same result", func(t *testing.T) {
		result, err := part1(testInput, 2)
		require.NoError(t, err)
		assert.Equal(t, "12", result)
	})
	t.Run("Part1 returns the same result", func(t *testing.T) {
		result, err := Part1()
		require.NoError(t, err)
		assert.Equal(t, "262", result)
	})
}

func TestHash(t *testing.T) {
	for _, tc := range []struct {
		l        string
		vs       []int
		expected int
	}{
		{"empty is 0", []int{0, 0, 0, 0}, 0},
		{"#... is 1", []int{1, 0, 0, 0}, 1},
		{"...# is 8", []int{0, 0, 0, 1}, 8},
		{"#..# is 9", []int{1, 0, 0, 1}, 9},
		{"empty is 0", []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{"#........ is 1", []int{1, 0, 0, 0, 0, 0, 0, 0, 0}, 1},
		{"#......## is 385", []int{1, 0, 0, 0, 0, 0, 0, 1, 1}, 385},
	} {
		t.Run(tc.l, func(t *testing.T) {
			actual := hash(tc.vs)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
