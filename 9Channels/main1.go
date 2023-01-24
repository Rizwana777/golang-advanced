package main

import (
	"fmt"
)

// Syntax to create a channel ch := make(chan type)
// For recieving data in channel ch <-  data
// For reading data from channel data := <- ch

func main() {

	c := make(chan int)

	for i := 0; i < 10; i++ {
		go foo(c, i)
	}

	//close(c)

	for i := range c {
		fmt.Println(i)
	}

}

func foo(ch chan int, value int) {
	// sending data to channels
	ch <- value * 2
}
