package aoc

type List[V any] struct {
	Head *Node[V]
}

type Node[V any] struct {
	prev  *Node[V]
	next  *Node[V]
	value V
}

func (n *Node[V]) Prev() *Node[V] {
	if n == nil {
		return nil
	}
	return n.prev
}

func (n *Node[V]) Next() *Node[V] {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *Node[V]) Value() V {
	if n == nil {
		return *new(V)
	}
	return n.value
}
