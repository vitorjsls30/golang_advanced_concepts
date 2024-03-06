// HttpRequest with context

// now, following up from the previous example, now we'll use context to control
// the maximum allowed time for a given request to process

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	api_url := "https://jsonplaceholder.typicode.com/users"

	// first, we declare a timeoutContext, which will set the maximum amoount of time
	// a given request can run
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	// deferring context cancelation
	defer cancel()

	// we use the created context in the http.NewRequestWithContext method
	req, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, api_url, nil)
	// if the request takes more than 200ms to perform, an error like the following will be shown:
	// ERR: could not process request: Get "https://jsonplaceholder.typicode.com/users": context deadline exceeded
	if err != nil {
		fmt.Printf("ERR: could not perform request: %s", err.Error())
		return
	}

	// processing request...
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("ERR: could not process request: %s", err.Error())
		return
	}

	// reading request data...
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERR: could not read response data: %s", err.Error())
	}

	// printing data length
	fmt.Printf("Request data length: %d", len(data))
}
