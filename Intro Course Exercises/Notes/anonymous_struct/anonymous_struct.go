package main

import "fmt"

// this is an anonymous struct. this is useful when you only
// need a struct once or twice, and do not wish to create a
// seperate named struct.

// anonymous structs can help to keep namespace "clean"
func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "john",
		last:  "brichetto",
		age:   27,
	}
	fmt.Printf("p1 is of type:%T\n", p1)
	fmt.Println(p1.first, p1.last, p1.age)

}
