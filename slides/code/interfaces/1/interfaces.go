package main

func main() {}

// START OMIT

// T is any type that has the listed methods
type T interface {
	F1() (int, error)
	F2(int) error
}

func f(t T) error {
	_, e := t.F1()
	e = t.F2(0)
	return e
}

// END OMIT
