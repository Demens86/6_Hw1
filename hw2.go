package main

import (
	"fmt"
)
type Employee struct {
	Name string
	Age int16
	Positon string
	Salary float64
}

func main() {
	/*
	type Employee struct {
		Name string
		Age int16
		Positon string
		Salary float64
	}
	*/
	x := Employee{
		Name: "Denis",
		Age: 28,
		Positon: "CEO",
		Salary: 999999.99,
	}
	x.info()
}

func (e Employee) info() {
	fmt.Printf("Name: %s | Age: %d | Position: %s | Salary: %.2f\n", e.Name, e.Age, e.Positon, e.Salary)
}