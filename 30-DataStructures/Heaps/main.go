package main

import "fmt"

type Heap struct {
	array []int
}

func (hp *Heap) HeapifyUp(index int) {
	parent := (index - 1) / 2

	if hp.array[parent] < hp.array[index] {
		hp.array[parent], hp.array[index] = hp.array[index], hp.array[parent]
		hp.HeapifyUp(parent)
	}
}

func (hp *Heap) HeapifyDown(index int) {

	largest := index
	left := (index * 2) + 1
	right := (index * 2) + 2
	length := len(hp.array)

	if left < length && hp.array[left] > hp.array[largest] {
		largest = left
	}

	if right < length && hp.array[right] > hp.array[largest] {
		largest = right
	}

	if largest != index {
		hp.array[index], hp.array[largest] = hp.array[largest], hp.array[index]
		hp.HeapifyDown(largest)
	}
}

func (hp *Heap) Delete() {

	if len(hp.array) == 0 {
		panic("Heap is empty")
	}

	lastElement := hp.array[len(hp.array)-1]
	hp.array = hp.array[:len(hp.array)-1]

	if len(hp.array) == 0 {
		return
	}

	hp.array[0] = lastElement
	hp.HeapifyDown(0)

}

func (hp *Heap) Insert(element int) {
	hp.array = append(hp.array, element)
	hp.HeapifyUp(len(hp.array) - 1)
}

func (hp *Heap) Empty() bool {
	if len(hp.array) == 0 {
		return true
	}

	return false
}

func (hp *Heap) Top() int {
	if len(hp.array) == 0 {
		panic("Heap is empty")
	}

	return hp.array[0]
}

func main() {
	heap := &Heap{}

	fmt.Println(heap)

	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(30)
	heap.Insert(3)
	heap.Insert(7)
	heap.Insert(11)
	heap.Insert(17)
	heap.Insert(19)

	for i := 1; i < 9; i++ {
		fmt.Println(heap.Top())
		heap.Delete()
	}
}
