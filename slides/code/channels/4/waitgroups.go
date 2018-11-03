package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// pass in 'i' so it's not caught by for loop
		go func(i int) {
			time.Sleep(time.Second)
			log.Printf("Done with %v\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println("Done with all")
}
