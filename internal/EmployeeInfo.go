package internal

import (
	"strings"
	"fmt"
)

func Info(name string) {
	var flag bool

	for _, emp := range Employees {
		//Сравнение без учета регистра
		if strings.EqualFold(emp.Name, name) {
			fmt.Printf("Имя: %s | Возраст: %d | Должность: %s | Зарплата: %.2f\n", emp.Name, emp.Age, emp.Position, emp.Salary)
			flag = true
			break
		}
	}

	if !flag {
		fmt.Println("Сотрудник c таким именем не найден")
	}
}