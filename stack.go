package stack

import (
	"sync"
)

// Stack ... Stack type
type Stack struct {
	arr  []interface{}
	lock sync.RWMutex
}

// New ... Create a new instance of a stack
func New() *Stack {
	return &Stack{}
}

// Len ... Total number of items stack is holding
func (s *Stack) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.arr)
}

// Empty ... Returns true if stack is empty
func (s *Stack) Empty() bool {
	return s.Len() == 0
}

// Top ...Returns the top most element
func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.arr[len(s.arr)-1]
}

// Push ...Pushes element on to the stack
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.arr = append(s.arr, v)
}

// Pop ...Pops the top element off the stack
func (s *Stack) Pop() {
	if !s.Empty() {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.arr = s.arr[:len(s.arr)-1]
	}
}
