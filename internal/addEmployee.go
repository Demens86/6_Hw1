package internal

import "fmt"

func AddEmployee() {
	var name, position string
	var age int16
	var salary float64

	fmt.Println("Введите имя сотрудника:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Errorf("Неккоретное имя сотрудника: %w", err)
	}

	fmt.Println("Введите возраст сотрудника:")
	_, err = fmt.Scan(&age)
	if err != nil {
		fmt.Errorf("Неккоретный возраст сотрудника: %w", err)
	}

	fmt.Println("Введите должность сотрудника:")
	_, err = fmt.Scanln(&position)
	if err != nil {
		fmt.Errorf("Неккоретная должность сотрудника: %w", err)
	}

	fmt.Println("Введите зарплату сотрудника:")
	_, err = fmt.Scan(&salary)
	if err != nil {
		fmt.Errorf("Неккоретная зарплата сотрудника: %w", err)
	}

	newEmployee :=  Employee{
		Name: name,
		Age: age,
		Position: position,
		Salary: salary,
	}

	Employees = append(Employees, newEmployee)

}