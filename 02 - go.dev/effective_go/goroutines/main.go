package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func main() {
	go say("Hello")
	say("World")
}
