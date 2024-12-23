package graph

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/pkg/utils/maps"
)

type Node[T comparable] interface {
	AddNeighbor(neighbor Node[T], weight int)
	GetNeighbors() map[T]int
	HasNeighbor(id T) bool
	Id() T
	RemoveNeighbor(id T)
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

func (n *BaseNode[T]) HasNeighbor(id T) bool {
	_, ok := n.neighbors[id]
	return ok
}

func (n *BaseNode[T]) RemoveNeighbor(id T) {
	delete(n.neighbors, id)
}

// String returns a string representation of the node
func (n *BaseNode[T]) String() string {
	neighbors := maps.ToSlice(n.neighbors, func(k T, v int) string {
		return fmt.Sprintf("  %v: %v", k, v)
	})
	return fmt.Sprintf("Node %v:\n%s", n.id, strings.Join(neighbors, "\n"))
}
