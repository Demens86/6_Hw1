package addEmployee

import (
	"EmployeeAccounting/internal"
	"fmt"
)

func AddEmployee(employees []internal.Employee) error {
	name, err := InputName()
	if err != nil {
		fmt.Println(err)
		return err
	}
	age, err := InputAge()
	if err != nil {
		fmt.Println(err)
		return err
	}
	position, err := InputPosition()
	if err != nil {
		fmt.Println(err)
		return err
	}
	salary, err := InputSalary()
	if err != nil {
		fmt.Println(err)
		return err
	}

	newEmployee :=  internal.Employee{
		Name: name,
		Age: age,
		Position: position,
		Salary: salary,
	}

	internal.Employees = append(internal.Employees, newEmployee)
	fmt.Printf("Сотрудник '%s' добавлен!\n", name)
	return nil
}