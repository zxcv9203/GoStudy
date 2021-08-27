package main

import (
	"fmt"
	"github.com/GoStudy/패키지/pack3"
)
func init() {
	fmt.Println("main.go init")
}
func main() {
	fmt.Println("main.go main start")
	pack3.Call()
}
