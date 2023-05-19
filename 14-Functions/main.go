package main

import "fmt"

func addFunc(a int, b int) int {
	return a + b
}

func main() {

	result := addFunc(3, 4)
	fmt.Println("Result is: ", result)

	summation, msg := addFunc2(1, 2, 3, 4, 5, 6)

	fmt.Printf("Result2 is: %v and message %v", summation, msg)

}

func addFunc2(values ...int) (int, string) {

	fmt.Printf("Type of values is: %T\n", values)

	calculate := 0
	for _, v := range values {
		calculate += v
	}

	return calculate, "Summation of elements"
}
