package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Learning Race-Around Conditions")

	var wg = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}
	//var RWmutex = &sync.RWMutex{}

	// Read/write mutex helps to lock the resource while reading it. It is similar to normal
	// mutex but it comes with locking while reading. And when any thread have to write, it throws
	// all others go routines which are reading the resource from the scope, update the resource
	// and then allows them to read it.

	var scores = []int{0}

	wg.Add(4)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Score 1")
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Score 2")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Score 3")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		//RWmutex.RLock()
		fmt.Printf("Reading...  %v\n", scores)
		//RWmutex.RUnlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(scores)
}
