package main

import (
	"fmt"
	"io"
	"net/http"
)

// interfaces describe behavior
// empty interfaces
// embed interfaces
// use stdlib interfaces
// Reader, Writer, Closer, Seeker, Sort, json.Marshaler
// look at
// 	json.RawMessage
//  os.File
//  Stringer
// 	http.Handler
//  Weirdness of ioutil.ReadAll and io.EOF
// 	Error

// Usecases: generalization, mocking and testing, etc.

// Pitfalls:
// Interfaces that take custom types, can they be avoided?
// Taking in huge interfaces when you only need a piece of it
// Hard to find interfaces, or what interfaces your code adheres to

// Uploader can handle file uploads
type Uploader interface {
	Upload(r io.Reader) (int, error)
}

// Downloader can download a file
type Downloader interface {
	Download(w io.Writer) (int, error)
}

type Deleter interface {
	// ...
}

type ServerError int

const (
	NotFound ServerError = 404
)

func (e ServerError) String() string {
	switch e {
	case NotFound:
		return "NotFound"
	}

	return "No Status Given"
}

func (e ServerError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(int(e))
	fmt.Fprint(w, e)
}
