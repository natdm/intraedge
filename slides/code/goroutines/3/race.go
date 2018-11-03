package main

import (
	"log"
	"sync"
)

func main() {

}

// START RACE
func hasRace() {
	m := make(map[string]bool)
	wg := &sync.WaitGroup{}

	for _, v := range []string{"Go", "Rust", "Scala"} {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			m[v] = true
		}(v)
	}

	wg.Wait()
	log.Println("done")
}

// END RACE

// START RACEFIX1
func fixRace1() {
	m := make(map[string]bool)
	done := make(chan struct{})
	write := make(chan string)
	wg := &sync.WaitGroup{}
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		for v := range write {
			m[v] = true
		}
	}()
	for _, v := range []string{"Go", "Rust", "Scala"} {
		wg.Add(1)
		go func(v string) { defer wg.Done(); write <- v }(v)
	}
	<-done
	log.Println("done")
}

// END RACEFIX1

// START RACEFIX2
func fixRace2() {
	m := make(map[string]bool)
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	for _, v := range []string{"Go", "Rust", "Scala"} {
		wg.Add(1)
		go func(v string) {
			mtx.Lock()
			m[v] = true
			mtx.Unlock()
			wg.Done()
		}(v)
	}

	wg.Wait()
	log.Println("done")
}

// END RACEFIX2
