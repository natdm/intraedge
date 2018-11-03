package main

type ptrInt *int

func pointers() {
	// value
	i := 5
	// reference
	p := &i
	// dereference
	p2 := *p
	_ = p2
	// make p nil
	p = nil
}

// accepting pointers
// modifying pointers
// pointer receivers
// when to use pointers
// nil vs null
