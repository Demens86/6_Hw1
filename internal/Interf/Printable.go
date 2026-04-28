package interf

import "fmt"

type Printable interface {
	Print()
}
type Employee struct {
	Name string
	Age int
	Position string
	Salary float64
}

func (emp Employee) Print() {
	fmt.Printf("Имя: %s | Возраст: %d | Должность: %s | Зарплата: %.2f\n", emp.Name, emp.Age, emp.Position, emp.Salary)
}

func PrintInfo(p Printable) {
	p.Print()
}