package graph

import (
	"math"

	"github.com/taskat/aoc/pkg/utils/containers/set"
)

// Dijkstra finds the shortest path from the start node to the end node
func (g Graph[ID]) Dijkstra(start Node[ID], goal func(ID) bool) ([]ID, int) {
	if goal(start.Id()) {
		return []ID{start.Id()}, 0
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
			path := []ID{minNodeId}
			for previousNode := previous[minNodeId]; previousNode != start.Id(); previousNode = previous[previousNode] {
				path = append([]ID{previousNode}, path...)
			}
			path = append([]ID{start.Id()}, path...)
			return path, minDistance
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
	return nil, -1
}
