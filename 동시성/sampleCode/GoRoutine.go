package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go func () {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func () {
		for i := -3; i < 0; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)
	fmt.Println("총 실행 시간: " + elapsedTime.String())
	time.Sleep(time.Second)
}