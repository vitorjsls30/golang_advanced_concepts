// Empty Interfaces
// Empty Interfaces are handy when you want pass any kind of parameter to a given function, which normally you wouldn't.

package main

import "fmt"

func getType(item interface{}) {
	fmt.Printf("Type of item: %T / Value of item : %v\n", item, item)
}

func main() {
	s := "Yo!! Let's try it out!"
	getType(s)

	x := 1234
	getType(x)
}
