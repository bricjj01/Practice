package main

import "fmt"

const (
	a = iota
	b
)

func main() {
	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
}
