package main

import (
	"fmt"
	"reflect"
)

type customError struct {
	code    string
	message string
}

func (e customError) Error() string {
	return e.code + ": " + e.message
}

func test() error {
	myError := customError{code: "1", message: "undefined error"}
	return myError
}

func main() {
	err := test()
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(err))
}
