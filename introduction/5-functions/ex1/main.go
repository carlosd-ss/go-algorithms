//- Create a function that returns an int
//- Create another function that returns an int and a string
//- Call the two functions
//- Demonstrate your results
package main

import "fmt"

func num(x int) int {
	return x
}
func alpha(x int, y string) (int, string) {
	return x, y
}

func main() {
	x := 10
	y := "abc"
	fmt.Println(num(x))
	fmt.Println(alpha(x, y))
}
