package main

import (
	"fmt"
	"strings"
)

// Functions are first class. Means recursion and closures are fair
// game. You can also do shorthand declaration in functions
func helloWorld() {
	// probably bad since it has no input and no output
	// reconsider what you're doing
}

type user struct {
	name string
}

// short variable names
func applyUserName(u user, name string) error {
	if len(u.name) > 0 {
		return fmt.Errorf("user has name of %s, need empty name", u.name)
	}
	u.name = name
	return nil
}

// Where could this crash? Add error handling
func userFirstName(u user) string {
	return strings.Split(u.name, " ")[0]
}

// fun with functions

func index(ss []string, s string) int {
	for i, v := range ss {
		if v == s {
			return i
		}
	}
	return -1
}

func contains(ss []string, s string) bool {
	return index(ss, s) >= 0
}

func mapStr(ss []string, fs ...func(string) string) []string {
	out := make([]string, len(ss))
	for i, v := range ss {
		val := v
		for _, f := range fs {
			val = f(v)
		}
		out[i] = val
	}
	return out
}

func enforeceRocks(s string) string {
	if !strings.HasSuffix(strings.ToLower(s), "rocks") {
		return strings.Split(s, " ")[0] + " rocks"
	}
	return s
}

var upperNames = mapStr([]string{
	"go rocks", "rust is tricky", "scala rocks"},
	strings.ToUpper, enforeceRocks,
)
