package main

import (
	"log"
	"time"
)

type ball struct {
	hits int
}

func main() {
	table := make(chan ball)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case ball := <-table:
				ball.hits++
				table <- ball
			case <-done:
				log.Println("ping done")
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case ball := <-table:
				ball.hits++
				table <- ball
			case <-done:
				log.Println("pong done")
				return
			}
		}
	}()

	table <- ball{}
	<-time.After(time.Second * 2)
	close(done)
	<-time.After(time.Second * 2)
	log.Println((<-table).hits)
}
