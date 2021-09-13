package main

import (
	"fmt"
	"time"
)

func main () {
	start := time.Now()
	func() {
		for i:=0; i < 1000; i++ {
			fmt.Println(i)
		} 
	}()

	elapsedTime := time.Since(start)

	fmt.Println("총 실행 시간: " + elapsedTime.String())

	time.Sleep(time.Second)
}