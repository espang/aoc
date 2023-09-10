package aoc_test

import (
	"testing"

	"github.com/espang/aoc/aoc"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	set := aoc.NewSet(2, 3, 4, 5)
	assert.False(t, set.Contains(1))

	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(4))

	assert.Equal(t, 5, set.Len())
	set.Add(4)
	assert.Equal(t, 5, set.Len())
}
