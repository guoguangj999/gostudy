package main

import "fmt"

type Animal interface {
	Eat()
	Move()
}

type Human struct {
	Name string
	Age  int
}

func (h Human) Eat() {
	say := fmt.Sprintf("I am %s,I can eat", h.Name)
	fmt.Println(say)
}

func (h Human) Move() {
	say := fmt.Sprintf("I am %s,I can move", h.Name)
	fmt.Println(say)
}

func main() {
	var animal Animal

	animal = Human{Name: "张三", Age: 20}

	animal.Eat()
	animal.Move()
}
