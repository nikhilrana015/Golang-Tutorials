package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)
	fmt.Printf("Type of time is: %T\n", currentTime)

	// fmt.Println("New Time is: ", currentTime.Format("2006-01-02 15:04:05 Monday"))

	// get the time in yyyy-mm-dd because reference time is in this format.
	//https://stackoverflow.com/questions/33119748/convert-time-time-to-string
	formattedTime := currentTime.Format("2006-01-02 15:04:05 Monday")
	fmt.Println("Fomatted Current time is: ", formattedTime)
	fmt.Printf("Type of time is: %T\n", formattedTime)

	createdDate := time.Date(1999, time.May, 15, 05, 43, 23, 43, time.UTC)
	fmt.Println("Current date is: ", createdDate)
	fmt.Printf("Type of time is: %T\n", formattedTime)

	formattedDate := createdDate.Format("02-01-2006 Monday")
	fmt.Println("Formatted date is: ", formattedDate)
	fmt.Printf("Type of time is: %T\n", formattedDate)

	// To build executable files for different operating systems
	// we should use the command: GOOS=linux go build

}
