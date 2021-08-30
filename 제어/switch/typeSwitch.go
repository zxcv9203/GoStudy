package main

import "fmt"


func main() {
	var t interface{} = "hello"
	switch s := t.(type) {
	case int:
		fmt.Println("int", s)
	case bool:
		fmt.Println("bool", s)
	case string:
		fmt.Println("string", s)
	default:
		fmt.Println("unknown type", s)
	}
}
