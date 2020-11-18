//- Starting with the following slice:
//- x: = [] int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//- Attach the value 52;
//- Attach the values 53, 54 and 55 to it using a single declaration;
//- Demonstrate the slice;
//- Attach the following slice to it:
//- y: = [] int {56, 57, 58, 59, 60}
//- Demonstrate the slice x.
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
