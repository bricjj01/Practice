package main

import "fmt"

func main() {
	myMap := make(map[string][]string)

	myMap["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	myMap["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	myMap["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}

	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
