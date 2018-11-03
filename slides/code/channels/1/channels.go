package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		log.Println("sending")
		ch <- 5
	}()

	log.Println("receiving")
	out := <-ch
	close(ch)
	log.Println(out)
}
