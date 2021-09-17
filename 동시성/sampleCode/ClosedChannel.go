package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 2)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch <- 10
		ch <- 11
		wg.Done()
	}()
	
	wg.Wait()
	//close(ch)
	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 11
	fmt.Println(<-ch) // 0
	//fmt.Println(<-ch) // 0
}