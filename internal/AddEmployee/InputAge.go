package addEmployee

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func InputAge() (int, error) {
	var age int
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите возраст сотрудника: ")
		input, err := reader.ReadString('\n')
		if err != nil {
		return age, fmt.Errorf("Ошибка чтения возраста сотрудника: %v", err)
		}

		input = strings.TrimSpace(input)

		age, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введите целое число!")
			continue
		}
		if age <18 {
			fmt.Println("Минимальный возраст приема на работу 18 лет!")
			continue
		} else if age >= 65 {
			fmt.Println("Человеку пора на пенсию!")
			continue
		}
		break
	}
	return age, nil
}