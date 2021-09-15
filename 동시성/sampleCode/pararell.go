package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	start := time.Now()
	go func() {
		for i:=0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i:=10; i < 13; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i:=100; i < 103; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i:=1000; i < 1003; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i:=10000; i < 10003; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i:=100000; i < 100003; i++ {
			fmt.Println(i)
		}
	}()
	
	elapsedTime := time.Since(start)

	fmt.Println("총 실행 시간: " + elapsedTime.String())

	time.Sleep(time.Second)
}