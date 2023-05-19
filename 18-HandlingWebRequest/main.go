package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	response, err := http.Get("https://chess.com/")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("Response type is %T\n: ", response)
	fmt.Println("Status code is: ", response.StatusCode)

	// cookies := response.Cookies()

	// fmt.Printf("Response type is: %T\n", cookies)
	// fmt.Println("Length of cookies is: ", len(cookies))
	// fmt.Println("Cookies are: ", cookies)

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))
}
