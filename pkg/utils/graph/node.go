package graph

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/pkg/utils/maps"
)

type Node[T comparable] interface {
	AddNeighbor(neighbor Node[T], weight int)
	Id() T
	GetNeighbors() map[T]int
	fmt.Stringer
}

type BaseNode[T comparable] struct {
	id        T
	neighbors map[T]int
}

func NewBaseNode[T comparable](id T) *BaseNode[T] {
	return &BaseNode[T]{
		id: id,
	}
}

func NewBaseNodeAutoId() *BaseNode[int] {
	return &BaseNode[int]{
		id: NextId(),
	}
}

func (n *BaseNode[T]) Id() T {
	return n.id
}

func (n *BaseNode[T]) AddNeighbor(neighbor Node[T], weight int) {
	if n.neighbors == nil {
		n.neighbors = make(map[T]int)
	}
	n.neighbors[neighbor.Id()] = weight
}

func (n *BaseNode[T]) GetNeighbors() map[T]int {
	return n.neighbors
}

// String returns a string representation of the node
func (n *BaseNode[T]) String() string {
	neighbors := maps.ToSlice(n.neighbors, func(k T, v int) string {
		return fmt.Sprintf("  %v: %v", k, v)
	})
	return fmt.Sprintf("Node %v:\n%s", n.id, strings.Join(neighbors, "\n"))
}
