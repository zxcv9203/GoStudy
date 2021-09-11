package main

import "fmt"

func main() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
