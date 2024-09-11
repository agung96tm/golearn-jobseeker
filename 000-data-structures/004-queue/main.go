package main

import "fmt"

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	q, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(q)
		fmt.Println(queue)
	}
}

type Queue struct {
	queue []int
}

func NewQueue() Queue {
	return Queue{
		queue: []int{},
	}
}

func (q *Queue) Enqueue(i int) {
	q.queue = append(q.queue, i)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.queue) == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	if len(q.queue) == 1 {
		temp := q.queue[0]
		q.queue = []int{}
		return temp, nil
	}

	temp := q.queue[0]
	q.queue = q.queue[1:]
	return temp, nil
}
