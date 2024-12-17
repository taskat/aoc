package graph

import (
	"strings"
)

type Graph[ID comparable] struct {
	nodes map[ID]Node[ID]
}

func New[ID comparable]() *Graph[ID] {
	return &Graph[ID]{
		nodes: make(map[ID]Node[ID]),
	}
}

func (g *Graph[ID]) AddNode(node Node[ID]) {
	if g.HasNode(node.Id()) {
		return
	}
	g.nodes[node.Id()] = node
}

func (g *Graph[ID]) AddDirectedEdge(from, to ID, weight int) {
	g.nodes[from].AddNeighbor(g.nodes[to], weight)
}

func (g *Graph[ID]) AddEdge(a, b ID, weight int) {
	g.nodes[a].AddNeighbor(g.nodes[b], weight)
	g.nodes[b].AddNeighbor(g.nodes[a], weight)
}

func (g *Graph[ID]) GetNode(id ID) Node[ID] {
	return g.nodes[id]
}

func (g *Graph[ID]) GetNodes() map[ID]Node[ID] {
	return g.nodes
}

func (g *Graph[ID]) HasNode(id ID) bool {
	_, ok := g.nodes[id]
	return ok
}

// String returns a string representation of the graph
func (g *Graph[ID]) String() string {
	nodes := make([]string, 0, len(g.nodes))
	for _, node := range g.nodes {
		nodes = append(nodes, node.String())
	}
	return strings.Join(nodes, "\n")
}
