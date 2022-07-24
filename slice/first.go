package slice

import (
	"sync"

	"github.com/sgago/col"
	"github.com/sgago/col/err"
)

func firstWorker[T any](slice []T, predicate func(index int, value T) bool, start int, end int) (int, T, error) {
	var defaultType T

	for index, value := range slice[start:end] {
		if predicate(index, value) {
			return index, value, nil
		}
	}

	return NotFound, defaultType, &err.NotFound{}
}

func First[T any](slice []T, predicate func(index int, value T) bool) (int, T, error) {
	var notFoundValue T

	if len(slice) == 0 {
		panic("The slice is empty.")
	}

	if predicate == nil {
		return 0, slice[0], nil
	}

	max := maxElems

	workers := len(slice) / max

	if workers > maxWorkers {
		workers = maxWorkers
	}

	if workers < 2 {
		return firstWorker(slice, predicate, 0, len(slice))
	}

	pvs := make(chan col.PV[T], workers)
	wg := new(sync.WaitGroup)

	for i := 0; i < workers; i++ {
		wg.Add(1)

		start := i * max
		end := len(slice)

		if i < workers-1 {
			end = start + max
		}

		go func(s []T, start int, end int, result chan<- col.PV[T]) {
			defer wg.Done()
			i, v, e := firstWorker(s, predicate, start, end)

			if e == nil {
				result <- col.PV[T]{Priority: i, Val: v}
			} else {
				result <- col.PV[T]{Priority: NotFound, Val: notFoundValue}
			}
		}(slice, start, end, pvs)
	}

	wg.Wait()
	close(pvs)

	for pv := range pvs {
		if pv.Priority != NotFound {
			return pv.Priority, pv.Val, nil
		}
	}

	return NotFound, notFoundValue, &err.NotFound{}
}
