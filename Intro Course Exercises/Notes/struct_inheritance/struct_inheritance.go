package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	l2k bool
}

func main() {
	p1 := person{
		first: "john",
		last:  "brichetto",
		age:   27,
	}

	p2 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
			age:   32,
		},
		l2k: true,
	}

	fmt.Println("name:", p1.first, p1.last, "\nage:", p1.age, "\nlicense to kill:", false)
	fmt.Println("name:", p2.first, p2.last, "\nage:", p2.age, "\nlicense to kill:", p2.l2k)
}
