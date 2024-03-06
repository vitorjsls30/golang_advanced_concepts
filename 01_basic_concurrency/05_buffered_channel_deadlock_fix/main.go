// buffered channel deadlock
// a deadlock may occur if we try to add more than previously set items into a given channel.

// following the last example...
package main

import "fmt"

func main() {

	ch := make(chan string, 1)

	// as before, we need to set these values in another go routine...
	go func() {
		ch <- "Yo! That's the first message!"
		ch <- "Now the second!"
	}()

	message1 := <-ch
	message2 := <-ch

	fmt.Printf("The first message is: %s\n", message1)
	fmt.Printf("The second message is: %s\n", message2)
}
