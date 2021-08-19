package main

import "fmt"

// Интерфейсы - функции без реализации
type Vehicle interface {
	move()
}

func drive(v Vehicle) {
	v.move()
}

type Car struct {
}

type Aircraft struct {
}

func (c Car) move() {
	fmt.Println("car driving")
}

func (a Aircraft) move() {
	fmt.Println("air flying")
}

func main() {

	fmt.Println("Hello world")

	teslaVeh := Car{}
	boingVeh := Aircraft{}

	// teslaVeh.move()
	// boingVeh.move()

	drive(teslaVeh)
	drive(boingVeh)

}
