package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the marks of maths: ")

	// comma ok syntax. Behave as try and catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Marks: ", input)
	fmt.Printf("Type of marks variable is: %T\n", input)

	// Parsing our string input data into float of bit 64
	mathsMarks, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New maths marks: ", mathsMarks+5)
	}

}
