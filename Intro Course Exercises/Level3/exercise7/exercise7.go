package main

import "fmt"

func main() {
	x := "wowzers"

	if x == "oh no" {
		fmt.Println("that's bad")
	} else if x == "lol" {
		fmt.Println("that's funny")
	} else {
		fmt.Println("I don't know what to say to that")
	}
}
