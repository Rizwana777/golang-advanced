// Program to create a map and print its keys and values

package main

import "fmt"

func main() {

	// create a map for storing age of people
	// a := map[type]type{"key":"value"}
	age := map[string]int{
		"Rizwana": 20,
		"Naaz":    24,
		"XYZ":     30,
	}

	fmt.Println(age)

	fmt.Println("First valuse in a map : ", age["Rizwana"])

	// Update map

	age["XYZ"] = 29

	fmt.Println("Updated age is : ", age["XYZ"])
	fmt.Println("Updated map : ", age)
}
