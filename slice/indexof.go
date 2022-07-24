package slice

import (
	"sync"

	"github.com/sgago/col/err"
)

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

	if workers < 2 {
		return indexOfWorker(slice, value, 0, len(slice))
	}

	indexes := make(chan int, workers)
	wg := new(sync.WaitGroup)

	for i := 0; i < workers; i++ {
		wg.Add(1)

		start := i * max
		end := len(slice)

		if i < workers-1 {
			end = start + max
		}

		go func(s []T, start int, end int, result chan<- int) {
			defer wg.Done()
			index, e := indexOfWorker(s, value, start, end)

			if e == nil {
				result <- index
			} else {
				result <- NotFound
			}
		}(slice, start, end, indexes)
	}

	wg.Wait()
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
