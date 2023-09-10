package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDay1 to make sure that the code will not be broken by changes in any shared code.
func TestDay1(t *testing.T) {
	t.Run("Part1 returns the same result", func(t *testing.T) {
		result, err := Part1()
		require.NoError(t, err)
		assert.Equal(t, "262", result)
	})

	t.Run("Part2 returns the same result", func(t *testing.T) {
		result, err := Part2()
		require.NoError(t, err)
		assert.Equal(t, "131", result)
	})
}
