package queue

import (
	"fmt"
)

func Example() {
	// The initial queue capacity if you have it (zero also works)
	capacity := 3

	// Declare a queue of integers
	queue := New[int](capacity)

	// Put a value into the queue
	queue.Enqueue(1)

	// Push multiple values at once onto the queue, 2 is pushed first, then 3
	queue.EnqueueMany(2, 3)

	// Or create a queue and push many values in one line of code
	// queue := queue.New(capacity, 1, 2, 3)

	// Look at the top value on the queue
	fmt.Println(queue.Peek()) // Prints 1

	// Take values out of the queue
	fmt.Println(queue.Dequeue()) // Prints 1
	fmt.Println(queue.Dequeue()) // Prints 2
	fmt.Println(queue.Dequeue()) // Prints 3
	// Output:
	// 1
	// 1
	// 2
	// 3
}
