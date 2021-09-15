# test

Go는 간편하게 사용할 수 있는 테스트 프레임워크를 내장하고 있는데 go test 명령을 실행하여 테스트 코드를 실행합니다.

그덕분에 Go는 TDD 방법론으로 개발을 하는데 적합합니다.

> TDD(Test-Driven-Development) 방법론

TDD란 Test Driven Development 의 약자로 테스트 주도 개발이라고 합니다.

반복되는 테스트를 이용하는 소프트웨어 방법론으로, 작은 단위의 테스트 케이스를 작성하고 이를 통과하는 코드를 추가하는 단계를 반복하여 구현합니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/ff4d1bb1-e50d-4400-8252-6cb26a639f1a/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105822Z&X-Amz-Expires=86400&X-Amz-Signature=9566e3eecf2ea33e917f6c8b13f5b81ca056a920a483a98b02f5d5634ddc531f&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> TDD를 하는 장점

- 객체지향적인 코드를 개발할 수 있습니다.

    테스트 코드를 먼저 작성한다면 좀 더 명확한 기능과 구조를 설계할 수 있습니다.

    각각의 함수를 정의할 때 각각의 기능들에 대해서 철저히 구조화 시켜 코드를 작성하게 됩니다.

    그 이유는 테스트의 용이성을 위해 복잡한 기능을 한 함수에 모두 구현할 경우 테스트 방식이 복잡해지고 시간이 오래 걸리며 코드 수정이 되는 경우 테스트 코드를 재 사용할 수 없게 되기 때문입니다.

    자연스럽게 TDD의 목적인 코드의 재사용성을 보장하며 코드를 작성하게 됩니다.

- 설계의 문제점을 빠르게 찾아낼 수 있습니다.

    작은 단위로 나눠서 테스트를 해가면서 개발하기 때문에 최초 설계 안을 만족시키며 입출력 구조와 기능의 정의를 명확히 하게 되므로 설계의 구조적 문제를 바로 찾아내게 됩니다.

    실제로 테스트 코드를 작성해보면서 인터페이스나 클래스의 구조들을 많이 수정하게 됩니다.

    그리고 미리 테스트 시나리오를 작성해봄으로써 코드 개발전 기능을 구현하기 위한 예외상황들을 미리 확인해보고 테스트하게 되는 효과가 발생하여 예외코드를 작성하기 쉬워집니다.

- 디버깅 시간의 단축

    기본적으로 단위 테스트 기반의 테스트 코드를 작성하기 때문에 추후 문제가 발생하였을 때 각각의 모듈 별로 테스트를 진행해보면 문제의 지점을 쉽게 찾아낼 수 있습니다.

    만약 TDD 개발이 아니라면 특정 버그를 찾기 위해서 모든 영역의 코드를 살펴봐야 할 것입니다.

    문제가 발생할 수 있는 지점은 DB 영역, Application 영역, Data 영역, Memory 영역 등 다양하기 때문에 모든 영역을 통합 테스트하게 되면 쉽게 문제의 지점을 찾을 수 없게 됩니다.

    하지만 TDD 개발로 인해 각각의 단위 테스트를 진행하게 된다면 영역을 분할하여 쉽게 찾아낼 수 있을 것 같습니다.

- 유지보수의 용이성

    대부분의 개발자는 설계 및 코드 작성 시 기술적인 관점으로 바라보게 됩니다.

    기술적인 관점이 나쁜 것은 아니지만 기능 자체의 실현에 목적을 두기 때문에 코드가 복잡해지고 테스트가 어려워집니다.

    TDD 개발로 인해 항상 그 테스트 요소들이 사용자 관점으로 정의되고 진행되기 때문에 입력과 출력의 흐림이 명확해지고 추후 구조의 변경 및 소스 수정 시 구조를 쉽게 파악하고 빠른 수정이 가능해집니다.

    더불러 재사용 테스트도 쉽게 가능해집니다.

- 테스트 문서의 대체 가능

    대부분의 개발 프로젝트 진행 시 테스트를 진행하는 경우 단순 통합 테스트의 지나지 않습니다.

    즉, 내부적으로 개별적인 모듈들이 어떻게 테스트 되었는지 제공할 수 없습니다.

    하지만 TDD를 구현하게 될 경우에 테스팅을 자동화 시킴과 동시에 보다 정확한 테스트 근거를 산출할 수 있습니다.

> 자동 테스트가 중요한 이유

결국, 사용자는 개발자가 개발만 프로그램을 사용하고 코드가 문제 있을 경우 큰 문제가 발생할 수도 있습니다.

그렇기에 개발자는 만든 프로그램에 하나하나 테스트를하며 문제가 없는지 찾고 해결해야 합니다.

이러한 과정을 자동으로 테스트를 하면 시간을 훨씬 절약할 수 있습니다.

> go test를 이용한 테스트 작성하기

go test 명령을 사용하면 테스트 코드를 실행할 수 있습니다.

go test는 다음과 같이 사용합니다.

```go
go test <경로>
```

위의 명령을 입력하면 해당 경로의 모든 test 파일을 실행합니다.

test 파일의 이름은 `*_test.go` 같은 식으로 작성하며 `*` 는 어떤 문자가 와도 상관 없습니다.

예를들어 list.go 라는 파일의 테스트 파일은 list_test.go 이런식으로 작성합니다.

테스트 파일은 `testing` 이라는 표준 패키지를 사용하는데 먼저 testing 패키지를 import하고, 테스트 메서드를 작성합니다.

테스트 함수를 `Test~~~`와 같은 특별한 형태를 갖습니다.

예를들어 TestAdd 이런식으로 작성합니다.

또한 테스트함수는 매개변수로 *testing.T를 받습니다.

테스트 함수는 다음 형태가 기본적인 형태입니다.

```go
func TestAdd(t *testing.T) {
	//code
}
```

> go test로 테스트 작성해보기

문자열을 슬라이스로 여러개 받아 문자열을 하나로 합쳐서 콤마로 구분하고 마지막은 and로 구분하는 패키지를 하나 만들어보겠습니다.

예를들어 `hello` `world` `go` 라는 문자열을 받으면 `hello, world and go` 를 출력합니다.

파일 구조는 다음과 같습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/9fd75239-4fd3-4b08-9689-a89f43f0b63c/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105833Z&X-Amz-Expires=86400&X-Amz-Signature=caaf89a747233807841974257030f09beb7539ee199ad6dd2d40a9bb517fc71b&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

- addComma.go

```go
package comma

import "strings"

func JoinWithCommas(parses []string) string {
	result := strings.Join(parses[:len(parses)-1], ", ")
	result += parses[len(parses)-1]
	return result
}
```

- main.go

```go
package main

import (
	"fmt"
	"github.com/GoStudy/test/comma"
)

func main() {
	parses := []string{"hello", "world", "go"}
	fmt.Println(comma.JoinWithCommas(parses))
}
```

이번에는 JoinWithCommas함수가 잘 동작하는 지 확인하기 위해서 test 파일을 한번 작성해보겠습니다.

- addComma_test.go

```go
package comma

import "testing"

func TestTwoString(t *testing.T) {

}
func TestMultipleString(t *testing.T) {

}
```

`go test <경로>` 를 입력해서 위의 테스트 파일을 실행시키면 다음 메시지를 확인할 수 있습니다.
위의 구조와 같다면 `go test github.com/comma` 를 입력하면됩니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/a24ee6e3-77cf-4f84-9bf5-13c8aae17f3a/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105847Z&X-Amz-Expires=86400&X-Amz-Signature=579c2fada4d2ec83a99b59c7e24a19f4e662c7c9b2b57375a896e8e788f30e8b&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

이제 test 파일에 테스트 케이스를 직접 추가해봅시다.

```go
func TestTwoString(t *testing.T) {
	list := []string{"hello", "world"}
	want := "hello and world"
	got := JoinWithCommas(list)
	if got != want {
		t.Error("not matched value")
	}
}
func TestMultipleString(t *testing.T) {
	list := []string{"hello", "world", "go"}
	want := "hello, world and go"
	got := JoinWithCommas(list)
	if got != want {
		t.Error("not matched value")
	}
}
```

이번에도 go test를 이용해서 실행 시켜보면 잘 실행되는 것을 볼 수 있습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/a24ee6e3-77cf-4f84-9bf5-13c8aae17f3a/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105902Z&X-Amz-Expires=86400&X-Amz-Signature=4ac2f34f34c199f619a54e78a37e7cfbd0bb7750017f2582fbe23397a153a4e8&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

하지만 위의 함수는 사실 한가지 문제점이 있는데 문자가 하나 올경우 정상적인 동작을 하지 않는다는 것입니다.

```go
func TestOneString(t *testing.T) {
	list := []string{"hello"}
	want := "hello"
	got := JoinWithCommas(list)
	if got != want {
		t.Error("not matched value")
	}
}
```

위의 함수를 추가해서 다시 실행시켜보면 테스트는 실패하게 됩니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/fe5127f6-b817-4e3b-9261-c5ef2c6477ee/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105915Z&X-Amz-Expires=86400&X-Amz-Signature=823b24d85d42859ca8d1f3e7c3ac80abbcf1c19c61183687de0412f5342985a7&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

테스트가 실패한 이유는 문자열이 한개인 경우 and를 포함해서 출력하기 때문입니다.

go test를 이용하면 이렇게 간편하게 코드를 테스트해볼 수 있습니다.

> go test의 출력결과를 명확하게 표시해주기

하지만 위의 코드 같은 경우 에러 메시지가 불친절합니다.

어떤 케이스에서 문제가 생겼는지 테스트 함수에서 유추해볼 수밖에 없고 어떤식으로 출력이 되는지 확인할 방법도 없습니다.

좀 더 명확하게 작성하기 위해 Errof 메서드를 사용해 자세히 적을 수 있습니다.

addComma_test.go를 다음과 같이 수정을 한 번 해봅시다.

```go
package comma

import "testing"

func TestOneString(t *testing.T) {
	list := []string{"hello"}
	want := "hello"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
	}
}
func TestTwoString(t *testing.T) {
	list := []string{"hello", "world"}
	want := "hello and world"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
	}
}
func TestMultipleString(t *testing.T) {
	list := []string{"hello", "world", "go"}
	want := "hello, world and go"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
	}
}
```

go test로 실행시켜보면 에러 메시지를 좀 더 명확하게 확인할 수 있습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/ecd137cc-576a-40e4-8ea5-ad94fe161afa/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105927Z&X-Amz-Expires=86400&X-Amz-Signature=8093613e3b5dfa8a926f20351de7f7399f1ab059a3fc534ee2bf4b7af57d419d&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> helper 함수

테스트 파일에는 테스트 함수만 작성할 수 있는 것은 아닙니다.

테스트 코드간에 반복되는 코드를 줄이기 위해서 테스트 파일에서 해당 코드만 따로 헬퍼(helper)함수로 분리할 수 있습니다.

헬퍼 함수를 사용하는 방법은 간단한데, Test로 함수가 시작하지만 않으면 사용이 가능합니다.

이번에 만든 addComma_test.go 파일에서 에러메시지 부분만 따로 분리하고 싶을 때 다음과 같이 헬퍼함수로 만들어서 분리할 수 있습니다.

```go
package comma

import (
	"fmt"
	"testing"
)

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)

}
func TestOneString(t *testing.T) {
	list := []string{"hello"}
	want := "hello"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}
func TestTwoString(t *testing.T) {
	list := []string{"hello", "world"}
	want := "hello and world"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}
func TestMultipleString(t *testing.T) {
	list := []string{"hello", "world", "go"}
	want := "hello, world and go"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}
```

> go test의 유용한 옵션

go test는 플래그를 주어 테스트를 할때 몇가지 옵션을 줄 수 있습니다.

예를 들어 -v 옵션을 사용하면 모든 테스트 함수의 테스트 결과가 출력됩니다.

보통 통과한 테스트는 아무 값도 출력하지 않지만 -v 플래그를 사용하면 전부 확인할 수 있습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/637debd3-7111-4941-8642-3c408d3ccc9c/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105938Z&X-Amz-Expires=86400&X-Amz-Signature=e280033a8a7a39be4a573ff319224b1ecc55cafef9dda34fc947b88efe766aa9&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

테스트 파일의 특정 함수만 실행시키고 싶은 경우 `-run` 플래그로 특정 함수만 테스트할 수 있습니다.

예를들어 `-run One` 식으로 실행시키면 One 이라는 문자열이 포함된 함수만 실행됩니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/675a7955-4d63-4f32-afba-17cda1cf8395/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210915%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210915T105949Z&X-Amz-Expires=86400&X-Amz-Signature=869b03a9f859ccd91d68437c8c6a0dc6ac888285136fdee4c7b010edacc6a61b&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> 테이블 주도 테스트

```go
package comma

import (
	"fmt"
	"testing"
)

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)

}
func TestOneString(t *testing.T) {
	list := []string{"hello"}
	want := "hello"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}
func TestTwoString(t *testing.T) {
	list := []string{"hello", "world"}
	want := "hello and world"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}
func TestMultipleString(t *testing.T) {
	list := []string{"hello", "world", "go"}
	want := "hello, world and go"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}
```

현재 만들어둔 테스트 파일을 보면 겹치는 부분이 많이 존재합니다.

함수를 각각의 개별로 관리하는 대신 입력 데이터와 해당 입력에 대한 예상 값을 테이블 형태로 만들어 하나의 테스트 함수에서 테이블의 각 행을 테스트하는 방식으로 변경할 수 있습니다.

테이블 형태에 대한 표준은 없지만 흔히 사용되는 방법은 해당 테스트에서만 특수하게 사용할 새로운 타입을 정의하여 각 테스트에 대한 입력과 예상 출력 값을 저장하는 것입니다.

```go
package comma

import (
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{"hello"}, want: "hello"},
		testData{list: []string{"hello", "world"}, want: "hello and world"},
		testData{list: []string{"hello", "world", "go"}, want: "hello, world and go"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}
```

이렇게 코드를 작성하면 기존 코드보다 중복되는 부분도 훨씬 적고 데이터를 추가하기도 훨씬 쉽습니다.(함수 추가가 아닌 슬라이스에 케이스를 하나 넣으면 됨)