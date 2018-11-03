package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := iter(1, 3, 5, 7, 9, 11)
	c2 := iter(2, 4, 6, 8, 10, 12)
	c4 := iter(9, 12, 15, 30, 40, 50505, 30139, 1231)
	c5 := iter(3245, 1234, 142, 4325, 124)
	c6 := merge(merge(c1, c2), c4, c5)
	for v := range c6 {
		fmt.Println(v)
	}
}

// https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func iter(is ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, i := range is {
			c <- i
		}
		close(c)
	}()
	return c
}
