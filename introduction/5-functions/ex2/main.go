//- Create a function that receives a variable parameter of type int and returns the sum of all received ints;
//- Pass a value of type slice of int as an argument to the function;
//- Create another function, it must receive a value of type slice of int and return the sum of all elements of the slice;
//- Pass a value of type slice of int as an argument to the function.

package main

import "fmt"

func sum(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}
func sumslice(x []int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}

func main() {
	sli := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(sli...))
	fmt.Println(sumslice(sli))
}
