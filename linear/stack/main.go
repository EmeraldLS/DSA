package main

import "fmt"

type Stack struct {
	vals []int
}

func (s *Stack) Push(val int) {
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() {
	s.vals = append(s.vals[len(s.vals):], s.vals[:len(s.vals)-1]...)
}

func main() {
	var stack Stack

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Pop()
	stack.Pop()
	stack.Pop()

	fmt.Println(stack)
}
