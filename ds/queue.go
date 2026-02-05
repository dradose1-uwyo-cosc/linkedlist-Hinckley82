package ds

import (
	"errors"
)

type Queue struct {
	list LinkedList
}

func (q *Queue) Push(value string) { // add Tail node
	q.list.Insert(value)
}

func (q *Queue) Pop() (string, error) { // remove the Head
	if q.list.IsEmpty() {
		return "", errors.New("queue is empty")
	}

	value := q.list.Head.data
	err := q.list.RemoveAt(0)
	if err != nil {
		return "", err
	}

	return value, nil
}
