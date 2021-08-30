package main

import "fmt"

func main() {
	Loop:
		fmt.Println("it is Error")
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if j == 2 {
					break Loop
				}
				fmt.Println(i, j)
			}
		}
		fmt.Println("Hello Go")
}
