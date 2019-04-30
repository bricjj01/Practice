package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("This is a great switch statement.")
	case false:
		fmt.Println("I don't like this switch statement.")
	default:
		fmt.Println("Boolean logic doesn't apply here.")
	}
}
