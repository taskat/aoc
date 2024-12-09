package linkedlist

import "fmt"

// node represents a node in a linked list
type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

// String returns the string representation of the node
func (n *node[T]) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprint(n.value)
}
