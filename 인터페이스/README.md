# Interface

지금까지 살펴본 모든 데이터 타입은 구체적인 데이터 타입만을 나타냈습니다. 이전까지 살펴봤던 데이터 타입은 값에 대한 정확한 타입을 지정함과 동시에, 예를 들어 산술 연산, 인덱싱, 슬라이스의 append range 데이터 타입처럼 각각의 타입에 맞는 추가 행동을 제공했는데요. 정확히 해당 사용해야하는 데이터 타입이 무엇인지를 매번 명시할 수 있고, 이를 변수의 초기화 단계나 매개변수나 인자에 정확히 기입할 수 있다면 기존의 데이터 타입만으로도 충분합니다.

그런데 Go에는 인터페이스 유형이라고 하는 또 다른 유형이 있습니다. 앞서 언급하였던 구체적인 데이터 타입과 달리, 인터페이스는 추상 데이터 타입입니다. 이때 추상적이라는 의미는 해당 인터페이스에 속하는 구체적인 데이터 타입들에 대한 구체적인 구현을 모른채 사용되는 데이터 타입이라고만 일단 가볍게 넘어가겠습니다. 때문에 인터페이스를 사용하면 데이터 타입의 모든 메서드를 노출하는 대신, 지정한 특정 메서드만을 사용하게 할 수 있습니다.

인터페이스 유형은 구체적인 데이터 타입이 해당 인터페이스의 인스턴스로 간주되기 위해 보유해야만 하는 메서드 집합을 지정합니다. 특정 메서드가 구현되어 있지 않다면, 해당 데이터 타입은 인터페이스에 포함될 수 없습니다. 인터페이스의 대표적인 예시로 io.Writer가 있습니다. io.Writer는 파일, 메모리 버퍼, 네트워크 연결, HTTP 클라이언트, archiver, hasher 등을 포함하는 바이트를 사용하는 데이터 타입에 대해서 추상화를 가능하게 하기 때문에 자주 사용되는 인터페이스 중 하나입니다.

이외에도 [io](./io/io.go) 패키지에 존재하는 Reader는 바이트를 읽을 수 있는 모든 유형을 나타내며, Closer 값은 파일 또는 네트워크 연결과 같이 닫을 수 있는 모든 값을 의미합니다.

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
```

더 나아가서, 이러한 인터페이스를 또 포괄하는 새로운 인터페이스를 선언할 수도 있습니다. 이를 인터페이스 임페딩(embedding in interface)이라고 하며 다음과 같이 작성됩니다.

```go
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

기존의 struct를 활용한 embedding으로 앞선 예시를 다시 작성한다면 다음과 같습니다.

```go
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}
```

물론 두가지 스타일을 혼용하는 것도 가능한 문법입니다.

```go
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Writer
}
```

위의 세 선언 모두 같은 의미를 같습니다. 메서드가 나타나는 순서는 중요하지 않고 인터페이스 내부에 명시하고자 하는 메서드가 잘 들어있느냐가 중요합니다.

OOP의 Object란 객체의 상태와 기능을 포함합니다.
- 기능: 외부로 공개되는 공개 메서드와 내부에서만 사용되는 내부 메서드로 나뉩니다.
- 객체와 객체끼리는 서로 관계를 가질 수 있습니다. 만약, A객체와 함께 B객체의 메서드가 사용된다면 혹은 그 반대의 경우라면 서로 관계를 맺고 있다고 할 수 있다.
- 이떄 관계만 따로 정의한 것을 Interface라고 한다.

이때, 객체의 상태와 관계를 분리시키면 객체 간의 관계를 구성하는데 있어 훨씬 유연하게 코드를 작성할 수 있습니다. 이를 Decoupling이라고 하며 의미를 풀어보자면 종속성을 때어낸다 정도로 이해할 수 있겠습니다. 즉 추상화가 가능해집니다.

Golang의 인터페이스를 설명하기 위해 잼과 빵의 관계를 예시로 설명을 해보고자 합니다.

1. 빵과 딸기잼의 관계를 명명하고자 한다.
2. 빵에게 있어 딸기잼은 자신에게 발리는 존재이다.
3. 이를 함수로 표현한다면 '빵에 딸기잼을 바른다.'라는 함수를 작성할 수 있다.
4. 만약 여기서 오렌지잼이 추가되었다고 하자.
5. '빵에 오렌지 잼을 바른다' 라는 함수를 작성할 수 있다.
6. 사과잼이 추가되고, 포도잼도 추가되고, ... 먹을 수 있는 무수한 잼들이 추가될 예정이다.
7. 그렇다면 잼이 추가되는 매 순간마다 해당 함수를 추가할 것인가?


사실 빵에게 있어서 오렌지 잼을 바르는 행위는 딸기잼을 바르는 행위와 동일하다. `무엇을 바른다` 라는 것이 중요하지, `무엇이` 어떠어떠한 특성을 지녔는지는 중요하지 않다. 오직 바를 수 있다를 중점적으로 본다.

인터페이스를 사용한다면, 오렌지 잼을 추가하던, 파인애플 잼을 추가하던, 관계를 맺고자 하는 대상이 `발린다` 라는 기능(메서드)만 정의되어 있다면 기존의 인터페이스를 활용하여 해당 객체와의 관계를 맺을 수 있다. 즉 하나의 인터페이스를 정의함으로써 잼이 추가될 때마다 이를 추가해야할 필요성이 사라지는 것이다.

상태와 기능을 서로 분리함으로써 객체의 확장성이 늘이고 추상화를 시켰다는 점이 포인트다.

절차가 아닌 객체 하나하나에 집중해서, 해당 객체의 특성 혹은 다른 객체와의 관계를 통해 코드를 써내려간다.

Go의 인터페이스의 의미는 직관적으로 해당 함수를 포함하는 타입이라면 (구조체를 포함하는 개념) 이를 또 다시 하나의 타입으로 묶을 수 있게 된다는 것이다.

이를 통해, 추상화를 이루는 것이다.

인터페이스의 정의 부분은 다음과 같다
```go 
type [인터페이스 이름] interface {
	메소드의 이름 (입력 인자) 출력 타입	// 메소드 1
	메소드의 이름 (입력 인자) 출력 타입 // 메소드 2
	메소드의 이름 (입력 인자) 출력 타입 // 메소드 3
	...
}
```
구조를 보면 확실히 이해가 더욱 용이하다. 해당 인터페이스에 속할 수 있는 타입은, 해당 메소드가 모두 구현되어 있어야 한다.

> 이떄 메소드가 구현되어 있다는 뜻의 의미는 해당하는 메소드에 대해서 각 타입이 핸들러로 존재함을 의미한다.

다음은 인터페이스를 사용하지 않는 코드이다.

```go

type Bread struct {
	value string
}

type StrawberryJam struct {
}

type OrangeJam struct {
}

type SpoonOfStrawberryJam struct {
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return " + Orange"
}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry"
}

func (b *Bread) String() string {
	return "bread" + b.value
}

func (j *StrawberryJam) getOneSpoon() *SpoonOfStrawberryJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) getOneSpoon() *SpoonOfOrangeJam {
	return &SpoonOfOrangeJam{}
}

func (b *Bread) PutJamStraw(jam *StrawberryJam) {
	spoon := jam.getOneSpoon()
	b.value += spoon.String()
}

func (b *Bread) PutJamOrange(jam *OrangeJam) {
	spoon := jam.getOneSpoon()
	b.value += spoon.String()
}
```

해당 코드를 인터페이스를 활용한 코드와 비교해보자.

```go
type SpoonOfJam interface {
	String() string
}

type Jam interface {
	getOneSpoon() SpoonOfJam
}

func (j *StrawberryJam) getOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) getOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.getOneSpoon()
	b.value += spoon.String()
}
```

다음은 인터페이스를 활용하며 달라지거나 수정된 부분이다.

```go
	var jam *SpoonOfOrangeJam
	jam = &SpoonOfOrangeJam{}

	fmt.Println((jam))
	jam.String()
```

## Empty Interface

인터페이스가 메소드를 포함하지 않는 경우, 이를 Empty Interface라고 부르게 됩니다. 

```go
Interface{} // empty Interface
```

해당 인터페이스는 메소드를 포함하지 않기 때문에, 모든 타입을 포괄하는 것으로 이해할 수 있습니다.

다양한 함수에서 매개변수를 들여올 경우 타입에 관계없이 인자를 받고 싶을때 Empty Interface를 사용하게 됩니다.

예를 들어 fmt의 Println의 함수 프로토타입은 다음과 같습니다.

```go
func Println(a ...Interface{}) (n int, err error)
```

매개변수 정의를 `a ...Interface`로 했기 때문에 가변 인자를 들여올 수 있음과 동시에 어떠한 타입도 들여올 수 있게 되었음을 확인할 수 있습니다.

## 타입 단언 (Type Assertion )

타입 단언은 인터페이스 값에 적용되는 연산이며 문법적으로는 다음과 같이 사용됩니다.

```go
x.(T)
```

형식 어설션은 피연산자의 동적 타입이 단언한 타입과 호환되는지를 확인합니다.


만약 단언한 형식 T가 구체적인 타입이라면, 타입 단언은 x의 동적 타입이 T로 변환 가능한지를 확인합니다. 만약 이 검사가 성공한다면, 타입 단언의 결과는 인터페이스로 추상화된 타입이 명시한 구체적인 타입으로 변환됩니다. 만약 검사가 실패한다면, panic에 빠지게 됩니다.

```go
var w io.Writer
w = os.Stdout
f := w.(*os.File)			// success: f == os.Stdout
c := w.(*byte.Buffer) 		// panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
```
[코드 실행](https://play.golang.org/)

타입 단언이 사용되는 대표적인 사례는 정적 타입에서 함수를 사용하고 싶지만, 접근이 제한되어 이를 동적 타입으로 변형하고 싶은 경우 입니다. 언어로 표현하는데는 한계가 있어 코드로 예시를 들어보겠습니다.

```go
package main

import (
	"fmt"
)

// Shape ...
type Shape interface{
	Area() float64
}

// Object ...
type Object interface{
	Volume() float64
}

// Cube ...
type Cube struct{
	side float64
}


// Area ...
func (c Cube) Area() float64{
	return 6 * (c.side * c.side)
}

// Volume ...
func (c Cube) Volume() float64{
	return c.side * c.side * c.side
}

func main() {
	c:= Cube{3}
	var s Shape = c
	var o Object = c
	fmt.Println("volume of interface of type Shape is", s.Area())
	fmt.Println("volume of interface of type Object is", o.Volume())
}

//Result
volume of interface of type Shape is 54
volume of interface of type Object is 27
```

코드를 살펴보면 main 함수에 두가지 변수가 존재하는 것을 확인할 수 있습니다. 이때 s는 Shape 인터페이스를 정적 타입으로 가지며, o는 Object 인터페이스를 정적 타입으로 가지게 됩니다.

이 경우 만약 s에서 정의되지 않은 메소드, 예를 들어 Volume()을 사용하고자 하는 경우 에러가 발생합니다.

이때 정적타입이 필요합니다. 정적 타입을 통해 인터페이스들의 동적 타입인 Cube를 뽑아낸다면, 앞서 인터페이스로 구분되었던 경계가 없어져 비로소 Cube의 모든 메소드를 사용할 수 있게 됩니다.

다음과 같이 사용할 수 있습니다.

```go
func main() {
	var s Shape = Cube{3}
	c := s.(Cube)
	fmt.Println("volume of interface of type Shape is", c.Area())
	fmt.Println("volume of interface of type Object is", c.Volume())
}

///Result
volume of interface of type Shape is 54
volume of interface of type Object is 27
```

앞서 언급했던 바와 같이, 타입 단언을 통해 정적 타입(인터페이스)에서 동적 타입으로의 변환을 거쳐 Volume을 사용할 수 있게 되었습니다.

만약 타입 단언이 허용되지 않은 동작으로 규정되는 경우, 즉 타입 단언이 실패하는 경우에는 설명 초반부에 설명하였듯 panic 상태로 돌입하며 대입을 진행하는 경우 nil 값이 들어가게 됩니다. 두번째 에러를 받는 변수를 동시에 대입시키면서 타입 단언이 성공적으로 완료 되었는지를 확인할 수도 있습니다.

[코드 예시](https://play.golang.org/p/FwJ9d-PjQEB)

이를 다른 관점에서 볼 경우, 해당 정적 타입이 명시한 인터페이스를 구현하는지를 확인한다고도 볼 수 있습니다. 해당 타입이 인터페이스를 구현하고 있는 경우, 타입 단언이 성공하게 되므로 ok 변수를 별도로 받는다면 True 값을 받게 됩니다. 반대로 타입 단언이 실패한다면 false 값을 받게되고, 이는 타입 단언을 하고자 하는 타입이 해당 인터페이스를 구현하고 있지 않다는 의미가 됩니다.

```go
m, ok := val.(json.Marshaler)
```
m에는 타입 단언이 성공했을 경우 값이 담기게 될 것이고, ok 에는 변환의 성공 여부가 담기게 될 것입니다.

단순히, 해당 정적 타입이 인터페이스를 구현하고 있는지의 여부만을 활용한다면 [공백 식별자](../공백식별자/README.md)를 통해 다음과 같이 표현할 수 있습니다.

```go
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```

다만, 몇몇 인터페이스 검사는 컴파일 타임에 발생합니다. 예를 들어 앞선 예시 중 `*os.File`, `io.Reader` 사이의 관계에서는 애초에 인터페이스 타입이 맞지 않는다면 컴파일조차 불가능하게 됩니다.

```go
var w io.Writer
w = os.Stdout
f := w.(*os.File)			// success: f == os.Stdout
c := w.(*byte.Buffer) 		// panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
```
[코드 실행](https://play.golang.org/)

## Type Embedding

여러가지 두가지 이상의 인터페이스를 병합하는 새로운 인터페이스를 만드는 것을 Type Embedding이라고 부릅니다.

```go
package main

import (
	"fmt"
)

// Shape ...
type Shape interface{
	Area() float64
}

// Object ...
type Object interface{
	Volume() float64
}

// Material ...
type Material interface{
	Shape
	Object
}

// Cube ...
type Cube struct{
	side float64
}

// Area ...
func (c Cube) Area() float64{
	return 6 * (c.side * c.side)
}

// Volume ...
func (c Cube) Volume() float64{
	return c.side * c.side * c.side
}

func main() {
	c:= Cube{3}
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("dynamic type and value of m of static type Material is %T and %v\n", m, m)
	fmt.Printf("dynamic type and value of s of static type Shape is %T and %v\n", s, s)
	fmt.Printf("dynamic type and value of o of static type Object is %T and %v\n", o, o)
}

//result
dynamic type and value of m of static type Material is main.Cube and {3}
dynamic type and value of s of static type Shape is main.Cube and {3}
dynamic type and value of o of static type Object is main.Cube and {3}
```
해당 예시에서 Cube는 Area와 Volume 메소드를 구현하기 때문에, Shape와 Object 인터페이스를 구현합니다. Material 인터페이스는 앞서 나열한 두 인터페이스를 포함하기 때문에, Cube는 Material 또한 구현하게 됩니다.

