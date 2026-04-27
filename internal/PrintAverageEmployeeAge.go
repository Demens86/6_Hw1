package internal

import "fmt"

func PrintAverageEmployeeAge() {
	sumAge, count := 0, 0

	for _, emp := range Employees {
		sumAge += emp.Age
		count++		
	}

	fmt.Printf("Средний возраст всех сотрудников: %d\n", sumAge/count)
}