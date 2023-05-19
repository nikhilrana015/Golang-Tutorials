package main

import "fmt"

func main() {

	/*
		What does defer do it works in a stack way. LIFO. So in terms of
		thinking, you can assume any statement that starts with defer keyword that
		should be put inside the queue and those defer statements executed just
		before the termination of the program.
	*/

	defer fmt.Println("Hello")
	defer fmt.Println("It's me")
	fmt.Println("Nikhil is here")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
