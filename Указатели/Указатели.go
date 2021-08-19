package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// // Указатели - объекты, в значении содержатся адрес на другой объект
	// // Определяется с помощю символа звездочки
	// var d *int
	// var i = 1
	// // Для присвоения, применяется операция & перед именем переменной
	// d = &i
	// *d = 3
	// fmt.Println(i)
	// // fmt.Println(d) //0xc000012078

	// // Для получения значения указателя пользуемся операцией *
	// var r = *d

	// fmt.Println(r) //1

	// // Можно менять значение переменной по адресу указателя
	// *d = 2
	// fmt.Println(i)

	// // можно использовать сокращенную форму
	// f := 3
	// dd := &f
	// fmt.Println(f)
	// fmt.Println(*dd)

	// пустой указатель имеет значение nil
	// var gp *float64
	// fmt.Println(*gp) // ! Ошибка
	// if *gp != nil {
	// 	fmt.Println(*gp)
	// }

	// // Переменные это именованные объекты в памяти.
	// // С помощью функции new() можно создавать безимянные объекты и также размещаются в памяти, и не имеют имени как переменные
	// noName := new(int)
	// // по умолчанию значение 0
	// fmt.Println(*noName)
	// *noName = 8
	// fmt.Println(*noName)

	// // Указатели как параметры функции
	// d := 5
	// fmt.Println("d before: ", d)
	// changeValue(&d)
	// fmt.Println("d before: ", d)

	// Указатели как результат функции
	// p1 := createPointer(7)
	// fmt.Println(*p1)

}

func changeValue(x *int) {
	*x = (*x) * (*x)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}
