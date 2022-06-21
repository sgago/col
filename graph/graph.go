package graph

type graph[T any] struct {
	vertices [][]T
}

func NewGraph[T any](vertices [][]T) graph[T] {
	return graph[T]{vertices: vertices}
}

// func (g *graph[T]) getNeighbors(value T) []T {
// 	//return g.vertices[value]
// }
