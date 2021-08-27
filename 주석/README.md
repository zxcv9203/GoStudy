# 주석

## 지원하는 코드 스타일 
1. `/* */` 블럭 주석 (C style)
2. `//` 한줄 주석 (C++ style)

모든 패키지는 패키지 구문 이전에 블럭 주석 형태의 패키지 주석이 있어야 한다. 여러 파일로 구성된 패키지의 경우, 패키지 주설은 어느 파일이든 상관없이 하나의 파일에 존재하면 되고 해당 주석이 사용된다. 패키지 주석을 작성할 떄는, 패키지 전체를 소개하면서, 전체 패키지에 관련된 정보가 들어가야 한다. 패키지 주석은 godoc 문서의 처음에 나타나게 되므로 이후의 자세한 사항을 작성해도 좋다.

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

패키지가 단순하다면 패키지 주석 또한 간단할 수 있다.
```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

주석을 더 깔끔하게 표현하기 위해 줄을 긋는 식의 지나친 포맷은 필요하지 않다. 

오히려 생성된 출력이 고정폭의 폰트로 주어지지 않을 수 있으므로, 스페이스나 정렬등에 의존하지 않아도 된다. 

주석의 가독성 부분은 gofmt가 그랬던 것처럼 godoc이 처리한다. 

주석은 해석되지 않는 일반 텍스트이므로 HTML이나 _this_ 같은 주석은 사용하지 않는 것이 좋다.

godoc은 수정하는 한가지는 들여쓰기된 텍스트를 고정폭의 폰트로 보여주는 것으로, 프로그램 코드 조각 등에 적합하다. [fmt package](https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/fmt/doc.go)의 패키지주석은 좋은 예이다.

문맥에 따라, godoc은 적절한 리포맷이 되지 않을 수 있다. 그래서 코드를 보기에 직관적으로 잘 작성해야 한다.(정확한 스펠링, 구두법, 문장구조, 긴문장의 최소화)

패키지에서 최상위 선언의 바로 앞에 있는 주석이 그 선언의 문자주석으로 처리된다. 패키지 내부에서 최상위 선언 바로 이전의 주석은 그 선언을 위한 doc comment이다. 프로그램에서 exported되는 모든 (대문자로 시작되는) 이름은 doc comment가 필요하다.

doc comment은 매우 다양한 자동 프레젠테이션들을 가능케 하는 완전한 문장으로 작성될 때 가장 효과적이다. 첫 문장은 선언된 이름으로 시작하는 한 줄 짜리 문장으로 요약되어야 한다.

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

모든 문서 설명이 설명하는 항목의 이름으로 시작하는 경우 go 도구의 doc 하위 명령을 사용하여 grep를 통해 출력을 실행할 수 있다. "Compile"이라는 이름은 기억하지 못하지만 정규식에 대한 구문 분석 함수를 찾고 있어서 명령을 실행했다는 상황을 가정해보자.

- 다음과 같이 명령어를 입력할 수 있다.
```
$  go doc -all regexp | grep -i parse
```

만약 패키지 시작 지점의 모든 doc comment가 "This function..." 하고 실행된다면, grep 명령어가 유효하지 않겟지만, doc comment가 패키지 별로 맨 앞에 패키지 명을 쓰도록 규칙을 정한다면 찾고자 하는 단어를 찾을 수도 있을 것이다.

- 명령어 결과
```
$ go doc -all regexp | grep -i parse
    Compile parses a regular expression and returns, if successful, a Regexp
    MustCompile is like Compile but panics if the expression cannot be parsed.
    parsed. It simplifies safe initialization of global variables holding
```

GO 언어의 선언 구문읜 그룹화가 가능하다. 단일 doc comment는 연관된 상수 혹은 변수가 포함된 그룹을 소개할 수 있을 것이다. 전체 선언문이 제시되기에 이러한 주석은 형식적일 수 있다.

> 그룹화는 아이템들 간의 관계일 수 있다. 
> 
> (mutex에 의해 보호되고 있다는 점에서 그룹을 이루는 변수들)
>```go
>var (
>    countLock   sync.Mutex
>    inputCount  uint32
>    outputCount uint32
>    errorCount  uint32
>)
>```