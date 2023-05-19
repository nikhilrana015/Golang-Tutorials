package main

import "fmt"

func main() {

	runs := 9

	if runs == 50 {
		fmt.Println("Half-century")
	} else if runs == 100 {
		fmt.Println("Century")
	} else {
		fmt.Println(runs)
	}

	// initializing and also handling in if-else

	if isCar := true; isCar {
		fmt.Println("Its a car")
	} else {
		fmt.Println("Its not a car")
	}
}
