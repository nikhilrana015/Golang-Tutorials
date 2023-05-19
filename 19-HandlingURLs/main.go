package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://www.youtube.com/results?search_query=golang"
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("Type of url is: %T\n", result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.RawPath)

	// It parses the rawQuery
	qParams := result.Query()
	fmt.Printf("The type of queryParams: %T\n", qParams)
	fmt.Printf("Possible values inside it: %+v\n", qParams)

	// create a url if we have chunks of information
	// always sends as a reference not a copy
	constructURL := &url.URL{
		Scheme:   "https",
		Host:     "youtube.com",
		Path:     "/results",
		RawQuery: "search_query=nikhil",
	}

	fmt.Printf("Type of constructURL is: %T\n", constructURL)
	fmt.Printf("My construct url is: %+v\n", constructURL)

	urlString := constructURL.String()
	fmt.Println("urlString: ", urlString)
}
