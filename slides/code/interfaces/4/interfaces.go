package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// START SERVERREAD OMIT
func main() {
	r, err := http.Get("http://httpbin.org/get&test=hello")
	if err != nil {
		panic(err) // OMIT
	}
	defer r.Body.Close()

	ct, err := lineCt(r.Body)
	if err != nil {
		panic(err) // OMIT
	}
	log.Printf("got %v lines\n", ct)
}

// END SERVERREAD OMIT
func lineCt(r io.Reader) (int, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	ct := 0
	for _, v := range bs {
		if v == '\n' {
			ct++
		}
	}
	return ct, nil
}
