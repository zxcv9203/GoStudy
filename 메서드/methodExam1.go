package main

import "fmt"

type width int
type height int

func (w width) double() width {
	return w * 2
}
func (h height) double() height {
	return h * 2
}
func main() {
	w := width(100)
	h := height(300)

	fmt.Println(w.double())
	fmt.Println(h.double())
}
