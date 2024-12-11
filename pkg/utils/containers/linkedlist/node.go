package linkedlist

import "fmt"

type Node[T any] interface {
	IsFirst() bool
	IsLast() bool
	SetValue(T)
	Value() T
	fmt.Stringer
}

// node represents a node in a linked list
type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

// IsFirst returns true if the node is the first node in the list
func (n *node[T]) IsFirst() bool {
	return n.prev == nil
}

// IsLast returns true if the node is the last node in the list
func (n *node[T]) IsLast() bool {
	return n.next == nil
}

// SetValue sets the value of the node
func (n *node[T]) SetValue(value T) {
	n.value = value
}

// String returns the string representation of the node
func (n *node[T]) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprint(n.value)
}

// Value returns the value of the node
func (n *node[T]) Value() T {
	return n.value
}
