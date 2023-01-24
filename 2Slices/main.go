package main

import (
	"fmt"
)

func main() {
	a := [5]int{55, 56, 57, 58, 59}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)

	fmt.Println("Another way of creating slice")

	c := []int{3, 4, 5} //creates and array and returns a slice reference
	fmt.Println(c)

	fmt.Println("Modified slice")
	c[1] = 7

	fmt.Println(c)

	fmt.Println("Iterate over a slice using for each loop")
	for _, v := range c {
		fmt.Println(v)
	}
}
