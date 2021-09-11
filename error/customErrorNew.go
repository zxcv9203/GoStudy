package main

import (
	"errors"
	"fmt"
	"reflect"
)

func test() error {
	return errors.New("oh.. error")
}
func main() {
	err := test()
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(err))
}
