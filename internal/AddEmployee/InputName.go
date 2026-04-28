package addEmployee

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
)

func InputName() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя сотрудника: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return name, fmt.Errorf("Ошибка чтения имени: %v", err)
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return name, errors.New("Имя сотрудника не может быть пустым!")
	}
	return name, nil
}