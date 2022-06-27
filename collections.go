package col

// A KeyValue data structure.
type KV[K any, T any] struct {
	Key K
	Val T
}

// A PriorityValue data structure.
type PV[T any] struct {
	Priority int
	Val      T
}
