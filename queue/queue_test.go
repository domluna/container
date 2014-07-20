package queue

import "testing"

var table = []interface{}{10, 22, "hello"}

func TestStack(t *testing.T) {
	q := New()

	for _, v := range table {
		q.Push(v)
	}

	for _, e := range table {
		v := q.Pop()
		if v != e {
			t.Errorf("expected %v, got %v", e, v)
		}
	}

	v := q.Pop()
	if v != nil {
		t.Errorf("expected %v, got %v", nil, v)
	}
}
