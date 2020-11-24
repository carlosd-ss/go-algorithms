//- Demonstrate the functioning of a closure.
//- That is: create a function that returns another function, where this other function makes use of a variable in addition to its scope.

package main

import "fmt"

func main() {
	a := i()
	b := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
