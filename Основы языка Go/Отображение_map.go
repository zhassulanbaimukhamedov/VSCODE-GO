package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// map[K]V или Отображение - представляет ссылку на хэш-таблицу(структуру данных), где каждый элемент представляет пару "ключ-значение"

	//var people map[string]int

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	//	fmt.Println(people["Tom"])

	// if val, has := people["Tom"]; has {
	// 	fmt.Println(val)
	// }

	// for key, val := range people {
	// 	fmt.Println(key, val)
	// }

	// make - создает пустую хэш-таблицу
	//people := make(map[string]int)
	delete(people, "Bob")
	fmt.Println(people)

}
