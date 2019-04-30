package main

import "fmt"

func main() {
	x := "favSport"

	switch x {
	case "okaySport":
		fmt.Println("Basketball is an okay sport.")
	case "greatSport":
		fmt.Println("Football is a great sport.")
	case "favSport":
		fmt.Println("Underwater basket weaving is everyone's favorite sport.")
	default:
		fmt.Println("Sports are boring, let's write code instead!")
	}
}
