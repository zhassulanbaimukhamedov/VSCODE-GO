package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	add(1, 2, 3)
	add(1, 2, 3, 4)
	add(1, 2, 3, 4, 5)
	add([]int{1, 2, 3}...)
	var a = add2(4, 5)
	fmt.Println(a)
	var z = add3(1, 2)
	fmt.Println(z) 
}

func add(numbers ...int) { // многоточие ставится если нет определенное количество значении
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

//
func add2(x, y int) int {
	return x + y
}

// именованные возвращаемые результаты
func add3(x, y int) (z int) {
	z = x + y
	return
}
