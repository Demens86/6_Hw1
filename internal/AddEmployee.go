package internal

import (
	"bufio"
	"fmt"
	"os"
)

func AddEmployee() {
	sc := bufio.NewReader(os.Stdin)

	newEmployee :=  Employee{
		Name: name,
		Age: age,
		Position: position,
		Salary: salary,
	}

	Employees = append(Employees, newEmployee)

}