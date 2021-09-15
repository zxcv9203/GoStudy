package main

import (
	"fmt"

	"github.com/GoStudy/test/comma"
)

func main() {
	parses := []string{"hello"}
	fmt.Println(comma.JoinWithCommas(parses))
}
