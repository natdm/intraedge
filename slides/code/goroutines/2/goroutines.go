package main

import (
	"errors"
	"log"
	"time"
)

func callAPI(url string) error { return errors.New("boom") }

// START OMIT
// now reports an error if failed
func reportError(e error) error {
	time.Sleep(time.Second)
	log.Println(e)
	return errors.New("failed")
}

func main() {
	err := callAPI("")
	if err != nil {
		go func() {
			err2 := reportError(err)
			if err2 != nil {
				panic(err2)
			}
		}()
	}
	// ends
}

// END OMIT
