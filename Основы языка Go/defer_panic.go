package main

import "fmt"

func main() {
	// // defer - оператор, позволяет выполнить функцию в конце программы, не важно где вызывается.
	// defer finish()
	// defer fmt.Println("Defer 2")
	// fmt.Println("Hello world")
	// fmt.Println("start")
	// fmt.Println("work")

	// panic - выход из программы в случае ошибок
	// fmt.Println(divide(4, 2))
	// fmt.Println(divide(4, 0))
	// fmt.Println("finish")
}
func finish() {
	fmt.Println("finish")
}
func divide(x, y float64) float64 {
	if y == 0 {
		panic("division by 0")
	}
	return x / y
}
