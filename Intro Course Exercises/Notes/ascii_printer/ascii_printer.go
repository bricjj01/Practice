package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("i value: %v\tUnicode: %#U\t ASCII character: %c\n", i, i, i)
	}
}
