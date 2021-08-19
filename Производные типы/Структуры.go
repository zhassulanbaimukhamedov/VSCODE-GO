package main

import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// Вложенные структуры
type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age  int
	contact
}

// Хранение ссылки на структуру того же типа
type node struct {
	value int
	//next  node // Подобное определение будет неправильным. Вместо этого поле должно представлять указатель на структуру
	next *node
}

// рекурсивный вывод списка
func printNodeValue(n *node) {
	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func main() {

	fmt.Println("Hello world")

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}
	first.next = &second // указатель первой ветки присваивает вторую ветку
	second.next = &third // указатель второй ветки присваивает третью ветку

	var current *node = &first // указатель current присваивает значение первой ветки
	for current != nil {
		fmt.Println(current.value)
		current = current.next //
	}

	// var tom = person{
	// 	name: "Tom",
	// 	age:  24,
	// 	contact: contact{
	// 		email: "tom@gmail.com",
	// 		phone: "+12345678910",
	// 	},
	// }

	// tom.email = "supertom@gmail.com"
	// fmt.Println(tom.email)
	// fmt.Println(tom.phone)

	// tom := person{name: "Tom", age: 23}
	// var tomPointer *person = &tom
	// tomPointer.age = 29

	// fmt.Println(tom.age)
	// tomPointer.age = 32
	// fmt.Println(tom.age)

	// var bob person = person{name: "Bob", age: 24}

	// fmt.Println(bob)
	// fmt.Println(tom.name)

}
