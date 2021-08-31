package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name := ""

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	d, err := f.Stat()
	fmt.Println(d)
	if err != nil {
		log.Fatal(err)
	}
}