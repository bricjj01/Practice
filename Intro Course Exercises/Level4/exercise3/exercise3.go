package main

import "fmt"

func main() {
	xi := []int{
		1, 2, 4, 8, 16,
		32, 64, 128, 256, 512,
	}
	fmt.Println(xi[:5])
	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:7])
}
