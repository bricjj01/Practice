package main

import "fmt"

func main() {
	a := 64

	fmt.Printf("The variable a is equal to:\n")
	fmt.Printf("\t%d\t in decimal\n", a)
	fmt.Printf("\t%b\t in binary\n", a)
	fmt.Printf("\t%#x\t in hexadecimal\n", a)

	b := a << 1

	fmt.Printf("The variable b is equal to:\n")
	fmt.Printf("\t%d\t in decimal\n", b)
	fmt.Printf("\t%b in binary\n", b)
	fmt.Printf("\t%#x\t in hexadecimal\n", b)

	c := a >> 1

	fmt.Printf("The variable c is equal to:\n")
	fmt.Printf("\t%d\t in decimal\n", c)
	fmt.Printf("\t%b\t in binary\n", c)
	fmt.Printf("\t%#x\t in hexadecimal\n", c)
}
