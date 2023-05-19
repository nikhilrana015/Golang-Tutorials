package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// responsible for creating the file.
	file, err := os.Create("hobby.txt")

	fmt.Printf("Type of file is: %T\n", file)

	if err != nil {
		panic(err)
	}

	data := "I love to play cricket."

	length, err := io.WriteString(file, data)

	if err != nil {
		panic(err)
	}

	fmt.Println("Total length is: ", length)

	readFile("hobby.txt")

	defer file.Close()

}

func readFile(fileName string) {

	dataInBytes, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data present inside the file is: ", dataInBytes)
}
