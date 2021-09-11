# 에러처리

Go는 여러개의 값을 리턴할 수 있는데 그 덕분에 일반적인 리턴 값(하나만 리턴 가능)에 비해 상세한 에러 내용을 제공하기 쉽게 만들어 줍니다.

> Error Type

Go는 내장 타입으로 `error` 라는 `interface` 타입을 갖습니다.

Go 에러는 이 `error` 인터페이스를 통해서 주고 받게 되는 데, 이 interface는 다음과 같은 하나의 메서드를 갖습니다.

```go
type error interface {
	Error() string
} 
```

> error 타입을 커스텀 해보자

개발자는 다음과 같이 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있습니다.

```go
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
```

`test` 함수는 `error` 반환 값을 가지기 때문에 `customError`를 받으면 `customError`의 `Error` 메서드를 실행시켜  `e.code + ": " + e.message` 를 반환 값으로 돌려주게 됩니다.

다음 코드 같은 경우 실행시키면 다음과 같이 출력 값이 나옵니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/9fea5adf-6faa-40f1-9c3a-5807064b1a4e/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210911%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210911T165346Z&X-Amz-Expires=86400&X-Amz-Signature=6bc539936f9cf455d2d1eae3b69d9117f81bfb34f9062c5b78671d881be4566e&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

첫 줄에는 `e.code + ": " + e.message` 값이 나오고 두번째 줄은 `customError`가 나오는 것을 볼 수 있습니다.

위의 방법말고도 errors 패키지의 New 메서드를 이용해서 다음과 같이 처리할 수도 있습니다.

```go
package main

import (
	"errors"
	"fmt"
	"reflect"
)

func test() error {
	return errors.New("oh.. error") // 에러 메시지를 반환(문자열)
}
func main() {
	err := test()
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(err))
}
```

다음과 같은 출력값을 얻을 수 있습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/519d8019-b238-46d3-b851-5b5512d079aa/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210911%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210911T165424Z&X-Amz-Expires=86400&X-Amz-Signature=c8e54ad9cd36e33b29fe5896c03e40bbcc44a0a800e3923fde7c1f17e1112d9b&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> 실제 에러처리해보기

만약 결과와 에러가 함께 반환되는 함수일 경우, 에러가 `nil`인지를 체크해서 에러가 발생했는지를 알 수 있습니다.

예를 들어, `os.Open()`함수는  다음과 같은 함수 원형을 갖습니다.

```go
func Open(name string) (*File, error)
```

첫번째는 파일 포인터를 반환하고 두번째는 error 타입(error 인터페이스)를 반환합니다.

그래서 이 경우 두 번째 error를 체크해서 nil이 반환되면 에러가 없는 것이고 nil이 아니면 err.Error()로 부터 해당 에러를 알 수 있게 됩니다.

다음 예제를 통해 어떤 방식으로 처리되는지 볼 수 있습니다.

```go
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
```

notExistFile 이라는 이름의 파일이 존재하지 않으면 에러를 출력하고 아니라면 파일 이름을 출력해줍니다.

존재하지 않을 경우 다음과 같이 메시지를 출력하고 종료합니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/1f761152-dc11-4616-bfe3-3cb7b1a550a0/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210911%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210911T165437Z&X-Amz-Expires=86400&X-Amz-Signature=a65c9ac7708115e0acb746ea89893f0141a1b91a95b2156823cbea38aa64b16f&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

log.Fatal은 에러 로그를 출력 후 프로그램을 완전히 종료시키는 함수입니다.

위의 방법말고도 error의 Type을 체크해서 에러 타입 별로 별도의 에러 처리를 하는 방식도 있습니다.

다음은 switch를 통해 에러 타입별로 처리를 다르게 하는 예제입니다.

```go
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
```

`switch` 문에서 `err.(type)` 의 방식(타입 단언)으로 타입 체크를 합니다.

위의 경우는 에러가 발생하지 않았을 경우 `default`, 발생한 에러가 `myError` 일경우 `case myError`로 가고 나머지 에러는 `case error` 로 갑니다.

이런 경우 다양한 상황에 대해 좀 더 유연한 에러처리가 가능합니다. 

> defer

Go 언어의 defer 키워드는 특정 함수나 메서드를 마지막에 실행하게 합니다.

일반적으로 defer는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 clean-up 작업을 위해 사용됩니다.

defer는 보통 반드시 실행해야 하는 구문이 있을 때 사용됩니다.

다음은 defer를 사용해야 하는 상황입니다.

```go
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
	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stat)
	closeFile(f)
}
```

만약 위의 코드와 같이 파일을 열고 다하고 난 후 파일을 닫을 경우 문제가 발생할 수 있습니다.

예를 들어 `f.Stat()` 에서 파일의 정보를 읽어오는데 실패했을 경우 에러를 출력하고 종료하는데 그러면 파일이 닫히지 않고 종료되게 됩니다.

이런 상황이 한 두번이라면 상관없지만 열린파일은 계속 운영체제의 자원을 점유하기 때문에 시간이 지나면서 많은 파일이 쌓이게 되면 프로그램에 문제가 발생하게 되고 심지어는 시스템 전체의 성능의 저하를 초래할 수도 있습니다.

이때 `defer` 키워드를 이용하여 종료되기 전에 `f.close` 를 호출하게끔 변경해주면 문제가 해결됩니다.

```go
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
```

만약 defer를 여러번 쓸 경우 스택처럼 동작합니다.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
```

위의 예제 같은 경우는 c b a 가 출력됩니다.

> panic

panic은 코드를 컴파일할 때는 문제가 없었지만 실제 실행해보니 문제가 발생하는 런타임 에러가 발생시 panic이 발생합니다.

예를들어서 배열의 범위를 벗어난 탐색이나 잘못된 연산을 하는 경우 등 발생할 수 있습니다.

이 panic을 직접 발생시킬수도 있습니다.

Go의 내장함수인 `panic()` 함수는 호출하게 되면 함수는 런타임 에러를 발생시키면서 종료됩니다.

패닉이 발생하면 호출 스택 목록이 패닉 에러 메시지에 나오게 됩니다.

패닉 함수는 빈 인터페이스로 단일 인자를 받으며 문자열로 변환되는 경우 패닉 로그 메시지의 일부로 출력됩니다.

패닉을 실행하면 다음과 같은 메시지를 볼 수 있습니다.

```go
package main

func three() {
	panic("so deep")
}

func two() {
	three()
}

func one() {
	two()
}

func main() {
	one()
}
```

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/4e290e92-3263-4742-a8f0-4a890e42c8fe/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210911%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210911T165450Z&X-Amz-Expires=86400&X-Amz-Signature=2346ae442cd7282eeb5ce91a56b6471191888894c379514ce6de28c8ab77f35d&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

만약 defer를 panic 전에 호출하면 defer 구문을 전부 실행시키고 나서 종료합니다.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("exit...")
	fmt.Println("execute...")
	panic("crash program")
}
```

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/fb4dbbf6-7c2d-44ff-ba55-3b6fb2e6b27c/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210911%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210911T165506Z&X-Amz-Expires=86400&X-Amz-Signature=7c8f229111eeaa5dc32cee65b0649d2d85b7a2d39422a1be6d755ad121c21a1c&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> recover

패닉 같은 경우 잘 사용하면 굉장히 간단하게 에러 처리를 해줄 수 있습니다.

하지만 패닉은 복잡한 스택 트레이스와 함께 프로그램을 중단시킵니다.

사용자에게 복잡한 스택 트레이스 메시지는 필요하지 않을뿐더러 간단하게 표현하는 것이 보기 좋을 것입니다.

Go에는 패닉 상태에 빠진 프로그램을 복구할 수 있는 `recover`라는 내장 함수가 존재합니다.

`recover` 함수를 사용하면 패닉에 빠졌을 때 프로그램이 종료되지 않고 계속해서 실행할 수 있습니다.

```go
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
```

위의 코드는 매개변수로 받은 name을 열고 닫는 프로그램이고 만약 파일을 열 수 없으면 패닉을 호출합니다.

하지만 defer 를 이용해서 recover를 호출하기 때문에  패닉이 발생하더라도 메인 함수는 끝까지 실행됩니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/c08a2f19-f19a-4407-aa3b-c04a19b97c2a/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210911%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210911T165516Z&X-Amz-Expires=86400&X-Amz-Signature=23544935176f2598d89e32e9e8bd75d7f443a9036e87cbf6d6ea62cc2aae8db3&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

실행하면 다음과 같은 결과가 나옵니다.

closeFile 같은 경우는 패닉 이후에 defer 선언이 되었기 때문에 실행되지 않습니다.

즉, recover를 하더라도 패닉이 발생한 함수의 이후 부분은 실행시켜주지 않습니다.

또한 recover로 복구했을 때 이유를 알 수 없는 에러 값이 반환되지 않은 정말 예상하지 못한 오류가 나올 때를 대비해서 타입 단언을 통해 추가적인 에러처리를 하는 것이 좋습니다.