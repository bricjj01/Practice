package main

import "fmt"

func main() {
	jb := struct {
		name  map[string]string
		likes []string
	}{
		name: map[string]string{
			"first": "John",
			"last":  "Brichetto",
		},
		likes: []string{"golang", "league of legends", "learning"},
	}

	fmt.Printf("%T", jb)
	fmt.Println(jb)
}
