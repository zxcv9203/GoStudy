package pack2

import (
	"fmt"
	"github.com/GoStudy/패키지/pack1"
)

func init() {
	fmt.Println("pack2 init")
}
func Call() {
	fmt.Println("pack2 called")
	pack1.Call()
}