package aoc_test

import (
	"testing"

	"github.com/espang/aoc/aoc"
	"github.com/stretchr/testify/assert"
)

func TestSliceDrop(t *testing.T) {
	var emptyInts []int
	assert.Equal(t, emptyInts, aoc.SliceDrop(2, []int{1}))
	assert.Equal(t, emptyInts, aoc.SliceDrop(2, []int{1, 2}))
	assert.Equal(t, []int{3}, aoc.SliceDrop(2, []int{1, 2, 3}))
}
