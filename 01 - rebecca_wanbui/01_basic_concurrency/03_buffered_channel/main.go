// buffered channels
// buffered channels are a way of concurrently execute different operations by sending and receiving information
// of a certaing element type for a given capacity.
// They are great when you need some sort of mechanism to limit concurrent usage and to manage the amount of work
// your system is allowed to do, to avoid it's overwhelming.

// following up from the previous examples...
package main

import "fmt"

func main() {

	ch := make(chan string, 1) // we set a new channel, this time with a starting capacity of 1 item

	go func() {
		ch <- "Yo! Now we're running in a buffered chan!!"
	}()

	message := <-ch
	fmt.Printf("Received message was: %s", message)
}
