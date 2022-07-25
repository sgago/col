// Package slice provides go-routine-backed
// functions for searching and modifying unsorted slices.
package slice

const (
	// The not found index value.
	NotFound int = -1

	// The default maximum number of worker values.
	DefaultMaxSearchWorkers int = 4

	// The default maximum number of search elements
	// before creating a new worker.
	DefaultMaxSearchLength int = 100_000
)

var (
	maxElems   = DefaultMaxSearchLength
	maxWorkers = DefaultMaxSearchWorkers
)

func GetMaxSearchLength() int {
	return maxElems
}

func SetMaxSearchLength(length int) {
	if length > 0 {
		maxElems = length
	} else {
		maxElems = DefaultMaxSearchLength
	}
}

func GetMaxSearchWorkers() int {
	return maxWorkers
}

func SetMaxSearchWorkers(workers int) {
	if workers > 0 {
		maxWorkers = workers
	} else {
		maxWorkers = DefaultMaxSearchWorkers
	}
}

// Clear removes all elements from a slice
// and maintains the slice's current capacity.
func Clear[T any](slice []T) []T {
	return make([]T, 0, cap(slice))
}

func RemoveFirst[T any](slice []T) []T {
	return slice[1:]
}

func RemoveLast[T any](slice []T) []T {
	return slice[:len(slice)-1]
}

// Swap flips the element values at indexA and indexB.
func Swap[T any](slice []T, indexA int, indexB int) []T {
	temp := slice[indexA]
	slice[indexA] = slice[indexB]
	slice[indexB] = temp

	return slice
}

func Contains[T comparable](slice []T, value T) bool {
	_, e := IndexOf(slice, value)

	return e == nil
}

func Any[T any](slice []T, predicate func(index int, value T) bool) bool {
	if predicate == nil {
		return len(slice) > 0
	}

	for index, value := range slice {
		if predicate(index, value) {
			return true
		}
	}

	return false
}

func All[T any](slice []T, predicate func(index int, value T) bool) bool {
	if predicate == nil {
		panic("The predicate cannot be nil.")
	}

	for index, value := range slice {
		if !predicate(index, value) {
			return false
		}
	}

	return true
}
