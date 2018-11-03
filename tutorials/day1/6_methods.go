package main

import "fmt"

type user2 struct {
	name       string
	age        int
	profession string
}

// notice this is the exact same as applyUserName?
func (u user2) applyName(name string) error {
	if len(u.name) > 0 {
		return fmt.Errorf("user has name of %s, need empty name", u.name)
	}
	u.name = name
	return nil
}

// add methods to types in 2_types.go
