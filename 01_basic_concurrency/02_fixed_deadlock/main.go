package main

import "fmt"

func main() {
	ch := make(chan string)

	// declaring an annonymous function that is immediately called in a new goroutine
	// allow us to correctly retrieve the value put ino the channel
	go func() {
		ch <- "Yo!! Now this is working!!"
	}()

	// retrieving here now works!
	message := <-ch
	fmt.Printf("The received message was: %s", message)
}
