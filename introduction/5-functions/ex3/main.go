//- Use the defer declaration in a way that demonstrates that its execution only occurs at the end of the context to which it belongs.

package main

import "fmt"

func main() {
	defer fmt.Println("First")
	fmt.Println("Second")
}
