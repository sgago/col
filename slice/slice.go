package slice

import (
	"sync"

	"github.com/sgago/collections/err"
)

const notFound int = -1

func First[T any](slice []T) T {
	return slice[0]
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
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

func indexOf[T comparable](slice []T, value T) (int, error) {
	for index, val := range slice {
		if value == val {
			return index, nil
		}
	}

	return notFound, &err.NotFound{}
}

func IndexOf[T comparable](slice []T, value T) (int, error) {
	max := 10_000

	workers := len(slice) / max

	if workers == 0 {
		return indexOf(slice, value)
	}

	indexes := make(chan int, workers)

	for i := 0; i < workers; i++ {
		indexOfWg.Add(1)

		start := i * max
		end := len(slice)

		if i < workers-1 {
			end = start + max
		}

		go func(w int, s []T, result chan<- int) {
			defer indexOfWg.Done()
			index, e := indexOf(s, value)

			if e == nil {
				result <- index + w*max
			} else {
				result <- notFound
			}
		}(i, slice[start:end], indexes)
	}

	indexOfWg.Wait()
	close(indexes)

	result := notFound

	for index := range indexes {
		if index != notFound && (result == notFound || index < result) {
			result = index
		}
	}

	if result != notFound {
		return result, nil
	}

	return notFound, &err.NotFound{}
}
