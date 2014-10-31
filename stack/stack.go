package stack

import "container/list"

// Stack is a stack data structure.
type Stack struct {
	list *list.List
}

// New creates a new Stack.
func New() *Stack {
	return &Stack{
		list: list.New(),
	}
}

// Push inserts an element at the top of the stack.
func (s *Stack) Push(v interface{}) {
	s.list.PushFront(v)
}

// Pop removes an element from the top of the stack.
func (s *Stack) Pop() interface{} {
	e := s.list.Front()
	if e != nil {
		return s.list.Remove(e)
	}
	return nil
}
