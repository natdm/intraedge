package models

import "io"

// Storage has all the necessary functionality for storage
type Storage interface {
	Save(name string, r io.Reader) error
	File(name string, w io.Writer) error
	List() (names []string, err error)
}

type NamesResponse struct {
	Names []string `json:"names"`
}
