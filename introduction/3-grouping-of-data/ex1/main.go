//- Using a compound literal:
//- Create an array that supports 5 values of type int
//- Assign values to your indexes
//- Use range and demonstrate the array values.
//- Using format printing, demonstrate the type of the array.

package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Println("Indice", i, "Valor", v)
	}
	fmt.Printf("%T", array)

}
