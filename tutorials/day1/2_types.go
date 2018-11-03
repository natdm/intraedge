package main

import "math/cmplx"

// Verbose assignment
// var <Name> <Type> = <Value>
var i int = 0

// If type is 'obvious', shorthanding is OK bny removing inferred type
// Strings are read-only byte slices
var s = ""

// Default is zero value. Need to assign type.
var b bool

var (
	i32  int32  = 0
	i64  int64  = 0
	ui   uint   = 0
	ui32 uint32 = 0
	bt   byte   = 0

	// Arrays are a type with a length.
	// [n]T{}
	// or var name [n]T
	// slices and arrays can be ranged over
	arrayIntA [2]int = [2]int{7, 2}
	arrayIntB        = [2]int{7, 2}
	// Other ways to declare arrays

	// Slices are *growable arrays*
	sliceInt = []int{7, 2}
	// maps can be ranged over
	mapStrings = map[string]string{}

	cmplex complex128 = cmplx.Sqrt(-5 + 12)

	// functions are types too
	xByX func(int) int = func(x int) int {
		return x * x
	}

	// shorthanded function
	yByY = func(y int) int {
		return y * y
	}
)

// MAPS, SLICES, ARRAYS
