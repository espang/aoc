package day20

import (
	"testing"

	"github.com/espang/aoc/aoc"
	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	t.Run("Move after", func(t *testing.T) {
		node1 := &Node{value: 1}
		node2 := Append(node1, 2)
		node3 := Append(node2, 3)
		node4 := Append(node3, 4)
		node1.prev = node4
		node4.next = node1

		assert.Equal(t, []int{1, 2, 3, 4}, aoc.Map(func(n *Node) int { return n.value }, AllNodes(node1)))

		MoveAfter(node1, node3)
		assert.Equal(t, []int{1, 4, 2, 3}, aoc.Map(func(n *Node) int { return n.value }, AllNodes(node1)))
	})
	t.Run("Move after 2", func(t *testing.T) {
		node1 := &Node{value: 1}
		node2 := Append(node1, 2)
		node3 := Append(node2, 3)
		node1.prev = node3
		node3.next = node1

		assert.Equal(t, []int{1, 2, 3}, aoc.Map(func(n *Node) int { return n.value }, AllNodes(node1)))

		MoveAfter(node1, node3)
		assert.Equal(t, []int{1, 2, 3}, aoc.Map(func(n *Node) int { return n.value }, AllNodes(node1)))
		MoveAfter(node1, node2)
		assert.Equal(t, []int{1, 3, 2}, aoc.Map(func(n *Node) int { return n.value }, AllNodes(node1)))
	})
}
