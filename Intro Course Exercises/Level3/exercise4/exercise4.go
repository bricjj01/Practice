package main

import "fmt"

func main() {
	i := 1991
	for {
		fmt.Println(i)
		i++
		if i == 2020 {
			break
		}
	}
}
