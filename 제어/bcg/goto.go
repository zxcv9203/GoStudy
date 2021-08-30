package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		goto ERROR1
	} else if a == 2 {
		goto ERROR2
	} else if a == 3 {
		goto ERROR1
	}
	fmt.Println("Success")
	return
ERROR1:
	fmt.Println("Error 1")
	return
ERROR2:
	fmt.Println("Error 2")
	return
}
