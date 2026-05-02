package main

import (
	"EmployeeAccounting/internal"
	"EmployeeAccounting/internal/AddEmployee"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("--- Мини-система учета сотрудников компании ---")
		fmt.Println("1. Добавить сотрудника")
		fmt.Println("2. Вывести информацию о сотруднике")
		fmt.Println("3. Вывести список всех сотрудников")
		fmt.Println("4. Вывести средний возраст всех сотрудников")
		fmt.Println("5. Вывести среднюю зарплату по должностям")
		fmt.Println("6. Выход")
		fmt.Println("Выберите пункт меню (1-6):")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка чтения ввода пункта меню: %v\n", err)
		}
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addEmployee.AddEmployee(internal.Employees)
		case "2":
			fmt.Println("Введите имя сотрудника: ")
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Ошибка чтения имени: %v\n", err)
				continue
			}
			name = strings.TrimSpace(name)
			err = internal.Info(name)
			if err != nil {
				fmt.Printf("%v", err)
			}
		case "3":
			internal.PrintAllEmployees(internal.Employees)
		case "4":
			internal.PrintAverageEmployeeAge()
		case "5":
			internal.PrintAverageSalaryPosition(internal.Employees)
		case "6":
			return
		default:
			fmt.Println("Неверный пунк меню!")

		}
	}
}
