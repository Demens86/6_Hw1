package main

import (
	"fmt"
	"EmployeeAccounting/internal"
)

func main() {
	
	internal.Employees = append(internal.Employees,internal.Employee{
		Name: "Denis",
		Age: 28,
		Position: "CEO",
		Salary: 999999.99,
	})
	//internal.AddEmployee()
	internal.Info("fedor")
	printAllEmployees(internal.Employees)
}

func printAllEmployees(employees []internal.Employee) {
	fmt.Println("\n--- Список всех сотрудников ---")
	for i, emp := range employees {
		fmt.Printf("%d. Имя: %s, Возраст: %d, Должность: %s, Зарплата: %.2f\n",
			i+1, emp.Name, emp.Age, emp.Position, emp.Salary)
	}
	fmt.Println()
}

/*
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
*/