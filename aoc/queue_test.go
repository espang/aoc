package aoc_test

import (
	"testing"

	"github.com/espang/aoc/aoc"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	queue := aoc.NewQueue[int]()
	assert.False(t, queue.NotEmpty())

	queue.Push(1)
	assert.True(t, queue.NotEmpty())

	queue.Push(2)
	assert.True(t, queue.NotEmpty())

	assert.Equal(t, 1, queue.Pop())
	assert.True(t, queue.NotEmpty())
	assert.Equal(t, 2, queue.Pop())
	assert.False(t, queue.NotEmpty())
}
