package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// lib := library{"Book1", "Book2", "Book3"}
	// lib.print()

	// tom := person{"Tom", 23}
	// tom.print()
	// tom.eat("pizza")

	tom := person{"Tom", 23}
	fmt.Println(tom.age)
	topPointer := &tom
	topPointer.updateAge(33)
	fmt.Println(tom.age)
}

// Методы представляют функцию определенного типа
// Методы определяется также как и функции за исключением, что в методах надо указать получателя(receiver)
// Получатель - это параметр. Параметр того типа, к которому привязан метод

// type library []string

// func (lib library) print() {
// 	for _, val := range lib {
// 		fmt.Println(val)
// 	}
// }

// Методы структуры
// type person struct {
// 	name string
// 	age  int
// }

// func (p person) print() {
// 	fmt.Println(p.name)
// 	fmt.Println(p.age)
// }

// func (p person) eat(meal string) {
// 	fmt.Println(p.name, " eat ", meal)
// }

// Методы указателей
// при вызове метода объект структуры передает значение, а не сам объект
type person struct {
	name string
	age  int
}

func (p *person) updateAge(newAge int) {
	(*p).age = newAge
}
