//Create a "person" struct
//- Create a function called mudeMe that has * person as a parameter. This function should change a value stored in the * person address.

package main

import "fmt"

type person struct {
	name string
}

func change(p *person) {
	p.name = "Cofee"
	fmt.Println(p.name)
}

func main() {
	batman := person{name: "Batman"}
	change(&batman)
}
