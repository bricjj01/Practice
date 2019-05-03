package main

import "fmt"

func main() {
	myMap := make(map[string][]string)

	myMap["Brichetto"] = []string{"Peanut Butter Swirl",
		"Eagles Touchdown Sundae", "Sweet Cream"}

	myMap["우정윤"] = []string{"눅차"}

	for k, v := range myMap {
		fmt.Println(k)
		for i, a := range v {
			fmt.Println(i, a)
		}
	}
}
