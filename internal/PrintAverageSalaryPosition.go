package internal

import "fmt"

func PrintAverageSalaryPosition(employees []Employee) {
	sumSalaryPosition := make(map[string]float64)
	countEmployeePosition := make(map[string]float64)

	//Суммирование зарплат по должности и подсчет кол-ва сотрудников
	//на этой должности
	for _, emp := range employees {

		sumSalaryPosition[emp.Position] += emp.Salary
		countEmployeePosition[emp.Position]++
	}

	averageSalaryPosition := make(map[string]float64, len(sumSalaryPosition))

	for position, sumSalary := range sumSalaryPosition {
		averageSalaryPosition[position] = sumSalary / countEmployeePosition[position]
	}

	fmt.Println("Вывод средней зарплаты по должности:")
	for position, salary := range averageSalaryPosition {
		fmt.Printf("%s: %.2f\n", position, salary)
	}
}