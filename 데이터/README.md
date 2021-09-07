# 배열

일반적인 프로그래밍 언어의 배열에 해당한다. 

일종의 같은 데이터 유형을 공유하는 값의 모음이라고 생각할 수 있다. 

앞으로 등장하는 예제를 보며, Go에서의 배열의 특징을 살펴보도록 하자.

### Go의 배열은 포인터가 아니라 그 자체로 값이다. 
```go
r := [...]string{"aa", "bb", "cc", "dd"}
fmt.Printf("Value of r: %d\n", r)
fmt.Printf("Size of r: %d\n", len(r))
```
출력을 해보면 전체 크기의 배열이 출력되는 것을 볼 수 있다.
```
Value of r: ["aa", "bb", "cc", "dd"]
Size of r: 100
```
* r에서 활용한 배열 정의를 배열 Literal이라고 부르기도 한다. 배열의 선언과 동시에 고정된 값을 넣어 선언해주는 것을 의미한다. 


* ... 은 일반적으로 크기가 정해지지 않아 값이 변할 수 있는(가변) 부분에서 사용된다. 
  
  해당 상황에서는 값을 굳이 명시하지 않고, 뒤에 명시된 원소들의 개수를 기반으로 자동으로 배열의 크기를 할당받기 위해 사용하였다.
  
### Go에서는 배열를 정의하는 방법이 다양하다. 

size를 정의하면서 사용자가 값을 넣어줄 수도 있고, elipsis(...)을 통해서 컴파일러가 자동으로 배열에 들어간 원소를 기준으로 크기를 캐치하도록 할 수도 있다.

```go
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
d := [2]int{1, 2}
e := [3]int{1, 2}
```

a와 b와 d에는 모두 정확히 동일한 배열가 들어있다. 

c의 경우에는 원소가 다른 부분이 존재하고(배열 전체가 하나의 값이므로 원소가 하나만 달라도 다르다), e는 배열의 크기가 다른 것들과 다르기 때문에 다른 배열이다.

 특히나 e와 다른 배열을 비교하고자 하는 시도는 애초에 타입이 달라서 불가능하다.

```go
fmt.Printf("a == e: %d\n\n", a == e) // cannot compare [2]int == [3]int
```

### Pass by Value vs Pass by Pointer

C가 그러한 것처럼, 외부에서 들어온 매개변수를 함수 내부에서 수정하고 싶은 경우 직접 해당 주소에 접근해야 힌다.

```go
func byte_zero_one(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func byte_zero_two(ptr *[32]byte) {
	*ptr = [32]byte{}
}
```

두 함수 중 `byte_zero_one`에 해당하는 함수는 내부적으로 `ptr`의 값들을 0으로 초기화 시켰지만, 해당 변경 사항이 외부로 전달되지 않는다. 

이 경우에 외부로 변경 사항을 내보내려면 반환이라는 추가 작업이 필요하다.

반대로 `byte_zero_two`의 경우, 직접 배열의 포인터에 접근했기 때문에 원본 값에 영향을 미칠 수 있다. 

주소를 참조하는 경우, 참조된 주소는 배열의 맨 처음 값을 나타낸다.

다만, 매개변수의 형태에서 볼 수 있듯 배열의 크기가 지정되어 있다. 

만약 배열의 크기가 다르다면 해당 예시에서 볼 수 있듯 해당 함수를 사용할 수 없는 경우가 발생한다. 

때문에 Go에서는 단순히 배열을 사용하기 보다는 슬라이스라는 데이터 타입을 더 자주 활용하는 편이다.

# Slice

슬라이스를 간단하게 표현하자면 배열의 일부분이라고 할 수 있다. 배열과 마찬가지로 인덱싱이 가능하며 길이가 존재한다. 배열과 달리 길이를 변경할 수 있다는 점이 특징이다.

예제 코드를 코드를 보며 하나씩 이해해보자.
```go
slice1 := []int{0, 1, 2, 3}
slice2 := append(slice1, 4, 5)
slice3 := make([]int, 2)
copy(slice3, slice2)
fmtPrintln(slice1, slice2)
```
가장 기본적인 슬라이스의 사용방법이다. 

프로그램을 실행하면 "0, 1, 2, 3" 값의 slice1과 "0, 1, 2, 3, 4, 5"의 값을 가진 slice2가 있다. 

append 함수를 통해 기존 슬라이드를 가져오고 그 뒤에 4, 5를 추가하여 새 슬라이드를 만들었다.

make는 동적으로 할당된 배열을 만들기 위한 go의 built-in 함수이다.

`p []int = make([]int, 0, 5)` 라는 표현식에서, 

- []int: 할당하고자 하는 타입
- 0은 length
- 5는 capacity를 의미한다.

이와 유사하게 Go에는 동적으로 할당할 때 사용하는 `new`라는 함수가 존재한다.

new 함수를 사용한 예시를 보면서 알아보자

```go

type ListNode struct {
     Val int
     Next *ListNode
}

func add newNode(nodePtr *ListNode) {
	if head == nil {
		nodePtr := new(ListNode)
	} else {
		nodePtr.Next = new(ListNode)
	}
}
```

다음의 [StackOverflow에서의 질문 및 토론]()에서 알 수 있듯, 이것이 stack 영역에 할당되느냐 heap에 할당되느냐는 go에서 그렇게 중요한 주제가 아니다. 해당 토론에서 모두가 입을 모아, 변수가 힙에 할당되는지 스택에 할당되는지 보다 그냥 사용하라고 한다. 중요한 것은 Make를 언제 사용해야하고 new를 언제 사용해야하냐는 것이다. 

Make와 new의 차이

1) make는 만들고자 했던 데이터 타입을 반환하는 반면, new는 포인터를 반환한다.
2) make는 메모리를 할당하고 명시된 length, capacity에 맞춰 초기화하지만, new는 메모리를 할당하고 단순히 해당 타입에 맞춰서만 초기화가 이뤄진다.

channel, slice, map은 go의 데이터 타입 중에서 동시에 여러 값을 가질 수 있는 성격을 가지고 있다. 때문에 해당 타입들을 사용하기 위해서는 사용하기 이전에 초기화가 요구된다. new도 초기화가 이뤄지기는 하지만Make는 아래 세 개의 타입을 위해 등장한 새로운 키워드다.

  - channel
  - slice
  - map


## slice는 배열에 대한 참조다.

```go
	d := []string{"r", "o", "a", "d"}
	e := d[2:]
	fmt.Println(e)
	// e == []string{'a', 'd'}
	e[1] = "m"
	// e == []string{'a', 'm'}
	fmt.Println(e)
	// d == []string{'r', 'o', 'a', 'm'}
	fmt.Println(d)

	/********
	Output
	---------
	[a d]
	[a m]
	[r o a m]
	*********/
```

해당 코드를 보면, 기존의 배열에서 잘라온 슬라이스를 값을 변경했을 때 기존의 배열이 어떻게 변하는지가 나타난다. 결론부터 말하자면 잘라온 슬라이스에서 변형이 이뤄지면 기존의 배열도 변화한다. 이는 슬라이스가 복사된 새로운 데이터를 만드는게 아닌, 기존 배열에 특정 부분을 가리키는 포인터로 동작하기 때문이다.

```go
s = s[2:4]
```
![](https://go.dev/blog/slices-intro/slice-2.png)

때문에 기존의 다른 데이터 타입의 경우, 

## 슬라이스가 빈 경우를 확인하는 방법

슬라이스가 비었는지를 확인하는 가장 좋은 방법은 슬라이스의 크기가 0인지를 확인하는 것이다.
```go
	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil

	// slice가 비었는지 확인하는 적절한 방법
	if len(s) == 0 {
		print("s is empty")
	}
```

코드의 주석에서 볼 수 있듯, 슬라이스가 nil인지 여부만으로 판단하기에는 모든 경우를 포괄하지 못한다.

앞서 간단하게 살펴보았듯, 슬라이스에는 append 함수가 존재한다.
```go
	s = make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p, %d\n", cap(s), len(s), s, s)
	}
```
이때, append를 할 경우, cap과 len의 변화에 주목해보자. 

len의 경우 문자 하나가 추가될 때마다 하나씩 추가된다.

cap의 경우 문자가 하나씩 추가되다가 cap의 상한을 만나게된다면 2배가 된다.

이를 코드로 표현하자면 다음과 같다.

```go
func appendSingleInt(x []int, y int) []int {
		var z []int
		zlen := len(x) + 1
		if zlen <= cap(x) {
			// 더 작성될 공간이 남았다면, 슬라이스를 연장한다.
			z = x[:zlen]
		} else {
			// 더 작성될 공간이 없다면, 배열의 크기를 재할당한다.
			zcap := zlen
			if zcap < 2*len(x) {
				zcap = 2 * len(x)
			} 
			z = make([]int, zlen, zcap)
			copy(z, x)
		}
		z[len(x)] = y
		return z
}
```
+ 참고 append를 활용하여 여러 원소를 한번에 넣는 것도 가능하다.
```go
s = append(s, 1, 2)		// 1과 2를 추가
s = append(s, s...) 	// s에 s를 한번더 추가
```

# Map

키값과 대응하는 값으로 이루어진 Map은 해시 테이블로부터 비롯된 자료 구조이다. Go에서는 Map을 기본 자료형으로 제공하고 있으며, 다음과 같이 정의한다.

```go
ages := make(map[string]int)
```

make를 통해 map을 할당하였다. 이때, 먼저오는 string은 키(key)로, 따라오는 int는 값(Value)로 사용된다.

map에서 키와 값을 추가하는 방법은 간단하다.
```go
ages["alice"] = 32
ages["charlie"] = 34
```

다음과 같이 직관적인 문법으로 새로운 키-값 쌍을 추가할 수 있다.

위의 선언과 대입을 map literal로 표현하면 다음과 같이 표현이 가능하다.

```go
agesMapLiteral := map[string]int {
	"alice": 32,
	"charlie": 34,
}
```

키-값 쌍을 제거하기 위해서는 built-in 함수인 delete를 사용한다.

```go
delete(ages, "charlie")
```

Map에 대해 주소값을 가져오려는 시도는 컴파일 에러로 간주된다.
```go
_ = &agesMapLiteral["bob"]
```

이는 Go의 Map의 특정 값이 가르키는 주소값이 일정하지 않기 때문에 발생한다. 

Go의 맵은 해시 맵으로, Go에서는 삽입에 의해 맵의 키-쌍 값이 추가됨에 따라 더 많은 버킷을 제공하도록 조정한다. 

이때 맵의 키-쌍은 새롭고 더 큰 버킷 배열에 증분되어 복사된다. 

반대로 키-쌍 값이 삭제된다면 기존의 공간이 회수된다. 

때문에 맵의 특정 키-값에 대해서 고정된 주소를 얻고자 하는 것은 올바르지 않은 동작으로 정의된다.

ages에 ["alice", "bob", "mia", "dana"]가 들어있다고 가정하고 다음 코드를 여러번 실행하는 경우, 실행 순서가 계속해서 변화하게 된다.

[source](https://stackoverflow.com/questions/38168329/why-are-map-values-not-addressable)

```go
fmt.Println("\nLoop")
for name, age := range ages{
	fmt.Printf("%s\t%d\n", name, age)
}

/************
	alice   32 
	bob     1
	mia     1
	dana    1

	bob     1
	mia     1
	dana    1
	alice   32

	dana    1
	alice   32
	bob     1
	mia     1
*************/
```

# struct

go에는 동일하거나 다른 유형의 필드가 포함된 구조체 유형이 존재한다.  구조체는 기본적으로 논리적 의미 또는 구성을 갖는 명명된 필드의 모음이며, 각 필드에는 특정 유형이 있다.

일반적으로 구조체 유형은 사용자 정의 데이터 유형을 나타낸다. 이미 내장된 데이터 유형으로 불충분한 경우, 구조체라는 사용자 정의 데이터 유형을 정의하여 사용할 수 있다.

이해를 돕기 위해, 게시할 블로그 게시물이 있다고 가정해보자.

```go
type blogPost struct {
	author 	string 	// 필드 1
	title	string 	// 필드 2
	postId	int		// 필드 3
}
```

구조체 정의의 예시이다. blogPost라는 구조체 내에 3개의 다른 필드를 정의하였다.

이렇게 정의된 구조체는 다음과 같이 사용될 수 있다.

```go
package main

import "fmt"

type blogPost struct {
	author string
	title  string
	postId int
}

func main() {
	var b blogPost // 구조체 초기화, 별도의 값을 지정하지 않는다면 zero value

	fmt.Println(b) // zero value 출력
	b = blogPost{
		author: "youngmki",
		title:  "Go programming basic",
		postId: 12345, // }를 다음 줄로 분리한다면 맨 마지막에도 컴마(,)를 표시해줘야 한다.
	}
	fmt.Println(b)
}
```
[코드 실행](https://play.golang.org/p/42qWonnag6I)

new 키워드를 사용하여 b에 포인터를 담는다면 다음과 같이 코드를 작성할 수 있다.

```go
package main

import "fmt"

type blogPost struct {
	author string
	title  string
	postId int
}

func main() {
	b := new(blogPost) // b := &(blogPost{})로 작성할 수도 있다.

	fmt.Println(b) // zero value 출력
	b = &blogPost{
		author: "youngmki",
		title:  "Go programming basic",
		postId: 12345, // }를 다음 줄로 분리한다면 맨 마지막에도 컴마(,)를 표시해줘야 한다.
	}
	fmt.Println(b)
}
```

`=` 기호를 통해 접근할 수도 있지만, dot `.`을 통해서도 접근할 수 있다.

```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        var b blogPost // blogPost 타입에 해당하는 b를 생성한다.
        b.author = "youngmki"
        b.title = "Go programming basic"
        b.postId = 12345

        fmt.Println(b)  

        b.author = "Chinedu" 
        fmt.Println("Updated Author's name is: ", b.author)           
}
```

new 키워드를 통해 선언된 구조체 포인터 변수도 dot `.`을 통해 해당 구조체 변수에 바로 접근할 수 있다.

```go
	b.postId = 13333
```

```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        b := blogPost{"youngmki", "Go programming basic", 12345}
        fmt.Println(b)        
}
```

다른 타입과 마찬가지로 리터럴로 간단하게 표현할 수도 있다.

만약 스코프 내에서만 유효한 구조체를 만들고 싶다면 다음과 같은 문법을 사용한다.

```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {

        // main 함수 내에서 구조체 선언 및 초기화
        b := struct {
          author  string
          title   string
          postId  int  
         }{
          author: "youngmki",
          title:"Go programming basic",
          postId: 12345,
        }

        fmt.Println(b)           
}
```

## 중첩된 struct 필드

다음과 같은 구조체 두개가 있다고 하자

```go
type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}
```

구조체의 필드에는 X, Y, Radius가 중복된다.

해당 경우에는 구조체 `Wheel`을 아래와 같이 사용할 수 있을 것이다. 
```go
var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5
w.Spokes = 20
```

구조체에서 유사성과 반복성을 발견한다면, 다음과 같이 공통적인 부분을 고려하는 것이 편리할 수 있다.

```go
type Point struct {
	X, Y int
}

type Circle struct {
	Center 	Point
	Radius 	int
}

type Wheel struct {
	Cir	Circle
	Spokes 	int
}
```

이렇게 구조체를 작성하게 된다면, 훨씬 명확하게 필드에 접근이 가능하지만, 깊게 위치한 필드에 접근이 번거로워진다는 단점이 있다.
```go
var w Wheel
w.Cir.Center.X = 8
w.Cir.Center.Y = 8
w.Cir.Radius = 5
w.Spokes = 20

// w.X << 직접적인 접근은 불가능해진다.
```
[실행 링크](https://play.golang.org/p/8mmlw7mMZPZ)

## 익명 필드로 구성된 구조체

직전의 중첩된 구조체를 다음과 같이 작성할 수도 있다.
```go
type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius 	int
}

type Wheel struct {
	Circle
	Spokes 	int
}
```

이는 다음과 같이 초기화 된다.
```go
w := Wheel{Circle{Point{8, 8}, 5}, 20}

or

w := Wheel {
	Circle: Circle{
		Point:	Point{X: 8, Y: 8}, 
		Radius:	5,
		},
	Spokes: 20,
}
```

이처럼 중첩된 구조체를 내부에서부터 하나씩 채워나감으로써 익명 필드로 구성된 구조체를 사용할 수 있다.

익명 함수로 작성된 경우, Wheel의 타입을 갖는 w가 곧바로 X에 접근이 가능하다.

```go
w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	
	w.X = 20 // w.Cir.Point.X 와 동일한 표현
```
주의해야 할 점은, 익명 필드의 경우 한 구조체 내에 동일한 타입이 두번 이상 등장할 수 없다는 것이다. 익명 필드는 타입을 기준으로 인식이 되기 때문에, 서로 다른 필드의 타입 중복은 모호할 수 밖에 없기 때문이다.

```go
type Wheel struct {
	Circle
	Circle		// ERROR
	Spokes 	int
}
```
[실행 링크](https://play.golang.org/p/NgGZb8BSr8D)


# JSON
----

Go언어에서 struct는 JSON과 긴밀한 연관을 갖는다. Go의 `encoding/json` 이라는 모듈을 통해 구조체 형식에서 JSON 형식으로의 [마셜링](https://ko.wikipedia.org/wiki/%EB%A7%88%EC%83%AC%EB%A7%81_(%EC%BB%B4%ED%93%A8%ED%84%B0_%EA%B3%BC%ED%95%99))과 반대로 JSON 파일을 알맞게 정의된 구조체 타입으로 변환해주는 디마셜링을 지원한다.

## 마셜링

구조체의 형식을 JSON 형태로 변환하는 과정이다.

기존의 struct 정의 부분과 약간의 차이가 존재하고 `Marshal` 이라고 하는 메서드를 사용한다.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string 	`json:"title"`		// 주의: 작은 따옴표가 아니라 back-tip(`) 이다
	Author string  	`json:"author"`		
}

func main() {
	book := Book{Title: "Learning Concurrency in golang", Author: "Youngmki"}	// 구조체 초기화
	
	// %#v 해당하는 값을 생성하는 소스 코드 스니펫
	fmt.Printf("%#v\n\n", book)
	
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
} 
```
Output

```
main.Book{Title:"Learning Concurrency in Python", Author:"Elliot Forbes"}

{"title":"Learning Concurrency in Python","author":"Elliot Forbes"}
```

마셜링 이전과 이후의 차이에 주목하자. struct 형태로 존재하던 값이 JSON 파일에서 읽을 수 있도록 변환되었다.

마셜링을 진행할 때, 좀 더 우리에게 친숙한 JSON 형식으로 변환하기 위해서는 `MarshalIndent` 함수를 사용한다.

> 참고로, MarshalIndent의 두번째 인수는 매 줄마다 prefix할 값("")이 들어가고, 세번째 인수에는 각 사이를 구분해줄 값("	")이 들어간다.

```go
	byteArray, err := json.MarshalIndent(book, "", "	")
```

으로 변환해주었고, 결과는 다음과 같다.

Output
```
main.Book{Title:"Learning Concurrency in Python", Author:"Elliot Forbes"}

{
	"title": "Learning Concurrency in Python",
	"author": "Elliot Forbes"
}
```

메서드의 반환값으로는, 1) JSON 호환 string 데이터와 2) err 값이 있다.
[실행 링크](https://play.golang.org/p/ZASjJNMOyDP)

## Un-마셜링

이번엔 JSON에서 호환가능한 형식에서 구조체 형태로의 언마셜링을 진행하려고 한다.

원리는 간단하다. JSON 형식의 데이터를 `Unmarshal` 메서드를 활용해서, 해당 메서드의 두번째 인자에 넣어준다.

```go
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}

type Info struct {
	Description string `json:"desc"`
}

func main() {
	jsonString := `{
		"name":"battery sensor", 
		"capacity":40, 
		"time":"2021-09-05", 
		"info": {
			"desc": "a sensor reading"
		}
	}`
	fmt.Println(reflect.TypeOf("Type of jsonString" + jsonString))

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	// %+v, 값이 구조체인 경우 구조체의 필드명까지 포함
	fmt.Printf("%+v\n", reading)	
}
```

코드를 차근차근 살펴보자면, 먼저 메인함수에서 json 파일을 읽는 대신, 그와 동일한 역할을 해줄 string 형식의 값(`jsonString`)을 생성했다.

그 아래 `SensorReading` 타입으로 변수 `reading`을 생성한다. 

해당 타입은 내가 생성한 구조체로 Info라는 구조체와 중첩되어 있다.

이때, 받고자 하는 JSON 파일의 구조(여기서는 변수 jsonString)와 정의한 구조체의 정의가 동일하다는 점에 주목하자. 

`SeonsorReading` 타입의 `reading`이라는 변수에는 `Unmarshal` 함수를 통해 `JSON 형식`에서 앞서 정의한 `SensorReading` 구조체 형식으로의 변환된 값이 저장된다.

[실행 링크](https://play.golang.org/p/HvljExEP_Vg)
