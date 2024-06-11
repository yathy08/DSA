package main

import (
	"fmt"
	"errors"
)

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("queue underflow")
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value, nil
}

func (q *Queue) Display() {
	fmt.Println("Queue elements:", q.elements)
}

func main() {
	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Display()

	value, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dequeued value:", value)
	}
	queue.Display()

	queue.Enqueue(40)
	queue.Display()

	value, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dequeued value:", value)
	}
	queue.Display()
}
