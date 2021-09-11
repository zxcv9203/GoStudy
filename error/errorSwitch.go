package main

import (
	"errors"
	"fmt"
	"log"
)

type MyError struct {
	code    string
	message string
}

func (e MyError) Error() string {
	return e.code + ": " + e.message
}

func test() error {
	my := MyError{code: "1", message: "undefined error"}
	return my
}

func main() {
	//err := test()
	err := errors.New("aa")
	switch err.(type) {
	case MyError:
		log.Print("my Error") // myError 타입 일경우
	case error:
		log.Fatal("default error") // myError 타입이 아닌 다른 에러일 경우
	default:
		fmt.Println("ok") // 에러가 없을 경우
	}
}
