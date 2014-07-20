package stack

import "testing"

var table = []interface{}{10, 22, "hello"}

func TestStack(t *testing.T) {
	s := New()

	for _, v := range table {
		s.Push(v)
	}

	v := s.Pop()
	if v != table[2] {
		t.Errorf("expected %v, got %v", table[2], v)
	}
	v = s.Pop()
	if v != table[1] {
		t.Errorf("expected %v, got %v", table[1], v)
	}
	v = s.Pop()
	if v != table[0] {
		t.Errorf("expected %v, got %v", table[0], v)
	}
	v = s.Pop()
	if v != nil {
		t.Errorf("expected %v, got %v", nil, v)
	}
}
