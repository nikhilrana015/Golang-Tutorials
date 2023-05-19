package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	seed := time.Now().UnixNano()
	fmt.Println("Seed value is: ", seed)
	rand.New(rand.NewSource(seed))

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Number is one, you can open")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
		fallthrough
	case 4:
		fmt.Println("Number is 4")
		fallthrough
	case 5:
		fmt.Println("Number is 5")
	case 6:
		fmt.Println("Number is 6 and you can open and you can play again")
	default:
		fmt.Println("Don't know which number it is")
	}

	// fallthrough means it will move to the next case and print it. And stop
	// when it didn't find the fallthrough

}
