package main

import "fmt"

type Stack struct {
	arr []int
	len int
}

func (s *Stack) Push(data int) {

	s.arr = append(s.arr, data)
	s.len += 1
}

func (s *Stack) Pop() {

	if s.len == 0 {
		panic("Stack is empty")
	}

	s.arr = s.arr[:s.len-1]
	s.len -= 1
}

func (s *Stack) Top() int {
	if s.len == 0 {
		panic("Stack is empty")
	}

	topElement := s.arr[s.len-1]
	return topElement
}

func (s *Stack) Empty() bool {
	if s.len == 0 {
		return true
	}

	return false
}

func main() {
	ss := &Stack{arr: []int{}, len: 0}
	ss.Push(2)
	ss.Push(4)
	ss.Push(5)
	fmt.Println(ss.Top())
	ss.Pop()
	ss.Pop()
	fmt.Println(ss.Top())
	ss.Pop()
	ss.Pop()
}
