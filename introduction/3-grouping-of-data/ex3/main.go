//- Based on the previous exercise, use slicing to demonstrate the values:
//- From the first to the third slice item (including the third item!)
//- From the fifth to the last item of the slice (including the last item!)
//- From the second to the seventh slice item (including the seventh item!)
//- From the third to the penultimate item of the slice (including the penultimate item!)
//- Challenge: obtain the same result as above using the len () function to determine the penultimate item

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}
