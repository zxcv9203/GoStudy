# Go 패키지

Go는 패키지를 통해 코드의 모듈화, 코드의 재사용 기능을 제공합니다.

Go는 패키지를 사용해서 작은 단위의 컴포넌트를 작성하고, 이러한 작은 패키지들을 활용해서 프로그램을 작성할 것을 권장합니다.

컴포넌트란 프로그래밍에 있어 재사용이 가능한 각각의 독립된 모듈을 의미합니다.

Go는 실제 프로그램 개발에 필요한 많은 패키지들을 표준 라이브러리로 제공합니다.

이런 표준 라이브러리 패키지들은 `GOROOT/pkg` 안에 존재합니다. `GOROOT`는 환경변수이며, Go를 설치한 디렉터리를 가리킵니다. 따로 지정해주지 않으면 보통 Go 설치시 자동으로 추가됩니다.

예를들어, 윈도우에서 Go를 설치하면 기본적으로 `C:\go` 에 설치되며 `GOROOT`는 해당 경로를 가리킵니다.

Go에서 사용하는 표준 패키지는 [https://golang.org/pkg](https://golang.org/pkg) 에 자세히 설명되어있습니다.

> main 패키지

일반적으로 패키지는 라이브러리로서 사용되지만 `main`이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식됩니다.

패키지 명이 `main`이면 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(excutable) 프로그램으로 만듭니다.

그리고 이 `main` 패키지 안의 `main()` 함수가 프로그램의 시작점이됩니다.

패키지를 공유 라이브러리로 만들때는 `main` 패키지나 `main` 함수를 사용하면 안됩니다.

> 패키지 Import

현재 프로그램에서 다른 패키지를 사용하기 위해서는 `import` 를 사용하여 패키지를 포함 시켜야 합니다.

예를들어 Go의 표준 라이브러리인 `fmt` 패키지를 사용하기 위하여 `import "fmt"` 와 같이 해당 패키지를 포함시킬 것을 선언해 줍니다.

```go
package main

func main() {
  fmt.Println("Hello")
}
```

위의 코드와 같이 해당 패키지를 선언하지 않고 사용하려고 할 경우 패키지를 찾지 못해 에러가 발생합니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/ea56414c-9d16-4e5a-a769-18e7bf7bf0ac/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124032Z&X-Amz-Expires=86400&X-Amz-Signature=7230e9c5ee56f176f0d9a2b362992ab1ec31ee0c0eae950c5921f46a850e7ce5&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

해당 에러는 다음과 같이 패키지를 임포트 시켜주면 해결되는 것을 볼 수 있습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/f56f57ed-8b3a-4dab-b21c-73a80078f9e0/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124053Z&X-Amz-Expires=86400&X-Amz-Signature=d857787177083e792aabbe715d4fc337be658ebf336abe6c80510ae761b20e28&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> 패키지 Import 경로

패키지를 import할 때 Go 컴파일러는 `GOROOT` 혹은 `GOPATH` 환경변수를 검색하는데, 표준패키지는 `GOROOT/pkg` 에서 찾고 사용자 패키지는 써드파티 패키지의 경우 `GOPATH/pkg`에서 패키지를 찾게 됩니다.

GOPATH란 Go 프로젝트를 진행하는데 사용하는 작업 공간의 루트를 지정하는 환경변수입니다.
ex ) GOPATH = /home/go

한 번 사용자가 직접 패키지를 만들고 안에 간단한 함수를 작성해보고 실행시켜 보겠습니다.

같은 폴더에는 서로 다른 패키지가 존재할 수 없으므로 다음 그림과 같은 형태로 만들어 보겠습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/4ddc5f5c-9e2a-40d2-9f0b-abba165fd363/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124112Z&X-Amz-Expires=86400&X-Amz-Signature=a6bf38bc4200bbb3e6ab937311c902b17de37dfd4b9bdd62c8f655f2cd8e4ee9&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

- path.go

    ```go
    package pack

    import "fmt"

    func hello() {
    	fmt.Println("Hello Go")
    }
    ```

- main.go

    ```go
    package main

    import "github.com/pack"

    func main() {
    	pack.hello()
    }
    ```

위의 코드의 `main.go`를 실행시키면 잘 되야 할 것 같지만 실제로는 다음과 같은 에러 메시지를 출력합니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/c3f7d855-ef83-401e-9ef3-209274143a8c/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124130Z&X-Amz-Expires=86400&X-Amz-Signature=dc6e67cb4f1c17dc6ec3a71d916a7638ea09057589a839c770878ab569e4e8ad&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

외부로 노출시키지 않은 이름은 참조할 수 없다고 나옵니다. 해당 문제는 간단하게 해결할 수 있습니다.

> 패키지 scope

패키지 내에는 함수, 구조체 인터페이스, 메서드 등이 존재하는데, 이들의 이름이 첫문자를 대문자로 시작하면 이는 외부로 노출 시킬 수 있습니다.

하지만 이름이 소문자로 시작하면 해당 패키지 내부에서만 사용할 수 있습니다.

즉, 위에서 코드를 `main.go`를 실행시켰을 때 에러가 발생한 이유는 사용하려는 함수가 첫 글자가 소문자로 작성되서 외부로 노출되지 않았기 때문입니다.

위의 `main.go`와 `path.go`를 다음과 같이 수정하고 다시 실행 시켜봅시다.

- main.go

    ```go
    package main

    import "github.com/pack"

    func main() {
    	pack.Hello()
    }
    ```

- path.go

    ```go
    package pack

    import "fmt"

    func Hello() {
    	fmt.Println("Hello Go")
    }
    ```

실행 시켜보면 다음과 같이 잘 실행되는 것을 볼 수 있습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/b020cedf-0e1e-4771-9e1b-8a75daca8e6e/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124153Z&X-Amz-Expires=86400&X-Amz-Signature=f39a908cedd34c9dc5aa83a8770883adc3d35d726a5a4706f3b3db53924a086f&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> init 함수

개발자가 패키지를 작성할 때 패키지 실행 시 처음으로 호출되는 init() 함수를 작성할 수 있습니다.

init() 함수 같은 경우 패키지가 로드되면서 실행되는 함수로 별도의 호출 없이 자동으로 호출됩니다.

다음은 init() 함수 예제입니다.

```go
package main

import "fmt"

var n int

func init() {
	n = 10
}

func main() {
	fmt.Println(n)
}
```

위의 코드를 실행 시켜보면 init 함수를 호출하지 않았음에도 n의 값이 10으로 변경되는 것을 알 수 있습니다.

이처럼 `init()` 을 사용하면 따로 호출하지 않아도 제일 먼저 실행됩니다.

그럼 여러 패키지에서 `init()` 을 사용할 때 순서는 어떻게 진행될까요?

Go 프로그램은 항상 main() 함수로 시작이됩니다. 만약 main 패키지가 다른 패키지를 임포트하고 있으면, 임포트된 각각의 패키지를 먼저 불러옵니다.

임포트된 패키지에서 또 다른 패키지를 임포트하고 있으면 패키지를 불러옵니다.

임포트 되는 모든 패키지를 불러온 후에 main() 함수가 실행됩니다.

간단한 예제를 통해 직접 확인해 보겠습니다.

다음과 같은 구조와 코드를 가진 프로젝트가 있다고 가정해봅시다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/be2407e7-e3ce-48df-9009-822a177153b7/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124214Z&X-Amz-Expires=86400&X-Amz-Signature=bb725fe654399377fe0b41bc7d4846132ce5c2d6a717f9c3c1c6f61657457ade&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

- pack1.go

    ```go
    package pack1

    import (
    	"fmt"
    )

    func init() {
    	fmt.Println("pack1 init")
    }
    func Call() {
    	fmt.Println("pack1 called")
    }
    ```

- pack2.go

    ```go
    package pack2

    import (
    	"fmt"
    	"github.com/pack1"
    )

    func init() {
    	fmt.Println("pack2 init")
    }
    func Call() {
    	fmt.Println("pack2 called")
    	pack1.Call()
    }
    ```

- pack3.go

    ```go
    package pack3

    import (
    	"fmt"
    	"github.com/pack2"
    )

    func init() {
    	fmt.Println("pack3 init")
    }
    func Call() {
    	fmt.Println("pack3 called")
    	pack2.Call()
    }
    ```

- main.go

    ```go
    package main

    import (
    	"fmt"
    	"github.com/pack3"
    )
    func init() {
    	fmt.Println("main.go init")
    }
    func main() {
    	fmt.Println("main.go main start")
    	pack3.Call()
    }
    ```

위의 main.go를 실행시켜보면 다음과 같은 결과가 나옵니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/bb26ad81-3500-4fa2-9a6d-5954a8fc6318/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124243Z&X-Amz-Expires=86400&X-Amz-Signature=caa256d6d54e1de36ffe8c96c8309d9184d9fd5d5de860cbf0efa1108ead03e2&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

위의 순서를 init 순서를 그림으로 표현하면 다음과 같습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/127553c0-6426-4488-953d-46ad6dad8292/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T124320Z&X-Amz-Expires=86400&X-Amz-Signature=18d4a63277fbc8b2a09151b23d975d3915e02798961e98b41a46807c32250763&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

1. main.go를 실행시키면 임포트된 패키지(pack3)를 먼저 가져옵니다.
2. pack3.go는 임포트된 패키지 pack2를 가져옵니다.
3. pack2.go는 임포트된 패키지 pack1을 가져옵니다.
4. pack1.go는 init() 함수를 실행합니다.
5. pack2.go는 init() 함수를 실행합니다.
6. pack3.go는 init() 함수를 실행합니다.
7. main.go는 init()함수를 실행합니다.
8. main.go에서 main 함수를 시작합니다.

주의사항 : 패키지를 처음 import 할때만 init을 실행합니다.

위에서는 패키지를 사용하지 않아 Call 함수를 만들어 사용하도록 했는데 임포트한 패키지를 사용하지 않고 init만 쓰고 싶을 때는 `alias` 를 이용하면 됩니다.

> 패키지 alias

패키지를 불러올때 안의 내용은 사용하지 않고 init 함수만 호출하는 방법은 없을까요?

이럴 때는 언더바(`_`)를 사용해서 `init()` 만 호출할 수 있습니다.

```go
import (
	_ "github.com/pack3"
)

func main() {
	
}
```

원래는 패키지를 사용하면 에러가 발생했지만 에러가 발생하지 않고 `init()` 을 잘 호출하는 것을 볼 수 있습니다.

또한 패키지 alias는 패키지 이름이 동일할때 alias를 이용해서 별칭을 지어 중복을 피해줄 수 있습니다.

```go
// 예시를 들기 위한 코드입니다.
import (
	mongo "database/mongo/db"
	mysql "database/mysql/db"
)
func main() {
	mongo.conn()
	mysql.conn()
}
```
