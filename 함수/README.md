# 다중 반환값

Golang에서는 단일 함수에서 한 번에 여러 값을 반환할 수 있다. 

## 동일한 반환형

두 함수 모두 같은 의미로써 두개의 정수형 리턴이 이뤄질 것을 선언하고 있다.

```go
func func_name(a, b int) (int, int) {
	return a, b
}

func func_name1(a, b int) (a, b int) {
	return a, b
}
```

## 서로 다른 반환형

물론 반환형이 일치할 필요는 없다. 

```go
func func_name2(a, b int) (float64, int) {
	return a, b
}
```

## 이름이 지정된 반환형

미리 리턴할 값에 대해서 이름을 붙여놓고 이를 활용할 수도 있다. 아래 예시에서 훨씬 간결하게 코드가 작성됨을 확인할 수 있다.

```go
func someFunc() (int, error) {
    var r int
    var e error
    ok := someOtherFunc(&r)  // contrived, I admit
    if !ok {
        return r, someError()
    }
    return r, nil
}

func someFunc() (r int, e error) {
    ok := someOtherFunc(&r)
    if !ok {
        e = someError()
    }
    return
}
```

일반적으로 알려진 named return type의 장점은 다음과 같다.

- 반환하고자하는 값의 역할에 따라 이름을 지정할 수 있어 일종의 문서 역할을 한다. (godoc에서 활용)
- 별도의 선언이 요구되지 않고, 자동으로 초기화된다.
- 매개변수가 go의 스택에 저장될 때 명명되어, 어셈블리어 코드를 인식할때 더 용이하다.[추가설명](./함수스택.md)


## 값과 함께 오류 반환

여러 반환형을 가지면, 오류를 처리할때 상당히 유용하게 활용할 수 있다. 

다음의 예에서 볼 수 있듯 하나의 함수 내에서 값과 에러라는 두가지 다른 반환형을 가질 수 있다는 점은 상당한 이점으로 작용한다.
```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

func capitalize(name string) (string, error) {
    if name == "" {
        return "", errors.New("no name provided")
    }
    return strings.ToTitle(name), nil
}

func main() {
    name, err := capitalize("sammy")
    if err != nil {
        fmt.Println("Could not capitalize:", err)
        return
    }

    fmt.Println("Capitalized name:", name)
}
```

에러가 발생하는 경우, main 함수에서 `name`은 빈 문자열을 갖고, 에러 문자열만을 갖게 된다.

## 익명 함수와 클로저

익명함수는 단어에서도 알 수 있듯이 '이름이 없는 함수'이다. 본래 함수의 이름은 해당 함수만이 갖는 고유한 시그니처 중의 하나로, 특히, 가독성의 측면에서 '해당 함수가 왜 필요하고 어떠한 기능을 하는지'를 대략적으로 표현하여 코드 가독성에 큰 역할을 하게 된다. 그렇다면 왜 굳이 이름이 없는 함수를 작성하여 가독성이 떨어지는 코드를 짤 수 있게 한 것인지 의문 발생하는데...

그 이유는 프로그램 속도에 있다. 프로그램을 작성하며 함수를 선언할 때, 해당 선언은 프로그램 전역으로 초기화되며 이는 곧 메모리의 점유를 의미한다. 선언 이후 이를 다시 찾아서 호출하는 과정은 프로그램 실행 속도에 영향을 주는 또다른 직접적인 영향이 된다.

이러한 단점을 보완하기 위해 '익명 함수'가 요구된다. 바로 위의 capitalize라는 함수를 익명함수로 변환한 예시를 보자.

```go
func capitalize(name string) (string, int, error) {
    handle := func(err error) (string, int, error) {
        return "", 0, err
    }

    if name == "" {
        return handle(errors.New("no name provided"))
    }

    return strings.ToTitle(name), len(name), nil
}
```

## defer

defer문은 일종의 함수의 실행에 관련된 일종의 제어문이라고 생각할 수 있다. if, else, for, while 다양한 제어문들이 코드의 실행 자체를 막는 느낌이라면, defer는 해당 함수를 스택에 저장해놓고 특정 시점이 되면 해당 함수를 실행한다.

defer는 특정 함수나 메소드를 프로그램 가장 마지막에 실행하도록 하는 제어 구문이다.

예를 들어 보자면
```go
package main

import (  
    "fmt"
)

func finished() {  
    fmt.Println("마지막에 출력")
}

func largest(nums []int) {  
    defer finished()    
    fmt.Println("시작할때 출력")
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println(nums, "에 들어있는 가장 큰 숫자는", max)
}

func main() {  
    nums := []int{78, 109, 2, 563, 300}
    largest(nums)
}
```

```
시작할때 출력 
[78 109 2 563 300]에 들어있는 가장 큰 숫자는 563  
마지막에 출력 
```

defer를 통해 리소스를 정리하는 방식은 Go에서 매우 일반적인 방식이다. 
아래의 코드는 defer를 사용하지 않고, 실행 중 파일을 생성하고 쓰는 간단한 프로그램이다.

defer를 좀더 이해하기 위해 아래 코드를 보자면
```go
package main

import (  
    "fmt"
)

func printA(a int) {  
    fmt.Println("defer가 사용된 경우의 변수 상태: ", a)
}
func main() {  
    a := 5
    defer printA(a)
    a = 10
    fmt.Println("defer보다 코드상으로 뒤에 위치할 떄의 변수 상태:", a)
}
```

다음과 같은 결과를 얻을 수 있다.
```go
defer가 사용된 경우의 변수 상태: 10  
defer보다 코드상으로 뒤에 위치할 떄의 변수 상태: 5 
```
결과에서 볼 수 있듯 defer는 argument는 키워드가 실행될 때를 기준으로 평가한다는 특징을 갖는다.

만약 defer가 여러번 사용된다면 순서는 어떻게 될까?

```go
package main

import (  
    "fmt"
)

func main() {  
    name := "Naveen"
    fmt.Printf("Original String: %s\n", string(name))
    fmt.Printf("Reversed String: ")
    for _, v := range []rune(name) {
        defer fmt.Printf("%c", v)
    }
}
```

![](https://golangbot.com/content/images/2020/02/defer-stack.png)

defer를 사용하면 각 호출은 스택에 저장된다. 최종적으로는 LIFO 방식으로 각 호출이 처리되어, 위의 프로그램에서는 역방향으로 문자를 출력하는 결과가 나오게 된다.

```go
Original String: Naveen  
Reversed String: neevaN  
```
```go
package main

import (
    "io"
    "log"
    "os"
)

func main() {
    if err := write("README.md", "README.md 파일입니다"); err != nil {
        log.Fatal("파일을 쓰는데 실패했습니다:", err)
    }
}

func write(fileName string, text string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    _, err = io.WriteString(file, text)
    if err != nil {
        return err
    }
    file.Close()
    return nil
}
```

먼저, 해당 프로그램에 작성된 에러 처리를 집중적으로 살펴보자. 파일을 열고, 파일에 쓰고 하는 과정에서 문제가 생기면, write는 err을 리턴하고 이를 받은 메인 함수가 로그를 기록하며 프로그램이 종료된다.

이때 문제는 중간에 에러가 발생하여 프로그램이 로그를 기록하는 경우, 리소스가 충분히 해제되지 않는다는 점이다. io.WriteString 함수가 비정상적으로 동작하여 에러 메세지를 리턴하는 경우, 이미 열린 파일을 닫지 않은 채 프로그램이 종료된다.

물론 io.WriteString의 에러를 처리하는 구문에 Close 함수를 추가하여 해결할 수는 있다.
```go
package main

import (
    "io"
    "log"
    "os"
)

func main() {
    if err := write("readme.txt", "This is a readme file"); err != nil {
        log.Fatal("failed to write file:", err)
    }
}

func write(fileName string, text string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    _, err = io.WriteString(file, text)
    if err != nil {
        file.Close()
        return err
    }
    file.Close()
    return nil
}
```
다만 이렇게 에러를 처리하는 것은 항상 어딘가 리소스를 처리하는 과정에서 버그를 발생시킬 수 있다는 가능성이 존재한다. 물론, 이는 당연히 코드를 작성하는 프로그래머의 책임이지만, defer를 사용하면 리소스에 대한 버그를 훨씬 효율적으로 처리할 수 있게 된다.

다시, 위의 프로그램을 defer를 사용하여 변경한다면 다음과 같이 작성할 수 있다.

```go
package main

import (
    "io"
    "log"
    "os"
)

func main() {
    if err := write("readme.txt", "This is a readme file"); err != nil {
        log.Fatal("failed to write file:", err)
    }
}

func write(fileName string, text string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = io.WriteString(file, text)
    if err != nil {
        return err
    }
    return nil
}
```

defer가 사용된 라인에 주목하자. defer를 사용함으로써 write라는 함수가 종료되기전에 항상 Close가 사용되어야 함을 컴파일러에게 알리게 되었다. 이제 더 많은 코드를 추가하고 이후 함수가 종료되는 더 많은 분기를 만들더라도, WriteString이 종료되면 언제나 Close가 따라올 것이다.

다만, 이렇게 되면 defer가 실행되는 라인에서 평가되는 Close는 문서가 열리기 전의 Close() 이므로 명백히 말하면 기존의 Close()와는 다르게 된다.

Go에서는 프로그램의 동작에 영향을 주지 않고 Close()를 두번 호출하는 것을 안전하고 허용되는 관행으로 처리한다. 물론 이게 이상적인 방법은 아니지만 프로그램에 문제가 없는 것이 우선이므로 이렇게 작성한다.
```go
package main

import (
    "io"
    "log"
    "os"
)

func main() {
    if err := write("readme.txt", "This is a readme file"); err != nil {
        log.Fatal("failed to write file:", err)
    }
}

func write(fileName string, text string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = io.WriteString(file, text)
    if err != nil {
        return err
    }

    return file.Close()
}
```