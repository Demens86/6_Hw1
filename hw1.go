// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var x, y int64
// 	_, _ = fmt.Scan(&x, &y)
// 	fmt.Println(sumDiv(x, y))
// 	fmt.Println(oper(x,y,mult))
// }

// func sumDiv(x int64, y int64) (int64, int64, int64)  {
// 	sum := x+y
// 	div:= x/y
// 	waste:= x - div
// 	return sum, div, waste
// }

// func mult(x int64, y int64) int64 {
// 	return x*y
// }
// func oper(x, y int64, operation func(int64, int64) int64) int64 {
// 	return operation(x, y)
// }