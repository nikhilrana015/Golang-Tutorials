package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct {
}

func (d Dog) speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) speak() string {
	return "Meow!"
}

type Hen struct {
}

func (h Hen) speak() string {
	return "????"
}

func main() {

	animals := []Animal{Dog{}, Cat{}, Hen{}}
	for _, animal := range animals {
		fmt.Println(animal.speak())
	}
}
