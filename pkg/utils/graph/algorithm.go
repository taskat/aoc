package graph

import (
	"math"

	"github.com/taskat/aoc/pkg/utils/containers/set"
)

// NodesOfBestPaths finds all the shortest paths from the start node to the end node
func (g Graph[ID]) NodesOfBestPaths(start Node[ID], goal func(ID) bool) []ID {
	if goal(start.Id()) {
		return []ID{start.Id()}
	}
	distances := make(map[ID]int)
	for id := range start.GetNeighbors() {
		distances[id] = math.MaxInt
	}
	distances[start.Id()] = 0
	previous := make(map[ID]ID)
	canReach := make(map[ID][]ID)
	neighbors := make(map[ID]int)
	for id, distance := range start.GetNeighbors() {
		neighbors[id] = distance
		previous[id] = start.Id()
		canReach[id] = []ID{start.Id()}
	}
	visited := set.New[ID]()
	visited.Add(start.Id())
	var goalId ID
	goalFound := false
	for len(neighbors) > 0 {
		var minNodeId ID
		minDistance := math.MaxInt
		for id, distance := range neighbors {
			newDistance := distances[previous[id]] + distance
			if newDistance < minDistance {
				minNodeId = id
				minDistance = newDistance
			}
		}
		if goal(minNodeId) && !goalFound {
			goalId = minNodeId
			goalFound = true
		}
		visited.Add(minNodeId)
		distances[minNodeId] = minDistance
		for node, distance := range g.GetNode(minNodeId).GetNeighbors() {
			if !visited.Contains(node) {
				neighbors[node] = distance
				previous[node] = minNodeId
				canReach[node] = append(canReach[node], minNodeId)
			}
			if visited.Contains(node) && distances[node] == distances[minNodeId]+distance {
				canReach[node] = append(canReach[node], minNodeId)
			}
		}
		delete(neighbors, minNodeId)
	}
	nodes := set.New[ID]()
	nodes.Add(goalId)
	previousNodes := canReach[goalId]
	for len(previousNodes) > 0 {
		currentNode := previousNodes[0]
		previousNodes = previousNodes[1:]
		if nodes.Contains(currentNode) {
			continue
		}
		nodes.Add(currentNode)
		if currentNode == start.Id() {
			continue
		}
		currentPrevs := canReach[currentNode]
		for _, prev := range currentPrevs {
			previousNodes = append(previousNodes, prev)
		}
	}
	return nodes.ToSlice()
}

// Dijkstra finds the shortest path from the start node to the end node
func (g Graph[ID]) Dijkstra(start Node[ID], goal func(ID) bool) Path[ID] {
	if goal(start.Id()) {
		return NewPath([]ID{start.Id()}, 0)
	}
	distances := make(map[ID]int)
	for id := range start.GetNeighbors() {
		distances[id] = math.MaxInt
	}
	distances[start.Id()] = 0
	previous := make(map[ID]ID)
	neighbors := make(map[ID]int)
	for id, distance := range start.GetNeighbors() {
		neighbors[id] = distance
		previous[id] = start.Id()
	}
	visited := set.New[ID]()
	visited.Add(start.Id())
	for len(neighbors) > 0 {
		var minNodeId ID
		minDistance := math.MaxInt
		for id, distance := range neighbors {
			newDistance := distances[previous[id]] + distance
			if newDistance < minDistance {
				minNodeId = id
				minDistance = newDistance
			}
		}
		if goal(minNodeId) {
			nodes := []ID{minNodeId}
			for previousNode := previous[minNodeId]; previousNode != start.Id(); previousNode = previous[previousNode] {
				nodes = append([]ID{previousNode}, nodes...)
			}
			nodes = append([]ID{start.Id()}, nodes...)
			return NewPath(nodes, minDistance)
		}
		visited.Add(minNodeId)
		distances[minNodeId] = minDistance
		for node, distance := range g.GetNode(minNodeId).GetNeighbors() {
			if !visited.Contains(node) {
				neighbors[node] = distance
				previous[node] = minNodeId
			}
		}
		delete(neighbors, minNodeId)
	}
	return NoPath[ID]()
}
