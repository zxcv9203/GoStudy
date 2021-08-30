package main

import (
	"fmt"
	"os"
)

func main() {
	name := "hello"
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	d, err := os.Stat(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	f.Close()
}