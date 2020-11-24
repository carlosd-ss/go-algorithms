//- Create a function that returns a function.
//- Assign the returned function to a variable.
//- Call the returned function.
package main

import "fmt"

func returnfunc() func() {
	return func() {
		fmt.Println("Hi, There!")
	}
}
func main() {
	x := returnfunc()
	x()
}
