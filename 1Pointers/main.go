package main

import "fmt"

func main() {
	var count int = 100

	p := &count

	fmt.Println("Address of count variable", p)

	v := *p

	fmt.Println("Value of count variable", v)

	// change value of v variable

	v = 200

	fmt.Println("Value of v after update", v)

	// Verify whether this changes the value of count variable or no

	fmt.Println("Verifying count value", count)

}
