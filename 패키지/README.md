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

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/ea56414c-9d16-4e5a-a769-18e7bf7bf0ac/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T073018Z&X-Amz-Expires=86400&X-Amz-Signature=58dc4030ebb2f71d28b1e330133dcf250231f66d7f035daca1659f445a39c33d&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

해당 에러는 다음과 같이 패키지를 임포트 시켜주면 해결되는 것을 볼 수 있습니다.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/f56f57ed-8b3a-4dab-b21c-73a80078f9e0/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T073047Z&X-Amz-Expires=86400&X-Amz-Signature=50b09a403059a73b03379a040208f377d56806b0b1648372a996b2396eccb536&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

> 패키지 Import 경로

패키지를 import할 때 Go 컴파일러는 `GOROOT` 혹은 `GOPATH` 환경변수를 검색하는데, 표준패키지는 `GOROOT/pkg` 에서 찾고 사용자 패키지는 써드파티 패키지의 경우 `GOPATH/pkg`에서 패키지를 찾게 됩니다.

GOPATH란 Go 프로젝트를 진행하는데 사용하는 작업 공간의 루트를 지정하는 환경변수입니다.
ex ) GOPATH = /home/go

한 번 사용자가 직접 패키지를 만들고 안에 간단한 함수를 작성해보고 실행시켜 보겠습니다.

같은 폴더에는 서로 다른 패키지가 존재할 수 없으므로 다음 그림과 같은 형태로 만들어 보겠습니다.

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/4ddc5f5c-9e2a-40d2-9f0b-abba165fd363/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T073109Z&X-Amz-Expires=86400&X-Amz-Signature=f0d9e016bdf748c9c2edaacaf841a96d1246c721dd12713f27a3eb6a11ba788d&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

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

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/c3f7d855-ef83-401e-9ef3-209274143a8c/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T073128Z&X-Amz-Expires=86400&X-Amz-Signature=8f3f64ff7cc4424b3b3a49f8f9afd59b1d9e427a442797857b38464ce3d53890&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

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

![Untitled](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/b020cedf-0e1e-4771-9e1b-8a75daca8e6e/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210827%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210827T073152Z&X-Amz-Expires=86400&X-Amz-Signature=b9fea2d8043cfdcfdd767c93b4824a02a9c69cf9be1c931e3a9b6eb904871def&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

---

추가하면 좋을거 같은것

go get (패키지 불러오기)

init 함수, alias