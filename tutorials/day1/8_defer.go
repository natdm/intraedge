package main

import (
	"net/http"
	"time"
)

func callAPI() {
	// calls API
}

func deferring() {
	defer callAPI()
	// do things
	// defer gets called last
}

func callAPIWithReturn() (int, error) {
	return http.StatusOK, nil
}

func deferringWithReturns() {
	defer func() {
		statusCode, err := callAPIWithReturn()
		if err != nil {
			// do things with err
			return
		}
		_ = statusCode
	}()

	// do logical things
}

// defer is called from last to first order when stacked. Popping off of stack
func stackedDeferrs() {
	defer println("first in")  // ran second
	defer println("second in") // ran first
	time.Sleep(time.Second)
}
