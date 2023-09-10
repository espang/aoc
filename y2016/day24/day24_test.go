package day24

import (
	"testing"

	"github.com/espang/aoc/aoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

func TestParseInput(t *testing.T) {
	walls, numbers := parseInput(testInput)
	assert.Equal(t, Walls{
		{true, true, true, true, true, true, true, true, true, true, true},
		{true, false, false, false, false, false, false, false, false, false, true},
		{true, false, true, true, true, true, true, true, true, false, true},
		{true, false, false, false, false, false, false, false, false, false, true},
		{true, true, true, true, true, true, true, true, true, true, true},
	}, walls)
	assert.Equal(t, map[int]aoc.Coordinate{
		0: {X: 1, Y: 1},
		1: {X: 3, Y: 1},
		2: {X: 9, Y: 1},
		3: {X: 9, Y: 3},
		4: {X: 1, Y: 3},
	}, numbers)

	moves := potentialMoves(walls, aoc.Coordinate{X: 3, Y: 1})
	assert.ElementsMatch(t, []aoc.Coordinate{
		{X: 4, Y: 1},
		{X: 2, Y: 1},
	}, moves)
}

func TestSee(t *testing.T) {
	seen := 0

	//looking for 0, 1, 2 and 3:
	search := setBit(setBit(setBit(setBit(0, 0), 1), 2), 3)
	assert.False(t, seen == search)
	seen = setBit(seen, 1)
	assert.False(t, seen == search)
	seen = setBit(seen, 1)
	assert.False(t, seen == search)
	seen = setBit(seen, 2)
	assert.False(t, seen == search)
	seen = setBit(seen, 3)
	assert.False(t, seen == search)
	seen = setBit(seen, 0)
	assert.True(t, seen == search)
}

func TestDay24(t *testing.T) {
	t.Run("part1 with test input returns the same result", func(t *testing.T) {
		result, err := part1(testInput)
		require.NoError(t, err)
		assert.Equal(t, "14", result)
	})
	t.Run("Part1 returns the same result", func(t *testing.T) {
		result, err := Part1()
		require.NoError(t, err)
		assert.Equal(t, "428", result)
	})
	t.Run("Part2 returns the same result", func(t *testing.T) {
		result, err := Part2()
		require.NoError(t, err)
		assert.Equal(t, "680", result)
	})
}
