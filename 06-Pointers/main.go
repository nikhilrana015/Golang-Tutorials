package main

import "fmt"

func main() {

	var number *int
	fmt.Println("Value of number: ", number)
	fmt.Printf("Type of number is: %T\n", number)

	newScore := 15
	var scorePtr = &newScore

	fmt.Println("Value of scorePtr: ", scorePtr)
	fmt.Println("address of newScore variable: ", &newScore)
	fmt.Println("Derefrencing the address: ", *scorePtr)
	fmt.Printf("Type of scorePtr is: %T\n", number)

	fmt.Println("***********************************")

	// Pointer of Pointer
	mathsMarks := 95
	var marksPtr1 = &mathsMarks
	var marksPtr2 = &marksPtr1

	fmt.Println("Addres of mathsMarks: ", &mathsMarks)
	fmt.Println("Value of marksPtr1: ", marksPtr1)
	fmt.Println("Derefrencing of marksPtr1: ", *marksPtr1)
	fmt.Println("Addres of marksPtr1: ", &marksPtr1)
	fmt.Println("Value of marksPtr2: ", marksPtr2)
	fmt.Println("Derefrencing of marksPtr2: ", *marksPtr2)
	fmt.Println("Double Derefrencing of marksPtr2: ", *(*marksPtr2))
}
