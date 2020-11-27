//- Make this code work: https://play.golang.org/p/j-EA6003P0
//- Using an anonymous self-executing function
//- Using buffer

package main

import "fmt"

func main() {
	// func
	c := make(chan int)
	go func() {
		c <- 50
	}()
	fmt.Println(<-c)

	// buffer
	d := make(chan int, 1)
	d <- 100
	fmt.Println(<-d)
}
