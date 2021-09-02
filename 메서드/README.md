# 메서드

Go 언어는 객체지향 프로그래밍(OOP)를 고유의 방식으로 지원합니다.

타 언어와 달리 OOP에서 클래스가 필드와 메서드를 갖는 것과 달리 Go언어에서는 사용자 지정 타입이 필드를 가지며 메서드는 별도로 분리되어 정의됩니다.

Go의 메서드는 특별한 형태의 함수입니다.

메서드는 함수를 정의할 때 func 키워드와 함수명 사이에 어떤 사용자 지정 타입을 위한 메서드인지를 표시하게 됩니다.

흔히 리시버(receiver)로 불리는 이 부분은 메서드가 속한  변수명과 사용자 지정 타입을 지정하는데 struct 변수명은 함수 내에서 매개변수 처럼 사용됩니다.

메서드의 기본적인 형태는 다음과 같습니다.

`func (<변수명> <사용자 지정 타입>) <메서드이름>(<메서드가 받을 매개변수>) (<반환 값>)`

다음 코드는 메서드를 사용한 예제입니다.

```go
package main

import "fmt"

type Rect struct {
	width, height float64
}

type Liter int

func (l *Liter) double() {
	*l *= 2
}

func (r *Rect) area() float64 {
	return r.width * r.height
}

func main() {
	water := Liter(3)
	rect := Rect{10.5, 22.3}
	water.double()

	fmt.Println(rect.area())
	fmt.Println(water)
}
```

위의 코드를 보면 struct가 아닌 int 타입의 자료형도 사용자 지정 타입으로 만든 후 사용할 수 있습니다.

메서드의 리시버는 포인터와 인터페이스를 제외한 모든 타입을 가질 수 있습니다.

그리고 메서드를 이용하면 함수명이 충돌하거나 너무 길어질 일 없이 사용할 수도 있습니다.

예를들어, int 타입으로 만든 사용자 지정 타입이 2개 있고 각각의 타입마다 원래 값에 2를 곱해서 반환해주는 프로그램을 만들고 싶다고 했을 때 메서드를 이용하면 함수의 중복과 함수명이 길어지는 것을 방지할 수 있습니다.

```go
package main

import "fmt"

type width int
type height int

func (w width) double() width {
	return w * 2
}

func (h height) double() height {
	return h * 2
}

func main() {
	w := width(100)
	h := height(300)

	fmt.Println(w.double())
	fmt.Println(h.double())
}
```

위의 코드처럼 메서드를 사용하지 않는다면 중복을 피하기 위해 `widthDouble()` `heightDouble()` 이런식으로 함수명이 길어졌겠지만 메서드를 사용하여 더 간결하게 작성할 수 있습니다.

그리고 외부로 만약 메서드를 노출시키고 싶을 경우 함수와 타입과 마찬가지로 맨 앞글자를 대문자로 변경하면 외부로 노출이 가능해집니다.

> 리시버 매개변수명 규칙

Go에서 리시버 매개변수의 이름에 아무 이름이나 사용할 수 있지만 가독성을 위해 하나의 타입에서 정의하는 모든 메서드에서는 동일한 리시버 이름을 사용하는 것이 좋습니다.

Go 개발자 들은 보통 리시버의 이름으로 리시버 타입의 첫 번째 문자를 소문자로 사용하는 컨벤션을 따릅니다.

> Pointer receiver(포인터 리시버) VS Value receiver (값 리시버)

값 리시버는 변수의 데이터를 복사하여 메서드에 전달하며, 포인터 리시버는 변수의 포인터만을 전달합니다.

값 리시버의 경우 만약 메서드내에서 그 값이 변경되더라도 호출자의 데이터는 변경되지 않는 반면 포인터 리시버는 메서드 내의 값 변경이 그대로 호출자에서 반영됩니다.

다음 예제를 실행시켜보면

```go
package main

import "fmt"

type Liter int

func (l *Liter) PtrDouble() {
	*l *= 2
}

func (l Liter) ValueDouble() {
	l *= 2
}
func main() {
	water := Liter(5)
	milk := Liter(3)

	water.PtrDouble()
	milk.ValueDouble()

	fmt.Println(water)
	fmt.Println(milk)
}
```

포인터를 보낸 water는 값을 보내는 milk는 값이 유지됩니다.

그리고 값 리시버에 포인터 변수를 보낼 경우 자동으로 포인터 변수의 값을 가져오고 포인터 리시버에 변수를 보낼 경우 자동으로 포인터로 변환해서 보냅니다.

위의 내용을 코드로 한번 구현해보겠습니다.

```go
package main

import "fmt"

type myString string

func (m myString) method() {
	fmt.Println("Method with value receiver")
}

func (m *myString) ptrMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := myString("a value")
	pointer := &value
	value.method()
	value.ptrMethod() // 값은 자동으로 포인터로 변환됩니다.
	pointer.method()  // 포인터가 가르키는 값을 자동으로 가져옵니다.
	pointer.ptrMethod()
}
```

하지만 위의 코드는 컨벤션을 어기고 있는데,

코드의 일관성을 위해 특정 타입의 메서드를 정의할 때 가급적이면 값 리시버와 포인터 리시버의 혼용은 피하고 두 타입 중 하나만 사용하는것을 권장하고 있습니다.