package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var rankingSignals = []string{"testing"}

func main() {
	// go greeter("hello")
	// greeter("world")

	var urls = []string{
		"https://pkg.go.dev",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://openai.com/blog/chatgpt",
	}

	for _, url := range urls {
		go getStatusCode(url)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(rankingSignals)

}

// func greeter(word string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(word)
// 	}
// }

func getStatusCode(url string) {

	defer wg.Done()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status code for given url %v is: %v\n", url, res.StatusCode)

	mutex.Lock()
	rankingSignals = append(rankingSignals, url)
	mutex.Unlock()
}
