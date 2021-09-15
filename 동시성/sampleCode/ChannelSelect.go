package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- 10
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- 20	
		}
	}()

	go func() {
		for {
			select {
				case a := <-ch1:
					fmt.Printf("ch1 데이터 %d 수신\n", a)
				case b := <-ch2:
					fmt.Printf("ch2 데이터 %d 수신\n", b)
			}
			
		}
	}()

	time.Sleep(5 * time.Second)
}