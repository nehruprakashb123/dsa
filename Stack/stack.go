package main

import (
	"errors"
	"fmt"
)

// Stack Implementation

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Display() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Stack (top to bottom): ")
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Printf("%v ", s.items[i])
	}
	fmt.Println()
}

func main() {
	stack := NewStack()

	//Push elements
	fmt.Println("Pushing elements: 10, 20, 30")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Display()
}

