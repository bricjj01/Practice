package main

import (
	"fmt"
)

const (
	// iota initializes as a zero valued int
	_ = iota
	// and increments every time it's called
	// iota = 1
	kb = 1 << (iota * 10)
	// iota = 2
	mb = 1 << (iota * 10)
	// iota = 3
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
