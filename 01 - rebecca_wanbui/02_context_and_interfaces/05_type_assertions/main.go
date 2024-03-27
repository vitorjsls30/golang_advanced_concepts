// knowing that we can use empty interfaces, to indicate that our methods accept any type of value, we can
// use type assertion to avoid mistakes. Ex:

package main

import "fmt"

func main() {
	var s interface{}

	// storing a string value into the empty interface item
	s = "Yo!! Now I'm a string!"

	// type assertion
	myStr := s.(string)
	fmt.Printf("The value is: %s", myStr)
}
