package main

import "fmt"

// In golang, if any variable or method starts with capital letter then
// it is considered as public and it can access from another packages as well.
const PiValue float64 = 3.147

func main() {

	var name string = "Nikhil Rana"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	var isEligible bool = true
	fmt.Println(isEligible)
	fmt.Printf("Variable is of type: %T \n", isEligible)

	// It is responsible to store the values from 0-255
	var score uint8 = 255
	fmt.Println(score)
	fmt.Printf("Variable is of type: %T \n", score)

	// It gives upto 5-6 decimal places after decimal point. Less accurate or less precise as cmprd to float64
	var percentage1 float32 = 29.7272781781728787287
	fmt.Println(percentage1)
	fmt.Printf("Variable is of type: %T \n", percentage1)

	// It is more accurate or more precise then float32.
	var percentage2 float64 = 29.7272781781728787287
	fmt.Println(percentage2)
	fmt.Printf("Variable is of type: %T \n", percentage2)

	// default values and aliases

	// default value of int is 0.
	var number int
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	// default value of bool is false.
	var isNumber bool
	fmt.Println(isNumber)
	fmt.Printf("Variable is of type: %T \n", isNumber)

	// default value of string is empty or "".
	var lastName string
	fmt.Println(lastName)
	fmt.Printf("Variable is of type: %T \n", lastName)

	// implicit type

	// Lexer implicitly decides the type of variable without mentioning it.
	var website = "youtube.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// Once lexer decide the type of variable it didn't change. Line below gives error.
	//website = 2

	// var-type
	// This is walrus operator. It is used only inside the methods and functions
	// It can't used globally while defining variables.
	sports := "cricket"
	fmt.Println(sports)
	fmt.Printf("Variable is of type: %T \n", sports)

	//constants
	fmt.Println(PiValue)
	fmt.Printf("Variable is of type: %T \n", PiValue)
}
