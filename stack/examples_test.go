package stack

import (
	"fmt"
)

func Example() {
	// The initial stack capacity if you have it (zero also works)
	capacity := 3

	// Declare a stack of integers
	stack := New[int](capacity)

	// Put a value onto the stack
	stack.Push(1)

	// Push multiple values at once onto the stack, 2 is pushed first, then 3
	stack.PushMany(2, 3)

	// Or create a stack and push many values in one line of code
	// stack := stack.New(capacity, 1, 2, 3)

	// Look at the top value on the stack
	fmt.Println(stack.Peek()) // Prints 3

	// Take values out of the stack
	fmt.Println(stack.Pop()) // Prints 3
	fmt.Println(stack.Pop()) // Prints 2
	fmt.Println(stack.Pop()) // Prints 1

	// Output:
	// 3
	// 3
	// 2
	// 1
}
