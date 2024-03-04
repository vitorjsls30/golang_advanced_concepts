package main

import "fmt"

func main() {
	ch := make(chan string)
	ch <- "Yo!! A simple channel!!"
	message := <-ch

	fmt.Printf("The received message was: %s", message)
}
