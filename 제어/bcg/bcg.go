package main

import "fmt"

func main() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	fmt.Println(a)
END:
	fmt.Println("End")
}
