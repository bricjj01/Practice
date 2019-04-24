package main

import "fmt"

//create new variable type
type johnstype int

var x johnstype

func main() {
	fmt.Printf("\n%v", x)
	fmt.Printf("\n%T", x)
	x = 42
	fmt.Printf("\n%v", x)
}
