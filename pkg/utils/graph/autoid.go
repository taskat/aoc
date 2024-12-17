package graph

var nextId int = 1

func NextId() int {
	id := nextId
	nextId++
	return id
}
