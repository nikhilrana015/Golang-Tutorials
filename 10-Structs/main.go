package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	age       int64
}

func main() {
	// structs are the alternative versions of classes.
	user_1 := User{"Nikhil", "Rana", "nikhilrana@gmail.com", 24}
	fmt.Printf("User 1 details are: %+v\n", user_1)
	fmt.Printf("My firstName is %v and lastName is %v", user_1.FirstName, user_1.LastName)
}
