package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, "divided by 4 yields the remainder", i%4)
	}
}
