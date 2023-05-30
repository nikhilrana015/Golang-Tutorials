package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) InsertAtEnd(data int) {

	node := &Node{Val: data, Next: nil}

	if l.Head == nil {
		l.Head = node
		l.Length += 1
		return
	}

	curr := l.Head

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = node
	curr = node
	l.Length += 1
}

func (l *LinkedList) DeleteAtEnd() {

	if l.Head == nil {
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		l.Length -= 1
		return
	}

	prev := &Node{Val: -1, Next: nil}
	curr := l.Head

	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
	l.Length -= 1

}

func (l *LinkedList) InsertAtFront(data int) {

	node := &Node{Val: data, Next: nil}

	node.Next = l.Head
	l.Head = node
	l.Length += 1
}

func (l *LinkedList) DeleteAtFront() {

	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next
	l.Length -= 1
}

func (l *LinkedList) PrintList() {

	curr := l.Head
	for curr != nil {
		fmt.Printf("%v --> ", curr.Val)
		curr = curr.Next
	}

	fmt.Println("")
}

func main() {

	list := &LinkedList{
		Head:   nil,
		Length: 0,
	}

	list.InsertAtFront(1)
	list.InsertAtEnd(2)
	list.InsertAtFront(3)
	list.InsertAtEnd(4)
	list.InsertAtEnd(5)
	list.InsertAtFront(100)
	list.DeleteAtEnd()
	list.DeleteAtEnd()
	list.DeleteAtFront()
	list.DeleteAtFront()
	list.DeleteAtFront()
	list.DeleteAtFront()
	list.PrintList()

}
