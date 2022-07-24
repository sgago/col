package slice

const (
	NotFound                int = -1
	DefaultMaxSearchWorkers int = 4
	DefaultMaxSearchLength  int = 100_000
)

var (
	maxElems   = DefaultMaxSearchLength
	maxWorkers = DefaultMaxSearchWorkers
)

func SetMaxSearchLength(length int) {
	if length > 0 {
		maxElems = length
	} else {
		maxElems = DefaultMaxSearchLength
	}
}

func SetMaxSearchWorkers(workers int) {
	if workers > 0 {
		maxWorkers = workers
	} else {
		maxWorkers = DefaultMaxSearchWorkers
	}
}

func Clear[T any](slice []T) []T {
	return make([]T, 0, cap(slice))
}

func RemoveFirst[T any](slice []T) []T {
	return slice[1:]
}

func RemoveLast[T any](slice []T) []T {
	return slice[:len(slice)-1]
}

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
