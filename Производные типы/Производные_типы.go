package main

import "fmt"

// Для определения именованного типа используем оператор type

type mile int
type kilometer int
type library []string

func main() {

	fmt.Println("Hello world")

	// var distance mile = 5
	// distanceToEnemy(distance)
	// // var distance2 kilometer = 5
	// //  fmt.distanceToEnemy(distance2) // ! Ошибка

	// var myLibrary library = library{"Book1", "Book2", "Book3"}
	// printBooks(myLibrary)

}

func distanceToEnemy(distance mile) {
	fmt.Println(distance, " миль")
}

func printBooks(lib library) {

	for _, value := range lib {

		fmt.Println(value)
	}
}
