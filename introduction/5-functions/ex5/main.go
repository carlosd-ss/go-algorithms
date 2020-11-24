//- Create a "square" type
//- Create a "circle" type
//- Create an "area" method for each type that calculates and returns the area of the figure
//- Circle area: 2 * π * radius
//- Square area: side * side
//- Create a "figure" type that defines as interface any type that has the "area" method
//- Create an "info" function that receives a "figure" type and returns the figure area
//- Create a "square" value
//- Create a "circle" value
//- Use the "info" function to demonstrate the "square" area
//- Use the "info" function to demonstrate the "circle" area

package main

import (
	"fmt"
	"math"
)

type square struct {
	lado float64
}

func (q square) area() {
	result := q.lado * q.lado
	fmt.Println("Square area:", result)
}

type circle struct {
	raio float64
}

func (c circle) area() {
	result := math.Pi * 2 * c.raio
	fmt.Println("Circle area:", result)
}

type info interface {
	area()
}

func measure(i info) {
	i.area()
}

func main() {

	x := square{
		lado: 15.0,
	}
	y := circle{
		raio: 5.25,
	}
	measure(x)
	measure(y)
}
