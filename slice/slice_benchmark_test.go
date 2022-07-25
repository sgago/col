package slice

import "testing"

func BenchmarkFirstWorker(b *testing.B) {

	values := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		values = append(values, i)
	}

	for i := 0; i < b.N; i++ {
		firstWorker(values, func(index int, value int) bool { return i == b.N-1 }, 0, b.N-1)
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
