package main

import (
	"errors"
	"fmt"
)

// custom error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide through zero")
	}

	return a / b, nil
}

func main() {
	num, err := divide(100, 0)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Println("Number: ", num)
	}
}
