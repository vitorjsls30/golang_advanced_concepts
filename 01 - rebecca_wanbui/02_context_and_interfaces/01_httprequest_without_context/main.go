// Contexts
// from the docs: package context defines the Context type, which carries deadlines, cancellation signals and other
// request-scoped values across API boundaries and between processes

package main

import (
	"fmt"
	"io"
	"net/http"
)

// let's say we wann fetch some data from an API
func main() {

	api_url := "https://jsonplaceholder.typicode.com/users"

	// first, we create our new request object
	req, err := http.NewRequest(http.MethodGet, api_url, nil)
	if err != nil {
		fmt.Printf("ERR: could not fetch url data: %s", err.Error())
		return
	}

	// after that, we process the request itself
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("ERR: could not process request: %s", err.Error())
		return
	}
	// if no problemns occurred, we defer the response.Body closing
	defer res.Body.Close()

	// we read all the received data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERR: could not read request response: %s", err.Error())
		return
	}

	// and finally, we print the data length
	fmt.Printf("Response length: %d", len(data))
}
