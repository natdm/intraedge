package main

import (
	"errors"
	"log"
	"time"
)

func callAPI(url string) error { return errors.New("boom") }

func reportError(e error) {
	time.Sleep(time.Second)
	log.Println(e)
}

func main() {
	err := callAPI("")
	if err != nil {
		go reportError(err)
	}
	// ends
}
