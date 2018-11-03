package main

import "fmt"

// http://www.golangbootcamp.com/book/types
// https://golang.org/doc/effective_go.html#embedding

// name is a persons name, with the first, middle, andlast name as a string.
type name string

func n() {
	// this should be inferred, `var n =, or n :=`
	var n name = name("Nathan Hyland")
	fmt.Printf("%s\n", n)
}

type creditCard int
type ssn int
type people map[name]person
type friends []person
type ssnFunc func(ssn) []creditCard

// converting user defined types to underlying types
// types and methods

// Why create custom types? Readability and methods.
