package slice

import "testing"

func BenchmarkIndexOfWorker(b *testing.B) {

	values := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		values = append(values, i)
	}

	for i := 0; i < b.N; i++ {
		indexOfWorker(values, b.N-1)
	}
}

func BenchmarkIndexOfC(b *testing.B) {

	values := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		values = append(values, i)
	}

	for i := 0; i < b.N; i++ {
		IndexOf(values, b.N-1)
	}
}
