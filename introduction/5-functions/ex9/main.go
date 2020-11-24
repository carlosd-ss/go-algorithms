//Callback: pass a function as an argument to another function.

package main

import "fmt"

func text() {
	fmt.Println("Hellow!")
}

func ex(x func()) {
	fmt.Println("Exec functions:")
	x()
}
func main() {
	ex(text)
}
