package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	var URL string = "https://catfact.ninja/fact"

	makeGetRequest(URL)
	makePostRequestJSONRequest(URL)
	makePostRequestFormRequest(URL)

}

func makeGetRequest(myUrl string) {

	fmt.Println("Making a get request")
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	// This is another way to have data in string type by writing strings.Builder.
	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	byteCount, err := responseString.Write(content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Len of the content is: ", byteCount)
	fmt.Println(responseString.String())

	// This is one of the easy way to see thorugh the content in english
	// form by converting them into string type.
	// fmt.Println(content)
	// fmt.Println(string(content))

}

func makePostRequestJSONRequest(myUrl string) {

	fmt.Println("Making a post request in json form")
	requestBody := strings.NewReader(`
		{
			"name":"Nikhil Rana",
			"age":"23",
			"Degree":"B.Tech",
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	byteCount, err := responseString.Write(content)

	if err != nil {
		panic(err)
	}

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}

func makePostRequestFormRequest(myUrl string) {

	fmt.Println("Make a post request in form way")
	data := url.Values{}
	data.Add("firstName", "Nikhil")
	data.Add("lastName", "rana")
	data.Add("age", "24")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	byteData, err := responseString.Write(content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Byte Data is: ", byteData)
	// Treating json as a string
	fmt.Println(responseString.String())

}
