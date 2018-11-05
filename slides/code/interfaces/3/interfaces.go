package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// START READER OMIT

// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// ... 2000 years later...
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
type Reader interface {
	Read(p []byte) (int, error)
}

// END READER OMIT

// START READALL OMIT

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(r io.Reader) ([]byte, error) {
	//  omitted
	return nil, nil // OMIT
}

// END READALL OMIT

type File struct{}

// START FILEREADFN OMIT

// Read reads up to len(b) bytes from the File.
// It returns the number of bytes read and any error encountered.
// At end of file, Read returns 0, io.EOF.
func (f *File) Read(b []byte) (n int, err error) {
	// omitted
	return 0, nil // OMIT
}

// END FILEREADFN OMIT

// START FILEREAD OMIT
func main() {
	// omit errors and Close for brevity
	f, _ := os.Open("./input.txt")
	ct, _ := lineCt(f)
	fmt.Printf("lines: %v\n", ct)
}

func lineCt(r io.Reader) (int, error) {
	bs, _ := ioutil.ReadAll(r)
	ct := 0
	for _, v := range bs {
		if v == '\n' {
			ct++
		}
	}
	return ct, nil
}

// END FILEREAD OMIT
