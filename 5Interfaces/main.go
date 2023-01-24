package main

import "fmt"

type Vehicle interface {
	Mileage() float64
	AverageSpeed() float64
}

type car struct {
	distance float64
	fuel     float64
	time     float64
}

func (b car) Mileage() float64 {
	return b.distance / b.fuel
}

func (b car) AverageSpeed() float64 {
	return b.distance / b.time
}

func main() {
	var m Vehicle
	m = car{100, 50.9, 7.9}
	q := car{80, 77.8, 2.9}

	fmt.Printf("Type of m %T\n", m)
	fmt.Printf("Type of 1 %T\n", q)

	fmt.Println("Mileage of first ride : ", m.Mileage())
	fmt.Println("Mileage of first ride : ", q.Mileage())

	//fmt.Printf("%0.2f", m.Mileage())
}
