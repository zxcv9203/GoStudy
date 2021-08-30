# 제어

Go언어의 제어구조는 C와 연관성이 있지만 중요한 점에서 차이가 있습니다.

Go언어에서는 `do`나 `while` 반복문이 존재하지 않으며, 단지 좀 더 일반화 된 `for`, 좀 더 유연한 `switch` 가 존재합니다.

`if` 와 `switch`는 선택적으로 `for` 와 같은 초기화 구문을 받을 수 있습니다.

`break` 와 `continue` 구문들은 선택적으로 어느것을 멈추거나 계속할지 식별하기 위해서 라벨을 받을 수 있습니다.

또한 `타입 switch`와 `다방향 통신 멀티플렉서`, `select` 의 새로운 제어 구조가 포함되어 있습니다.

표현식에 괄호는 필요하지 않으며 제어문에 내용은 항상 중괄호로 구분해야 합니다.

> if

`if` 문은 주어진 조건이 맞다면 `{}` 안의 코드를 실행시킵니다.

중괄호를 의무적으로 사용해야 하며 사용하지 않을 경우 에러가 발생합니다.
if문의 표현식에는 괄호가 사용되지 않습니다.

Go 언어에서 if문의 간단한 예제는 다음과 같습니다.

```go
package main

import "fmt"

func main() {
	var a int = 3
	if a != 0 {
		fmt.Println(a)
	}
}
```

> if 표현식의 타입

Go에서는 표현식의 값이 무조건 `bool` 타입이어야 합니다.

예를 들어, 다음과 같은 사용은 불가능합니다.

```go
package main

import "fmt"

func main() {
	var a int
	if a {
		fmt.Println(a)
	}
}
```

위의 코드를 실행시켜보면 다음과 같이 bool 타입이 아니기때문에 if를 사용할 수 없다고 나옵니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/3dff7089-0f47-49cf-ba5d-b54db0066a38/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210830%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210830T050226Z&X-Amz-Expires=86400&X-Amz-Signature=ef99c06b7a20c0419a7633e50f3a84a2788c92360f74d211139a3ae6ff393268&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> else if if

if 문은 else if 혹은 else 문과 함께 사용할 수 있습니다.

else if 문은 if 조건문이 거짓일 때 다시 다른 if 표현식을 검사하는 데 사용되며, else 문은 이전의 if문, else if문 이 모두 거짓일 때 실행됩니다.

else if와 else는 이전의 if 구문들을 닫는 중괄호와 한줄에서 선언되어야 하며 if 문과 마찬가지로 표현식과 시작하는 중괄호가 같은 줄에 있어야 합니다.

if .. else if .. else 구문의 예제는 다음과 같습니다.

```go
package main

import "fmt"

func main() {
	var a int = 3
	if a == 1 {
		fmt.Println("a = 1")
	} else if a == 2 {
		fmt.Println("a = 2")
	} else {
		fmt.Println("a != 1 && a != 2")
	}
}
```

위의 코드를 실행 시켜보면 if와 else if가 모두 거짓이기 때문에 마지막 else가 실행되는 것을 알 수 있습니다.

> if 표현식 안의 초기화

if 문에서 표현식을 사용하기 이전에 변수를 초기화하여 함께 사용할 수 있습니다.

여기서 주의해야 할 점은 초기화한 변수는 if문 블럭 혹은 if .. else 블럭 안에서만 사용할 수 있습니다.

```go
package main

import "fmt"

func main() {
	if a := 3 * 2; a < 3 {
		fmt.Println(a)
	}
	// fmt.Println(a) -> 에러 발생
}
```

만약 if문 범위 밖에 있는 주석 처리한 코드를 주석 해제하고 실행시키면 다음과 같이 에러가 발생합니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/aa33efe6-91aa-48db-a73b-7151a714c98c/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210830%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210830T050248Z&X-Amz-Expires=86400&X-Amz-Signature=2807ac7ac9456b5dad06c351d4024a2336aab5a02876c5ba88384d58fdce4d30&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> 불필요한 else 피하기

Go의 여러 라이브러리들을 보게되면 if 구문이 다음 구문으로 진행되지 않을 때, 즉 `break` `continue` `goto` `return` 으로 인해서 구문이 종료될 경우 불필요한 else는 생략되는 것을 볼 수 있습니다.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	name := "hello"
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	d, err := os.Stat(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	f.Close()
}
```

위의 코드는 다음의 순서로 실행됩니다.

1. hello라는 파일을 읽어옵니다.
2. 읽기에 실패할 경우 에러 메시지를 출력하고 종료합니다.
3. 읽기에 성공한 경우 파일을

열기에 실패하면 에러 메시지를 출력하고 종료하고 열기에 성공하면 `if`문 이후의 코드를 실행시키 때문에 else를 사용할 필요가 없습니다.

> for

Go언어에서 for 반복문은 C언어와 비슷하지만 일치하지는 않는다.

Go에서 for는 while 처럼 동작할 수 있고, 따라서 do-while 이 없습니다.

Go에서 for는 다음과 같이 사용할 수 있습니다.

```go
for i; i < 5; i++ {
	//code
}

for a < 5 {
	// code
}

for {                 // 무한루프
	// code
}
```

단축 변수 선언 `:=` 은 반복문에서 index 변수 선언을 쉽게 만들어 줍니다.

```go
sum := 0

for i := 0; i < 10; i++ {
	sum += i
}
```

만약 array, slice, string, map, channel로 부터 읽어 들이는 반복문을 작성한다면, range 구문이 이 반복문을 관리해줄 수 있습니다.

```go
for index, value := range array {
	fmt.Printf("array[%d] = %d", index, value)
}
```

.첫번째는 인덱스나 키가 들어가고 2번째는 해당 인덱스의 값이 들어갑니다.

만약 range 안에 첫 번째 아이템만이 필요하다면 두번째는 작성하지 않고 다음과 같이 사용할 수 있습니다.

```go
for i := range arr {
	fmt.Println(i)
}
```

만약 range 안에서 두번째 아이템만이 필요하다면 공백 식별자 (`_`)를 사용하여 첫 번째를 버릴수 있습니다.

```go
for _, value := range arr {
	fmt.Println(value)
}
```

string의 경우 range는 UTF-8 파싱에 의한 개별적인 유니코드 문자를 처리하는데 유용합니다.

잘못된 인코딩은 하나의 바이트를 제거하고 U+FFFD 룬 문자로 대체할 것입니다.

```go
for pos, char := range "日本\x80語" { // \x80 은 합법적인 UTF-8 인코딩이다
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

위의 코드는 다음과 같이 출력됩니다.

```go
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

추가적으로 Go언어는 콤마(`,`) 연산자가 없으며 `++`, `--` 는 표현식이 아니라 명령문입니다.

따라서 for문 안에서 여러개의 변수를 사용하려면 병렬 할당(parallel assignment)을 사용해야 합니다.

```go
// 배열 a를 반대로 만들기
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	a[i], a[j] = a[j], a[i]
}
```

> break, continue, goto

경우에 따라 for 루프내에서 처음으로 돌아가거나, 반복문을 바로 빠져나오거나 임의의 위치로 이동시키고 싶을 때가 있습니다.

- 처음으로 돌아가고 싶을때는 `continue`

  반목문에서 특정 부분 이하는 실행하지 않고 넘어가려면 continue 키워드를 사용해야 합니다.
  간단한 예제를 통해서 `continue` 를 사용해 보겠습니다.

    ```go
    package main

    import "fmt"

    func main() {
    	for i := 0; i < 5; i++ {
    		if i == 2 {
    			continue
    		}
    		fmt.Println(i)
    	}
    }
    ```

  위의 코드는 i를 0에서 5까지 증가시키면서 `fmt.Println(i)` 를 출력하는 반복문입니다.
  하지만 위에 `if i == 2` 일때 `continue` 를 하므로 0, 1, 3, 4만 출력되는 것을 볼 수 있습니다.

  다음과 같이 `continue` 키워드에 레이블을 지정할 수도 있습니다.

    ```go
    package main

    import "fmt"

    func main() {
    	Loop:
    		for i := 0; i < 3; i++ {
    			for j := 0; j < 3; j++ {
    				if j == 2 {
    					continue Loop
    				}
    				fmt.Println(i, j)
    			}
    		}
    		fmt.Println("Hello, Go")
    }
    ```

  중첩된 for 반복문에서 특정 부분 이하는 실행하지 않고 넘어갈 때 `continue` 키워드에 레이블을 지정하면 편리합니다.

  주의해야 할점은 레이블과 for 키워드 사이에 다른 코드가 존재하면 에러가 발생합니다.

    ```go
    package main

    import "fmt"

    func main() {
    	Loop:
    		fmt.Println("a")
    		for i := 0; i < 3; i++ {
    			for j := 0; j < 3; j++ {
    				if j == 2 {
    					continue Loop
    				}
    				fmt.Println(i, j)
    			}
    		}
    		fmt.Println("Hello, Go")
    }
    ```

  ![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/7ad9153b-3a8b-4dfa-9e7b-5b506b47c43e/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210830%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210830T050315Z&X-Amz-Expires=86400&X-Amz-Signature=66687ac8bf5de2469d9e677d28271011094c3307cd53bc17e4ebcb3aeb4f1090&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

  `continue` 같은 경우 for에서만 사용할 수 있습니다.

- 반복문을 바로 빠져나오고 싶을때는 `break`

  for 키워드에 아무것도 설정하지 않으면 무한 루프가 됩니다. 이때, 반복문을 중단하려면 중괄호 블록안에 조건을 정하고 `break` 키워드를 사용하면 됩니다.

    ```go
    package main

    import "fmt"

    func main() {
    	i := 0
    	for {
    		if i > 4 {
    			break
    		}
    		fmt.Println(i)
    		i++
    	}
    }
    ```

  다음과 같이 break 키워드에 레이블을 지정할 수도 있습니다.

    ```go
    package main

    import "fmt"

    func main() {
    	Loop:
    		for i := 0; i < 3; i++ {
    			for j := 0; j < 3; j++ {
    				if j == 2 {
    					break Loop
    				}
    				fmt.Println(i, j)
    			}
    		}
    		fmt.Println("Hello Go")
    }
    ```

  주의해야할 점은 `continue` 와 마찬가지로 레이블과 for 키워드 사이에 다른 코드가 있으면 안 됩니다.

    ```go
    package main

    import "fmt"

    func main() {
    	Loop:
    		fmt.Println("it is Error")
    		for i := 0; i < 3; i++ {
    			for j := 0; j < 3; j++ {
    				if j == 2 {
    					break Loop
    				}
    				fmt.Println(i, j)
    			}
    		}
    		fmt.Println("Hello Go")
    }
    ```

  `break` 같은 경우에는 for 루프 이외에 switch문이나 select 문에서도 사용할 수 있습니다.

- 임의의 위치로 이동시키고 싶을때는 `goto` 를 사용합니다.

  `goto` 는 정해진 레이블로 바로 이동합니다. 보통 프로그래밍에서 goto 키워드는 권장되지 않지만 일부 상황에서는 코드를 간단하게 만들 수 있으므로 적절히 활용하는 것이 중요합니다.

  goto 키워드는 `goto <LABEL>` 과 같은 형식으로 사용됩니다. 여기서 LABEL(레이블)은 변수 이름을 짓는 규칙과 같습니다.

    ```go
    goto LABEL // 이동할 레이블을 지정합니다.
    LABEL:
    		// 실행할 코드를 작성합니다.
    ```

  먼저 if 조건문으로 에러 처리를 해줄때 중복되는 코드가 있으면 `goto` 를 사용하면 중복 코드 없이 에러를 처리할 수 있습니다.

    ```go
    package main

    import "fmt"

    func main() {
    	var a int = 1

    	if a == 1 {
    		fmt.Println("Error 1")        // 중복
    		return
    	} else if a == 2 {
    		fmt.Println("Error 2")
    		return
    	} else if a == 3 {
    		fmt.Println("Error 1")       // 중복
    		return
    	}
    	fmt.Println("Success")
    }
    ```

    ```go
    package main

    import "fmt"

    func main() {
    	var a int = 1

    	if a == 1 {
    		goto ERROR1
    	} else if a == 2 {
    		goto ERROR2
    	} else if a == 3 {
    		goto ERROR1
    	}
    	fmt.Println("Success")
    	return
    ERROR1:
    	fmt.Println("Error 1")
    	return
    ERROR2:
    	fmt.Println("Error 2")
    	return
    }
    ```

  `goto` 같은 경우에는 for 루프와 관계없이 사용할 수 있습니다.

아래 코드는 `continue`, `break`, `goto`를 같이 사용한 코드 예제입니다.

```go
package main

import "fmt"

func main() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	fmt.Println(a)
END:
	fmt.Println("End")
}
```

위의 코드를 실행시키면 End가 출력됩니다.

> switch

여러 값을 비교해야 하는 경우 혹은 다수의 조건을 체크해야하는 경우사용합니다.

if, else if를 조건문으로 나열하는 것보다 switch문을 사용하는게 좀 더 간단하게 조건을 표현할 수 있고 이는 Go언어 다운 표현방식입니다.

표현식은 상수이거나 정수일 필요가 없고 case 구문은 위에서부터 끝까지 해당 구문이 true가 아니면 계속 값을 비교합니다.

switch의 구조는 다음과 같습니다.

```go
switch 변수 {
	case 값1 :
				// 값1일 때 실행할 코드를 작성합니다.
	case 값2 :
				// 값2일 때 실행할 코드를 작성합니다.
	default:
				// 모든 case에 해당하지 않을 때 실행할 코드를 작성합니다.
}
```

한가지 특이한 점은 Go 언어의 switch 분기문은 C, C++와는 달리 case에서 break 키워드를 생략합니다.

간단한 예제를 작성해보겠습니다.

```go
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
```

switch의 i 값을 입력했고 입력한 i 값을 비교해서 0 ~ 4 사이의 숫자라면 해당하는 숫자를 출력하고 아니라면 잘못된 값을 입력했다고 알려줍니다.

case는 위에서 언급한 것 처럼 상수나 정수가 아닌 문자열 같은 값이 올 수도 있습니다.

```go
package main

import "fmt"

func main() {
	s := "world"
	switch s {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("일치하는 문자열이 없습니다.")
	}
}
```

또한 복수개의 case 값이 있을 경우 콤마(`,` )로 연결해서 나열할 수 있습니다.

```go
package main

import "fmt"

func main() {
	n := 4

	switch n {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3,4:
		fmt.Println(3, 4)
	default:
		fmt.Println("match failed")
	}
}
```

위의 코드를 실행시켜보면 n이 4이기 때문에 case 3, 4에 걸려서 3, 4가 출력되는 것을 볼 수 있습니다.

- switch에서 break 하기

  Go에서는 switch문에서 break로 끝을 내주지 않아도 되지만 중간에 break를 넣어서 일정 코드만 실행시키게 할 수 있습니다.

    ```go
    package main

    import "fmt"

    func main() {
    	n := 2

    	switch n {
    	case 1:
    		fmt.Println(1)
    	case 2:
    		fmt.Println(2)
    		break
    		fmt.Println("no print")
    	case 3,4:
    		fmt.Println(3, 4)
    	default:
    		fmt.Println("match failed")
    	}
    }
    ```

- switch fallthrough

  특정 case의 문장을 실행한 뒤 다음 case의 문장을 실행하고 싶을 때는 fallthrough 키워드를 사용합니다.

  마치 C, C++의 switch 분기문에서 break 키워드를 생략한 것처럼 동작합니다. (컴파일러가 자동으로 break를 추가하기 때문)

  단, 맨마지막 case에는 fallthrough 키워드를 사용할 수 없습니다.

    ```go
    package main

    import "fmt"

    func main() {
    	i := 3
    	
    	switch i {
    	case 4:
    		fmt.Println("4 이상")
    		fallthrough
    	case 3:
    		fmt.Println("3 이상")
    		fallthrough
    	case 2:
    		fmt.Println("2 이상")
    		fallthrough
    	case 1:
    		fmt.Println("1 이상")
    	}
    }
    ```

> 타입 switch

스위치 구문은 인터페이스 변수의 동적 타입을 확인하는데 사용될 수도 있습니다.

이러한 스위치는 타입 단언(Type assertion)의 문법을 사용하되 괄호안에 키워드 `type`을 사용합니다.

```go
package main

import "fmt"

func main() {
	var t interface{} = "hello"
	switch s := t.(type) {
	case int:
		fmt.Println("int", t)
	case bool:
		fmt.Println("bool", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("unknown type", t)
	}
}
```