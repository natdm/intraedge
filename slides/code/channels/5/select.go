package main

import (
	"log"
	"math/rand"
	"time"
)

var s1 = rand.New(rand.NewSource(time.Now().UnixNano()))

// START OMIT
func main() {
	for i := 0; i < 10; i++ {
		select {
		case b := <-fn(1000):
			log.Printf("got %v", b)
		case <-fn(1000):
			log.Println("got one")
		case <-time.After(time.Second):
			log.Println("timed out")
			// default:
			// 	log.Println("default")
		}
	}
}

func fn(i int) <-chan bool {
	out := make(chan bool)
	go func() {
		time.Sleep(time.Duration(s1.Intn(i)))
		out <- true
	}()
	return out
}

// END OMIT
