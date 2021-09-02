package main

import "fmt"

type myString string

func (m myString) method() {
	fmt.Println("Method with value receiver")
}

func (m *myString) ptrMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := myString("a value")
	pointer := &value
	value.method()
	value.ptrMethod()
	pointer.method()
	pointer.ptrMethod()
}