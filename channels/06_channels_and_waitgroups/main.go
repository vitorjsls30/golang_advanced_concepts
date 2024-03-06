package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(10)
}

func main() {

	// declaring an un-buffered channel
	ch := make(chan int)

	// declaring our min goroutine
	go func() {
		// creating a wait group to help us sync the following operations
		wg := sync.WaitGroup{}

		// spawing a new goroutine for each loop
		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func() {
				// deferring wg.Done call
				defer wg.Done()
				// finally sending the rdn number into the channel
				ch <- doWork()
			}()
		}

		// waiting for all the wait groups to finish...
		wg.Wait()
		// closing the channel
		close(ch)
	}()

	// listing the result
	fmt.Println("Result is:")
	for n := range ch {
		fmt.Println(n)
	}
}
