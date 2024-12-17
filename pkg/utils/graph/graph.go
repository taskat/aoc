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

type Path[ID comparable] struct {
	nodes []ID
	cost  int
}

func NewPath[ID comparable](nodes []ID, cost int) Path[ID] {
	return Path[ID]{nodes, cost}
}

func NoPath[ID comparable]() Path[ID] {
	return Path[ID]{nil, -1}
}

func (p *Path[ID]) AddNode(node ID, cost int) {
	p.nodes = append(p.nodes, node)
	p.cost += cost
}

func (p *Path[ID]) AddFirstNode(node ID, cost int) {
	p.nodes = append([]ID{node}, p.nodes...)
	p.cost += cost
}

func (p Path[ID]) Copy() Path[ID] {
	nodes := make([]ID, len(p.nodes))
	copy(nodes, p.nodes)
	return Path[ID]{nodes, p.cost}
}

func (p Path[ID]) Cost() int {
	return p.cost
}

func (p Path[ID]) LastNode() ID {
	return p.nodes[len(p.nodes)-1]
}

func (p Path[ID]) Nodes() []ID {
	return p.nodes
}
