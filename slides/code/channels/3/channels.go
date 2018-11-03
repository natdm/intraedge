package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func main() {
	wg := &sync.WaitGroup{}
	totals := make(chan int)
	for nums := range generateInts(4) {
		wg.Add(1)
		go addSlice(nums, wg, totals)
	}

	go func() {
		wg.Wait()
		close(totals)
	}()

	for total := range totals {
		log.Printf("got total %v\n", total)
	}
}

func addSlice(s []int, wg *sync.WaitGroup, total chan<- int) {
	defer wg.Done()
	var n int
	for _, v := range s {
		n += v
	}
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	total <- n
}

// generateInts sends random integers down a channel
func generateInts(batchSize int) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		var b []int
		for _, n := range makeRandom() {
			b = append(b, n)
			if len(b) == batchSize {
				out <- b
				b = []int{}
			}
		}
	}()
	return out
}

// END OMIT
func makeRandom() []int {
	nums := make([]int, 100)
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := 0; i < 100; i++ {
		nums[i] = r.Intn(100)
	}
	return nums
}
