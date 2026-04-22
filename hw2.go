package main

import (
	"fmt"
)
type Employee struct {
	Name string
	Age int16
	Position string
	Salary float64
}

var employees []Employee

func main() {
	/*
	type Employee struct {
		Name string
		Age int16
		Positon string
		Salary float64
	}
	*/
	employees = append(employees,Employee{
		Name: "Denis",
		Age: 28,
		Position: "CEO",
		Salary: 999999.99,
	})
	employees = addEmployee(employees, "Vasya", 23, "Developer", 54334.55)
	printAllEmployees(employees)
}

func (e Employee) info() {
	fmt.Printf("Name: %s | Age: %d | Position: %s | Salary: %.2f\n", e.Name, e.Age, e.Position, e.Salary)
}

func addEmployee(employees []Employee, name string, age int16, position string, salary float64) []Employee{
	newEmployee :=  Employee{
		Name: name,
		Age: age,
		Position: position,
		Salary: salary,
	}
	employees = append(employees, newEmployee)
	return employees
}

func printAllEmployees(employees []Employee) {
	fmt.Println("\n--- Список всех сотрудников ---")
	for i, emp := range employees {
		fmt.Printf("%d. Имя: %s, Возраст: %d, Должность: %s, Зарплата: %.2f\n",
			i+1, emp.Name, emp.Age, emp.Position, emp.Salary)
	}
	fmt.Println()
}