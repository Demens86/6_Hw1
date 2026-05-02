package internal

import (
	"strings"
	"fmt"
	"errors"
)

func Info(name string) (error) {
	var flag bool

	for _, emp := range Employees {
		//Сравнение без учета регистра
		if strings.EqualFold(emp.Name, name) {
			fmt.Printf("Имя: %s\nВозраст: %d\nДолжность: %s\nЗарплата: %.2f\n", emp.Name, emp.Age, emp.Position, emp.Salary)
			flag = true
			break
		}
	}

	if !flag {
		//fmt.Println("Сотрудник c таким именем не найден")
		return errors.New("Сотрудник c таким именем не найден!\n")
	}
	return nil
}