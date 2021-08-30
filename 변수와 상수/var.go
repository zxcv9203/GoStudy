package main

import (
	"fmt"
	"reflect"
)

var (
	myInt int
	myFloat float64
	myBool bool
	myString string
)

func main() {
	var a int
	var b float64 = 10.5

	a = 10
	b = 42.42

	var c, d int
	c, d = 24, 42
	var e, f float64 = 0.42, 4.2
	var g = 3
	fmt.Printf("a = %d, b = %f\n", a, b)
	fmt.Printf("c = %d, d = %d\n", c, d)
	fmt.Printf("e = %f, f = %f\n", e, f)
	fmt.Printf("g = %d, g의 type = %s", g, reflect.TypeOf(g))
}
