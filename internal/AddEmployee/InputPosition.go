package addEmployee

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
)

func InputPosition() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите должность сотрудника: ")
	position, err := reader.ReadString('\n')
	if err != nil {
		return position, fmt.Errorf("Ошибка чтения должности: %v", err)
	}
	position = strings.TrimSpace(position)
	if position == "" {
		return position, errors.New("Должность сотрудника не может быть пустой!")
	}
	return position, nil
}