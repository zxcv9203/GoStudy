# 캡슐화와 임베딩

> 설정자 메서드

특정 연, 월, 일을 저장하고 싶을 때 사용자 지정 타입으로 구조체를 만들어서 다음과 같이 선언해서 사용하면 됩니다.

```go
package main

type Date struct {
	Year int
	Month int
	Day int
}

func main() {
	date := Date{Year: 2021, Month: 9, Day:7}
	fmt.Println(date)
}
```

위의 코드는 잘 동작하지만 문제점이 있는데 실제 존재하지 않는 요일들도 값으로 들어가진다는 것입니다.

예를들어, 연, 월, 일에 음수는 존재하지 않지만 현재는 음수를 직접 집어 넣을 수 있습니다.

그래서 값을 집어넣기전에 이게 유효한 값인지 확인하는 과정이 필요하며, 컴퓨터 과학에서는 이 과정을 `데이터 유효성 검증` 이라고 합니다.

구조체의 각 필드에 값을 검증하고 설정하기 위한 메서드들을 추가해보겠습니다.

이런 역할을 하는 메서드를 `설정자 메서드(setter method)`라고 부릅니다.

Go의 컨벤션에 따라서 보통 `Set<methodname>`의 형태로 선언합니다.

다음과 같이 유효한 범위내의 값만 설정할 수 있도록 설정자 메서드를 만들어 봅시다.

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year int
	Month int
	Day int
}

func (d Date) SetYear(year int) error {
	if year <= 0 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d Date) SetMonth(month int) error {
	if month <= 0 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d Date) SetDay(day int) error {
	if day <= 0 {
		return errors.New("invalid day")
	}
	d.Year = day
	return nil
}
func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(9)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
```

하지만 위의 코드를 실행시켜보면 실제로 값이 저장이 되지 않고 제로 값으로 초기화 된 상태입니다.

왜 메서드에서 값을 변경해주었는데 값이 계속 제로 값일까요?

포인터로 값을 전달하지 않으면 값의 복사만 이루어지기 때문에 실제로 값을 변경하기 위해서는 포인터 값을 보내야하기 때문입니다.

그렇다면 값이 실제로 변경되는지 확인하기 위해 Date 타입을 포인터로 받아봅시다.

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year int
	Month int
	Day int
}

func (d *Date) SetYear(year int) error {           // 리시버 매개변수를 포인터로 받음
	if year <= 0 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {          // 리시버 매개변수를 포인터로 받음
	if month <= 0 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {               // 리시버 매개변수를 포인터로 받음
	if day <= 0 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}
func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(9)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
```

이제 값이 잘 변경되는 것을 확인할 수 있습니다.

> 접근자 메서드

설정자 메서드를 이용하면 안전하게 값을 집어 넣을 수 있지만 현재는 구조체에 값을 직접 할당하면 데이터가 유효한지 검사를 할 수 없습니다.

따라서 설정자 메서드를 통해서만 필드를 설정할 수 있도록 필드를 보호하는 방법이 필요합니다.

그러기 위해서 사용자 지정 타입을 별도 패키지로 분리한 다음 필드를 모두 노출시키지 않는 방법을 통해서 막을 수 있습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/cd6d5295-67e6-4890-a6e6-131542a14465/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210907%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210907T063935Z&X-Amz-Expires=86400&X-Amz-Signature=648f2e629f0cc4f76955db917dc6a45f7e9c623dd75a62823b424b71d8ae1fca&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

- date.go

```go
package Date

import "errors"

type Date struct {
	year int
	month int
	day int
}

func (d *Date) SetYear(year int) error {
	if year <= 0 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month <= 0 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day <= 0 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
```

- main.go

```go
func main() {
	date := calendar.Date{}
	date.year = 2019
	date.month = 14
	date.day = 50
	fmt.Println(date)
	date = calendar.Date{year: 0, month:0, day: -2}
}
```

메인을 실행시켜보면 필드를 소문자로 변경하여 더이상 접근이 불가능하기 때문에 에러가 발생합니다.

이제 설정자 메서드로만 값을 설정할 수 있게 바꾸어주었습니다.

하지만 구조체 필드를 외부로 노출하지 않을경우 구조체 각각의 필드만 접근을 불가능 해집니다.

그래서 구조체의 필드 또는 변수의 값을 가져오는 것이 주요목적인 메서드인 `접근자 메서드(getter method)`를 추가해야합니다.

접근자 메서드를 구현하는 것은 매우 간단한데 단순히 필드의 값을 반환해주기만 하면 됩니다.

그리고 Go에서는 컨벤션에 따라 보통 접근자 메서드 이름에는 접근하고자 하는 필드나 변수의 이름과 동일한 이름을 사용합니다.

```go
package Date

import "errors"

type Date struct {
	year int
	month int
	day int
}

func (d *Date) SetYear(year int) error {
	if year <= 0 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) Year() int {                    // 접근자 메서드
	return d.year
}

func (d *Date) SetMonth(month int) error {
	if month <= 0 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) Month() int {                   // 접근자 메서드
	return d.month
}

func (d *Date) SetDay(day int) error {
	if day <= 0 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
func (d *Date) Day() int {                   // 접근자 메서드
	return d.day
}
```

main에서 접근자 메서드를 이용해서 값을 전달해보면 잘 되는 것을 볼 수 있습니다.

```go
package main

import (
	"fmt"
	"github.com/GoStudy/캡슐화와 임베딩/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(9)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)
	fmt.Println(date.Month)
	fmt.Prinltn(date.Day)
}
```

> 캡슐화

프로그램의 어느 한 영역에 있는 데이터를 다른 코드로부터 숨기는 것을 캡슐화라고 합니다.

이는 Go만의 기법이나 특징은 아닙니다.

캡슐화는 잘못된 데이터로부터 코드를 보호하는 데 사용할 수 있기 때문에 중요하게 다뤄지며, 또한 데이터에 직접 접근할 수 없기 때문에 캡슐화 된 영역을 수정할 때 다른 코드에 미치는 영향에 대해서도 걱정할 필요가 없습니다.

다른 많은 프로그래밍 언어에서는 데이터를 클래스 내에서 캡슐화합니다. (클래스는 Go의 타입과 유사한 개념이지만 동일하지는 않습니다)

반면 Go에서는 데이터를 패키지 내에서 캡슐화하며 노출되지 않은 변수, 구조체 필드, 함수 및 메서드를 사용해 구현합니다.

캡슐화는 Go 보다는 다른 언어에서 훨씬 더 자주 사용됩니다.

어떤 언어들은 심지어 직접 접근해도 괜찮은 경우임에도 모든 필드에 대해 접근자 및 설정자를 정의하는게 컨벤션인 경우도 있습니다.

하지만 Go 개발자는 필드 데이터의 유효성 검증이 필요한 경우와 같이 꼭 필요한 경우에만 캡슐화를 사용하는 경향이 있습니다.

Go에서는 필드 캡슐화가 필요 없다고 생각되면 일반적으로 필드를 외부에 노출시키고 직접 접근하는 것이 좋습니다.

> 임베딩

Go 언어는 클래스를 제공하지 않으므로 상속 또한 제공하지 않습니다.

하지만 구조체에서 임베딩(Embedding)을 사용하면 상속과 같은 효과를 낼 수 있습니다.

만약 기념일을 저장하는 Event라는 타입을 만든다고 합시다.

해당 타입에는 연, 월, 일과 어떤 기념일인지 적는 내용이 필요하다고 했을 때, 연, 월, 일 부분은 아까 만든 Date 타입을 이용해서 다음과 같이 선언할 수 있습니다.

```go
type Event struct {
	desc string
	Date
}
```

위와 같은식으로 구조체 임베딩을 할 경우 마치 Event 구조체에서 Date의 필드를 선언한 것 처럼 접근할 수 있습니다.

예를들어, 구조체로 선언했을 경우 Event.Date.year 이런식으로 접근해야 하지만 임베딩을 할 경우 Event.year 이런식으로 접근할 수 있습니다.

여기서 주의해야 할점은 숨겨진 필드나 메서드는 승격되지 않기때문에 사용할 수 없습니다.

만약, 임베딩을 했을 때 두 구조체간에 겹치는 메서드(이름이 동일한 메서드)가 존재할 경우 임베딩을 사용한 구조체가 사용됩니다.