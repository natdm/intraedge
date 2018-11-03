package main

import (
	"fmt"
	"math/rand"
)

func _if() {
	if rand.Int() > 10 {
		fmt.Println("number was greater than 10")
	}

	// scope of I ends after if/else statement, unless assigned elsewhere
	if i := rand.Int(); i > 10 {
		fmt.Printf("%v was greater than 10", i)
	}
}

func _switch() {
	// no break needed, not a fallthrough
	x := 5

	switch x {
	case 5:
		{
			println("was 5")
		}
	default:
		{
			println("default")
		}
	}

	// omit brackets, not needed

	fn := func() int {
		return 10
	}

	fn2 := func() int {
		panic("nope")
	}

	// 5 runs, 10 runs, fn2() will not run
	switch x = fn(); x {
	case 5:
		println("was 5")
	case 10:
		fmt.Printf("was %v", x)
	case fn2():
		fmt.Println("got to fn2")
	default:
		println("awkward")
	}

	// switching off of nothing is shorthanded if/else
	switch {
	case x < 10:
		fmt.Println("x is less than 10")
	case x == 11:
		fmt.Println("x is exactly 11")
	default:
		fmt.Println("x is greater than 11")
	}
}

// Go has only a for loop. No while or dowhile.
func _forloops() {
	// init, condition, post
	for i := 0; i < 10; i++ {
		fmt.Printf("at %v\n", i)
	}

	sum := 0
	// just the condition (since init and post are optioonal)
	// It's now a while loop
	for sum < 100 {
		sum += 10
	}
	fmt.Printf("sum is at %v", sum)

	sum = 0
	// loop forever, or until a break
	for {
		sum += 10
		if sum > 100 {
			break
		}
	}
}
