# 채널

Golang은 고루틴을 사용하여 병렬 및 동시 프로그래밍을 가능하게 합니다.

앞서 concurrency를 통해 고루틴이 무엇인지, 어떻게 이를 사용할 수 있는지 알아보았습니다.

이번에 알아볼 채널은 고루틴 간의 통신 채널을 의미합니다.

즉, 고루틴이 내부에서 통신하기 위한 매커니즘을 제공합니다. 

## 채널에 대한 소개

채널은 많은 유연성을 제공하고 동시성과 관련된 문제를 해결합니다. 채널은 다음과 같이 요약할 수 있습니다.

1. 채널은 의사 소통을 위한 매커니즘입니다.
2. 채널은 다른 고루틴에 매개변수로 보낼 수 있습니다.
3. 채널은 Publisher and Subscriber 모델로 작동합니다.

### golang을 사용한 메모리 격리

채널의 등장 이전에는 애플리케이션이 데이터를 전역적으로 저장하는데 사용되었으며 다른 스레드/하위 프로세스에 의한 데이터 조작에 위험이 존재했습니다. 이를 데이터 무결성 검사등을 통해 해결할 수 있겠지만, 상당히 번거롭습니다.

이처럼 전역 메모리를 사용하여 서로 다른 쓰레드 혹은 하위 프로세스 간에 데이터 동기화가 이뤄진다면 여러 문제가 유발될 수 있습니다.

Golang은 정보 흐름을 위한 안전한 파이프라인을 채널을 통해 제공합니다. 이때 하나의 하위 프로세스만 해당 데이터에 접근할 수 있습니다. go 채널 내에서 사용 가능한 데이터는 하나의 고루틴에만 접근할 수 있습니다. 이때 데이터를 보내는 사람과 받는 사람 간의 일대일 관계가 형성됩니다.

## 1. 가장 기본적인 채널
```go
package main
import "fmt"

func main() {
	dataChannel := make(chan string)
	fmt.Println(<-dataChannel)
}
```

위의 코드를 살펴보자면

1. 새로운 객체를 만들기 위해 `make` 키워드를 사용
2. `chan`(채널) 이라는 객체 타입을 지정
3. 데이터의 타입을 채널에 의해 반환받는 `string`으로 지정

`dataChannel := make(chan string)`

해당 라인에서 고루틴 내에서 데이터 스트링을 전송하는데 사용할 수 있는 `channel` 유형의 객체를 만들었습니다.

`fmt.Println(<-dataChannel)`

해당 라인에서 채널이 데이터를 수신할 때까지 계속 기다리라고 main 함수에 지시하고 있습니다. 위의 프로그램에서는 채널에 데이터를 보내는 다른 고루틴이 없습니다. 이때 사용가능한 다른 채널이 없기 때문에 해당 프로그램에서는 데이터 수신을 계속 기다리는 교착 상태(Deadlock)에 빠지게 됩니다.
![](https://miro.medium.com/max/689/1*jEJUSd7sX8I08cEGRs-b9Q.png)

## 2. 채널에 데이터 추가

위의 코드에서는 채널을 만들고 채널에서 데이터를 사용할 수 있을 때까지 기다렸습니다. 이때 데이터를 사용할 수 없었기 때문에 교착 상태에 빠진 모습을 볼 수 있었습니다.

교착 상태의 해소를 시도하기 위해 아래의 코드는 채널에 데이터를 공급하고자 합니다.

```go
package main

import "fmt"

func main() {
	dataChannel := make(chan string)
	dataChannel <- "Some Sample Data"
	fmt.Println(<-dataChannel)
}
```

위의 코드에서 채널에 데이터 샘플을 추가하는 모습을 볼 수 있습니다. 그럼 이제 채널에서 데이터를 수신할 수 있는 걸까요?

해당 코드에 대한 출력입니다.
![](https://miro.medium.com/max/641/1*2lYBq2zLVxcKNS5j4rRobQ.png)

출력 결과에서 볼 수 있듯 다시 교착 상태에 빠진 것을 볼 수 있습니다.

위의 코드에서는 채널에 데이터를 추가하고 있습니다. 데이터를 보내자마자 다른 쓰레드가 해당 데이터를 받을 수 있을 때까지 쓰레드가 차단됩니다. 채널에서 데이터를 수신할 수 있는 다른 쓰레드가 없기 때문에 이제 메인 쓰레드가 차단됩니다. 마찬가지로 교착 상태가 발생한 상황입니다.

## 교착 상태 해결

기본적으로 채널에는 저장 용량이 없습니다. 즉, 메시지가 발신자에서 수신자로 즉시 전달되어야 함을 의미합니다.

수신자가 없으면 메시지는 계속 발신자에게 붙어 있습니다.

채널에 버퍼를 추가하여 내부적으로 일부 메세지를 저장할 수 있는 용량을 제공한다면 해당 데이터가 처리되지 않더라도 보낸 사람이 작업을 계속할 수 있습니다.

즉, 버퍼링된 채널을 활용하면 교착상태의 해결이 가능합니다.


## 3. 버퍼링된 채널 만들기

데이터가 반대쪽에서 추출되지 않더라도 쓰레드가 계속 실행되도록 하기 위해서는 채널에 버퍼를 제공해야 합니다.

버퍼를 추가하면 채널에 저장 용량이 제공되며 Pub -> Sub로 넘어갈때, 이를 즉시 사용하고 폐기하지 않아도 됩니다. 코드를 통해 살펴보면 다음과 같습니다.

```go
package main

import "fmt"

func main() {
	dataChannel := make(chan string, 3)
	dataChannel <- "Some Sample Data"
	dataChannel <- "Some Other Sample Data"
	dataChannel <- "Buffered Channel"
	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)
}
```

위의 코드에서는 버퍼가 있는 채널을 만들고 있습니다. 

make의 두번째 인자인 3은 채널이 세 개의 문자열을 저장할 수 있음을 의미합니다.

위의 코드에서 채널은 버퍼링이 되어있기 때문에 채널에 더 많은 데이터를 추가한다면 기본 채널이 좀 더 멀리 이동하게 됩니다.

버퍼가 가득 차면 데이터를 폐기해야 합니다. 그렇지 않으면 교착 상태가 발생합니다.

따라서 버퍼를 추가하면 쓰레드가 채널에 데이터를 저장할 수 있습니다. 이 데이터는 나중에 실행 중에 삭제되어 교착 상태를 제거할 수 있습니다. 출력은 다음과 같습니다.

![](https://miro.medium.com/max/614/1*qO8dp2Aim7ydOpY1Nm4reg.png)

## 버퍼링이 없이는 채널을 쓸 수 없나?

그렇지 않습니다. 버퍼링을 사용한 것은 채널의 속성 중 수신자는 송신자를 항상 기다린다는 특성 때문입니다. 기본적으로 버퍼가 존재하지 않는 Go 채널은 서로를 기다리는 과정에서 현재 진행 중인 쓰레드를 잠시 차단합니다. 그리고 진행 중인 다른 버퍼, 즉 고루틴에서 채널에 적절한 조치를 취해주게 되면 차단이 비로소 풀리고 계속해서 프로그램이 진행되게 되는 것이죠.

만약 채널에 버퍼링이 되어있다면, 하나의 송신이 이뤄졌다고 해당 쓰레드가 차단되지는 않습니다. 본래 채널은 버퍼링처리가 되어 있지 않다면, 수신 받으면 곧바로 송신을 해야 합니다. 데이터 파이프 라인으로써 채널 자체적으로는 전달해주는 역할만 할뿐 데이터를 적재할 수 없습니다. 일단 데이터가 적재된다면, 채널은 해당 쓰레드를 차단합니다. 더 쌓일 수 있는 여지를 막아버게 되죠. 근데 여기서 적재된 것을 해소해주지 않는다면 계속해서 기다리기만 할 것입니다. 바로 여기서 데드락이 발생하죠.

그런데 만약 버퍼링처리를 한다면 버퍼링한 개수만큼은 계속 데이터를 적재할 수 있습니다. 이 말을 좀 더 풀어쓰면, 버퍼가 다 채워지기 전까지는 쓰레드가 차단되지 않는다는 의미입니다. 만약 버퍼가 다 찼는데도 이를 해소(송신)하지 않는다면, 또 다시 데드락이 발생합니다.

이러한 특성을 활용한다면 특정 고루틴이 끝날때까지 특정 라인을 넘어가지 못하도록 강제할 수 있습니다.
```go
package main

import "fmt"

func main() {
	DataChannel := make(chan bool)

	go func() {
		for i := 0; i < 123; i++ {
			fmt.Println(i)
		}
		DataChannel <- true
	}()
	<-DataChannel
}
```

일반적인 고루틴이라면, 메인 쓰레드와 별도로 실행이 되기 때문에 당연히 해당 고루틴의 완료 여부와 무관하게 메인 쓰레드는 계속해서 진행하게 될 것입니다. 혹시 쓰레드와 관련해서 프로그래밍을 진행해보신 경험이 있다면, `wait` 혹은 `sleep`과 같은 메서드를 호출해서 강제로 특정 쓰레드의 완료를 기다리셨던 경험이 있을 것입니다. 

golang에서는 채널을 활용한다면 훨씬 깔끔하게 특정 쓰레드의 완료를 기다릴 수 있습니다. 채널의 특성 때문입니다. go의 채널은 수신자와 송신자가 서로를 기다리는 속성을 가지고 있습니다. 때문에 위의 코드에서는 `DataChannel`에 송신이 이뤄지기 전에는 수신이 불가하므로 익명함수가 온전히 종료된 이후에야 메인 쓰레드가 계속해서 진행될 수 있습니다.

Go의 채널에서는 서로를 기다릴때 현재 진행 중인 쓰레드를 잠시 차단합니다. 

여기까지가 채널의 기본적인 개념이였습니다. 데이터 파이프라인으로써 채널은 go의 동시성 프로그래밍의 핵심 요소입니다. 해당 파트를 이해하는데 OS 지식을 요구하는 부분이 있어 어렵다고 느낄 부분이 있지만, 그럼에도 채널은 go 언어의 핵심이므로 잘 이해할 필요성이 있다고 생각합니다.

# 자주 헷갈리는 상황들

아래는 Channel을 사용하면서 발생하는 대표적인 오류 사례들입니다. 사실 정리를 한다고 했으나, 무엇보다도 이를 익히는 가장 좋은 방법은 실제 프로그래밍 경험이라고 생각합니다. 간단하게 짚고 넘어가고 이후 관련해서 문제가 될때마다 다시 보면 좋을 거 같습니다.

### range는 채널이 close 되어야 끝난다.

채널에서 데이터를 receive하는 방법으로 range가 있습니다. 이때, range는 채널이 close 되어야 끝이 나게 됩니다. 즉, 채널을 close해주지 않는다면 영원히 기다리면서 교착 상태가 발생하게 됩니다.

```go
ch := make(chan int, 1)
ch <- 101
for value := range ch {
	fmt.Println(value)
	// close(ch)
}
```

## close된 채널에는 send 할 수 없다.
```go
ch := make(chan int)
close(ch)
ch <- 1
```
이미 닫힌 채널을 통해 데이터를 전송하려는 시도는 panic을 발생시킵니다.

## close된 채널에서 receive 할 수 있다.

```go
ch := make(chan int, 2)

var wg sync.WaitGroup
wg.Add(1)
go func() {
	ch <- 10
	ch <- 11
	wg.Done()
}()

wg.Wait()
close(ch)
fmt.Println(<-ch) // 10
fmt.Println(<-ch) // 11
fmt.Println(<-ch) // 0
```

반대로 close된 채널에서 recieve는 가능하다.

## Select

### select 문은 무한루프는 아니지만, case가 올 때까지 기다린다.

채널의 특성이 select와 결합하여 발생한다고 이해하면 됩니다. select를 만나고 select의 조건이 채널의 송신과 관련된다면, 해당 채널이 수신되기 전까지는 select문이 계속 기다리게 됩니다.

만약 default가 있다면, 얘기가 달라지긴 하겠습니다. (채널에 수신이 이뤄지지 않더라도 기다리지 않고 default case 진행)

```go
ch := make(chan int)
takeSomeTime := func()  {
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()
}

start := time.Now()
takeSomeTime()
		
select {
	case <-ch:
		fmt.Println(time.Since(start)) // 2.004948752s
}
```

`takeSomeTime`은 2초간 대기 후에 채널로 데이터를 보내는 함수입니다. 해당 데이터를 보내야 case는 값을 receive할 수 있습니다. 만약 select문에서 <-ch 발생을 기다리지 않고 넘어갔다면 Output이 아무것도 출력되지 않겠죠. 하지만 <-ch가 발생할 때까지 기다렸기 때문에 대략 2초 후에 출력되었습니다.

## Select에서 case는 순차적으로 실행되지 않는다.

```go
c1 := make(chan interface{})
close(c1)
c2 := make(chan interface{})
close(c2)

var c1Count, c2Count int
for i := 1000; i >= 0; i-- {
	select {
	case <-c1:
		c1Count++
	case <-c2:
		c2Count++
	}
}

fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
```

```go
Output
c1Count: 519
c2Count: 482
```
select에서 case는 무작위로 실행됩니다. 각 case는 모두 비슷한 확률로 실행됩니다. (균일한 의사 무작위 선택, 여기서는 50%)

앞서 설명했듯이, close()를 해도 receive할 수 있기 때문에 위의 예는 문제 없이 작동하게 되는 모습입니다.

## case에 함수가 있다면, 그 함수가 끝난 후 다른 케이스를 검사한다.
```go
select {
	case <-ch:
	case ch <- doSomething():
}
```
channel로 값이 오는 것을 기다리는 것 외에도 channel에 값을 전달하는 식으로도 select문을 이용할 수 있습니다.
이때, doSomething이 끝날 때까지 다른 case는 검사하지 않습니다.

```go
longFunction := func() interface{} {
	defer fmt.Println("end long function")
	fmt.Println("start long function")
	time.Sleep(time.Second)
	return nil
}

shortFunction := func() interface{} {
	defer fmt.Println("end short function")
	fmt.Println("start short function")
	return nil
}

chan1 := make(chan interface{},10)
chan2 := make(chan interface{},10)

for i := 0; i < 10; i++ {
	select {
		case chan1 <- longFunction():
		case chan2 <- shortFunction():
	}
}
```
output
```
start long function
end long function
start short function
end short function
start long function
end long function
start short function
end short function
start long function
...
```

소요시간이 긴 long function이 끝난 이후 다른 case의 short function이 실행되는 모습을 볼 수 있습니다. select case에서 값을 send하는 경우는 함수 실행이 종료되어야 그 다음 케이스를 검사한다고 이해할 수 있겠습니다.


## case에 있는 함수는 끝까지 실행하지만, 적절하지 않은 상황이라면 다른 case를 실행한다.

```go
select {
	case chan1 <- funcA():
	case chan2 <- funcB():
}
```

위와 같은 상황에서 funcA를 먼저 처리한다고 가정해봅시다. 그러면 funcA는 끝까지 실행이 되겠죠. 그런데 끝까지 실행하고 보니 chan1이 데이터를 받지 못하는 상황(버퍼가 꽉찬 채널)일 수 있습니다. 그러면 funcB()로 차례가 넘어가게 됩니다.

```go
chan1 := make(chan interface{})
chan2 := make(chan interface{},10)

functionA := func() interface{} {
	defer fmt.Println("end a function")
	fmt.Println("start a function")
	return nil
}

functionB := func() interface{} {
	defer fmt.Println("end b function")
	fmt.Println("start b function")
	return nil
}

for i := 0; i < 10; i++ {
	select {
	case chan1 <- functionA():
		fmt.Println("case A running")
	case chan2 <- functionB():
		fmt.Println("case B running")
	}
}
```

위의 예제를 보면 channel에 데이터를 보내기는 하지만 받지는 않습니다. ch1의 경우 버퍼가 없기 떄문에 데이터를 보낼 수 없고, ch2의 경우에는 버퍼가 넉넉하게 있기 때문에 데이터를 보낼 수 있는 상황입니다. 위 코드를 실행하면 다음과 같은 결과가 나오게 됩니다.

Output
```go
start a function
end a function

start b function
end b function

case B running

start a function
end a function

start b function
end b function

case B running
```

case의 functionA(), functionB() 모두 실행되지만 실질적으로 채널에 데이터를 보내는 것은 functionB()만 가능하다는 것을 볼 수 있었습니다.