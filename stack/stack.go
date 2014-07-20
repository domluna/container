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

func (s *Stack) Push(v interface{}) {
	s.list.PushFront(v)
}

func (s *Stack) Pop() interface{} {
	e := s.list.Front()
	if e != nil {
		return s.list.Remove(e)
	}
	return nil
}
