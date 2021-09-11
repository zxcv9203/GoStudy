package main

import (
	"fmt"
	"log"
	"os"
)

func closeFile(f *os.File) {
	fmt.Println("Closing File...")
	f.Close()
}
func main() {
	f, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)
	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stat)
}
