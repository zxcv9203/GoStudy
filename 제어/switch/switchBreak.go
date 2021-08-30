package main

import "fmt"

func main() {
	n := 2

	switch n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		break
		fmt.Println("no print")
	case 3,4:
		fmt.Println(3, 4)
	default:
		fmt.Println("match failed")
	}
}
