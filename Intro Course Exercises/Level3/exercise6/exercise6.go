package main

import "fmt"

func main() {
	x := "lol"

	switch x {
	case "ouch", "owie":
		fmt.Println("Hey, stop that!")
	case "lol", "haha":
		fmt.Println("Wow, that's funny!")
	default:
		fmt.Println("Beep boop")
	}
}
