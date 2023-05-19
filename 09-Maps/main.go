package main

import "fmt"

func main() {

	rollMap := make(map[string]string)

	rollMap["1"] = "Nikhil"
	rollMap["2"] = "Ajay"
	rollMap["3"] = "Sneha"

	fmt.Println("RollMap is: ", rollMap)
	fmt.Printf("Type of rollMap: %T\n", rollMap)

	delete(rollMap, "3")

	fmt.Println("Map after deletion of key 3: ", rollMap)

	for keys, values := range rollMap {
		fmt.Printf("Key is %v, value is %v\n", keys, values)
	}

	var colors map[string]string    // 1st-way
	cars := make(map[string]string) // 2nd-way
	newColors := map[string]string{ // 3rd-way
		"red":   "000fff",
		"white": "000000",
	}

	fmt.Printf("%+v", colors)
	fmt.Println("\n", cars)
	fmt.Println("\n", newColors)
}
