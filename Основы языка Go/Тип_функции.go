package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// var f func(int, int) int = add
	// var z = f(3, 4)
	// fmt.Println(z)

	// f := add
	// fmt.Println(f(3, 4))
	// f = multiple
	// fmt.Println(f(3, 4))
	// f1 := display
	// f1("hello func display")

	// Функции как параметры других функции
	// action(3, 4, add)
	// action(3, 4, multiple)

	// slice := []int{-2, 4, 3, -1, 7, -4, 23}
	// sumOfEvens := sum(slice, isEven)
	// fmt.Println(sumOfEvens)
	// sumOfPositive := sum(slice, isPositive)
	// fmt.Println(sumOfPositive)

	// Функция как результат другой функции
	// f := selectFn(1)
	// fmt.Println(f(3, 4))
	// f = selectFn(3)
	// fmt.Println(f(3, 4))

}

func isEven(n int) bool {
	return n%2 == 0
}

func isPositive(n int) bool {
	return n > 0
}

func sum(a []int, criteria func(int) bool) int {
	result := 0
	for _, v := range a {
		if criteria(v) {
			result += v
		}
	}
	return result
}

func action(a int, b int, operation func(int, int) int) {
	result := operation(a, b)
	fmt.Println(result)
}

// функция суммирования
func add(x int, y int) int {
	return x + y
}

// функция умножения
func multiple(x int, y int) int {
	return x * y
}

// функция вычитания
func substract(x int, y int) int {
	return x - y
}

// функция выборки функции как результат
func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return substract
	} else {
		return multiple
	}

}
func display(message string) {
	fmt.Println(message)
}
