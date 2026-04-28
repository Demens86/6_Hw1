package addEmployee

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func InputSalary() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	var salary float64
	for {
		fmt.Print("Введите зарплату сотрудника: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return salary, fmt.Errorf("Ошибка чтения зарплаты!")
		}

		input = strings.TrimSpace(input)

		salary, err = strconv.ParseFloat(input,64)
		if err != nil {
			fmt.Println("Зарплата некорретная, введите снова!")
			continue
		}
		if salary <0 {
			fmt.Println("Зарплата не может быть отрицательной!")
			continue
		}
		break
	}
	return salary, nil
}