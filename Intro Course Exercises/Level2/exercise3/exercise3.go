package main

import "fmt"

// because this constant is typed it can only be used where strings, AND ONLY STRINGS, can be used
const alpha string = "alpha"

// this constant is untyped, it can be used with a new type whose underlying type is string
const omega = "omega"

func main() {
	fmt.Println("This is the   typed constant alpha:", alpha)
	fmt.Println("This is the untyped constant omega:", omega)
}
