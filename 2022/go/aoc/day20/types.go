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

func SwapBackward(n *Node) {
	nextNode := n.next
	prevNode := n.prev
	n.prev = prevNode.prev
	n.next = prevNode
	prevNode.prev.next = n
	prevNode.prev = n
	prevNode.next = nextNode
	nextNode.prev = prevNode
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
