package main

import (
	"fmt"
)

func main() {
	var a [3]int //int array with length 3
	fmt.Println(a)

	fmt.Println("Initialize array values")
	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	fmt.Println("Short hand declaration")

	b := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(b)

}
