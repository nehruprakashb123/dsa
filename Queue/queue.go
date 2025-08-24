package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Add the items from the rear of queue
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Remove the items from front of queue
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Front returns the front item without removing it
func (q *Queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	item := q.items[0]
	return item, nil
}

// Rear returns rear item without removing it
func (q *Queue) Rear() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	return q.items[len(q.items)-1], nil
}

// IsEmpty return if the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

// Display prints the queue values
func (q *Queue) Display() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	for _, item := range q.items {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

func main() {
	queue := NewQueue()

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Display()
	queue.Dequeue()
	queue.Display()
	front, _ := queue.Front()
	fmt.Println(front)

	//Dequeuing the elements
	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		fmt.Printf("Dequeued: %v \n", item)
		queue.Display()
	}
}
