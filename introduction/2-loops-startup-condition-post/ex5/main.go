//- Create a program that uses the switch statement, where the switch statement is a string-type variable with the identifier "EsporteFavorito".

package main

import "fmt"

func main() {
	sport := "Bikecross"

	switch sport {
	case "Bikecross":
		fmt.Println("Bike")
	case "Futebol":
		fmt.Println("Futebol")
	case "Volei":
		fmt.Println("Volei")
	}

}
