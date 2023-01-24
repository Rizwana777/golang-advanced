package main

import (
	"fmt"
	"sync"
)

// Syntax to create a channel ch := make(chan type)
// For recieving data in channel ch <-  data
// For reading data from channel data := <- ch

var wg sync.WaitGroup

func foo(ch chan int, value int) {
	defer wg.Done()
	// sending data to channels
	ch <- value * 2
}

func main() {

	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(c, i)
	}
	wg.Wait()
	fmt.Println("wait for all gorotines to finish")
	close(c)

	for i := range c {
		fmt.Println(i)
	}

}
