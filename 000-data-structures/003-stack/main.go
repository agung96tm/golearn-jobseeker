package main

import (
	"errors"
)

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// stack.Push(99)
	// removedStack, err := stack.Pop(0)
	// err := stack.Insert(99, 1)
}

type Stack struct {
	stack []int
}

func NewStack() Stack {
	return Stack{
		stack: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack is empty")
	}

	temp := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]

	return temp, nil
}

func (s *Stack) Insert(val int, loc int) error {
	if loc < 0 || loc > len(s.stack) {
		return errors.New("loc out of range")
	}

	s.stack = append(s.stack[:loc], append([]int{val}, s.stack[loc:]...)...)
	return nil
}
