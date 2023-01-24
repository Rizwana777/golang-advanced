package main

import (
	"fmt"
	"time"
)

func Add(a int, b int, i int) {
	c := a + b

	time.Sleep(5 * time.Millisecond)

	fmt.Println(i, " Time : ", time.Since(time.Now()))
	fmt.Println("Add : ", c)
}

func main() {
	go Add(4, 4, 1)
	Add(8, 7, 2)
}
