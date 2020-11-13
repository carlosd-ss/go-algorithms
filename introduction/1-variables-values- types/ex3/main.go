//- Using the solution from the previous exercise:
//1. In package-level scope, assign the following values to variables:
//1. for "x" assign 42
//2. for "y" assign "James Bond"
//3. for "z" set it to true
//2. In the main function:
//1. Use fmt.Sprintf to assign all of these values to a single variable. Do this string assignment to a variable named "s" using the short declaration operator.
//2. Demonstrate the variable "s"

package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
	fmt.Println(s)
}
