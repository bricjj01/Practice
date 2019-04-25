package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
	)

	fmt.Println("it's the year:", 2019-a)
	fmt.Println("last year was:", 2019-b)
	fmt.Println("the year before that was:", 2019-c)
	fmt.Println("the year before that was:", 2019-d)
}
