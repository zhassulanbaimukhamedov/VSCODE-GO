package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// Анонимные функции - это функции, которым не назначено имя.
	// Анонимные функции позволяют определить действие там, где оно применяется.
	// f := func(x, y int) int { return x + y }
	// fmt.Println(f(3, 4))
	// переменную f можно присвоить любую функцию, которая соответствует типу func(int, int) int

	// Анонимная функция как аргумент функции
	// action(3, 4, func(x, y int) int { return x + y })
	// action(3, 4, func(x, y int) int { return x * y })

	// Анонимная функция как результат функции
	// f := selectFn(1)
	// fmt.Println(f(3, 4))
	// f = selectFn(3)
	// fmt.Println(f(3, 4))

	// Преимуществом анонимных функции является то, что они имеют доступ к окружению, в котором они вызываются
	// f := square()
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

}

func action(x, y int, operation func(int, int) int) {
	result := operation(x, y)
	fmt.Println(result)
}

func selectFn(a int) func(int, int) int {
	if a == 1 {
		return func(x, y int) int { return x + y }

	} else if a == 2 {
		return func(x, y int) int { return x - y }

	} else {
		return func(x, y int) int { return x * y }
	}
}

// функция square определяет локальную переменную x и возвращает анонимную функцию. Анонимная функция увеличивает значение и возвращает квадрат переменной x
func square() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}
