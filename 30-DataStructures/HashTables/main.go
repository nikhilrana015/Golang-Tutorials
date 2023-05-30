package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type HashTable struct {
	size  int
	array []*Node
}

func (hashtables *HashTable) HashFunction(key int) int {
	return key % hashtables.size
}

func (hashtables *HashTable) Insert(value int) {
	getKey := hashtables.HashFunction(value)
	newNode := &Node{Val: value, Next: nil}

	if hashtables.IsPresent(value) {
		return
	}

	newNode.Next = hashtables.array[getKey]
	hashtables.array[getKey] = newNode
}

func (hashtables *HashTable) Delete(value int) {
	getKey := hashtables.HashFunction(value)
	curr := hashtables.array[getKey]

	if curr == nil {
		return
	}

	var prev *Node
	prev = nil

	for curr != nil && curr.Val != value {
		prev = curr
		curr = curr.Next
	}

	if prev == nil {
		hashtables.array[getKey] = curr.Next
	} else if curr != nil {
		prev.Next = curr.Next
	}

}

func (hashtables *HashTable) IsPresent(value int) bool {
	getKey := hashtables.HashFunction(value)
	curr := hashtables.array[getKey]

	for curr != nil && curr.Val != value {
		curr = curr.Next
	}

	if curr == nil {
		return false
	}

	return true
}

func main() {

	fmt.Println("Let's go with Hash-Tables")
	hashtables := &HashTable{size: 10, array: make([]*Node, 10)}

	for i := 1; i < 16; i++ {
		hashtables.Insert(i)
	}

	// hashtables.Delete(55)
	// fmt.Println(hashtables.IsPresent(15))
	// hashtables.Insert(4)
	// hashtables.Insert(3)
	// hashtables.Insert(2)
	// hashtables.Insert(1)

	hashtables.Delete(15)
	hashtables.Delete(101)

	for i := 0; i < 10; i++ {
		fmt.Printf("Values at given index %v: ", i)
		curr := hashtables.array[i]

		for curr != nil {
			fmt.Printf("%v ", curr.Val)
			curr = curr.Next
		}

		fmt.Println("")
	}

}
