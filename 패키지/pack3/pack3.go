package pack3

import (
	"fmt"
	"github.com/GoStudy/패키지/pack2"
)

func init() {
	fmt.Println("pack3 init")
}
func Call() {
	fmt.Println("pack3 called")
	pack2.Call()
}