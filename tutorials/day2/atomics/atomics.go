package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64
	go func() {
		for i := 0; i < 1000000000000; i++ {
			go func() {
				<-time.After(time.Millisecond)
				atomic.AddUint64(&ops, 1)
			}()
		}
	}()

	for i := 0; i < 1000; i++ {
		fmt.Printf("ops: %v\n", atomic.LoadUint64(&ops))
	}
}
