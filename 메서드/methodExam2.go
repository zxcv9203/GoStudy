package main

import "fmt"

type Liter int

func (l *Liter) PtrDouble() {
	*l *= 2
}

func (l Liter) ValueDouble() {
	l *= 2
}
func main() {
	water := Liter(5)
	milk := Liter(3)

	water.PtrDouble()
	milk.ValueDouble()

	fmt.Println(water)
	fmt.Println(milk)
}
