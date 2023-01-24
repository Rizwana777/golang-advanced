package main

import "fmt"

// Syntax to create a channel ch := make(chan type)
// For recieving data in channel ch <-  data
// For reading data from channel data := <- ch

func main() {

	c := make(chan int)

	go foo(c, 4)
	go foo(c, 2)

	v1 := <-c
	v2 := <-c

	fmt.Println(v1)
	fmt.Println(v2)

}

func foo(ch chan int, value int) {
	// sending data to channels
	ch <- value * 2
}
