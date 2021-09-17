package main

import "fmt"

func main() {
	DataChannel := make(chan bool)

	go func() {
		for i := 0; i < 123; i++ {
			fmt.Println(i)
		}
		DataChannel <- true
	}()
	fmt.Println(<-DataChannel)
}