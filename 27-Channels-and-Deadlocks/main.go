package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and deadlock")

	// https://stackoverflow.com/questions/11943841/what-is-channel-buffer-size
	// these are the buffered channels
	myChannel := make(chan int, 5)
	wg := &sync.WaitGroup{}

	// Another way to declare channel.
	//var myChannel2 chan string

	// In channel, it says that you can't pass a value to the channel until someone is listening to it.
	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)

	// receiver function or RECEIVE ONLY
	// receiver function never has a right to close the channel
	// if we close the channel, then we still access the values present inside it.
	go func(myChannel <-chan int, wg *sync.WaitGroup) {

		// val, isChannelOpen := <-myChannel
		// fmt.Println(isChannelOpen)
		// fmt.Println(val)

		for val := range myChannel {
			// _, isChannelOpen := <-myChannel
			// fmt.Printf("Channel Open: %v, and value: %v\n", isChannelOpen, val)
			fmt.Println("Value: ", val)
		}

		// fmt.Println(<-myChannel)
		// fmt.Println(<-myChannel)
		wg.Done()

	}(myChannel, wg)

	// sender function or SEND ONLY
	// sender function has a right to close the channel
	// if we close the channel and then tried to send values to the channel it creates a panic or error.
	go func(myChannel chan<- int, wg *sync.WaitGroup) {

		//close(myChannel)
		myChannel <- 5
		myChannel <- 6
		close(myChannel)

		//myChannel <- 6
		wg.Done()

	}(myChannel, wg)

	wg.Wait()
}
