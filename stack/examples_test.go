package stack

import (
	"fmt"
)

func Example() {
	// The initial stack capacity if you have it (zero also works)
	capacity := 3

	// Declare a stack of integers
	stack := New[int](capacity)

	// Put values into the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Or create and push range in one line of code
	// stack.New(capacity, 1, 2, 3)

	// Look at the top value on the stack
	fmt.Println(stack.Peek())

	// Take values out of the stack
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	// Output:
	// 3
	// 3
	// 2
	// 1
}
