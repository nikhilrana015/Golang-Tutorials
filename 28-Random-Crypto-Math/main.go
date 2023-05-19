package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// random number
	// For random number generation there are 2 packages: Crypto and math package.

	// Random number generation using math/rand package
	//rand.Seed(time.Now().UnixNano())
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Println("Number is: ", rand.Intn(6))

	// Using crypto/rand package
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		panic(err)
	}
	fmt.Println(randomNumber)

}
