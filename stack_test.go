package stack

import (
	"testing"
)

func TestCreation(t *testing.T) {
	stack := New()
	if stack == nil {
		t.Error("stack created is nil")
	}
}

func TestTop(t *testing.T) {
	stack := New()
	if stack.Top() != nil {
		t.Error("stack top should have been nil")
	}
}

func TestPush(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)

	if stack.Empty() {
		t.Error("stack should not be empty")
	}

	if stack.Top() != 2 {
		t.Error("stack top value should be 2")
	}
}

func TestPop(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)

	if stack.Empty() {
		t.Error("stack should not be empty")
	}

	stack.Pop()
	if stack.Top() != 1 {
		t.Error("stack top should be 1.")
	}

	stack.Pop()
	if !stack.Empty() {
		t.Error("stack should be empty.")
	}
}
