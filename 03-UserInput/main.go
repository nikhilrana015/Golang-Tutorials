package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the marks of maths: ")

	// comma ok syntax. Behave as try and catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Marks: ", input)
	fmt.Printf("Type of marks variable is: %T", input)

}
