package main

import (
	"fmt"
	"os"
)

func recoverPanic() {
	r := recover()
	if r == nil {
		return
	}
	err, ok := r.(error)
	if ok {
		fmt.Println(err) // 에러 타입으로 반환이 가능할 경우 에러 메시지를 출력합니다.
	} else {
		panic(r) // 예상되지 않은 에러일 경우 다시 패닉 발생
	}
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	f.Close()
}

func openFile(name string) {
	defer recoverPanic()
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer closeFile(f) // 패닉 발생시 호출되지 않습니다.
}

func main() {
	openFile("Invalid.txt")
	fmt.Println("exit")
}
