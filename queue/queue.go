package queue

import "container/list"

// Queue is a queue data structure.
type Queue struct {
	list *list.List
}

// New creates a new Queue.
func New() *Queue {
	return &Queue{
		list: list.New(),
	}
}

// 
func (q *Queue) Push(v interface{}) {
	q.list.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	e := q.list.Front()
	if e != nil {
		return q.list.Remove(e)
	}
	return nil
}
