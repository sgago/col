package slice

import (
	"sync"

	"github.com/sgago/col/err"
)

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

func First[T any](slice []T, predicate func(index int, value T) bool) (T, error) {
	if len(slice) == 0 {
		panic("The slice is empty.")
	}

	if predicate == nil {
		return slice[0], nil
	}

	for index, value := range slice {
		if predicate(index, value) {
			return value, nil
		}
	}

	var notFoundValue T

	return notFoundValue, &err.NotFound{}
}

func Last[T any](slice []T, predicate func(index int, value T) bool) (T, error) {
	if len(slice) == 0 {
		panic("The slice is empty.")
	}

	if predicate == nil {
		return slice[len(slice)-1], nil
	}

	for index, value := range slice {
		if predicate(index, value) {
			return value, nil
		}
	}

	var notFoundValue T

	return notFoundValue, &err.NotFound{}
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

var indexOfWg sync.WaitGroup

func indexOfWorker[T comparable](slice []T, value T, start int, end int) (int, error) {
	for index, val := range slice[start:end] {
		if value == val {
			return index + start, nil
		}
	}

	return NotFound, &err.NotFound{}
}

func IndexOf[T comparable](slice []T, value T) (int, error) {
	max := maxElems

	workers := len(slice) / max

	if workers > maxWorkers {
		workers = maxWorkers
	}

	if workers == 0 {
		return indexOfWorker(slice, value, 0, len(slice))
	}

	indexes := make(chan int, workers)

	for i := 0; i < workers; i++ {
		indexOfWg.Add(1)

		start := i * max
		end := len(slice)

		if i < workers-1 {
			end = start + max
		}

		go func(s []T, start int, end int, result chan<- int) {
			defer indexOfWg.Done()
			index, e := indexOfWorker(s, value, start, end)

			if e == nil {
				result <- index
			} else {
				result <- NotFound
			}
		}(slice, start, end, indexes)
	}

	indexOfWg.Wait()
	close(indexes)

	result := NotFound

	for index := range indexes {
		if index != NotFound && (result == NotFound || index < result) {
			result = index
		}
	}

	if result != NotFound {
		return result, nil
	}

	return NotFound, &err.NotFound{}
}
