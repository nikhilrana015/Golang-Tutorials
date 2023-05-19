package main

import (
	"fmt"
)

func main() {

	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"}

	// for i := 0; i < len(months); i++ {
	// 	fmt.Println(months[i])
	// }

	// for idx := range months {
	// 	fmt.Println(months[idx])
	// }

	// for-each type loop
	for idx, month := range months {
		fmt.Printf("%v-month is: %v\n", idx+1, month)
	}

	number := 1

	// type of while loop. Golang doesn't contain while as a keyword or while loop.
	for number < 10 {

		if number == 5 {
			number++
			continue
		} else if number == 7 {
			//break
			goto nikhil
		} else {
			fmt.Println(number)
		}
		number++
	}

	// use any-name to define our go-to tag
nikhil:
	fmt.Println("Jumping to go-to statement")

}
