//- Create and use an anonymous function.
package main

import "fmt"

func main() {
	func() {
		fmt.Println("-----------Go-------------")
	}()
}
