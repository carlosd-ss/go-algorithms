//- Using a compound literal:
//- Create a slice of type int
//- Assign 10 values to it
//- Use range to demonstrate all these values.
//- And use format printing to demonstrate your type.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range slice {
		fmt.Println("Indice", i, "Valor", v)
	}
	fmt.Printf("%T", slice)

}
