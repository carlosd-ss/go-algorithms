//- Using the previous solution, place the values of type "person" on a map, using the surnames as key.
//- Demonstrate the map values using range.
//- The different flavors must be demonstrated using another range, within the previous range.

package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Flavors  []string
}

func main() {

	meumapa := make(map[string]Person)
	meumapa["Robot"] = Person{
		Name:     "Mr",
		Lastname: "Robot",
		Flavors:  []string{"Açai", "Strawberry", "Chocolate"},
	}
	meumapa["Smith"] = Person{
		Name:     "Will",
		Lastname: "Smith",
		Flavors:  []string{"Açai", "Strawberry", "Chocolate"},
	}

	for _, v := range meumapa {
		fmt.Println("My name is", v.Name, "and my favorite ice creams are:")
		for _, v := range v.Flavors {
			fmt.Println("–", v)
		}
	}
}
