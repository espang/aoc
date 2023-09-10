package day20

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	prev, next *Node
	value      int
}

func Append(to *Node, value int) *Node {
	newNode := &Node{value: value}
	to.next = newNode
	newNode.prev = to
	return newNode
}

func MoveAfter(toMove *Node, pos *Node) {
	if toMove.prev == pos || toMove == pos {
		return
	}

	if toMove.next == pos {
		SwapForwad(toMove)
		return
	}

	prev := toMove.prev
	next := toMove.next

	toMove.next = pos.next
	toMove.prev = pos

	pos.next.prev = toMove
	pos.next = toMove

	next.prev = prev
	prev.next = next
}

func MoveBefore(toMove *Node, pos *Node) {
	MoveAfter(toMove, pos.prev)
}

func SwapForwad(n *Node) {
	nextNode := n.next
	prevNode := n.prev

	n.prev = nextNode
	n.next = nextNode.next
	nextNode.next.prev = n
	nextNode.next = n
	nextNode.prev = prevNode
	prevNode.next = nextNode
}

func PrintAll(n *Node) {
	values := []string{
		strconv.Itoa(n.value),
	}
	for current := n.next; current != n; current = current.next {
		values = append(values, strconv.Itoa(current.value))
	}
	fmt.Println(strings.Join(values, ", "))
}

func MakeRing(values []int) *Node {
	if len(values) == 0 {
		panic("expect non empty list")
	}
	head := &Node{value: values[0]}
	nodes := []*Node{
		head,
	}
	current := head
	for _, n := range values[1:] {
		nextNode := Append(current, n)
		nodes = append(nodes, nextNode)
		current = nextNode
	}
	// make it a ring by connecting the last with the first node
	head.prev = current
	current.next = head
	return head
}

func AllNodes(head *Node) []*Node {
	nodes := []*Node{}
	for current := head; current.next != head; current = current.next {
		nodes = append(nodes, current)
	}
	nodes = append(nodes, head.prev)
	return nodes
}

func FindNodeByValue(head *Node, val int) *Node {
	for current := head; current.next != head; current = current.next {
		if current.value == val {
			return current
		}
	}
	return nil
}

func Move(amount int, head *Node) *Node {
	current := head
	for i := 0; i < amount; i++ {
		current = current.next
	}
	for i := 0; i >= amount; i-- {
		current = current.prev
	}
	return current
}
