package stack

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {

	s := New(3, 1, 2, 3)

	a := s.Pop()
	b := s.Pop()
	c := s.Pop()

	if a != 3 || b != 2 || c != 1 {
		t.Error("Stack did not push and pop elements correctly.")
	}
}

func TestIsEmpty_WithElements_IsFalse(t *testing.T) {
	s := New(1, 1)

	if s.IsEmpty() {
		t.Error("Stack is empty but IsEmpty return true.")
	}
}

func TestIsEmpty_WithNoElements_IsTrue(t *testing.T) {
	s := New[int](1)

	if !s.IsEmpty() {
		t.Error("Stack is empty but IsEmpty return false.")
	}
}
