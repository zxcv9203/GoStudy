package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("notExistFile")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Name())
}
