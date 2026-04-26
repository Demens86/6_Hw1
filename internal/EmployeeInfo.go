package internal

import (
	"strings"
	"fmt"
)

func Info(name string) {
	var flag bool

	for _, emp := range Employees {
		//Точное сравнение с учетом регистра
		// if strings.Contains(emp.Name, name) {
		// 	fmt.Printf("Name: %s | Age: %d | Position: %s | Salary: %.2f\n", emp.Name, emp.Age, emp.Position, emp.Salary)
		// 	flag = true
		// 	break
		// }
		//Сравнение без учета регистра
		if strings.Contains(strings.ToLower(emp.Name), strings.ToLower(name)) {
			fmt.Printf("Name: %s | Age: %d | Position: %s | Salary: %.2f\n", emp.Name, emp.Age, emp.Position, emp.Salary)
			flag = true
			break
		}
		
	}

	if !flag {
		fmt.Println("Сотрудник c таким именем не найден")
	}

}