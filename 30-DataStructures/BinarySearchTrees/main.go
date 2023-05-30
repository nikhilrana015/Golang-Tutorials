package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(val int) {

	if bst.root == nil {
		bst.root = &Node{Data: val, Left: nil, Right: nil}
		return
	}

	if bst.root.Data > val {
		if bst.root.Left == nil {
			newNode := &Node{Data: val, Left: nil, Right: nil}
			bst.root.Left = newNode
		} else {
			temp := bst.root
			bst.root = temp.Left
			bst.Insert(val)
			bst.root = temp
		}
	} else {
		if bst.root.Right == nil {
			newNode := &Node{Data: val, Left: nil, Right: nil}
			bst.root.Right = newNode
		} else {
			temp := bst.root
			bst.root = temp.Right
			bst.Insert(val)
			bst.root = temp
		}
	}
}

func (bst *BinarySearchTree) FindVal(data int) bool {

	if bst.root == nil {
		return false
	}

	if bst.root.Data == data {
		return true
	}

	temp := bst.root
	leftSide, rightSide := false, false

	if bst.root.Data > data {

		bst.root = temp.Left
		leftSide = bst.FindVal(data)
	} else {
		bst.root = temp.Right
		rightSide = bst.FindVal(data)
	}

	bst.root = temp
	return leftSide || rightSide
}

func (bst *BinarySearchTree) PrintBST() {

	if bst.root == nil {
		return
	}

	temp := bst.root

	// going to left
	bst.root = temp.Left
	bst.PrintBST()

	// Print the Value Inorder of BST always gives ascending order
	fmt.Printf("%v ", temp.Data)

	// going to right
	bst.root = temp.Right
	bst.PrintBST()

	bst.root = temp

}

func main() {
	bst := &BinarySearchTree{root: nil}

	bst.Insert(100)
	bst.Insert(80)
	bst.Insert(70)
	bst.Insert(900)
	bst.Insert(281)
	bst.Insert(21)
	bst.Insert(1)

	fmt.Println(bst.root)
	fmt.Println(bst.FindVal(190))
	fmt.Println(bst.root)
	bst.PrintBST()

}
