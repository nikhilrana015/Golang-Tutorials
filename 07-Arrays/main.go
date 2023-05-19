package main

import "fmt"

func main() {

	// Arrays
	var sports [5]string

	sports[0] = "cricket"
	sports[1] = "football"
	sports[4] = "badminton"

	fmt.Println(sports)
	fmt.Println("Len of sports array: ", len(sports))

	var marksSheet [4]int
	marksSheet[0] = 100
	marksSheet[3] = 95

	fmt.Println(marksSheet)
	fmt.Println("Len of marksSheet array: ", len(marksSheet))

	// Declaration and initialization at the same time
	var countries = []string{"India", "Australia", "Norway", "Finland"}
	fmt.Println(countries)
	fmt.Println("Len of countries array: ", len(countries))
}
