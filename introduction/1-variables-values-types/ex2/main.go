//- Use var to declare three variables. They must have package-level scope. Do not assign values to these variables. Use the following identifiers and types for these variables:
//1. Identifier "x" must have type int
//2. Identifier "y" must have type string
//3. Identifier "z" must be of type bool
//- In the main function:
//1. Demonstrate the values of each identifier
//2. The compiler has assigned values for these variables. What are these values called?

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
	//The compiler has assigned values for these variables. What are these values called?
	//-Zero values

}
