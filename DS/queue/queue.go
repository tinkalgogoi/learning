package queue

import "fmt"

type Queue struct {
	items []string
}

func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() {
	if len(q.items) > 0 {
		fmt.Print(q.items[0]) // First element
		q.items = q.items[1:] // Dequeue
	} else {
		fmt.Print("Underflow")
	}
}
