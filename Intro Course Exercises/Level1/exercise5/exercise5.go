package main

import "fmt"

//create new variable type
type johnstype int

var x johnstype
var y int

func main() {
	//print variable information
	fmt.Printf("\n%v", x)
	fmt.Printf("\n%T", x)
	//change the variable
	x = 42
	fmt.Printf("\n%v", x)
	//convert to underlying type
	y = int(x)
	fmt.Printf("\n%v", y)
	fmt.Printf("\n%T", y)
}
