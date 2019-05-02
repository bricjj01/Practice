package main

import "fmt"

func main() {
	arr := [5]int{345, 65784, 24536, 2457, 94765}
	fmt.Println(arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
