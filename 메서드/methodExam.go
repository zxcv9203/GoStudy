package main

import "fmt"

type Rect struct {
	width, height float64
}

type Liter int

func (l *Liter) double() {
	*l *= 2
}

func (r *Rect) area() float64 {
	return r.width * r.height
}

func main() {
	water := Liter(3)
	rect := Rect{10.5, 22.3}
	water.double()

	fmt.Println(rect.area())
	fmt.Println(water)
}
