//- Assign a function to a variable.
//- Call this function.

package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Gopher")
	}
	x()
}
