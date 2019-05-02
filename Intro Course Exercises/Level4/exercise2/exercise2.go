package main

import "fmt"

func main() {
	xi := []int{
		1, 2, 4, 8, 16,
		32, 64, 128, 256, 512,
	}
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", xi)
}
