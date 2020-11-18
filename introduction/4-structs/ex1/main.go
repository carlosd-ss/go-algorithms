//- Create a type "person" with underlying type struct, which can contain the following fields:
//- Name
//- Surname
//- Favorite flavors of ice cream
//- Create two "person" values and demonstrate these values, using range on the slice that contains the ice cream flavors.
package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Flavors  []string
}

func main() {
	p1 := Person{
		Name:     "Mr",
		Lastname: "Robot",
		Flavors:  []string{"Açai", "Strawberry", "Chocolate"},
	}
	p2 := Person{
		Name:     "Will",
		Lastname: "Smith",
		Flavors:  []string{"Açai", "Strawberry", "Chocolate"},
	}

	fmt.Println(p1.Name, p1.Lastname)
	for _, v := range p1.Flavors {
		fmt.Println("\t", v)
	}

	fmt.Println(p2.Name, p2.Lastname)
	for _, v := range p2.Flavors {
		fmt.Println("\t", v)
	}

}
