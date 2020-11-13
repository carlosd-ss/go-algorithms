//- Demonstrate the rest of the division by 4 of all numbers between 10 and 100
package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}
