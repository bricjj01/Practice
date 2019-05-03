package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first: "John",
		last:  "Brichetto",
		iceCream: []string{"Peanut Butter Swirl,", "Eagles Touchdoown Sundae,",
			"Sweet Cream"},
	}

	p2 := person{
		first:    "우정윤",
		last:     "",
		iceCream: []string{"녹차맛"},
	}

	fmt.Println(p1.first, p1.last, "likes the following ice cream flavors:")
	for _, v := range p1.iceCream {
		fmt.Printf("%v\n", v)
	}

	fmt.Println(p2.first, "양은 다음과 같은 아이스크림 맛을 좋아해요:")
	for _, v := range p2.iceCream {
		fmt.Printf("%v\n", v)
	}

}
