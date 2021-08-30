package main

import "fmt"

func main() {
	var a int = 3
	if a == 1 {
		fmt.Println("a = 1")
	} else if a == 2 {
		fmt.Println("a = 2")
	} else {
		fmt.Println("a != 1 && a != 2")
	}
}
