package slice

import "testing"

func BenchmarkIndexOfWorker(b *testing.B) {

	values := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		values = append(values, i)
	}

	for i := 0; i < b.N; i++ {
		indexOfWorker(values, b.N-1, 0, b.N-1)
	}
}

func BenchmarkIndexOfConcurrentDefaults(b *testing.B) {

	maxElems = DefaultMaxSearchLength
	maxWorkers = DefaultMaxSearchWorkers

	values := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		values = append(values, i)
	}

	for i := 0; i < b.N; i++ {
		IndexOf(values, b.N-1)
	}
}

func BenchmarkFirstWorker(b *testing.B) {

	values := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		values = append(values, i)
	}

	for i := 0; i < b.N; i++ {
		predicateWorker(values, func(index int, value int) bool { return i == b.N-1 }, 0, b.N-1)
	}
}

func BenchmarkFirstConcurrentDefaults(b *testing.B) {

	values := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		values = append(values, i)
	}

	for i := 0; i < b.N; i++ {
		First(values, func(index int, value int) bool { return i == b.N-1 })
	}
}
