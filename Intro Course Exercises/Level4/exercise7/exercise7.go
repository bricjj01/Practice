package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	xxs := [][]string{xs1, xs2}
	for _, xs := range xxs {
		for _, s := range xs {
			fmt.Println(s)
		}
	}
}
