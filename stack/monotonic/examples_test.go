package monotonic

import (
	"fmt"

	"github.com/sgago/col"
)

func Example() {
	// The initial capacity, if you have it (zero also works)
	capacity := 3

	monostack, _ := New[int](Decreasing, capacity) // This stack is always decreasing

	pv1 := col.PV[int]{Priority: 3, Val: 3}
	pv2 := col.PV[int]{Priority: 1, Val: 1}
	pv3 := col.PV[int]{Priority: 6, Val: 6}

	monostack.Push(pv1) // Stack: 3

	// Pushing 1 keeps the values in decreasing value
	// This maintains the monotonic condition
	monostack.Push(pv2) // Stack: 3 1

	// Look at the top value in the monotonic stack
	fmt.Println(monostack.Peek())

	// Pushing 6 would break the monotonic condition
	// A stack of values 3, 1, 6 are NOT always decreasing
	// So we pop any values less than 6 off the stack
	// Push will return the popped values 1 and 3
	fmt.Println(monostack.Push(pv3)) // Stack: 6

	// Like a normal stack, we can pop values too
	fmt.Println(monostack.Pop()) // Stack: Empty

	// Output:
	// {1 1}
	// [{1 1} {3 3}]
	// {6 6}
}
