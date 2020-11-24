//- Create a struct type "person" that contains the fields:
//- name
//- surname
//- age
//- Create a "person" method that demonstrates the full name and age;
//- Create a "person" value;
//- Use the method created to demonstrate this value.

package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (p person) info() {
	fmt.Println("My name is", p.name, p.surname, "and my age is", p.age)
}

func main() {
	batman := person{
		name:    "Batman",
		surname: "Dark",
		age:     30,
	}
	batman.info()
}
