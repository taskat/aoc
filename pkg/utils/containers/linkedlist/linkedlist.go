package linkedlist

import (
	"fmt"
	"strings"
)

// LinkedList represents a linked list
type LinkedList[T any] struct {
	first  *node[T]
	last   *node[T]
	length int
}

// New creates a new linked list
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// FromSlice creates a linked list from a slice
func FromSlice[T any](values []T) *LinkedList[T] {
	ll := New[T]()
	for _, v := range values {
		ll.Insert(v, ll.length)
	}
	return ll
}

// Clear clears the list
func (l *LinkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

// Get returns the value at the given index. If the index is out of bounds, it
// panics.
func (l *LinkedList[T]) Get(index int) T {
	return l.getNode(index).value
}

// getNode returns the node at the given index. If the index is out of bounds,
// it panics.
func (l *LinkedList[T]) getNode(index int) *node[T] {
	if index < 0 || index >= l.length {
		panic("index out of bounds")
	}
	if index < l.length/2 {
		current := l.first
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current
	}
	current := l.last
	for i := l.length - 1; i > index; i-- {
		current = current.prev
	}
	return current
}

// Insert inserts a value at the given index. If the index is out of bounds,
// it panics.
func (l *LinkedList[T]) Insert(value T, index int) {
	if index < 0 || index > l.length {
		panic("index out of bounds")
	}
	node := &node[T]{value: value}
	if l.length == 0 {
		l.first = node
		l.last = node
	} else if index == 0 {
		node.next = l.first
		l.first.prev = node
		l.first = node
	} else if index == l.length {
		node.prev = l.last
		l.last.next = node
		l.last = node
	} else {
		current := l.getNode(index)
		node.prev = current.prev
		node.next = current
		current.prev.next = node
		current.prev = node
	}
	l.length++
}

// InsertFirst inserts a value at the beginning of the list
func (l *LinkedList[T]) InsertFirst(value T) {
	l.Insert(value, 0)
}

// InsertLast inserts a value at the end of the list
func (l *LinkedList[T]) InsertLast(value T) {
	l.Insert(value, l.length)
}

// Length returns the length of the list
func (l *LinkedList[T]) Length() int {
	return l.length
}

// Remove removes the value at the given index. If the index is out of bounds,
// it panics.
func (l *LinkedList[T]) Remove(index int) {
	if index < 0 || index >= l.length {
		panic("index out of bounds")
	}
	if l.length == 1 {
		l.first = nil
		l.last = nil
	} else if index == 0 {
		l.first = l.first.next
		l.first.prev = nil
	} else if index == l.length-1 {
		l.last = l.last.prev
		l.last.next = nil
	} else {
		current := l.getNode(index)
		current.prev.next = current.next
		current.next.prev = current.prev
	}
	l.length--
}

// RemoveFirst removes the first value
func (l *LinkedList[T]) RemoveFirst() {
	l.Remove(0)
}

// RemoveLast removes the last value
func (l *LinkedList[T]) RemoveLast() {
	l.Remove(l.length - 1)
}

// Set sets the value at the given index. If the index is out of bounds, it
// panics.
func (l *LinkedList[T]) Set(index int, value T) {
	l.getNode(index).value = value
}

// String returns the string representation of the list
func (l *LinkedList[T]) String() string {
	current := l.first
	var str strings.Builder
	for i := 0; i < l.length; i++ {
		fmt.Fprint(&str, current)
		if i < l.length-1 {
			fmt.Fprint(&str, " -> ")
		}
		current = current.next
	}
	return str.String()
}

// ToSlice returns the list as a slice
func (l *LinkedList[T]) ToSlice() []T {
	slice := make([]T, l.length)
	current := l.first
	for i := 0; i < l.length; i++ {
		slice[i] = current.value
		current = current.next
	}
	return slice
}
