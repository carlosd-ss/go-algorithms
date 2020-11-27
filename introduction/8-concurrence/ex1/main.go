//- In addition to the main goroutine, create two other goroutines.
//- Each additional goroutine must make a separate print.
//- Use waitgroups to make your goroutines end before the program ends.

package main

import (
	"fmt"
	"sync"
	"time"
)

func x(wg *sync.WaitGroup) {
	defer wg.Done()
	count := 7
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func y(wg *sync.WaitGroup) {
	defer wg.Done()
	sl := []string{"A", "B", "C", "D", "E", "F", "G"}
	for _, v := range sl {
		time.Sleep(1 * time.Second)
		fmt.Println(v)
	}
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go x(wg)
	go y(wg)
	wg.Wait()
}
