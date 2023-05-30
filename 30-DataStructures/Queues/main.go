package main

import "fmt"

type Queue struct {
	arr []int
	len int
}

func (q *Queue) Push(data int) {
	q.arr = append(q.arr, data)
	q.len += 1
}

func (q *Queue) Pop() {

	if q.len == 0 {
		panic("Queue is empty")
	}

	q.arr = q.arr[1:]
	q.len -= 1
}

func (q *Queue) Front() int {

	if q.len == 0 {
		panic("Queue is empty")
	}

	topElement := q.arr[0]
	return topElement
}

func (q *Queue) Empty() bool {
	if q.len == 0 {
		return true
	}

	return false
}

func main() {
	q := &Queue{arr: []int{}, len: 0}

	q.Push(1)
	fmt.Println(q.Front())
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Pop()
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
}
