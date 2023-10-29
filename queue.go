package main

import "errors"

type Queue struct {
	queue  []int
	size   int
	IsFlex bool
}

func NewQueue(size int, isFlex bool) Queue {

	if isFlex {
		return Queue{
			queue:  make([]int, 0),
			size:   size,
			IsFlex: isFlex,
		}
	}

	return Queue{
		queue:  make([]int, size),
		size:   size,
		IsFlex: isFlex,
	}
}

func (q *Queue) Push(item int, autoPop bool) error {
	if len(q.queue) >= q.size && !q.IsFlex {
		if autoPop {
			q.Pop()
			q.queue = append(q.queue, item)
			return nil
		}
		return errors.New("Queue overflowing")
	} else {
		q.queue = append(q.queue, item)
	}
	return nil
}

func (q *Queue) Pop() (int, error) {
	if len(q.queue) == 0 {
		return 0, errors.New("Queue underflow")
	}
	itemToPop := q.queue[0]
	q.queue = q.queue[1:]
	return itemToPop, nil
}
