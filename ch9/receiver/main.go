package main

import "fmt"

type MyType string
type Number int

func (n *Number) Double() {
	*n *= 2
}

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()

	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)

	value2 := MyType("a value")
	pointer := &value2
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()
}
