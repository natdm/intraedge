package main

type unexported struct {
	name    string
	age     int
	male    bool
	friends []string
}

// SomeExported ...
type SomeExported struct {
	Name    string
	Age     int
	Male    bool
	friends []name
}

type person struct {
	name name
	age  int
}

// having an instance of struct vs embedding structs
// embedding interfaces -- amazing but advanced
// JSON/XML and other notations
