package main

import (
	"log"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	for num := range generateInts() {
		log.Println(num)
	}
}

// generateInts sends random integers down a channel
func generateInts() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range makeRandom() {
			out <- n
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
