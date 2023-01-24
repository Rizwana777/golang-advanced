package main

import "fmt"

func main() {
	var a int
	var b float32
	var c int

	a, b = 2, 5.5

	c = a + int(b)

	fmt.Println(c)
}
