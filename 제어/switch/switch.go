package main

import "fmt"

func main() {
	i := 0
	fmt.Print("0 ~ 4 사이의 수를 입력해주세요 : ")
	fmt.Scan(&i)
	switch i {
	case 0:
		fmt.Println("0을 입력하셨습니다.")
	case 1:
		fmt.Println("1을 입력하셨습니다.")
	case 2:
		fmt.Println("2를 입력하셨습니다.")
	case 3:
		fmt.Println("3을 입력하셨습니다.")
	case 4:
		fmt.Println("4를 입력하셨습니다.")
	default:
		fmt.Println("잘못된 값을 입력하셨습니다.")
	}
}
